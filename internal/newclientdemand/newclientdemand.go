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

// Get Request from Web Portal for a new client, store it on db
// Change db state of client to new demand
// POST from WebPortal

package newclientdemand

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jsenon/vpncentralmanager/internal/postclientconfig"
	"github.com/jsenon/vpncentralmanager/pkg/calc/randomstring"
	"github.com/jsenon/vpncentralmanager/pkg/db/dynamo"
	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
)

// Server struct are configuration of a vpn node server
type Server struct {
	conf []*pb.ConfigFileReq //nolint: megacheck, structcheck
}

// Item struct for client registration into db
type Item struct {
	Client     string `json:"Client"`
	ClientName string `json:"ClientName"`
	AddressVpn string `json:"AddressVpn"`
	PublicKey  string `json:"PublicKey"`
	Status     string `json:"Status"`
}

// GetClientDemand store in db a demand from client
func (s *Server) GetClientDemand(ctx context.Context, in *pb.ConfigFileReq) (*pb.AckWeb, error) {
	fmt.Println("In GetClientDemand")
	fmt.Println("Debug: ", in)
	sess, err := dynamo.ConnectDynamo()
	if err != nil {
		log.Fatalf("Error in connect: %v", err)
	}
	fmt.Println("Connected to dynamodb")
	svc := dynamodb.New(sess)

	//Prepare Item insertion
	idclient := randomstring.RandStringBytesMaskImprSrc(16)
	item := Item{
		Client:     idclient,
		ClientName: in.Hostname,
		PublicKey:  in.Keypublic,
		Status:     "New Demand",
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Error in marshall: %v", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("VPNCLIENT"),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Successfully added new client to VPNCLIENT table")

	// Debug
	fmt.Println("Item:", item)

	// POST To VPN Server
	postclientconfig.PostClientConf(idclient)

	return &pb.AckWeb{Ack: true}, nil

}
