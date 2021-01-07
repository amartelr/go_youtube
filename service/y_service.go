package service

import (
	"github.com/amartelr/go_youtube/entity"
	"github.com/amartelr/go_youtube/util"
)

// Service any
type Service interface {
	GetMysSubscriptions() (subscription []entity.Subscription, error util.RestErrorResponse)
	GetPlayList() (playList []entity.PlayList, error util.RestErrorResponse)
}
