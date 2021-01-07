package service

import (
	"log"

	"github.com/amartelr/go_youtube/entity"
	modelPlayL "github.com/amartelr/go_youtube/model/playlist"
	modelSus "github.com/amartelr/go_youtube/model/subscription"
	"github.com/amartelr/go_youtube/util"
	"github.com/go-resty/resty/v2"
)

const (
	endpoint = "https://youtube.googleapis.com/youtube/v3"
)

// NewYoutubeServ any
func NewYoutubeServ(param *entity.Parameters) Service {
	return &youtubeService{
		Param: param,
	}
}

type youtubeService struct {
	Param *entity.Parameters
}

func (repository *youtubeService) GetMysSubscriptions() (subscriptions []entity.Subscription, error util.RestErrorResponse) {

	url := endpoint + "/" + repository.Param.Model
	client := resty.New()

	var subscriptionslist = modelSus.Subscription{}

	_, err := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       repository.Param.Part,
			"mine":       repository.Param.Mine,
			"key":        repository.Param.Key,
			"maxResults": repository.Param.MaxResults,
			"order":      repository.Param.Order,
		}).ForceContentType("application/json").SetAuthToken(repository.Param.AccessToken).SetResult(&subscriptionslist).Get(url)

	if err != nil {
		log.Fatal(err)
	}
	for _, item := range subscriptionslist.Items {
		subscriptions = append(subscriptions, entity.Subscription{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ResourceID:  item.Snippet.ResourceID.ChannelID,
		})
	}
	return subscriptions, util.RestErrorResponse{}
}

func (repository *youtubeService) GetPlayList() (playList []entity.PlayList, error util.RestErrorResponse) {

	url := endpoint + "/" + repository.Param.Model
	client := resty.New()

	var modelPlaylist = modelPlayL.PlayList{}

	_, err := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       repository.Param.Part,
			"mine":       repository.Param.Mine,
			"key":        repository.Param.Key,
			"maxResults": repository.Param.MaxResults,
			"order":      repository.Param.Order,
		}).ForceContentType("application/json").SetAuthToken(repository.Param.AccessToken).SetResult(&modelPlaylist).Get(url)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range modelPlaylist.Items {
		playList = append(playList, entity.PlayList{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ResourceID:  item.ID,
		})
	}
	return playList, util.RestErrorResponse{}
}
