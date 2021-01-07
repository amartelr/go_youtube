package main

import (
	"fmt"
	"log"

	"github.com/amartelr/go_youtube/service"
	"github.com/amartelr/go_youtube/util"
)

func main() {
	params := util.LoadParams()

	myYoutubeServ := service.NewYoutubeServ(params)
	params.Model = "subscriptions"
	subsc, err := myYoutubeServ.GetMysSubscriptions()
	if err.Code > 0 {
		log.Fatal(err)
		return
	}

	for _, item := range subsc {
		println(fmt.Sprintf("%s: https://www.youtube.com/channel/%s", item.Title, item.ResourceID))
	}

	params.Model = "playlists"
	list, err := myYoutubeServ.GetPlayList()
	if err.Code > 0 {
		log.Fatal(err)
		return
	}

	for _, item := range list {
		println(fmt.Sprintf("%s: https://music.youtube.com/playlist?list=%s", item.Title, item.ResourceID))
	}
}
