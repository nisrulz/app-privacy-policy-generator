# /// script
# dependencies = [
#   "requests",
#   "chevron",
#   "markdown",
#   "pillow",
# ]
# ///

import os
import re
import json
import time
import requests
import chevron
import markdown
from PIL import Image
from datetime import datetime, timedelta
from pathlib import Path
from argparse import ArgumentParser

OWNER = "nisrulz"
REPO = "app-privacy-policy-generator"
ISSUE_NUMBER = "65"
API_URL = f"https://api.github.com/repos/{OWNER}/{REPO}/issues/{ISSUE_NUMBER}/comments"
HEADERS = {"User-Agent": REPO}

BASE = Path(__file__).parent
JSON_DIR = BASE / "comments_json"
PROFILE_DIR = BASE / "profile_pictures"
IMAGES_DIR = BASE / "downloaded_images"
TEMPLATE = BASE / "template.mustache"
OUTPUT = BASE / "reviews.html"
PUBLIC = BASE / "../../public"
ONE_WEEK = timedelta(weeks=1)


def parse_args():
    p = ArgumentParser(description="Generate reviews page from GitHub issue comments")
    p.add_argument("--force-fetch", action="store_true", help="Ignore cached JSON, re-fetch from GitHub")
    return p.parse_args()


def setup_dirs():
    for d in (JSON_DIR, PROFILE_DIR, IMAGES_DIR):
        d.mkdir(exist_ok=True)


def fetch_comments(force):
    comments = []
    page = 1
    while True:
        path = JSON_DIR / f"comments_page_{page}.json"
        data = None
        if not force and path.exists() and datetime.fromtimestamp(path.stat().st_mtime) > datetime.now() - ONE_WEEK:
            data = json.loads(path.read_text())
            print(f"  Loaded cached: {path.name}")
        else:
            data = fetch_page(page)
            if data is None:
                break
            path.write_text(json.dumps(data))
            print(f"  Fetched & saved: {path.name} ({len(data)} comments)")
            time.sleep(1)
        if not data:
            break
        comments.extend(data)
        page += 1
    comments.reverse()
    return comments


def fetch_page(page):
    resp = requests.get(API_URL, headers=HEADERS, params={"per_page": 100, "page": page})
    if resp.status_code == 403:
        print("  Rate limited, waiting 60s...")
        time.sleep(60)
        return fetch_page(page)
    if resp.status_code != 200:
        print(f"  Error {resp.status_code}")
        return None
    return resp.json()


def ensure_image(url, path):
    if path.exists():
        try:
            with Image.open(path) as img:
                if img.size[0] > 0:
                    return False
        except Exception:
            pass
    path.parent.mkdir(exist_ok=True)
    path.write_bytes(requests.get(url).content)
    return True


def resize_avatar(path):
    with Image.open(path) as img:
        img.resize((48, 48), Image.LANCZOS).save(path)


def format_reactions(raw):
    mapping = {"+1": "thumbs_up", "-1": "thumbs_down", "laugh": "laugh", "hooray": "hooray",
               "confused": "confused", "heart": "heart", "rocket": "rocket", "eyes": "eyes"}
    return {mapping[k]: v for k, v in raw.items() if k in mapping and v > 0}


def process_comment(c, avatar_dir, images_dir):
    c["reactions"] = format_reactions(c.get("reactions", {}))
    c["created_at_formatted"] = datetime.strptime(c["created_at"], "%Y-%m-%dT%H:%M:%SZ").strftime("%d %b %Y")
    c["body"] = markdown.markdown(c["body"])

    avatar_path = avatar_dir / f"{c['user']['login']}.png"
    avatar_url = c["user"]["avatar_url"]
    if ensure_image(avatar_url, avatar_path):
        resize_avatar(avatar_path)
        print(f"  Downloaded & resized avatar: {c['user']['login']}")
    c["user"]["avatar_local"] = str(avatar_path)

    for url in re.findall(r"https?://[^\s]+\.(?:png|jpg|jpeg|gif)", c["body"]):
        name = url.rsplit("/", 1)[-1].split("?")[0]
        img_path = images_dir / name
        if ensure_image(url, img_path):
            print(f"  Downloaded image: {name}")
        c["body"] = c["body"].replace(url, str(img_path))


def render(comments):
    html = chevron.render(TEMPLATE.read_text(), {"issue_number": ISSUE_NUMBER, "comments": comments})
    OUTPUT.write_text(html)
    print(f"\n  Generated: {OUTPUT}")


def copy_to_public():
    for src in (IMAGES_DIR, PROFILE_DIR):
        dst = PUBLIC / src.name
        dst.mkdir(exist_ok=True)
        for f in src.iterdir():
            if f.is_file() and f.name != ".DS_Store":
                (dst / f.name).write_bytes(f.read_bytes())
    (PUBLIC / OUTPUT.name).write_bytes(OUTPUT.read_bytes())
    print("  Copied assets → public/")


def main():
    args = parse_args()
    setup_dirs()
    print("Fetching comments...")
    comments = fetch_comments(args.force_fetch)
    print(f"Processing {len(comments)} comments...")
    for c in comments:
        process_comment(c, PROFILE_DIR, IMAGES_DIR)
    print("Rendering template...")
    render(comments)
    copy_to_public()
    print("Done.")


if __name__ == "__main__":
    main()
