package profile

import (
	"context"
	"fmt"
	"strconv"
	historywatchDom "test_dealls/src/business/domain/history_watch"
	likesDom "test_dealls/src/business/domain/likes"
	photoDom "test_dealls/src/business/domain/photo"
	profileDom "test_dealls/src/business/domain/profile"
	subscriptionDom "test_dealls/src/business/domain/subscription"
	"test_dealls/src/entity"
	"test_dealls/src/utils/appcontext"
	"test_dealls/src/utils/log"
)

type Interface interface {
	GetList(ctx context.Context, params entity.Profile) ([]entity.Profile, error)
	GetListPeople(ctx context.Context, params entity.Profile) (entity.Profile, error)
	GetDetail(ctx context.Context, params entity.Profile) (entity.Profile, error)

	Create(ctx context.Context, params entity.Profile) (entity.Profile, error)
	Update(ctx context.Context, params entity.Profile) (entity.Profile, error)
	Upgrade(ctx context.Context, params entity.Profile) (entity.Profile, error)
}

type Profile struct {
	log log.Interface
	dom domain
}

type domain struct {
	historyWatch historywatchDom.Interface
	likes        likesDom.Interface
	profile      profileDom.Interface
	photo        photoDom.Interface
	subscription subscriptionDom.Interface
}

func Init(log log.Interface, profileD profileDom.Interface, photoD photoDom.Interface, likesD likesDom.Interface, subsriptionD subscriptionDom.Interface, historyWatchD historywatchDom.Interface) Interface {
	return &Profile{
		log: log,
		dom: domain{
			historyWatch: historyWatchD,
			likes:        likesD,
			profile:      profileD,
			photo:        photoD,
			subscription: subsriptionD,
		},
	}
}

func (p *Profile) GetList(ctx context.Context, params entity.Profile) ([]entity.Profile, error) {
	p.log.Info(ctx, "Get list profile")

	listProfile, err := p.dom.profile.GetList(ctx, params)
	if err != nil {
		p.log.Error(ctx, fmt.Sprintf("error in GetList Profile, err: %s", err))
		return nil, err
	}

	return listProfile, nil
}

func (p *Profile) GetListPeople(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	p.log.Info(ctx, "Get list profile people")

	/*
		Condition:
		1. Tampilkan profile lain dengan gender yg berlawanan dengan user
		2. Jika user no subscribe maka batas harian 10 (like+skip)
		3. Jika sudah di like maka tidak boleh tampil kembali.
	*/

	// Get profile user
	userID, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	profile, err := p.dom.profile.GetDetailByUserID(ctx, entity.Profile{
		UserID: int64(userID),
	})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	// check limit daily dengan hitung di history watch
	listWatch, err := p.dom.historyWatch.GetListByProfileID(ctx, entity.HistoryWatch{
		Profile1: profile.ID,
	})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	mapWatch := map[int64]entity.HistoryWatch{}
	for _, v := range listWatch {
		mapWatch[v.Profile2] = v
	}

	// Get Subscribe
	listSubscription, err := p.dom.subscription.GetList(ctx, entity.Subscription{})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	mapSubsription := map[int64]entity.Subscription{}
	for _, v := range listSubscription {
		mapSubsription[v.ID] = v
	}

	if int64(len(listWatch)) > mapSubsription[profile.SubscriptionID].Value {
		p.log.Error(ctx, "telah melebihi limit harian")
		return entity.Profile{}, fmt.Errorf("telah melebihi limit harian")
	}

	// get profile people
	listProfile, err := p.dom.profile.SearchProfilePeople(ctx, entity.Profile{
		UserID: int64(userID),
		Gender: profile.Gender,
	})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	// Get Photo
	listPhoto, err := p.dom.photo.GetList(ctx, entity.Photo{})
	if err != nil {
		p.log.Error(ctx, fmt.Sprintf("error in GetList Likes, err: %s", err))
		return entity.Profile{}, err
	}

	mapsPhoto := map[int64][]string{}
	for _, v := range listPhoto {
		mapsPhoto[v.ProfileID] = append(mapsPhoto[v.ProfileID], v.PhotoURL)
	}

	// Get likes
	listLikes, err := p.dom.likes.GetListByLikerID(ctx, entity.Likes{
		LikerID: profile.ID,
	})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	mapLikes := map[int64]entity.Likes{}
	for _, v := range listLikes {
		mapLikes[v.LikedID] = v
	}

	// Tidak boleh ketemu yg sudah di like
	res := entity.Profile{}
	for _, v := range listProfile {
		if _, isOK := mapLikes[v.ID]; isOK {
			continue
		}

		if _, isOK := mapLikes[v.ID]; isOK {
			continue
		}

		if photos, isOK := mapsPhoto[v.ID]; isOK {
			v.Photo = photos
		}

		if value, isOK := mapSubsription[v.SubscriptionID]; isOK {
			if value.Name == "premium" {
				v.Label = "PREMIUM"
			} else {
				v.Label = "BASIC"
			}
		}

		res = v
	}

	// Create limit daily
	historyWatch, err := p.dom.historyWatch.Create(ctx, entity.HistoryWatch{
		Profile1:  profile.ID,
		Profile2:  res.ID,
		Label:     "",
		CreatedAt: appcontext.GetRequestStartTime(ctx),
	})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	res.HistoryWatchID = historyWatch.ID

	return res, nil
}

func (p *Profile) GetDetail(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	p.log.Info(ctx, "Get Profile by ID")

	id, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		p.log.Error(ctx, err.Error())
	}
	params.ID = int64(id)

	profile, err := p.dom.profile.GetDetail(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	return profile, nil
}

func (p *Profile) Create(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	p.log.Info(ctx, "Create profile")

	userID, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	params.UserID = int64(userID)
	params.CreatedAt = appcontext.GetRequestStartTime(ctx)
	profile, err := p.dom.profile.Create(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	return profile, nil
}

func (p *Profile) Update(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	p.log.Info(ctx, "Update profile")

	userID, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		return entity.Profile{}, err
	}

	params.UserID = int64(userID)
	params.UpdatedAt = appcontext.GetRequestStartTime(ctx)
	profile, err := p.dom.profile.Update(ctx, params)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	return profile, nil
}

func (p *Profile) Upgrade(ctx context.Context, params entity.Profile) (entity.Profile, error) {
	p.log.Info(ctx, "Upgrade profile")

	userID, err := strconv.Atoi(appcontext.GetUserIDAgent(ctx))
	if err != nil {
		return entity.Profile{}, err
	}

	profile, err := p.dom.profile.GetDetailByUserID(ctx, entity.Profile{
		UserID: int64(userID),
	})
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	profile.SubscriptionID = params.SubscriptionID
	profile.UpdatedAt = appcontext.GetRequestStartTime(ctx)
	_, err = p.dom.profile.Update(ctx, profile)
	if err != nil {
		p.log.Error(ctx, err.Error())
		return entity.Profile{}, err
	}

	return profile, nil
}
