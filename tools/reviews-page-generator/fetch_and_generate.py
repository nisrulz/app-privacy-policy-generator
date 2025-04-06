import os
import requests
import chevron
import markdown
import json
import time
import re
from PIL import Image
from datetime import datetime, timedelta
import argparse


# Parse command-line arguments
def parse_arguments():
    parser = argparse.ArgumentParser(
        description="Fetch GitHub issue comments and generate HTML page."
    )
    parser.add_argument(
        "--force-fetch",
        action="store_true",
        help="Force fetching of fresh comments from GitHub API",
    )
    return parser.parse_args()


# Get arguments
args = parse_arguments()
FORCE_FETCH = args.force_fetch  # Set based on script argument


# GitHub repo details and issue number
OWNER = "nisrulz"
REPO = "app-privacy-policy-generator"
ISSUE_NUMBER = "65"
URL = f"https://api.github.com/repos/{OWNER}/{REPO}/issues/{ISSUE_NUMBER}/comments"
HEADERS = {"User-Agent": f"{REPO}"}

# Directory constants
JSON_DIR = "comments_json"
PROFILE_PIC_DIR = "profile_pictures"
DOWNLOADED_IMAGES_DIR = "downloaded_images"
ONE_WEEK_AGO = datetime.now() - timedelta(weeks=1)


# Create necessary directories
def create_directories():
    os.makedirs(JSON_DIR, exist_ok=True)
    os.makedirs(PROFILE_PIC_DIR, exist_ok=True)
    os.makedirs(DOWNLOADED_IMAGES_DIR, exist_ok=True)


# Fetch comments from GitHub API or load from cached JSON files
def fetch_comments():
    params = {"per_page": 100, "page": 1}
    comments = []

    while True:
        json_file = os.path.join(JSON_DIR, f"comments_page_{params['page']}.json")
        fetch_from_network = should_fetch_from_network(json_file)

        if fetch_from_network:
            page_comments = fetch_comments_from_api(params)
            if not page_comments:
                break
            save_comments_to_file(json_file, page_comments)
        else:
            page_comments = load_comments_from_file(json_file)

        comments.extend(page_comments)
        params["page"] += 1

        if not page_comments:
            break

    comments.reverse()
    return comments


# Check if the comments should be fetched from network
def should_fetch_from_network(json_file):
    if FORCE_FETCH:
        return True  # Always fetch if force mode is enabled

    if os.path.exists(json_file):
        file_mod_time = datetime.fromtimestamp(os.path.getmtime(json_file))
        if file_mod_time < ONE_WEEK_AGO:
            return True
    else:
        return True
    return False


# Fetch comments from GitHub API
def fetch_comments_from_api(params):
    response = requests.get(URL, headers=HEADERS, params=params)
    if response.status_code == 403:
        print("⚠️ Rate limit exceeded. Waiting for 60 seconds...")
        time.sleep(60)
    elif response.status_code != 200:
        print(f"🔴 Error: Received status code {response.status_code}")
        return []
    return response.json()


# Save comments to a JSON file
def save_comments_to_file(json_file, page_comments):
    with open(json_file, "w") as f:
        json.dump(page_comments, f)
    print(f"✅ Fetched and saved {len(page_comments)} comments to {json_file}\n")
    time.sleep(1)  # Delay between each page fetch


# Load comments from a JSON file
def load_comments_from_file(json_file):
    with open(json_file, "r") as f:
        page_comments = json.load(f)
    print(f"\n✅ Loaded comments from {json_file}")
    return page_comments


# Check if the image is already downloaded
def is_image_already_downloaded(local_path, expected_url):
    if os.path.exists(local_path):
        try:
            with Image.open(local_path) as img:
                width, height = img.size
            file_size = os.path.getsize(local_path)
            return width > 0 and height > 0 and file_size > 0
        except Exception:
            return False
    return False


# Download and cache images
def download_and_cache_image(url, local_path):
    if is_image_already_downloaded(local_path, url):
        print(f"⚡ Image already exists, skipping download: {local_path}")
        return False

    response = requests.get(url)
    with open(local_path, "wb") as f:
        f.write(response.content)
    print(f"⬇️ Downloaded image from {url}")
    return True


# Process and download profile pictures
def process_profile_pictures(comments):
    for comment in comments:
        avatar_url = comment["user"]["avatar_url"]
        username = comment["user"]["login"]
        avatar_file = os.path.join(PROFILE_PIC_DIR, f"{username}.png")
        if download_and_cache_image(avatar_url, avatar_file):
            resize_image(avatar_file)
            print(f"✅ Resized profile picture for {username} to 48x48 pixels\n")
        comment["user"]["avatar_local"] = avatar_file


# Resize image to 48x48 pixels
def resize_image(avatar_file):
    with Image.open(avatar_file) as img:
        img = img.resize((48, 48), Image.LANCZOS)
        img.save(avatar_file)


# Process and download images in comment bodies
def process_comment_images(comments):
    image_pattern = re.compile(
        r"(https?://(?:[^\s/$.?#].[^\s]*)\.(?:png|jpg|jpeg|gif))"
    )
    for comment in comments:
        comment["created_at_formatted"] = datetime.strptime(
            comment["created_at"], "%Y-%m-%dT%H:%M:%SZ"
        ).strftime("%d %b %Y")
        comment["body"] = markdown.markdown(comment["body"])
        for image_url in image_pattern.findall(comment["body"]):
            image_name = os.path.basename(image_url)
            image_path = os.path.join(DOWNLOADED_IMAGES_DIR, image_name)
            download_and_cache_image(image_url, image_path)
            comment["body"] = comment["body"].replace(image_url, image_path)


# Render HTML template with comments
def render_html(comments):
    data = {"issue_number": ISSUE_NUMBER, "comments": comments}

    with open("template.mustache", "r") as f:
        template = f.read()
    html_content = chevron.render(template, data)

    with open("reviews.html", "w") as f:
        f.write(html_content)

    print("\n=========================================")
    print("\n🚀 Static webpage generated: reviews.html")


# Main function to execute the steps
def main():
    create_directories()
    comments = fetch_comments()
    process_profile_pictures(comments)
    process_comment_images(comments)
    render_html(comments)


if __name__ == "__main__":
    main()
