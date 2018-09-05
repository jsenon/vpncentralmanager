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

// Send configuration of new client to vpn server
// Change db state of client to ready when configuation is applied
// POST to VPN Server

package postclientconfig

import (
	"bytes"
	"context"
	"net"

	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jsenon/vpncentralmanager/config"
	"github.com/jsenon/vpncentralmanager/pkg/calc/nextip"
	"github.com/jsenon/vpncentralmanager/pkg/db/dynamo"

	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
	"google.golang.org/grpc"
)

// Use for demo Purpose: address of fake vpn server
const (
	address = "localhost:50051"
)

// Item struct for client registration into db
type Item struct {
	Client     string `json:"Client"`
	ClientName string `json:"ClientName"`
	AddressVpn string `json:"AddressVpn"`
	PublicKey  string `json:"PublicKey"`
	Status     string `json:"Status"`
}

// ItemKey Key in db
type ItemKey struct {
	Client string `json:"Client"`
}

// UpdateIPVPN item to update
type UpdateIPVPN struct {
	AddressVpn string `json:":v"`
}

// Define range of a client
// rangeip = "10.200.200.1/21"
const minipclient = "10.200.200.5"
const maxipclient = "10.200.206.254"

// TODO Contact all VPN Server

// PostToAll post configuration to all VPN Server
func PostToAll() {

}

// PostClientConf send config to VPN Server
func PostClientConf(ctx context.Context, idclient string) { // nolint: gocyclo
	_, span := trace.StartSpan(ctx, "(*Server).PostClientConf")
	defer span.End()

	// Contact VPN Server GRPC
	var conn *grpc.ClientConn
	log.Debug().Msg("Fake send to VPN Server")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Can't connect to VPN Server for %s", config.Service)
	}
	defer conn.Close() // nolint: errcheck
	client := pb.NewSendClientConfigClient(conn)

	// Connection to DynamoDB
	sess, err := dynamo.ConnectDynamo()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Can't connect to Dynamo Server for %s", config.Service)
	}
	svc := dynamodb.New(sess)
	out, err := dynamo.SearchDynamo(svc, "VPNCLIENT", idclient, "Client")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Can't search on Dynamo Server for %s", config.Service)
	}

	item := Item{}
	err = dynamodbattribute.UnmarshalMap(out.Item, &item)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error unmarshal for %s", config.Service)
	}

	// Calculate VPN Client IP
	scan, err := ScanDynamo(svc, "VPNCLIENT")
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error in scan for %s", config.Service)
	}
	n := net.ParseIP(minipclient)

	for _, res := range scan {
		if bytes.Compare(net.ParseIP(res.AddressVpn), n) > 0 && bytes.Compare(net.ParseIP(res.AddressVpn), net.ParseIP(maxipclient)) < 0 {
			n = net.ParseIP(res.AddressVpn)
		}
	}
	// Increment ip address
	// TODO : How to manage if IP Address has been deleted
	ippriv := nextip.NextIP(net.IP.String(n))
	allowediprange := ippriv + "/32"
	clientkeypub := item.PublicKey

	// Send Configuration to vpn server
	response, err := client.SendClientConfig(context.Background(), &pb.ConfigFileResp{Keypublic: clientkeypub, Allowedrange: allowediprange})
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error when calling GRPC Server for %s", config.Service)
	}
	log.Info().Msgf("Response from server: %b", response.Request)

	// Update in DB with key idclient
	// Change Status
	key, err := dynamodbattribute.MarshalMap(ItemKey{
		Client: idclient,
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error in marshal for %s", config.Service)
	}
	update, err := dynamodbattribute.MarshalMap(UpdateIPVPN{
		AddressVpn: ippriv,
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error in marshal for %s", config.Service)
	}
	err = dynamo.UpdateipvpnDynamo(svc, "VPNCLIENT", key, update)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Issue update status for %s", config.Service)
	}

}

// ScanDynamo scan and Unmarshal all records
func ScanDynamo(svc *dynamodb.DynamoDB, table string) ([]Item, error) {
	var records []Item
	err := svc.ScanPages(&dynamodb.ScanInput{
		TableName: aws.String(table),
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		recs := []Item{}
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
		if err != nil {
			log.Fatal().
				Err(err).
				Str("service", config.Service).
				Msgf("failed to unmarshal Dynamodb Scan Items for %s", config.Service)
		}
		records = append(records, recs...)
		return true // keep paging
	})
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Error in scan for %s", config.Service)
	}
	return records, err
}
