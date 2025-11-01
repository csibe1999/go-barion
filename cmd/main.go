package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/csibe1999/go-barion/pkg/barion"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-resty/resty/v2"
	"github.com/namsral/flag"
	"github.com/shopspring/decimal"
)

func main() {
	var (
		baseURL string
		posKey  string
	)

	flag.StringVar(&baseURL, "baseurl", "https://api.test.barion.com/v2", "Barion API base URL")
	flag.Parse()
	posKey, _ = os.LookupEnv("BARION_POS_KEY")
	platformEmail, _ := os.LookupEnv("PLATFORM_EMAIL")
	merchantEmail, _ := os.LookupEnv("MERCHANT_EMAIL")
	buyerEmail, _ := os.LookupEnv("BUYER_EMAIL")

	c := resty.New()
	c.SetDebug(true)
	client := barion.NewClient(baseURL, posKey, c)
	client.SetLogger(log.Print)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Post("/callback", client.CallbackHandler(func(state *barion.PaymentState) {
		log.Printf("PaymentState callback result: %v", state)
	}))

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(":3000", router); err != nil {
			log.Println(err)
		}
	}()

	payment := barion.PaymentRequest{
		POSKey:           posKey,
		PaymentType:      barion.Immediate,
		GuestCheckout:    true,
		PaymentRequestID: "fa-01",
		PayerHint:        buyerEmail,
		RedirectURL:      "https://merchanturl/Redirect?paymentId=xyz",
		CallbackURL:      "https://merchanturl/Callback?paymentId=xyz",
		Locale:           barion.HU,
		Currency:         barion.HUF,
		FundingSources:   []barion.FundingSources{barion.All},
		Transactions: []barion.PaymentTransaction{
			{
				POSTransactionID: "fa-01-01",
				Payee:            merchantEmail,
				Total:            decimal.NewFromInt(5000),
				PayeeTransactions: []barion.PayeeTransaction{
					{
						POSTransactionID: "TR-01-01-01",
						Payee:            platformEmail,
						Total:            decimal.NewFromInt(500),
						Comment:          "Marketplace commission: TR-01-01-01.",
					},
				},
				Items: []barion.Item{
					{
						Name:        "English lesson",
						Description: "Advanced Business English lesson from native speaker",
						Quantity:    decimal.NewFromInt(2),
						Unit:        "hour",
						UnitPrice:   decimal.NewFromInt(2500),
						ItemTotal:   decimal.NewFromInt(5000),
						SKU:         "ENG-ADV-NTV",
					},
				},
			},
		},
	}

	response, err := client.StartPayment(context.TODO(), &payment)
	if err != nil {
		log.Fatal(spew.Sdump(err))
	}
	spew.Dump(response)
	log.Println(response.PaymentID)

	paymentState, err := client.GetPaymentState(context.TODO(), response.PaymentID)
	if err != nil {
		log.Fatal(spew.Sdump(err))
	}
	spew.Dump(paymentState)

	wg.Wait()
}
