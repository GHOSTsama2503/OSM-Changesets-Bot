package main

import (
	"osm-changesets-bot/env"
	"osm-changesets-bot/internal"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	env.Load()

	// TODO: improve this

	for {
		latest, err := internal.GetLatestChangesetId()
		if err != nil {
			log.Error("error getting latest changeset", "err", err)
		}

		changesets, err := internal.NewChangesets(latest)
		if err != nil {
			log.Error("error getting new changesets", "err", err)
		}

		for _, changeset := range changesets {
			err := internal.SendToTelegram(changeset)
			if err != nil {
				log.Error("error sending changeset to telegram", "id", changeset.Id, "err", err)
				break
			}

			idInt, err := strconv.Atoi(changeset.Id)
			if err != nil {
				log.Error("changeset id must be an int", "id", changeset.Id, "err", err)
				break
			}

			err = internal.SetLatestChangesetId(idInt)
			if err != nil {
				log.Error("error setting latest changeset", "id", changeset.Id, "err", err)
				break
			}
		}

		time.Sleep(time.Duration(env.TaskInterval) * time.Second)
	}
}
