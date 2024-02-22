package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"osm-changesets-bot/env"
)

type InlineKeyboardButton struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type Message struct {
	ChatId                int                  `json:"chat_id"`
	Text                  string               `json:"text"`
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`
	DisableWebPagePreview bool                 `json:"disable_web_page_preview"`
}

func SendToTelegram(changeset Changeset) error {

	msgText := fmt.Sprintf("%s\n\n%s\n\n🟢 %s | 🟠 %s | 🔴 %s", changeset.Title, changeset.Description, changeset.Create, changeset.Modify, changeset.Delete)

	changesetBtn := InlineKeyboardButton{}
	changesetBtn.Text = "🌐 Changeset"
	changesetBtn.Url = fmt.Sprintf("https://www.openstreetmap.org/changeset/%s", changeset.Id)

	userBtn := InlineKeyboardButton{}
	userBtn.Text = "👤 User"
	userBtn.Url = fmt.Sprintf("https://www.openstreetmap.org/user/%s", url.QueryEscape(changeset.Username))

	osmChaBtn := InlineKeyboardButton{}
	osmChaBtn.Text = "🌐 OSMCha"
	osmChaBtn.Url = fmt.Sprintf("https://osmcha.org/changesets/%s", changeset.Id)

	overPassBtn := InlineKeyboardButton{}
	overPassBtn.Text = "🌐 Overpass"
	overPassBtn.Url = fmt.Sprintf("https://overpass-api.de/achavi/?changeset=%s", changeset.Id)

	var inline_keyboard [][]InlineKeyboardButton

	var firstRow []InlineKeyboardButton
	firstRow = append(firstRow, changesetBtn, userBtn)

	var secondRow []InlineKeyboardButton
	secondRow = append(secondRow, osmChaBtn, overPassBtn)

	inline_keyboard = append(inline_keyboard, firstRow, secondRow)

	markup := InlineKeyboardMarkup{}
	markup.InlineKeyboard = inline_keyboard

	data := Message{}
	data.ChatId = env.ChannelId
	data.Text = msgText
	data.ReplyMarkup = markup
	data.DisableWebPagePreview = true

	apiUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", env.BotToken)

	requestData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiUrl, "application/json; charset=UTF-8", bytes.NewBuffer(requestData))
	if err != nil {
		return err
	}

	// TODO: check response
	_ = resp

	return nil
}
