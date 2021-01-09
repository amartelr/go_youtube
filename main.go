package main

import (
	"fmt"
	"log"

	"github.com/amartelr/go_youtube/entity"
	"github.com/amartelr/go_youtube/service"
	playser "github.com/amartelr/go_youtube/service/playlist"
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
		Part:       "snippet,contentDetails",
		ChannelId:  "",
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
		println(fmt.Sprintf("%s: https://music.youtube.com/channel/%s", item.Title, item.ResourceID))
	}
	myplaylistServ := playser.NewPlaylistService(c)

	optionsList2 := entity.PlaylistListOptions{
		Part:       "snippet,contentDetails",
		ChannelId:  "",
		Mine:       "True",
		MaxResults: "3",
		Order:      "alphabetical",
	}

	playlist, err := myplaylistServ.PlaylistList(&optionsList2)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range playlist {
		println(fmt.Sprintf("%s: https://music.youtube.com/playlist?list=%s", item.Title, item.ResourceID))
	}

}
