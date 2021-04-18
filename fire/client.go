package fire

import (
	"context"
	"encoding/json"

	bca "github.com/ianeinser/bca-api-go"
)

//Client is used to invoke BCA FIRe API
type Client struct {
	Client      bca.APIImplementation
	AccessToken string
	CorporateID string
	AccessCode  string
	BranchCode  string
	UserID      string
	LocalID     string
}

//NewClient is used to initialize new fire.Client
func NewClient(config bca.Config) Client {
	return Client{
		CorporateID: config.FIReCorporateID,
		AccessCode:  config.AccessCode,
		BranchCode:  config.BranchCode,
		UserID:      config.UserID,
		LocalID:     config.LocalID,
		Client:      bca.NewAPI(config),
	}
}

//Account provides service transaction “Transaction to BCA’s Account” and also “Transfer to Other Bank”
func (c *Client) TeleTransferToAccount(ctx context.Context, ptr_ttAccountRequest *bca.TeleTransferAccountRequest) (*bca.TeleTransferAccountResponse, error) {
	var ttAccountResponse bca.TeleTransferAccountResponse

	jsonReq, err := json.Marshal(*ptr_ttAccountRequest)
	if err != nil {
		return &ttAccountResponse, err
	}

	path := "/fire/transactions/to-account"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttAccountResponse); err != nil {
		return &ttAccountResponse, err
	}

	return &ttAccountResponse, nil
}

//InquiryAccount provides service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
func (c *Client) InquiryAccount(ctx context.Context, ptr_ttInquiryAccountRequest *bca.InquiryAccountRequest) (*bca.InquiryAccountResponse, error) {
	var ttInquiryAccountResponse bca.InquiryAccountResponse

	jsonReq, err := json.Marshal(*ptr_ttInquiryAccountRequest)
	if err != nil {
		return &ttInquiryAccountResponse, err
	}

	path := "/fire/accounts"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttInquiryAccountResponse); err != nil {
		return &ttInquiryAccountResponse, err
	}

	return &ttInquiryAccountResponse, nil
}

//InquiryAccountBalance provides service to Inquiry balance for Vostro’s Account
func (c *Client) InquiryAccountBalance(ctx context.Context, ptr_inquiryAccountBalanceRequest *bca.InquiryAccountBalanceRequest) (*bca.InquiryAccountBalanceResponse, error) {
	var inquiryAccountBalanceResponse bca.InquiryAccountBalanceResponse

	jsonReq, err := json.Marshal(*ptr_inquiryAccountBalanceRequest)
	if err != nil {
		return &inquiryAccountBalanceResponse, err
	}

	path := "/fire/accounts/balance"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &inquiryAccountBalanceResponse); err != nil {
		return &inquiryAccountBalanceResponse, err
	}

	return &inquiryAccountBalanceResponse, nil
}

//TTInquiryTransaction provides service to Inquiry Transaction that has been submitted before
func (c *Client) InquiryTransaction(ctx context.Context, ptr_inquiryTransactionRequest *bca.InquiryTransactionRequest) (*bca.InquiryTransactionResponse, error) {
	var inquiryTransactionResponse bca.InquiryTransactionResponse

	jsonReq, err := json.Marshal(*ptr_inquiryTransactionRequest)
	if err != nil {
		return &inquiryTransactionResponse, err
	}

	path := "/fire/transactions"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &inquiryTransactionResponse); err != nil {
		return &inquiryTransactionResponse, err
	}

	return &inquiryTransactionResponse, nil
}

//TTCashTransfer provides service for transaction “Cash Transfer” to Non account holder
func (c *Client) TeleTransferCashTransfer(ctx context.Context, ptr_ttCashTransferRequest *bca.TeleTransferCashTransferRequest) (*bca.TeleTransferCashTransferResponse, error) {
	var ttCashTransferResponse bca.TeleTransferCashTransferResponse

	jsonReq, err := json.Marshal(*ptr_ttCashTransferRequest)
	if err != nil {
		return &ttCashTransferResponse, err
	}

	path := "/fire/transactions/cash-transfer"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttCashTransferResponse); err != nil {
		return &ttCashTransferResponse, err
	}

	return &ttCashTransferResponse, nil
}

//TTAmendCashTransfer provides service for Amendment “Cash Transfer” to Non account holder
func (c *Client) TeleTransferAmendCashTransfer(ctx context.Context, ptr_ttAmendCashTransferRequest *bca.TeleTransferAmendCashTransferRequest) (*bca.TeleTransferAmendCashTransferResponse, error) {
	var ttAmendCashTransferResponse bca.TeleTransferAmendCashTransferResponse

	jsonReq, err := json.Marshal(*ptr_ttAmendCashTransferRequest)
	if err != nil {
		return &ttAmendCashTransferResponse, err
	}

	path := "/fire/transactions/cash-transfer/amend"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttAmendCashTransferResponse); err != nil {
		return &ttAmendCashTransferResponse, err
	}

	return &ttAmendCashTransferResponse, nil
}

//TTCancelCashTransfer provides service for Cancellation “Cash Transfer” to Non account holder
func (c *Client) TeleTransferCancelCashTransfer(ctx context.Context, ptr_ttCancelCashTransferRequest *bca.TeleTransferCancelCashTransferRequest) (*bca.TeleTransferCancelCashTransferResponse, error) {
	var ttCancelCashTransferResponse bca.TeleTransferCancelCashTransferResponse

	jsonReq, err := json.Marshal(*ptr_ttCancelCashTransferRequest)
	if err != nil {
		return &ttCancelCashTransferResponse, err
	}

	path := "/fire/transactions/cash-transfer/cancel"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttCancelCashTransferResponse); err != nil {
		return &ttCancelCashTransferResponse, err
	}

	return &ttCancelCashTransferResponse, nil
}
