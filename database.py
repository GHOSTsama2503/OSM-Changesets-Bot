from env import DETA_KEY
from deta import Deta

db = Deta(DETA_KEY)
changesets = db.Base("changesets")


def already_parsed(changeset_id: int) -> bool:
    latest = changesets.get("latest")["id"]
    return changeset_id <= latest


def update_latest(changeset_id: int):
    changesets.update({"id": changeset_id}, "latest")
