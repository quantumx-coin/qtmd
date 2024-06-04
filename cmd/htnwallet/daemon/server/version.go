package server

import (
	"context"

	"github.com/quantumx-coin/qtmd/cmd/htnwallet/daemon/pb"
	"github.com/quantumx-coin/qtmd/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
