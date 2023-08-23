import logging

from deta import Deta

from .env import DETA_KEY


log = logging.getLogger(__name__)

db = Deta(DETA_KEY)
changesets = db.Base("changesets")


def already_parsed(changeset_id: int) -> bool:
    latest = changesets.get("latest")["id"]
    return changeset_id <= latest


def update_latest(changeset_id: int) -> bool:
    try:
        changesets.update({"id": changeset_id}, "latest")
        return True
    except:
        return False
