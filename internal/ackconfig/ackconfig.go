// Copyright © 2018 Julien SENON <julien.senon@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Get Ackknoledgement from newly vpn server when configuration is applied successfully
// Change db State from In Sync to ready
// POST from VPN Server

package ackconfig

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jsenon/vpncentralmanager/config"
	"github.com/jsenon/vpncentralmanager/pkg/db/dynamo"
	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
)

// Server struct are configuration of a vpn node server
type Server struct {
	conf []*pb.State //nolint: megacheck, structcheck
}

// Item document to store in db
type Item struct {
	Server     string `json:"Server"`
	ServerName string `json:"ServerName"`
	AddressVpn string `json:"AddressVpn"`
	AddressPub string `json:"AddressPub"`
	PublicKey  string `json:"PublicKey"`
	Status     string `json:"Status"`
}

// ItemKey Key in db
type ItemKey struct {
	Server string `json:"Server"`
}

// UpdateStatus item to update
type UpdateStatus struct {
	Status string `json:":s"`
}

// TODO: Check Status value send by VPN Server: In sync, Ready, Deleted

// GetAck receive ack from new node
func (s *Server) GetAck(ctx context.Context, in *pb.State) (*pb.AckNode, error) {
	_, span := trace.StartSpan(ctx, "(*Server).GetAck")
	defer span.End()

	fmt.Println("In Ack")
	fmt.Println("Debug: ", in)
	sess, err := dynamo.ConnectDynamo()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Can't connect to Dynamo Server for %s", config.Service)
	}
	svc := dynamodb.New(sess)
	out, err := dynamo.SearchDynamo(svc, "VPNSERVER", in.Serverid, "Server")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error in search on Dynamo Server for %s", config.Service)
	}
	item := Item{}
	err = dynamodbattribute.UnmarshalMap(out.Item, &item)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error unmarshal for %s", config.Service)
	}
	log.Info().Msgf("Old Status: %s", item.Status)

	// Change Status
	key, err := dynamodbattribute.MarshalMap(ItemKey{
		Server: in.Serverid,
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error marshal for %s", config.Service)
	}
	update, err := dynamodbattribute.MarshalMap(UpdateStatus{
		Status: in.Status,
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error marshal for %s", config.Service)
	}
	err = dynamo.UpdateStatusDynamo(svc, "VPNSERVER", key, update)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error update status on Dynamo for %s", config.Service)
	}

	// Check if correctly updated
	out, err = dynamo.SearchDynamo(svc, "VPNSERVER", in.Serverid, "Server")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error in search on Dynamo for %s", config.Service)
	}
	err = dynamodbattribute.UnmarshalMap(out.Item, &item)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error unmarshal for %s", config.Service)
	}
	log.Info().Msgf("New Status: %s", item.Status)
	return &pb.AckNode{Ack: true}, nil
}
