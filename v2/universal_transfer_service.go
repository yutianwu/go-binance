package binance

import (
	"context"
)

const (
	AccountTypeSpot       = "SPOT"
	AccountTypeUsdtFuture = "USDT_FUTURE"
)

// UniversalTransferService transfer to subaccount
type UniversalTransferService struct {
	c               *Client
	fromEmail       string
	toEmail         string
	fromAccountType string
	toAccountType   string
	asset           string
	amount          string
}

// ToEmail set toEmail
func (s *UniversalTransferService) ToEmail(toEmail string) *UniversalTransferService {
	s.toEmail = toEmail
	return s
}

func (s *UniversalTransferService) FromEmail(fromEmail string) *UniversalTransferService {
	s.fromEmail = fromEmail
	return s
}

func (s *UniversalTransferService) FromAccountType(fromAccountType string) *UniversalTransferService {
	s.fromAccountType = fromAccountType
	return s
}

func (s *UniversalTransferService) ToAccountType(toAccountType string) *UniversalTransferService {
	s.toAccountType = toAccountType
	return s
}

// Asset set asset
func (s *UniversalTransferService) Asset(asset string) *UniversalTransferService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *UniversalTransferService) Amount(amount string) *UniversalTransferService {
	s.amount = amount
	return s
}

func (s *UniversalTransferService) universalTransfer(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"fromEmail":       s.fromEmail,
		"fromAccountType": s.fromAccountType,
		"toEmail":         s.toEmail,
		"toAccountType":   s.toAccountType,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res *UniversalTransferResponse, err error) {
	data, err := s.universalTransfer(ctx, "/sapi/v1/sub-account/universalTransfer", opts...)
	if err != nil {
		return nil, err
	}
	res = &UniversalTransferResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UniversalTransferResponse define transfer to subaccount response
type UniversalTransferResponse struct {
	TranId       int64  `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}
