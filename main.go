package main

import (
	"fmt"
	"log"

	gserv "github.com/amartelr/go_youtube/service"
	"github.com/amartelr/go_youtube/util"
)

func main() {
	service, err := gserv.InitAuth()
	if err != nil {
		log.Fatal(err)
		return
	}

	params := util.LoadParams()
	myYoutubeServ := gserv.NewYoutubeServ(service, params)
	params.Model = "subscriptions"
	subscri, err := myYoutubeServ.GetMysSubscriptions()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range subscri {
		println(fmt.Sprintf("%s: https://www.youtube.com/channel/%s", item.Title, item.ResourceID))
	}

	params.Model = "playlists"
	list, err := myYoutubeServ.GetPlayList()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range list {
		println(fmt.Sprintf("%s: https://music.youtube.com/playlist?list=%s", item.Title, item.ResourceID))
	}
}
