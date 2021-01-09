package main

import (
	"fmt"
	"log"

	"github.com/amartelr/go_youtube/entity"
	"github.com/amartelr/go_youtube/service"
	srvchan "github.com/amartelr/go_youtube/service/channel"
	srvplay "github.com/amartelr/go_youtube/service/playlist"
	srvpl "github.com/amartelr/go_youtube/service/playlistitem"
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
		MaxResults: "3",
		Order:      "alphabetical",
	}
	subscri, err := mySubServ.SubscriptionList(&optionsList)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range subscri {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.ResourceID, item.Description))
	}

	myplaylistServ := srvplay.NewPlaylistService(c)

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
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.ResourceID, item.Description))
	}

	mychannelServ := srvchan.NewChannelService(c)

	optionsList3 := entity.ChannelListOptions{
		Part:       "snippet,contentDetails",
		ChannelId:  "",
		Mine:       "True",
		MaxResults: "3",
		Order:      "alphabetical",
	}

	channellist, err := mychannelServ.ChannelList(&optionsList3)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range channellist {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.ResourceID, item.Description))
	}

	myPlayListItemServ := srvpl.NewPlaylistitemService(c)

	optionsList4 := entity.PlayListItemListOptions{
		Part:       "snippet,contentDetails",
		PlaylistId: "LM",
		Mine:       "True",
		MaxResults: "3",
		Order:      "alphabetical",
	}

	playlistitems, err := myPlayListItemServ.PlaylistitemList(&optionsList4)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range playlistitems {
		println(fmt.Sprintf("%s;%s;%s", item.Title, item.VideoID, item.Description))

	}
}
