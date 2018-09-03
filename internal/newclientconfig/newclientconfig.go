// FAKE Server

package newclientconfig

import (
	"context"
	"fmt"

	"github.com/jsenon/vpncentralmanager/pkg/grpc/pb"
)

// Server struct are configuration of a vpn node server
type Server struct {
	conf []*pb.ConfigFileResp //nolint: megacheck, structcheck
}

// SendClientConfig simulate VPN Server
func (s *Server) SendClientConfig(ctx context.Context, in *pb.ConfigFileResp) (*pb.Request, error) {
	fmt.Println("Info received in fake VPN Server")
	fmt.Println("Debug: ", in)
	return &pb.Request{Request: true}, nil
}
