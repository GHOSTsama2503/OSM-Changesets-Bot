import asyncio
import json
import logging
from urllib.parse import quote

from aiohttp import ClientSession

from . import database, logger, env, feed


log = logging.getLogger(__name__)


async def send_message(data: dict) -> bool:
    data = json.dumps(data)

    async with ClientSession() as session:
        url = f"https://api.telegram.org/bot{env.BOT_TOKEN}/sendMessage"
        headers = {"Content-Type": "application/json"}

        try:
            async with session.post(url=url, data=data, headers=headers) as response:
                response.raise_for_status()
                return True
        except:
            log.warning("Error sending message", exc_info=True)
            return False


def create_data(changeset: dict) -> dict:
    changeset_url = f"https://www.openstreetmap.org/changeset/{changeset['id']}"
    user_url = f"https://www.openstreetmap.org/user/{quote(changeset['user'])}"

    overpass_url = f"https://overpass-api.de/achavi/?changeset={changeset['id']}"
    osmcha_url = f"https://osmcha.org/changesets/{changeset['id']}"

    inline_keyboard = [
        [
            {"text": "🌐 Changeset", "url": changeset_url},
            {"text": "👤 User", "url": user_url},
        ],
        [
            {"text": "🌐 OSMCha", "url": osmcha_url},
            {"text": "🌐 Overpass", "url": overpass_url},
        ],
    ]

    text = f"{changeset['title']}\n\n"
    text += f"{changeset['summary']}\n\n"
    text += f"{changeset['date']}\n"

    text += "🟢 {create} | 🟠 {modify} | 🔴 {delete}".format(
        create=changeset.get("create"),
        modify=changeset.get("modify"),
        delete=changeset.get("delete"),
    )

    data = {
        "chat_id": env.CHANNEL_ID,
        "text": text,
        "reply_markup": {"inline_keyboard": inline_keyboard},
        "disable_web_page_preview": True,
    }

    return data


async def changesets_handler(changesets: list[dict]):
    for changeset in changesets:
        data = create_data(changeset)

        sent = False
        while not sent:
            sent = await send_message(data=data)

        updated = False
        while not updated:
            updated = database.update_latest(changeset["id"])

        # avoid flood
        await asyncio.sleep(5)


async def main():
    log.info("Client Started!")

    while True:
        try:
            new_changesets = feed.new_changesets(env.FEED_URL)
            log.info(f"New Changesets: {len(new_changesets)}")

        except:
            log.error("Error getting new changesets:", exc_info=True)
            new_changesets = []

        if new_changesets:
            log.info("Working... :(")
            await changesets_handler(new_changesets)
            log.info("Finished! :D")

        log.info("Sleeping... zzz...")
        await asyncio.sleep(env.TASK_INTERVAL)


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except (KeyboardInterrupt, asyncio.exceptions.CancelledError):
        pass
    except:
        log.error("Unexpected error", exc_info=True)
