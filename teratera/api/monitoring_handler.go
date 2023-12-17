package api

import (
	"context"
	"runtime/debug"
	"slices"

	"connectrpc.com/connect"
	"github.com/minguu42/sandbox/teratera/gen/teraterapb/v1"
)

func (h handler) CheckHealth(context.Context, *connect.Request[teraterapb.CheckHealthRequest]) (*connect.Response[teraterapb.CheckHealthResponse], error) {
	revision := "xxxxxxx"
	if info, ok := debug.ReadBuildInfo(); ok {
		if i := slices.IndexFunc(info.Settings, func(s debug.BuildSetting) bool {
			return s.Key == "vcs.revision"
		}); i != -1 {
			revision = info.Settings[i].Value[:len(revision)]
		}
	}
	return connect.NewResponse(&teraterapb.CheckHealthResponse{Revision: revision}), nil
}
