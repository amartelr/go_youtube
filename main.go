package main

import (
	"fmt"
	"log"

	"github.com/amartelr/go_youtube/entity"
	"github.com/amartelr/go_youtube/service"
	subser "github.com/amartelr/go_youtube/service/subscription"
)

func main() {
	c, err := service.NewClient(
		"assets/client_secret.json",
		"assets/api_key.json",
		"youtube-go.json",
		"youtube.YoutubeReadonlyScope")
	if err != nil {
		log.Fatal(err)
		return
	}

	mySubServ := subser.NewSubscriptionService(c)

	optionsList := entity.SubscriptionListOptions{
		Part:       "snippet",
		Mine:       "True",
		MaxResults: "2",
		Order:      "alphabetical",
	}
	subscri, err := mySubServ.SubscriptionList(&optionsList)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range subscri {
		println(fmt.Sprintf("%s: https://www.youtube.com/channel/%s", item.Title, item.ResourceID))
	}
}
