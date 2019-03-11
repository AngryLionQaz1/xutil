package plugin

import "xutil/cmd/filter/common"

type Plugin struct {
	ctx *common.Context
}

func NewPlugin(ctx *common.Context) *Plugin {

	return &Plugin{ctx: ctx}

}

func loadPlguin(ctx *common.Context) {

}
