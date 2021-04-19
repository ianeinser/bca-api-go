package general

import (
	"context"
	"net/url"

	"github.com/ianeinser/bca-api-go"
)

//Client is used to invoke BCA General API
type Client struct {
	Client      bca.APIImplementation
	CorporateID string
	AccessToken string
}

//NewClient is used to initialize new general Client
func NewClient(config bca.Config) Client {
	return Client{
		CorporateID: config.CorporateID,
		Client:      bca.NewAPI(config),
	}
}

//FundTransfer is used to send fund transfer instructions to BCA using this service. The source of fund transfer must be from corporate’s own deposit account. The recipient may be any deposit account within BCA
func (c *Client) ForeignExchangeRate(ctx context.Context, ptr_foreignExchangeRate *bca.ForeignExchangeRateRequest) (*bca.ForeignExchangeRateResponse, error) {
	var foreignExchangeRateResponse bca.ForeignExchangeRateResponse

	path := "/general/rate/forex"

	currencyCode := (*ptr_foreignExchangeRate).CurrencyCode
	rateType := (*ptr_foreignExchangeRate).RateType

	v := url.Values{}
	v.Add("CurrencyCode", currencyCode)
	v.Add("RateType", rateType)
	path += "?" + v.Encode()

	if err := c.Client.Call("GET", path, c.AccessToken, nil, nil, &foreignExchangeRateResponse); err != nil {
		return &foreignExchangeRateResponse, err
	}
	return &foreignExchangeRateResponse, nil
}
