package domain

import (
	historywatch "test_dealls/src/business/domain/history_watch"
	"test_dealls/src/business/domain/likes"
	"test_dealls/src/business/domain/match"
	"test_dealls/src/business/domain/photo"
	"test_dealls/src/business/domain/profile"
	"test_dealls/src/business/domain/subscription"
	"test_dealls/src/business/domain/user"
	"test_dealls/src/utils/config"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

type Domain struct {
	HistoryWatch historywatch.Interface
	Likes        likes.Interface
	Match        match.Interface
	Photo        photo.Interface
	Profile      profile.Interface
	Subsription  subscription.Interface
	User         user.Interface
}

func Init(log log.Interface, db sql.Interface, cfg config.Application) *Domain {
	return &Domain{
		HistoryWatch: historywatch.Init(log, db),
		Likes:        likes.Init(log, db),
		Match:        match.Init(log, db),
		Photo:        photo.Init(log, db),
		Profile:      profile.Init(log, db),
		Subsription:  subscription.Init(log, db),
		User:         user.Init(log, db),
	}
}
