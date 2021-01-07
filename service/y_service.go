package service

import (
	"github.com/amartelr/go_youtube/entity"
)

// Service any
type Service interface {
	GetMysSubscriptions() (subscription []entity.Subscription, err error)
	GetPlayList() (playList []entity.PlayList, err error)
}
