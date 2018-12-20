package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nlopes/slack"
)

func main() {

	b, err := ioutil.ReadFile("slack-token.txt")
	if err != nil{
		panic(err)
	}

	api := slack.New(strings.TrimSpace(string(b)))
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "test message",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	channelID, timestamp, err := api.PostMessage("gitlab", slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

}
