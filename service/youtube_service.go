package service

import (
	"log"

	entity "github.com/amartelr/go_youtube/entity"
	channel "github.com/amartelr/go_youtube/model/channel"
	playlist "github.com/amartelr/go_youtube/model/playlist"
	playlistitem "github.com/amartelr/go_youtube/model/playlistitem"
	subscription "github.com/amartelr/go_youtube/model/subscription"
	"github.com/go-resty/resty/v2"
)

type YoutubeProvider interface {
	PlaylistList(options *entity.PlaylistListOptions) (playlist []entity.PlayList, err error)
}

type service struct {
	Client *entity.Client
}

// NewService any
func NewService(client *entity.Client) *service {
	return &service{
		Client: client,
	}
}

func (c *service) ListPlayList(options entity.PlaylistListOptions) (playlists []entity.PlayList, err error) {

	url := c.Client.EndPoint + "/playlists"

	var playlistlist = playlist.Playlist{}
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

func (c *service) ListSubscription(options entity.SubscriptionListOptions) (subscriptions []entity.Subscription, err error) {

	url := c.Client.EndPoint + "/subscriptions"

	var subscriptionslist = subscription.Subscription{}
	client := resty.NewWithClient(c.Client.HttpClient)

	_, eErr := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       options.Part,
			"mine":       options.Mine,
			"key":        c.Client.ApiKey,
			"maxResults": options.MaxResults,
			"order":      options.Order,
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

func (c *service) ListChannel(options entity.ChannelListOptions) (channels []entity.Channel, err error) {

	url := c.Client.EndPoint + "/channels"

	var channellist = channel.Channel{}
	client := resty.NewWithClient(c.Client.HttpClient)

	_, eErr := client.R().EnableTrace().
		SetQueryParams(map[string]string{
			"part":       options.Part,
			"mine":       options.Mine,
			"key":        c.Client.ApiKey,
			"maxResults": options.MaxResults,
			"order":      options.Order,
		}).SetResult(&channellist).Get(url)

	if eErr != nil {
		log.Fatal(eErr)
	}
	for _, item := range channellist.Items {
		channels = append(channels, entity.Channel{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			ResourceID:  item.ID,
		})
	}
	return channels, eErr
}

func (c *service) ListPlayListItem(options entity.PlayListItemListOptions) (playlistitems []entity.PlayListItem, err error) {

	url := c.Client.EndPoint + "/playlistItems"

	var playlistlist = playlistitem.Playlistitem{}
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
