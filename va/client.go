package va

import (
	"context"
	"net/url"

	bca "github.com/ianeinser/bca-api-go"
)

//Client is used to invoke BCA Virtual Account API
type Client struct {
	Client      bca.APIImplementation
	AccessToken string
	CompanyCode string
}

//NewClient is used to initialize new va.Client
func NewClient(config bca.Config) Client {
	return Client{
		CompanyCode: config.CompanyCode,
		Client:      bca.NewAPI(config),
	}
}

//VAInquiryStatusPayment is used to see the list of payment status that are owned by the customers. The data will be automatically queried between D-day (hari H) until D-2 day (H-2 / the day before yesterday), with maximum records returned are 10 rows
func (c *Client) VAInquiryStatusPayment(ctx context.Context, ptr_vaInquiryStatusPaymentRequest *bca.VAInquiryStatusPaymentRequest) (*bca.VAInquiryStatusPaymentResponse, error) {
	var vaInquiryStatusPaymentResponse bca.VAInquiryStatusPaymentResponse
	path := "/va/payments"

	v := url.Values{}
	v.Add("CompanyCode", (*ptr_vaInquiryStatusPaymentRequest).CompanyCode)

	if (*ptr_vaInquiryStatusPaymentRequest).CustomerNumber != "" {
		v.Add("CustomerNumber", (*ptr_vaInquiryStatusPaymentRequest).CustomerNumber)
	}

	if (*ptr_vaInquiryStatusPaymentRequest).RequestID != "" {
		v.Add("RequestId", (*ptr_vaInquiryStatusPaymentRequest).RequestID)

	}

	path += "?" + v.Encode()

	if err := c.Client.Call("GET", path, c.AccessToken, nil, nil, &vaInquiryStatusPaymentResponse); err != nil {
		return &vaInquiryStatusPaymentResponse, err
	}
	return &vaInquiryStatusPaymentResponse, nil

}
