/*
Copyright (C) Marck Tomack <marcktomack@tutanota.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	bot, _ = tb.NewBot(tb.Settings{
		Token:  "YOUR TOKEN HERE",
		Poller: &tb.LongPoller{Timeout: 25 * time.Minute},
	})
)

func onText() {
	bot.Handle(tb.OnText, func(m *tb.Message) {
		searchURL, _ := regexp.Compile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
		url := searchURL.FindString(m.Text)
		if url == "" {
			println("no link")
		} else {
			getReq, _ := http.Get(url)
			realLink := getReq.Request.URL.String()
			msg := strings.Replace(m.Text, url, realLink, -1)
			bot.Send(m.Sender, "<b>This is the real link:</b>\n\n"+msg, tb.ModeHTML)
		}

	})
}

func onVideo() {
	bot.Handle(tb.OnVideo, func(m *tb.Message) {
		searchURL, _ := regexp.Compile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
		url := searchURL.FindString(m.Caption)
		if url == "" {
			println("no link")
		} else {
			getReq, _ := http.Get(url)
			realLink := getReq.Request.URL.String()
			msg := strings.Replace(m.Caption, url, realLink, -1)
			bot.Send(m.Sender, "<b>This is the real link:</b>\n\n"+msg, tb.ModeHTML)
		}

	})
}

func onPhoto() {
	bot.Handle(tb.OnPhoto, func(m *tb.Message) {
		searchURL, _ := regexp.Compile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
		url := searchURL.FindString(m.Caption)
		if url == "" {
			println("no link")
		} else {
			getReq, _ := http.Get(url)
			realLink := getReq.Request.URL.String()
			msg := strings.Replace(m.Caption, url, realLink, -1)
			bot.Send(m.Sender, "<b>This is the real link:</b>\n\n"+msg, tb.ModeHTML)
		}

	})
}

func main() {
	onText()
	onVideo()
	onPhoto()
	bot.Start()

}
