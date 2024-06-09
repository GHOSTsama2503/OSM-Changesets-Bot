package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"osm-changesets-bot/env"
	"strings"
)

type LatestChangeset struct {
	Key         string `json:"key"`
	ChangesetId int    `json:"changeset_id"`
}

func apiUrl() (string, error) {
	detaKey := strings.Split(env.DetaKey, "_")
	if len(detaKey) != 2 {
		return "", errors.New("invalid deta key")
	}

	collectionId := detaKey[0]

	apiUrl := fmt.Sprintf("https://database.deta.sh/v1/%s/%s", collectionId, env.DetaBaseName)
	return apiUrl, nil
}

func SetLatestChangesetId(value int) error {
	url, err := apiUrl()
	if err != nil {
		return err
	}

	latest := LatestChangeset{}
	latest.ChangesetId = value

	data, err := json.Marshal(latest)
	if err != nil {
		return err
	}

	client := http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", env.DetaKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respData := LatestChangeset{}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return err
	}

	return nil
}

func GetLatestChangesetId() (int, error) {
	url, err := apiUrl()
	if err != nil {
		return 0, err
	}

	url = url + "/items/latest"
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", env.DetaKey)

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	data := LatestChangeset{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.ChangesetId, nil
}
