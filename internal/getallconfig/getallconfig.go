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

// Send all configuration to client or vpn server
// GET from VPN Server or Web Portal

package getallconfig

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jsenon/vpncentralmanager/pkg/db/dynamo"
	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
)

// Server struct are configuration of a vpn node server
type Server struct {
	conf []*pb.AllConfigFileResp //nolint: megacheck, structcheck
}

// ItemServer struct are server vpn item in DB
type ItemServer struct {
	Server     string `json:"Server"`
	ServerName string `json:"ServerName"`
	AddressVpn string `json:"AddressVpn"`
	AddressPub string `json:"AddressPub"`
	PublicKey  string `json:"PublicKey"`
	Status     string `json:"Status"`
}

// ItemClient struct are client item in DB
type ItemClient struct {
	Client     string `json:"Client"`
	ClientName string `json:"ClientName"`
	AddressVpn string `json:"AddressVpn"`
	PublicKey  string `json:"PublicKey"`
	Status     string `json:"Status"`
}

// GetAllConfig send all configuration
func (s *Server) GetAllConfig(ctx context.Context, in *pb.AllConfigFileReq) (*pb.AllConfigFileResp, error) {
	fmt.Println("Info received in GetAll")
	fmt.Println("Debug: ", in)

	// Case if vpn or client config
	fmt.Println("Type config asked:", in.Type)

	sess, err := dynamo.ConnectDynamo()
	if err != nil {
		fmt.Println("failed to connect to dynaoDB:", err)
	}
	svc := dynamodb.New(sess)

	var server *pb.Item
	var client *pb.Item

	switch typeconf := in.Type; typeconf {
	case "vpnserver":
		var records []ItemServer

		var serverarray []*pb.Item

		_ = svc.ScanPages(&dynamodb.ScanInput{
			TableName: aws.String("VPNSERVER"),
		}, func(page *dynamodb.ScanOutput, last bool) bool {
			recs := []ItemServer{}
			err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
			if err != nil {
				panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
			}
			records = append(records, recs...)
			return true // keep paging
		})
		for _, res := range records {
			server = &pb.Item{
				Id:         res.Server,
				Name:       res.ServerName,
				Addressvpn: res.AddressVpn,
				Addresspub: res.AddressPub,
				Publikey:   res.PublicKey,
				Status:     res.Status,
			}
			serverarray = append(serverarray, server)
		}

		return &pb.AllConfigFileResp{Items: serverarray}, nil

	case "client":
		var records []ItemClient

		var clientarray []*pb.Item

		_ = svc.ScanPages(&dynamodb.ScanInput{
			TableName: aws.String("VPNCLIENT"),
		}, func(page *dynamodb.ScanOutput, last bool) bool {
			recs := []ItemClient{}
			err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
			if err != nil {
				panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
			}
			records = append(records, recs...)
			return true // keep paging
		})
		for _, res := range records {
			client = &pb.Item{
				Id:         res.Client,
				Name:       res.ClientName,
				Addressvpn: res.AddressVpn,
				Addresspub: "",
				Publikey:   res.PublicKey,
				Status:     res.Status,
			}
			clientarray = append(clientarray, client)
		}
		return &pb.AllConfigFileResp{Items: clientarray}, nil

	default:
		fmt.Printf("Wrong type %s.", typeconf)
	}
	return &pb.AllConfigFileResp{Items: nil}, nil
}