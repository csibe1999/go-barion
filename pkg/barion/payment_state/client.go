package paymentstate

import (
	"fmt"

	"github.com/csibe1999/go-barion/pkg/barion"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	postKey string
	baseUrl string
	r       *resty.Client
}

func New(PaymentId string) (*barion.PaymentState, error) {
	return getC().New(PaymentId)
}

func getC() Client {
	return Client{postKey: barion.POSKey, baseUrl: barion.BaseUrl, r: resty.New()}
}

func (c Client) New(PaymentId string) (*barion.PaymentState, error) {
	response, err := c.GetPaymentState(PaymentId)
	return response, err
}

func (c *Client) GetPaymentState(PaymentId string) (*barion.PaymentState, error) {
	url := c.baseUrl + fmt.Sprintf("/v4/payment/%s/paymentstate", PaymentId)
	req := c.r.R()
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("x-pos-key", c.postKey)
	req.SetResult(barion.PaymentState{})
	req.SetError(barion.ErrorResponse{})
	res, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("1Error getting payment state for payment id %s: %v", PaymentId, err)
	}
	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*barion.ErrorResponse)
			x.Status = res.Status()
			return nil, fmt.Errorf("2Error getting payment state for payment id %s: %v", PaymentId, res.Error())
		}
		return nil, fmt.Errorf("3Error getting payment state for payment id %s: %v", PaymentId, &barion.ErrorResponse{
			Status: res.Status(),
		})
	}
	return res.Result().(*barion.PaymentState), nil
}
