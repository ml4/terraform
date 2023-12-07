// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package cloudplugin1

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform/internal/cloudplugin"
	"github.com/hashicorp/terraform/internal/cloudplugin/cloudproto1"
	"github.com/hashicorp/terraform/internal/terminal"
	"google.golang.org/grpc"
)

// GRPCCloudPlugin is the go-plugin implementation, but only the client
// implementation exists in this package.
type GRPCCloudPlugin struct {
	Streams *terminal.Streams

	plugin.GRPCPlugin
	Impl cloudplugin.Cloud1
}

// Server always returns an error; we're only implementing the GRPCPlugin
// interface, not the Plugin interface.
func (p *GRPCCloudPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("cloudplugin only implements gRPC clients")
}

// Client always returns an error; we're only implementing the GRPCPlugin
// interface, not the Plugin interface.
func (p *GRPCCloudPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errors.New("cloudplugin only implements gRPC clients")
}

// GRPCServer always returns an error; we're only implementing the client
// interface, not the server.
func (p *GRPCCloudPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("cloudplugin only implements gRPC clients")
}

// GRPCClient returns a new GRPC client for interacting with the cloud plugin server.
func (p *GRPCCloudPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCCloudClient{
		client:  cloudproto1.NewCommandServiceClient(c),
		context: ctx,
		streams: p.Streams,
	}, nil
}
