package main

import (
	"osm-changesets-bot/env"
	"osm-changesets-bot/internal"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	env.Load()

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

			erro := internal.SetLatestChangesetId(changeset.Id)
			if err != nil {
				log.Error("error setting latest changeset", "id", changeset.Id, "err", erro)
				break
			}

			// avoid flood
			time.Sleep(5 * time.Second)
		}

		time.Sleep(time.Duration(env.TaskInterval) * time.Second)
	}
}
