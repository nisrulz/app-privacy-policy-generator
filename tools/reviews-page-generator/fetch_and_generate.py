# /// script
# dependencies = [
#   "requests",
#   "chevron",
#   "markdown",
#   "pillow",
# ]
# ///

import re
import json
import time
import shutil
import requests
import chevron
import markdown
from PIL import Image
from concurrent.futures import ThreadPoolExecutor, as_completed
from datetime import datetime, timedelta
from pathlib import Path
from argparse import ArgumentParser

OWNER = "nisrulz"
REPO = "app-privacy-policy-generator"
ISSUE_NUMBER = "65"
API_URL = f"https://api.github.com/repos/{OWNER}/{REPO}/issues/{ISSUE_NUMBER}/comments"

REACTION_MAP = {
    "+1": "thumbs_up", "-1": "thumbs_down", "laugh": "laugh",
    "hooray": "hooray", "confused": "confused", "heart": "heart",
    "rocket": "rocket", "eyes": "eyes",
}

BASE = Path(__file__).parent
JSON_DIR = BASE / "comments_json"
PROFILE_DIR = BASE / "profile_pictures"
IMAGES_DIR = BASE / "downloaded_images"
TEMPLATE = BASE / "template.mustache"
OUTPUT = BASE / "reviews.html"
PUBLIC = BASE / "../../public"
ONE_WEEK = timedelta(weeks=1)
CONCURRENCY = 8


def setup_dirs():
    for d in (JSON_DIR, PROFILE_DIR, IMAGES_DIR):
        d.mkdir(exist_ok=True)


def format_reactions(raw):
    return {REACTION_MAP[k]: v for k, v in raw.items() if k in REACTION_MAP and v > 0}


def fetch_comments(session, force):
    comments = []
    page = 1
    while True:
        path = JSON_DIR / f"comments_page_{page}.json"
        use_cache = not force and path.exists() and datetime.fromtimestamp(path.stat().st_mtime) > datetime.now() - ONE_WEEK
        if use_cache:
            data = json.loads(path.read_text())
            print(f"  Loaded cached: {path.name}")
        else:
            data = fetch_page(session, page)
            if data is None:
                break
            path.write_text(json.dumps(data))
            print(f"  Fetched {path.name} ({len(data)} comments)")
            time.sleep(1)
        if not data:
            break
        comments.extend(data)
        page += 1
    comments.reverse()
    return comments


def fetch_page(session, page):
    for _ in range(3):
        resp = session.get(API_URL, params={"per_page": 100, "page": page})
        if resp.status_code == 403:
            print("  Rate limited, waiting 60s...")
            time.sleep(60)
            continue
        if resp.status_code != 200:
            print(f"  Error {resp.status_code}")
            return None
        return resp.json()
    return None


def download_avatars(session, comments):
    def dl(c):
        url = c["user"]["avatar_url"]
        path = PROFILE_DIR / f"{c['user']['login']}.png"
        if path.exists():
            return False
        path.write_bytes(session.get(url).content)
        Image.open(path).resize((48, 48), Image.LANCZOS).save(path)
        return True

    with ThreadPoolExecutor(CONCURRENCY) as pool:
        futures = {pool.submit(dl, c): c for c in comments}
        for f in as_completed(futures):
            if f.result():
                print(f"  Downloaded avatar: {futures[f]['user']['login']}")


def download_images(session, comments):
    seen = set()

    def dl(url):
        name = url.rsplit("/", 1)[-1].split("?")[0]
        path = IMAGES_DIR / name
        if path.exists() or url in seen:
            return None
        seen.add(url)
        path.write_bytes(session.get(url).content)
        return name

    urls = set()
    for c in comments:
        for url in re.findall(r"https?://[^\s]+\.(?:png|jpg|jpeg|gif)", c["body"]):
            urls.add(url.split("?")[0])

    with ThreadPoolExecutor(CONCURRENCY) as pool:
        for name in pool.map(dl, urls):
            if name:
                print(f"  Downloaded image: {name}")


def prepare_comments(comments):
    for c in comments:
        c["reactions"] = format_reactions(c.get("reactions", {}))
        c["created_at_formatted"] = datetime.strptime(c["created_at"], "%Y-%m-%dT%H:%M:%SZ").strftime("%d %b %Y")
        c["body"] = markdown.markdown(c["body"])
        c["user"]["avatar_local"] = str(PROFILE_DIR / f"{c['user']['login']}.png")
        for url in re.findall(r"https?://[^\s]+\.(?:png|jpg|jpeg|gif)", c["body"]):
            name = url.rsplit("/", 1)[-1].split("?")[0]
            c["body"] = c["body"].replace(url, str(IMAGES_DIR / name))


def render(comments):
    OUTPUT.write_text(chevron.render(TEMPLATE.read_text(), {
        "issue_number": ISSUE_NUMBER,
        "comments": comments,
        "total_comments": len(comments),
    }))
    print(f"  Generated: {OUTPUT}")


def copy_to_public():
    shutil.copytree(IMAGES_DIR, PUBLIC / IMAGES_DIR.name, dirs_exist_ok=True)
    shutil.copytree(PROFILE_DIR, PUBLIC / PROFILE_DIR.name, dirs_exist_ok=True)
    shutil.copy2(OUTPUT, PUBLIC / OUTPUT.name)
    print("  Copied assets → public/")


def main():
    args = ArgumentParser(description="Generate reviews page from GitHub issue comments")
    args.add_argument("--force-fetch", action="store_true", help="Ignore cached JSON, re-fetch from GitHub")
    args = args.parse_args()

    setup_dirs()

    session = requests.Session()
    session.headers.update({"User-Agent": REPO})

    print("Fetching comments...")
    comments = fetch_comments(session, args.force_fetch)
    print(f"Processing {len(comments)} comments...")

    download_avatars(session, comments)
    download_images(session, comments)
    prepare_comments(comments)

    print("Rendering template...")
    render(comments)
    copy_to_public()
    print("Done.")


if __name__ == "__main__":
    main()
