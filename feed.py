import re
from datetime import datetime
from time import mktime
from typing import Any

import feedparser

import database


def new_changesets(feed_url: str) -> list[dict] | None:
    parse = feedparser.parse(feed_url, sanitize_html=True)

    if not parse.entries or not isinstance(parse.entries, list):
        return

    entries: list[dict[str, Any]] = parse.entries
    changesets = []

    for info in entries[::-1]:
        changeset_id = int(re.search(r"changesets/(.+?)/", info.id).group(1))
        changeset = {}

        changeset["id"] = changeset_id
        changeset["url"] = info.id
        changeset["title"] = info.title

        summary = re.search(r"^(.+?)<br />", info.summary)
        changeset["summary"] = summary.group(1) if summary else ""

        changeset["user"] = re.search(r"by (.+)$", info.title).group(1)

        changeset["create"] = re.search(r"Create: ([0-9]+)", info.summary).group(1)
        changeset["modify"] = re.search(r"Modify: ([0-9]+)", info.summary).group(1)
        changeset["delete"] = re.search(r"Delete: ([0-9]+)", info.summary).group(1)

        date = datetime.fromtimestamp(mktime(info.published_parsed))
        changeset["date"] = date.strftime("%Y-%m-%d | %H:%M:%S")

        parsed = False
        while not parsed:
            try:
                if not database.already_parsed(changeset["id"]):
                    changesets.append(changeset)
                parsed = True
            except:
                pass

    return changesets
