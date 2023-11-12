package omikuji_service

import (
	"context"
	"math/rand"
	omikujiv1 "slash_mochi/gen/go/slash_mochi/v1/omikuji"

	"connectrpc.com/connect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type OmikujiService struct {
}

func NewOmikujiService() *OmikujiService {
	return &OmikujiService{}
}

func (s *OmikujiService) OpenOmikuji(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[omikujiv1.OmikujiResponse], error) {
	return connect.NewResponse(&omikujiv1.OmikujiResponse{
		Result: omikujiv1.OmikujiResponse_ResultLevel(rand.Intn(13)),
	}), nil
}
