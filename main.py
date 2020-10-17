"""Bot implementation using Telethon library"""

import os

from dotenv import load_dotenv
from telethon import TelegramClient, events, utils, types

load_dotenv()

API_ID = int(os.getenv("APP_ID"))
API_HASH = os.getenv("APP_HASH")
BOT_TOKEN = os.getenv("BOT_TOKEN")
APP_SESSION = os.getenv("APP_SESSION")

bot = TelegramClient(APP_SESSION, API_ID, API_HASH).start(bot_token=BOT_TOKEN)


@bot.on(events.NewMessage(pattern=r'(?i).*\b(hello|hi)\b'))
async def handler(event):
    keyboard = types.InlineKeyboardMarkup()
    sender = await event.get_sender()
    name = utils.get_display_name(sender)
    print(name, 'said', event.text, '!')


with bot:
    bot.run_until_disconnected()
