// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad-dotnet-driver/dotnet"
)

func main() {
	// Serve the plugin
	ver, err := dotnet.DotnetVersionInfo()
	println("dotnet", ver, err)
	//plugins.Serve(factory)
}

// factory returns a new instance of a nomad driver plugin
func factory(log hclog.Logger) interface{} {
	ctx, _ := context.WithCancel(context.Background())
	return dotnet.NewDriver(ctx, log)
}
