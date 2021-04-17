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

//TTAccount provides service transaction “Transaction to BCA’s Account” and also “Transfer to Other Bank”
func (c *Client) TTAccount(ctx context.Context, ptr_ttAccountRequest *bca.TTAccountRequest) (*bca.TTAccountResponse, error) {
	var ttAccountResponse bca.TTAccountResponse

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

//TTInquiryAccount provides service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
func (c *Client) TTInquiryAccount(ctx context.Context, ptr_ttInquiryAccountRequest *bca.TTInquiryAccountRequest) (*bca.TTInquiryAccountResponse, error) {
	var ttInquiryAccountResponse bca.TTInquiryAccountResponse

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

//TTInquiryAccountBalance provides service to Inquiry balance for Vostro’s Account
func (c *Client) TTInquiryAccountBalance(ctx context.Context, ptr_ttInquiryAccountBalanceRequest *bca.TTInquiryAccountBalanceRequest) (*bca.TTInquiryAccountBalanceResponse, error) {
	var ttInquiryAccountBalanceResponse bca.TTInquiryAccountBalanceResponse

	jsonReq, err := json.Marshal(*ptr_ttInquiryAccountBalanceRequest)
	if err != nil {
		return &ttInquiryAccountBalanceResponse, err
	}

	path := "/fire/accounts/balance"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttInquiryAccountBalanceResponse); err != nil {
		return &ttInquiryAccountBalanceResponse, err
	}

	return &ttInquiryAccountBalanceResponse, nil
}

//TTInquiryTransaction provides service to Inquiry Transaction that has been submitted before
func (c *Client) TTInquiryTransaction(ctx context.Context, ptr_ttInquiryTransactionRequest *bca.TTInquiryTransactionRequest) (*bca.TTInquiryTransactionResponse, error) {
	var ttInquiryTransactionResponse bca.TTInquiryTransactionResponse

	jsonReq, err := json.Marshal(*ptr_ttInquiryTransactionRequest)
	if err != nil {
		return &ttInquiryTransactionResponse, err
	}

	path := "/fire/transactions"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &ttInquiryTransactionResponse); err != nil {
		return &ttInquiryTransactionResponse, err
	}

	return &ttInquiryTransactionResponse, nil
}

//TTCashTransfer provides service for transaction “Cash Transfer” to Non account holder
func (c *Client) TTCashTransfer(ctx context.Context, ptr_ttCashTransferRequest *bca.TTCashTransferRequest) (*bca.TTCashTransferResponse, error) {
	var ttCashTransferResponse bca.TTCashTransferResponse

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
func (c *Client) TTAmendCashTransfer(ctx context.Context, ptr_ttAmendCashTransferRequest *bca.TTAmendCashTransferRequest) (*bca.TTAmendCashTransferResponse, error) {
	var ttAmendCashTransferResponse bca.TTAmendCashTransferResponse

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
func (c *Client) TTCancelCashTransfer(ctx context.Context, ptr_ttCancelCashTransferRequest *bca.TTCancelCashTransferRequest) (*bca.TTCancelCashTransferResponse, error) {
	var ttCancelCashTransferResponse bca.TTCancelCashTransferResponse

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
