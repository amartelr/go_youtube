package service

import (
	"log"
	"net/http"

	"github.com/amartelr/go_youtube/entity"
	modelPlay "github.com/amartelr/go_youtube/model/playlist"
	modelSus "github.com/amartelr/go_youtube/model/subscription"
	"github.com/go-resty/resty/v2"
)

const (
	endpoint = "https://youtube.googleapis.com/youtube/v3"
)

// NewYoutubeServ any
func NewYoutubeServ(client *http.Client, param *entity.Parameters) Service {
	return &youtubeService{
		Client: client,
		Param:  param,
	}
}

type youtubeService struct {
	Param  *entity.Parameters
	Client *http.Client
}

func (repository *youtubeService) GetMysSubscriptions() (subscriptions []entity.Subscription, err error) {

	url := endpoint + "/" + repository.Param.Model

	var subscriptionslist = modelSus.Subscription{}
	client := resty.NewWithClient(repository.Client)

	_, eErr := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       repository.Param.Part,
			"mine":       repository.Param.Mine,
			"key":        repository.Param.Key,
			"maxResults": repository.Param.MaxResults,
			"order":      repository.Param.Order,
		}).SetResult(&subscriptionslist).Get(url)

	if eErr != nil {
		log.Fatal(eErr)
	}
	for _, item := range subscriptionslist.Items {
		subscriptions = append(subscriptions, entity.Subscription{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ResourceID:  item.Snippet.ResourceID.ChannelID,
		})
	}
	return subscriptions, eErr
}

func (repository *youtubeService) GetPlayList() (playList []entity.PlayList, err error) {

	url := endpoint + "/" + repository.Param.Model
	client := resty.NewWithClient(repository.Client)

	var modelPlaylist = modelPlay.PlayList{}

	_, eErr := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       repository.Param.Part,
			"mine":       repository.Param.Mine,
			"key":        repository.Param.Key,
			"maxResults": repository.Param.MaxResults,
			"order":      repository.Param.Order,
		}).SetResult(&modelPlaylist).Get(url)

	if eErr != nil {
		log.Fatal(eErr)
	}

	for _, item := range modelPlaylist.Items {
		playList = append(playList, entity.PlayList{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ResourceID:  item.ID,
		})
	}
	return playList, eErr
}
