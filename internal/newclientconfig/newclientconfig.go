// FAKE Server

package newclientconfig

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"

	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
)

// Server struct are configuration of a vpn node server
type Server struct {
	conf []*pb.ConfigFileResp //nolint: megacheck, structcheck
}

// SendClientConfig simulate VPN Server
func (s *Server) SendClientConfig(ctx context.Context, in *pb.ConfigFileResp) (*pb.Request, error) {
	_, span := trace.StartSpan(ctx, "(*Server).SendClientConfig")
	defer span.End()

	span.Annotate([]trace.Attribute{
		trace.Int64Attribute("len", int64(len(in.Allowedrange))+int64(len(in.Keypublic))),
	}, "Data in")

	log.Info().Msg("Info received in fake VPN Server")
	log.Info().Msgf("Debug: %s", in)
	return &pb.Request{Request: true}, nil
}
