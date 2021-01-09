package service

import (
	"log"

	"github.com/amartelr/go_youtube/entity"
	model "github.com/amartelr/go_youtube/model/channel"
	"github.com/go-resty/resty/v2"
)

// Service any
type Service interface {
	ChannelList(options *entity.ChannelListOptions) (channel []entity.Channel, err error)
}

// NewYoutubeServ any
func NewChannelService(client *entity.Client) Service {
	return &channelService{
		Client: client,
	}
}

type channelService struct {
	Client  *entity.Client
	Options *entity.ChannelListOptions
}

func (c *channelService) ChannelList(options *entity.ChannelListOptions) (channels []entity.Channel, err error) {

	url := c.Client.EndPoint + "/channels"

	var channellist = model.Channel{}
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
