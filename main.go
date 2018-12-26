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
		Text:    "test message <https://gitlab.ozon.ru/bx/layout-api/-/jobs/810408|helm-check>",
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

	channelID, timestamp, err := api.PostMessage("gitlab", slack.MsgOptionText("Test message", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)

	user, err := api.GetUserByEmail("ilukyanov@ozon.ru")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("User:%s Full name:%s\n", user.Name, user.RealName)
	userName := "@" + user.Name
	userID, _, err1 := api.PostMessage(userName, slack.MsgOptionText("Test message", false), slack.MsgOptionAttachments(attachment), slack.MsgOptionUsername("ilukyanov"))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("Message successfully sent to user %s\n", userID )


}
