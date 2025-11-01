package callback

import (
	"log"
	"net/http"

	"github.com/csibe1999/go-barion/pkg/barion"
	paymentstate "github.com/csibe1999/go-barion/pkg/barion/payment_state"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	postKey string
	baseUrl string
	r       *resty.Client
}

func New(paymentStateHandler func(state *barion.PaymentState)) http.HandlerFunc {
	return getC().New(paymentStateHandler)
}

func getC() Client {
	return Client{postKey: barion.POSKey, baseUrl: barion.BaseUrl, r: resty.New()}
}

func (c Client) New(paymentStateHandler func(state *barion.PaymentState)) http.HandlerFunc {
	cb := c.CallbackHandler(paymentStateHandler)
	return cb
}

func (c *Client) CallbackHandler(paymentStateHandler func(state *barion.PaymentState)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["paymentId"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'PaymentId' is missing")
			http.Error(w, "Url Param 'PaymentId' is missing", http.StatusBadRequest)
			return
		}

		key := keys[0]

		state, err := paymentstate.New(key)
		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		go func() {
			paymentStateHandler(state)
		}()
		_, _ = w.Write([]byte("OK"))
		return
	})
}
