package api

import (
	"fmt"
	"github.com/alice52/archive/bili/util"
	"github.com/alice52/archive/common/global"
	"github.com/micro-services-roadmap/kit-common/kg"
	"github.com/micro-services-roadmap/kit-common/viperx"
	"github.com/micro-services-roadmap/kit-common/zapx"
	"testing"
)

func init() {
	viperx.InitViper("../config-local.yaml")
	kg.L = zapx.Zap(global.CONFIG.Zap)
}

func TestGetVideoInfo(t *testing.T) {
	info, err := logonFunc().VideoInfo("BV117411r7R1")
	if err != nil {
		t.Error(err)
		return
	}
	s := util.MustMarshal(info)
	fmt.Println(s)
}
