package usecase

import (
	"test_dealls/src/business/domain"
	historywatch "test_dealls/src/business/usecase/history_watch"
	"test_dealls/src/business/usecase/likes"
	"test_dealls/src/business/usecase/photo"
	"test_dealls/src/business/usecase/profile"
	"test_dealls/src/business/usecase/user"
	"test_dealls/src/utils/config"
	"test_dealls/src/utils/log"
)

type Usecase struct {
	HistoryWatch historywatch.Interface
	Likes        likes.Interface
	Photo        photo.Interface
	Profile      profile.Interface
	User         user.Interface
}

func Init(log log.Interface, d *domain.Domain, cfg config.Application) *Usecase {
	return &Usecase{
		HistoryWatch: historywatch.Init(log, d.HistoryWatch),
		Likes:        likes.Init(log, d.Likes, d.Profile, d.Photo, d.Match, d.HistoryWatch),
		Photo:        photo.Init(log, d.Photo),
		Profile:      profile.Init(log, d.Profile, d.Photo, d.Likes, d.Subsription, d.HistoryWatch),
		User:         user.Init(log, d.User),
	}
}
