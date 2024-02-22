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

type PutItems struct {
	Set struct {
		ChangesetId int `json:"changeset_id"`
	} `json:"set"`
}

func apiUrl() (string, error) {
	detaKey := strings.Split(env.DetaKey, "_")
	if len(detaKey) != 2 {
		return "", errors.New("invalid deta key")
	}

	collectionId := detaKey[0]

	apiUrl := fmt.Sprintf("https://database.deta.sh/v1/%s/%s/items/latest", collectionId, env.DetaBaseName)
	return apiUrl, nil
}

func SetLatestChangesetId(value int) error {
	url, err := apiUrl()
	if err != nil {
		return err
	}

	data := PutItems{}
	data.Set.ChangesetId = value

	reqBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := http.Client{}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(reqBody))
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
