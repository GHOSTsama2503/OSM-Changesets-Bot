package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	DetaBaseName string = "changesets"
)

var (
	BotToken      string
	ChannelId     int
	FeedUrl       string
	DetaKey       string
	TaskInterval  int
	RetryInterval int
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		panic("could not load the env file")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	if len(BotToken) == 0 {
		panic("bot token can not be empty")
	}

	channelIdEnv := os.Getenv("CHANNEL_ID")
	channelIdParsed, err := strconv.Atoi(channelIdEnv)
	if err != nil {
		panic("channel id must be an int")
	}
	ChannelId = channelIdParsed

	FeedUrl = os.Getenv("FEED_URL")
	if len(FeedUrl) == 0 {
		panic("feed url must be valid")
	}

	DetaKey = os.Getenv("DETA_KEY")
	if len(DetaKey) == 0 {
		panic("deta key can not be empty")
	}

	taskInterval := os.Getenv("TASK_INTERVAL")
	taskIntervalParsed, err := strconv.Atoi(taskInterval)
	if err != nil {
		panic("task interval must be an integer")
	}
	TaskInterval = taskIntervalParsed

	retryInterval := os.Getenv("RETRY_INTERVAL")
	retryIntervalParsed, err := strconv.Atoi(retryInterval)
	if err != nil {
		panic("retry interval must be an integer")
	}
	RetryInterval = retryIntervalParsed
}
