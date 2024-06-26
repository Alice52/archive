package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/c"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/gookit/goutil/jsonutil"
	"github.com/micro-services-roadmap/kit-common/kg"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type UserFavServiceIn struct{}

func (c *UserFavServiceIn) SyncUserFav() (err error) {
	folders, err := dal.Q.ArchivedFavFolder.Find()
	if err != nil {
		return err
	}

	for _, f := range folders {
		if err = DoSyncUserFav(f.ID); err != nil {
			kg.L.Error("sync user fav in folder: "+cast.ToString(f.Mid)+" error", zap.Error(err))
		}
	}

	return err
}

func DoSyncUserFav(id int64) (err error) {
	favs, err := api.LogonClient.UserFavOfFolder(id)
	if err != nil {
		return err
	}

	var vs []*model.ArchivedVideo
	for _, media := range favs.Data.Medias {
		m := &model.ArchivedFav{
			ID:       media.ID,
			Fid:      id,
			Bvid:     &media.Bvid,
			Cover:    &media.Cover,
			Duration: media.Duration,
			Ctime:    media.Ctime,
			FavTime:  media.FavTime,
			Title:    &media.Title,
			Intro:    &media.Intro,
			Type:     media.Type,
		}

		if media.Season != nil {
			sea := jsonutil.MustString(media.Season)
			m.Season = &sea
		}
		up := jsonutil.MustString(media.Upper)
		m.Upper = &up
		cnt := jsonutil.MustString(media.CntInfo)
		m.CntInfo = &cnt
		resp := jsonutil.MustString(media)
		m.Resp = &resp

		if err = dal.Q.ArchivedFav.Save(m); err != nil {
			kg.L.Error("sync upper error", zap.Error(err))
		}

		if len(media.Bvid) > 0 {
			vs = append(vs, &model.ArchivedVideo{
				Bvid:         media.Bvid,
				ArchivedType: c.ArchivedTypeFav})
		}
	}

	_ = UserVideoService.Merge(vs)
	return err
}
