from aiogram import Bot
from aiogram.types import InlineKeyboardButton, InlineKeyboardMarkup
import asyncio
import env
import feed
import logger
import logging
from urllib.parse import quote

bot = Bot(env.BOT_TOKEN)
log = logging.getLogger(__name__)


def create_message(changeset: dict) -> str:
    text = f"{changeset['title']}\n\n"
    text += f"{changeset['summary']}\n\n"
    text += f"{changeset['date']}\n"
    text += f"🟢 {changeset['create']} | 🟠 {changeset['modify']} | 🔴 {changeset['delete']}"
    return text


async def send_changeset(changeset: dict):
    osm_url = "https://www.openstreetmap.org"

    changeset_url = f"{osm_url}/changeset/{quote(str(changeset['id']))}"
    user_url = f"{osm_url}/user/{quote(changeset['user'])}"

    inline_keyboard = [[
        InlineKeyboardButton("🌐 Changeset", url=changeset_url),
        InlineKeyboardButton("👤 User", user_url)
    ]]
    reply_markup = InlineKeyboardMarkup(inline_keyboard=inline_keyboard)

    message = create_message(changeset)
    await bot.send_message(env.CHANNEL_ID, message, reply_markup=reply_markup)


async def worker():
    while True:
        new_changesets = feed.new_changesets(env.FEED_URL)
        log.info(f"New Changesets: {len(new_changesets)}")

        for changeset in new_changesets:
            await send_changeset(changeset)
            await asyncio.sleep(3) # avoid flood

        await asyncio.sleep(15 * 60) # task interval


async def main():
    loop = asyncio.get_running_loop()
    loop.create_task(worker())


if __name__ == "__main__":
    asyncio.run(main())
