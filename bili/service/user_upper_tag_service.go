package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/gookit/goutil/jsonutil"
	"github.com/micro-services-roadmap/kit-common/kg"
	"go.uber.org/zap"
)

type UserUpperTagServiceIn struct{}

func (c *UserUpperTagServiceIn) SyncUpperTags() (err error) {
	tags, err := api.LogonClient.MyUppersTags()
	if err != nil {
		return err
	}

	for _, tag := range tags.Data {
		m := &model.ArchivedUpsTag{
			TagID:  tag.Tagid,
			Count_: &tag.Count,
			Name:   &tag.Name,
			Tip:    &tag.Tip,
		}
		resp := jsonutil.MustString(tag)
		m.Resp = &resp
		if err = dal.Q.ArchivedUpsTag.Save(m); err != nil {
			kg.L.Error("sync upper tags error", zap.Error(err))
		}
	}

	return err
}
