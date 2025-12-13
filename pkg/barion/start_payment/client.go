package startpayment

import (
	"encoding/json"
	"fmt"

	"github.com/csibe1999/go-barion/pkg/barion"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	postKey     string
	baseUrl     string
	redirectUrl string
	callbackUrl string
	r           *resty.Client
}

// Creates a Checkout Session object.
// Creates a Checkout Session object.
func New(request *barion.PaymentRequest) (*barion.PaymentRequestResponse, error) {
	return getC().New(request)
}

func getC() Client {
	return Client{postKey: barion.POSKey, baseUrl: barion.BaseUrl, redirectUrl: barion.RedirectUrl, callbackUrl: barion.CallbackUrl, r: resty.New()}
}

func (c Client) New(request *barion.PaymentRequest) (*barion.PaymentRequestResponse, error) {
	response, err := c.StartPayment(request)
	return response, err
}

func (c *Client) StartPayment(request *barion.PaymentRequest) (*barion.PaymentRequestResponse, error) {
	url := c.baseUrl + "/v2/Payment/Start"
	request.POSKey = c.postKey
	request.CallbackURL = c.callbackUrl
	request.RedirectURL = c.redirectUrl

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req := c.r.R()
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(body)
	req.SetResult(&barion.PaymentRequestResponse{})
	req.SetError(&barion.ErrorResponse{})

	res, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("Error sending payment request %v: %s", request, err)
	}

	if res.IsError() {
		if res.Error() != nil {
			x := res.Error().(*barion.ErrorResponse)
			x.Status = res.Status()
		}
		return nil, fmt.Errorf("Error sending payment request %v: %s", request, res.Error().(*barion.ErrorResponse))
	}
	return res.Result().(*barion.PaymentRequestResponse), nil
}
