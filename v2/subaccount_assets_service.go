package binance

import (
	"context"
)

const (
	FutureTypesUsdt int = 1
	FutureTypesCoin int = 2
)

type SubAccountFutureAssetsService struct {
	c           *Client
	email       string
	futuresType int
}

// Email set email
func (s *SubAccountFutureAssetsService) Email(email string) *SubAccountFutureAssetsService {
	s.email = email
	return s
}

func (s *SubAccountFutureAssetsService) FuturesType(futuresType int) *SubAccountFutureAssetsService {
	s.futuresType = futuresType
	return s
}

func (s *SubAccountFutureAssetsService) subAccountFutureAssets(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"futuresType": s.futuresType,
		"email":       s.email,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubAccountFutureAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFutureAssetsResponse, err error) {
	data, err := s.subAccountFutureAssets(ctx, "/sapi/v2/sub-account/futures/account", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubAccountFutureAssetsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFutureAssetsResponse struct {
	FutureAccountResp struct {
		Assets []FutureAsset `json:"assets"`
	} `json:"futureAccountResp"`
}

type FutureAsset struct {
	Asset              string `json:"asset"`
	MarginBalance      string `json:"marginBalance"`
	MaxWithdrawBalance string `json:"maxWithdrawAmount"`
	UnrealizedProfit   string `json:"unrealizedProfit"`
	WalletBalance      string `json:"walletBalance"`
}
