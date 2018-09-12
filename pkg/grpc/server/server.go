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
	"net"
	"os"
	"runtime"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/reflection"

	ack "github.com/jsenon/vpncentralmanager/internal/ackconfig"
	adv "github.com/jsenon/vpncentralmanager/internal/advertise"
	gac "github.com/jsenon/vpncentralmanager/internal/getallconfig"
	ncc "github.com/jsenon/vpncentralmanager/internal/newclientconfig"
	ncd "github.com/jsenon/vpncentralmanager/internal/newclientdemand"

	"github.com/jsenon/vpncentralmanager/pkg/exporter/jaegerexporter"
	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
)

const (
	port = ":50051"
)

// Serve launch command serve
func Serve() {

	// ctx := context.Background()
	jaegerexporter.NewExporterCollector()
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	// ctx, span := trace.StartSpan(ctx, "Serve")
	// span.End()

	log.Info().Msg("Dynamo url: " + os.Getenv("urldynamo"))

	log.Info().Msg("Start GRPC Server")
	lis, err := net.Listen("tcp", port)
	log.Info().Msg("Listening GRPC on port:" + port)
	if err != nil {
		log.Error().Msgf("Error %s", err.Error())
		runtime.Goexit()
	}
	// Creates a new gRPC server
	// s := grpc.NewServer()
	if err = view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Error().Msgf("Error %s", err.Error())
		runtime.Goexit()
	}
	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	service.RegisterChannelzServiceToServer(s)

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
		log.Error().Msgf("Error %s", err.Error())
		runtime.Goexit()
	}
}
