# OSMCubaChangesetsBot

### Environment
- **BOT_TOKEN**: Telegram bot token, is obtained from [@BotFather](https://t.me/BotFather) (Create a new bot if required)
- **CHANNEL_ID**: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
- **DETA_KEY**: [Deta](https://deta.space/legacy) API key.
- **FEED_URL**: URL of the [OSMCha](https://osmcha.org) filter that you want to use.
- **TASK_INTERVAL**: Time in minutes between each feed parse, default to 15m.
- **RETRY_INTERVAL**: Time in seconds to wait after each error before running the task again.

### Deploy
You need the following packages installed before running the project: `git`, `python3`, `python3-pip`, `python3-venv`

Install packages (Ubuntu/Debian):
```sh
sudo apt install -y git python3 python3-pip python3-venv
```

Run the project:
```sh
git clone https://github.com/GHOSTsama2503/OSMCubaChangesetsBot
cd OSMCubaChangesetsBot

python3 -m venv .venv
source .venv/bin/activate

pip install -r requirements.txt

python3 main.py
```

### Code Style
- [Black Formatter](https://marketplace.visualstudio.com/items?itemName=ms-python.black-formatter)
- [isort](https://marketplace.visualstudio.com/items?itemName=ms-python.isort)
