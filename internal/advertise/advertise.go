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

// Get Config from a new vpn server, store it on db and send info to finish vpn server configuration
// Change db State of VPN Server to In sync
// POST from VPN Server

package advertise

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jsenon/vpncentralmanager/pkg/calc/nextip"
	"github.com/jsenon/vpncentralmanager/pkg/calc/randomstring"

	"github.com/jsenon/vpncentralmanager/pkg/db/dynamo"
	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
)

// Server struct are configuration of a vpn node server
type Server struct {
	conf []*pb.NodeConf //nolint: megacheck, structcheck
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

const rangeip = "10.200.200.1/21"
const minipserver = "10.200.207.5"
const maxipserver = "10.200.207.254"

// GetConfig retrieve config from new node
func (s *Server) GetConfig(ctx context.Context, in *pb.NodeConf) (*pb.RespNode, error) {
	sess, err := dynamo.ConnectDynamo()
	if err != nil {
		log.Fatalf("Error in connect: %v", err)
	}
	svc := dynamodb.New(sess)

	fmt.Println("Receive advertise from VPN Server: ", in.Hostname)

	// Check next available ip for new VPN Server
	scan, err := ScanDynamo(svc, "VPNSERVER")
	if err != nil {
		log.Fatalf("Error in scan: %v", err)
	}
	n := net.ParseIP(minipserver)

	for _, res := range scan {
		if bytes.Compare(net.ParseIP(res.AddressVpn), n) > 0 && bytes.Compare(net.ParseIP(res.AddressVpn), net.ParseIP(maxipserver)) < 0 {
			fmt.Println("biggest")
			n = net.ParseIP(res.AddressVpn)
		} else {
			fmt.Println("Error IP VPN Server")
		}
	}

	// increment ip address
	// TODO : How to manage if IP Address has been deleted
	ippriv := nextip.NextIP(net.IP.String(n))

	// Make Final test to check if IPVPN is not already take

	//Prepare Item insertion
	idserver := randomstring.RandStringBytesMaskImprSrc(16)
	item := Item{
		Server:     idserver,
		ServerName: in.Hostname,
		AddressVpn: ippriv,
		AddressPub: in.Ippublic,
		PublicKey:  in.Keypublic,
		Status:     "In sync",
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Error in marshall: %v", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("VPNSERVER"),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Error in put item: %v", err)
	}
	fmt.Println("Successfully added new server to VPNSERVER table")

	// Debug
	fmt.Println("Item:", item)

	// Return info to vpn server
	return &pb.RespNode{Ipprivate: ippriv, Allowedrange: rangeip, Vpnservername: idserver}, nil
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
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}
		records = append(records, recs...)
		return true // keep paging
	})
	if err != nil {
		log.Fatalf("Error in scan: %v", err)
	}
	return records, err
}
