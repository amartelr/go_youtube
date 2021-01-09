package service

import (
	"log"

	"github.com/amartelr/go_youtube/entity"
	model "github.com/amartelr/go_youtube/model/playlistitem"
	"github.com/go-resty/resty/v2"
)

// Service any
type Service interface {
	PlaylistitemList(options *entity.PlayListItemListOptions) (playlistitem []entity.PlayListItem, err error)
}

// NewYoutubeServ any
func NewPlaylistitemService(client *entity.Client) Service {
	return &playlistitemService{
		Client: client,
	}
}

type playlistitemService struct {
	Client  *entity.Client
	Options *entity.PlayListItemListOptions
}

func (c *playlistitemService) PlaylistitemList(options *entity.PlayListItemListOptions) (playlistitems []entity.PlayListItem, err error) {

	url := c.Client.EndPoint + "/playlistItems"

	var playlistlist = model.Playlistitem{}
	client := resty.NewWithClient(c.Client.HttpClient)

	_, eErr := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       options.Part,
			"playlistId": options.PlaylistId,
			"mine":       options.Mine,
			"key":        c.Client.ApiKey,
			"maxResults": options.MaxResults,
			"order":      options.Order,
		}).SetResult(&playlistlist).Get(url)

	if eErr != nil {
		log.Fatal(eErr)
	}
	for _, item := range playlistlist.Items {
		playlistitems = append(playlistitems, entity.PlayListItem{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			VideoID:     item.Snippet.ResourceID.VideoID,
		})
	}
	return playlistitems, eErr
}
