package barion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type logger func(...interface{})

type client struct {
	postKey string
	baseUrl string
	logger  logger
	r       *resty.Client
}

func NewClient(baseUrl string, postKey string) *client {
	r := resty.New()
	return &client{
		baseUrl: baseUrl,
		postKey: postKey,
		logger:  log.Print,
		r:       r,
	}
}

func (c *client) SetLogger(logger logger) {
	c.logger = logger
}

func (c *client) StartPayment(ctx context.Context, request *PaymentRequest) (*PaymentRequestResponse, error) {
	url := c.baseUrl + "/v2/Payment/Start"
	request.POSKey = c.postKey

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req := c.r.R()
	req.SetContext(ctx)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(body)
	req.SetResult(&PaymentRequestResponse{})
	req.SetError(&ErrorResponse{})

	res, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("Error sending payment request %v: %s", request, err)
	}

	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*ErrorResponse)
			x.status = res.Status()
		}
		return nil, fmt.Errorf("Error sending payment request %v: %s", request, res.Error().(*ErrorResponse))
	}
	return res.Result().(*PaymentRequestResponse), nil
}

func (c *client) GetPaymentState(ctx context.Context, PaymentId string) (*PaymentState, error) {
	url := c.baseUrl + fmt.Sprintf("/v4/payment/%s/paymentstate", PaymentId)
	req := c.r.R()
	req.SetContext(ctx)
	req.SetHeader("Content-Type", "application/json")
	req.SetQueryParam("POSKey", c.postKey)
	req.SetQueryParam("PaymentId", PaymentId)
	req.SetResult(PaymentState{})
	req.SetError(ErrorResponse{})
	res, err := req.Get(url)
	if err != nil {
		return nil, fmt.Errorf("1Error getting payment state for payment id %s: %v", PaymentId, err)
	}
	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*ErrorResponse)
			x.status = res.Status()
			return nil, fmt.Errorf("2Error getting payment state for payment id %s: %v", PaymentId, res.Error())
		}
		return nil, fmt.Errorf("3Error getting payment state for payment id %s: %v", PaymentId, &ErrorResponse{
			status: res.Status(),
		})
	}
	return res.Result().(*PaymentState), nil
}
