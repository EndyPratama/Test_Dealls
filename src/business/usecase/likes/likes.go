package likes

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	historywatchDom "test_dealls/src/business/domain/history_watch"
	likesDom "test_dealls/src/business/domain/likes"
	matchDom "test_dealls/src/business/domain/match"
	photoDom "test_dealls/src/business/domain/photo"
	profileDom "test_dealls/src/business/domain/profile"
	"test_dealls/src/entity"
	"test_dealls/src/utils/appcontext"
	"test_dealls/src/utils/log"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Likes) ([]entity.Likes, error)

	Approve(ctx context.Context, params entity.Likes) (entity.Likes, error)
	Skip(ctx context.Context, params entity.Likes) (entity.Likes, error)
	Delete(ctx context.Context, params entity.Likes) (entity.Likes, error)
}

type Likes struct {
	log log.Interface
	dom domain
}

type domain struct {
	historyWatch historywatchDom.Interface
	likes        likesDom.Interface
	match        matchDom.Interface
	profile      profileDom.Interface
	photo        photoDom.Interface
}

func Init(log log.Interface, likesD likesDom.Interface, profileD profileDom.Interface, photoD photoDom.Interface, matchD matchDom.Interface, historyWatchD historywatchDom.Interface) Interface {
	return &Likes{
		log: log,
		dom: domain{
			historyWatch: historyWatchD,
			likes:        likesD,
			match:        matchD,
			profile:      profileD,
			photo:        photoD,
		},
	}
}

func (l *Likes) GetList(ctx context.Context, params entity.Likes) ([]entity.Likes, error) {
	l.log.Info(ctx, "Get list likes")

	userID, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		return nil, err
	}

	profileLiker, err := l.dom.profile.GetDetailByUserID(ctx, entity.Profile{UserID: int64(userID)})
	if err != nil {
		l.log.Error(ctx, err.Error())
		return nil, err
	}

	// Get all profile
	listLikes, err := l.dom.likes.GetListByLikerID(ctx, entity.Likes{LikerID: profileLiker.ID})
	if err != nil {
		l.log.Error(ctx, err.Error())
		return nil, err
	}

	listProfile, err := l.dom.profile.GetList(ctx, entity.Profile{})
	if err != nil {
		l.log.Error(ctx, err.Error())
		return nil, err
	}

	mapProfile := map[int64]entity.Profile{}
	for _, v := range listProfile {
		mapProfile[v.ID] = v
	}

	listMatches, err := l.dom.match.GetListByProfile(ctx, entity.Matches{
		Profile1: profileLiker.ID,
		Profile2: profileLiker.ID,
	})
	if err != nil {
		l.log.Error(ctx, err.Error())
		return nil, err
	}

	mapMatchesProfile1 := map[int64]entity.Matches{}
	mapMatchesProfile2 := map[int64]entity.Matches{}
	for _, v := range listMatches {
		mapMatchesProfile1[v.Profile1] = v
		mapMatchesProfile2[v.Profile2] = v
	}

	for i, v := range listLikes {
		if profile, isOk := mapProfile[v.LikedID]; isOk {
			listLikes[i].LikedName = profile.Name
		}

		if matches, isOk := mapMatchesProfile1[v.LikerID]; isOk {
			listLikes[i].MatchesID = matches.ID
		}

		if matches, isOk := mapMatchesProfile1[v.LikedID]; isOk {
			listLikes[i].MatchesID = matches.ID
		}
	}

	return listLikes, nil
}

func (l *Likes) Approve(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	l.log.Info(ctx, "Approve likes")

	params.CreatedAt = appcontext.GetRequestStartTime(ctx)
	likes, err := l.dom.likes.Create(ctx, params)
	if err != nil {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	// Check if match
	match, err := l.dom.likes.CheckLikes(ctx, entity.Likes{
		LikerID: params.LikedID,
		LikedID: likes.LikerID,
	})
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	if match.LikerID != 0 && match.LikedID != 0 {
		_, err := l.dom.match.Create(ctx, entity.Matches{
			Profile1:   likes.LikerID,
			Profile2:   likes.LikedID,
			Matched_at: appcontext.GetRequestStartTime(ctx),
			CreatedAt:  appcontext.GetRequestStartTime(ctx),
		})
		if err != nil {
			l.log.Error(ctx, err.Error())
			return entity.Likes{}, err
		}
	}

	// Update history watch
	_, err = l.dom.historyWatch.Update(ctx, entity.HistoryWatch{
		ID:        params.HistoryWatchID,
		Label:     "Like",
		UpdatedAt: appcontext.GetRequestStartTime(ctx),
	})
	if err != nil {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	return likes, nil
}

func (l *Likes) Skip(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	l.log.Info(ctx, "Skip likes")

	// Update history watch
	_, err := l.dom.historyWatch.Update(ctx, entity.HistoryWatch{
		ID:        params.HistoryWatchID,
		Label:     "Skip",
		UpdatedAt: appcontext.GetRequestStartTime(ctx),
	})
	if err != nil {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	return params, nil
}

func (l *Likes) Delete(ctx context.Context, params entity.Likes) (entity.Likes, error) {
	l.log.Info(ctx, "Delete likes")

	dataLike, err := l.dom.likes.GetDetail(ctx, params)
	if err != nil {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	if dataLike.ID == 0 {
		l.log.Error(ctx, "data likes not found")
		return entity.Likes{}, fmt.Errorf("data likes not found")
	}

	likes, err := l.dom.likes.Delete(ctx, params)
	if err != nil {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	matched, err := l.dom.match.GetDetail(ctx, entity.Matches{
		ID: params.MatchesID,
	})
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
		l.log.Error(ctx, err.Error())
		return entity.Likes{}, err
	}

	if matched.ID != 0 {
		_, err = l.dom.match.Delete(ctx, entity.Matches{
			ID: matched.ID,
		})
		if err != nil {
			l.log.Error(ctx, err.Error())
			return entity.Likes{}, err
		}
	}

	return likes, nil
}
