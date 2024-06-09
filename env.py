from dotenv import load_dotenv
import os

load_dotenv()

BOT_TOKEN = os.getenv("BOT_TOKEN")
CHANNEL_ID = int(os.getenv("CHANNEL_ID"))
DETA_KEY = os.getenv("DETA_KEY")
FEED_URL = os.getenv("FEED_URL")