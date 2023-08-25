<h1 align="center"> OSM Changesets Bot</h1>

Easy way to see the [changesets](https://wiki.openstreetmap.org/wiki/Changeset) in a certain area in [Telegram](http://telegram.org). Send the messages directly to you, or create a channel and share it with other users.

### Set up the environment variables

Environment variables are necessary, to declare them just create a `.env` file at the root of the project (you can use another method if you prefer).

-   **BOT_TOKEN**: Telegram bot token, is obtained from [@BotFather](https://t.me/BotFather) (Create a new bot if required)
-   **CHANNEL_ID**: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
-   **DETA_KEY**: [Deta](https://deta.space) API key.
-   **FEED_URL**: URL of the [OSMCha](https://osmcha.org) filter that you want to use.
-   **TASK_INTERVAL**: Time in minutes between each feed parse, default to 15m.
-   **RETRY_INTERVAL**: Time in seconds to wait after each error before running the task again.

### Deploy using Docker Compose

Firstly, you need to install [Docker](https://www.docker.com) on your computer. Assuming you have already done that, then proceed with the following steps.

Create a file called `compose.yaml` at the root of the project with the following content:

```yaml
version: "3.8"

services:
    bot:
        image: ghcr.io/ghostsama2503/osm-changesets-bot:latest
        container_name: "osmbot"

        environment:
            BOT_TOKEN: $BOT_TOKEN
            CHANNEL_ID: $CHANNEL_ID
            DETA_KEY: $DETA_KEY
            FEED_URL: $FEED_URL
            TASK_INTERVAL: $TASK_INTERVAL
            RETRY_INTERVAL: $RETRY_INTERVAL
```

To finish, run the container.

```sh
docker compose up -d
```

### Deploy manually

You need to install `python` (it is usually installed in most of the most popular distros) and `git`. To install them, use the following command.

```sh
apt install git python3 python3-pip python3-venv
```

Clone the repository.

```sh
git clone https://github.com/GHOSTsama2503/OSM-Changesets-Bot
```

Switch to the project folder

```sh
cd OSM-Changesets-Bot
```

Create and activate a virtual environment to install the dependencies.

```sh
python3 -m venv .venv && source .venv/bin/activate
```

Install the dependencies.

```sh
pip install -r requirements.txt
```

Run the project.

```sh
python3 -m src
```

</details>

### Code style

-   [Black Formatter](https://marketplace.visualstudio.com/items?itemName=ms-python.black-formatter)
-   [isort](https://marketplace.visualstudio.com/items?itemName=ms-python.isort)
