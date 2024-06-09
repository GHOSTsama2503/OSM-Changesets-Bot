package internal

import (
	"errors"
	"osm-changesets-bot/env"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

func NewChangesets(latest int) ([]Changeset, error) {
	var changesets []Changeset
	parser := gofeed.NewParser()

	feed, err := parser.ParseURL(env.FeedUrl)
	if err != nil {
		return changesets, errors.New("error getting new changesets")
	}

	for i := 0; i < len(feed.Items); i++ {
		item := feed.Items[i]

		titleSplited := strings.Split(item.Title, " by ")
		if len(titleSplited) != 2 {
			return changesets, errors.New("unexpected changeset title")
		}

		username := titleSplited[1]

		titleStart := strings.Split(titleSplited[0], " ")
		if len(titleStart) != 2 {
			return changesets, errors.New("unexpected changeset title")
		}

		id := titleStart[1]
		splitedDescription := strings.Split(item.Description, "<br>")

		descriptionParts := len(splitedDescription)
		if descriptionParts == 0 || descriptionParts > 3 {
			return changesets, errors.New("unexpected changeset description")
		}

		var description string
		var changes []string

		changesRegex, err := regexp.Compile(`Create: (\d+), Modify: (\d+), Delete: (\d+)`)
		if err != nil {
			return changesets, err
		}

		for _, line := range splitedDescription {
			changesMatch := changesRegex.FindStringSubmatch(line)

			if len(changesMatch) == 4 {
				changes = changesMatch

			} else if description == "" && len(changes) == 0 {
				description = line

			} else {
				// example value: Changeset flagged for: New mapper
				_ = line
			}
		}

		changeset := Changeset{}
		changeset.Id = id
		changeset.Title = item.Title
		changeset.Description = description
		changeset.Create = changes[1]
		changeset.Modify = changes[2]
		changeset.Delete = changes[3]
		changeset.Username = username

		changesets = append(changesets, changeset)
	}

	return changesets, nil
}
