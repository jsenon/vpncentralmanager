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

package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	ack "github.com/jsenon/vpncentralmanager/internal/ackconfig"
	adv "github.com/jsenon/vpncentralmanager/internal/advertise"
	gac "github.com/jsenon/vpncentralmanager/internal/getallconfig"
	ncc "github.com/jsenon/vpncentralmanager/internal/newclientconfig"
	ncd "github.com/jsenon/vpncentralmanager/internal/newclientdemand"

	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// Serve launch command serve
func Serve() {
	fmt.Println("Start GRPC Server")
	lis, err := net.Listen("tcp", port)
	fmt.Println("Listening GRPC on port:", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	// s := grpc.NewServer()
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatalf("Error registering grpc: %v", err)
	}
	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	reflection.Register(s)

	// Server to register new VPN Server
	pb.RegisterAdvertiseServer(s, &adv.Server{})

	// Server to activate a new VPN Server
	pb.RegisterAckConfigServer(s, &ack.Server{})

	// Server to register a new client
	pb.RegisterNewClientDemandServer(s, &ncd.Server{})

	// Server to retrieve all config
	pb.RegisterRetrieveConfigServer(s, &gac.Server{})

	// Fake VPN Server
	pb.RegisterSendClientConfigServer(s, &ncc.Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error listening grpc: %v", err)
	}
}
