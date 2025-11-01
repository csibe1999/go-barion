package barion

import (
	"github.com/go-resty/resty/v2"
)

type logger func(...interface{})

type Client struct {
	postKey string
	baseUrl string
	logger  logger
	r       *resty.Client
}
