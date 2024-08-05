// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dotnet

import (
	"context"
	"github.com/TibaGroup/nomad-dotnet-driver/dotnet"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of a nomad driver plugin
func factory(log hclog.Logger) interface{} {
	ctx, _ := context.WithCancel(context.Background())
	return dotnet.NewDriver(ctx, log)
}
