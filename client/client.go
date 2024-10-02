// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package client is a containerz grpc client.
package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	cpb "github.com/openconfig/gnoi/containerz"
)

var (
	// Dial is the dailer to use to read containerz
	Dial = grpc.DialContext
)

// Client is a grpc containerz client.
type Client struct {
	cli cpb.ContainerzClient
}

// NewClient builds a new containerz client with username and password for authentication.
func NewClient(ctx context.Context, addr, username, password string) (*Client, error) {
	// Create metadata with username and password
	md := metadata.Pairs(
		"username", username,
		"password", password,
	)

	// Attach metadata to context
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Use grpc.WithInsecure() to avoid using TLS
	conn, err := Dial(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %v", err)
	}

	return &Client{
		cli: cpb.NewContainerzClient(conn),
	}, nil
}
