package main

import (
	"fmt"
	"log"

	"github.com/amartelr/go_youtube/entity"
	"github.com/amartelr/go_youtube/service"
	yt "github.com/amartelr/go_youtube/service"
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

	ytserv := yt.NewService(c)

	playlist, err := ytserv.ListPlayList(entity.PlaylistListOptions{
		Part:       "snippet,contentDetails",
		ChannelId:  "",
		Mine:       "True",
		MaxResults: "1",
		Order:      "alphabetical",
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range playlist {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.ResourceID, item.Description))
	}

	subscri, err := ytserv.ListSubscription(entity.SubscriptionListOptions{
		Part:       "snippet,contentDetails",
		ChannelId:  "",
		Mine:       "True",
		MaxResults: "1",
		Order:      "alphabetical",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range subscri {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.ResourceID, item.Description))
	}

	channellist, err := ytserv.ListChannel(entity.ChannelListOptions{
		Part:       "snippet,contentDetails",
		ChannelId:  "",
		Mine:       "True",
		MaxResults: "1",
		Order:      "alphabetical",
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range channellist {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.ResourceID, item.Description))
	}

	playlistitems, err := ytserv.ListPlayListItem(entity.PlayListItemListOptions{
		Part:       "snippet,contentDetails",
		PlaylistId: "LM",
		Mine:       "True",
		MaxResults: "1",
		Order:      "alphabetical",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range playlistitems {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.VideoID, item.Description))

	}
}
