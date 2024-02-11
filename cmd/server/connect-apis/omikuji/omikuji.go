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
	rawLevel := rand.Intn(10000)
	resultLevel := omikujiv1.OmikujiResponse_TAIRA

	if rawLevel < 100 {
		resultLevel = omikujiv1.OmikujiResponse_DAIKYOU
	} else if rawLevel < 200 {
		resultLevel = omikujiv1.OmikujiResponse_SUEKYOU
	} else if rawLevel < 300 {
		resultLevel = omikujiv1.OmikujiResponse_HANNKYOU
	} else if rawLevel < 400 {
		resultLevel = omikujiv1.OmikujiResponse_SHOUKYOU
	} else if rawLevel < 500 {
		resultLevel = omikujiv1.OmikujiResponse_KYOU
	} else if rawLevel < 1700 {
		resultLevel = omikujiv1.OmikujiResponse_TAIRA
	} else if rawLevel < 2900 {
		resultLevel = omikujiv1.OmikujiResponse_SUESHOUKICHI
	} else if rawLevel < 4100 {
		resultLevel = omikujiv1.OmikujiResponse_SUEKICHI
	} else if rawLevel < 5400 {
		resultLevel = omikujiv1.OmikujiResponse_HANNKICHI
	} else if rawLevel < 6600 {
		resultLevel = omikujiv1.OmikujiResponse_SHOUKICHI
	} else if rawLevel < 7800 {
		resultLevel = omikujiv1.OmikujiResponse_CHUUKICHI
	} else if rawLevel < 9000 {
		resultLevel = omikujiv1.OmikujiResponse_KICHI
	} else {
		resultLevel = omikujiv1.OmikujiResponse_DAIKICHI
	}

	return connect.NewResponse(&omikujiv1.OmikujiResponse{
		Result: omikujiv1.OmikujiResponse_ResultLevel(resultLevel),
	}), nil
}
