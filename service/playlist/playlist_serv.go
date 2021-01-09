package service

import (
	"log"

	"github.com/amartelr/go_youtube/entity"
	model "github.com/amartelr/go_youtube/model/playlist"
	"github.com/go-resty/resty/v2"
)

// Service any
type Service interface {
	PlaylistList(options *entity.PlaylistListOptions) (playlist []entity.PlayList, err error)
}

// NewYoutubeServ any
func NewPlaylistService(client *entity.Client) Service {
	return &playlistService{
		Client: client,
	}
}

type playlistService struct {
	Client  *entity.Client
	Options *entity.PlaylistListOptions
}

func (c *playlistService) PlaylistList(options *entity.PlaylistListOptions) (playlists []entity.PlayList, err error) {

	url := c.Client.EndPoint + "/playlists"

	var playlistlist = model.Playlist{}
	client := resty.NewWithClient(c.Client.HttpClient)

	_, eErr := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       options.Part,
			"mine":       options.Mine,
			"key":        c.Client.ApiKey,
			"maxResults": options.MaxResults,
			"order":      options.Order,
		}).SetResult(&playlistlist).Get(url)

	if eErr != nil {
		log.Fatal(eErr)
	}
	for _, item := range playlistlist.Items {
		playlists = append(playlists, entity.PlayList{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ResourceID:  item.ID,
		})
	}
	return playlists, eErr
}
