package service

import (
	"log"

	"github.com/amartelr/go_youtube/entity"
	model "github.com/amartelr/go_youtube/model/subscription"
	"github.com/go-resty/resty/v2"
)

// Service any
type Service interface {
	SubscriptionList(options *entity.SubscriptionListOptions) (subscription []entity.Subscription, err error)
}

// NewYoutubeServ any
func NewSubscriptionService(client *entity.Client) Service {
	return &subscriptionService{
		Client: client,
	}
}

type subscriptionService struct {
	Client  *entity.Client
	Options *entity.SubscriptionListOptions
}

func (c *subscriptionService) SubscriptionList(options *entity.SubscriptionListOptions) (subscriptions []entity.Subscription, err error) {

	url := c.Client.EndPoint + "/subscriptions"

	var subscriptionslist = model.Subscription{}
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
