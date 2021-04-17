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
	var ptr_ttAccountResponse *bca.TTAccountResponse

	jsonReq, err := json.Marshal(*ptr_ttAccountRequest)
	if err != nil {
		return ptr_ttAccountResponse, err
	}

	path := "/fire/transactions/to-account"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttAccountResponse); err != nil {
		return ptr_ttAccountResponse, err
	}

	return ptr_ttAccountResponse, nil
}

//TTInquiryAccount provides service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
func (c *Client) TTInquiryAccount(ctx context.Context, ptr_ttInquiryAccountRequest *bca.TTInquiryAccountRequest) (*bca.TTInquiryAccountResponse, error) {
	var ptr_ttInquiryAccountResponse *bca.TTInquiryAccountResponse

	jsonReq, err := json.Marshal(*ptr_ttInquiryAccountRequest)
	if err != nil {
		return ptr_ttInquiryAccountResponse, err
	}

	path := "/fire/accounts"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttInquiryAccountResponse); err != nil {
		return ptr_ttInquiryAccountResponse, err
	}

	return ptr_ttInquiryAccountResponse, nil
}

//TTInquiryAccountBalance provides service to Inquiry balance for Vostro’s Account
func (c *Client) TTInquiryAccountBalance(ctx context.Context, ptr_ttInquiryAccountBalanceRequest *bca.TTInquiryAccountBalanceRequest) (*bca.TTInquiryAccountBalanceResponse, error) {
	var ptr_ttInquiryAccountBalanceResponse *bca.TTInquiryAccountBalanceResponse

	jsonReq, err := json.Marshal(*ptr_ttInquiryAccountBalanceRequest)
	if err != nil {
		return ptr_ttInquiryAccountBalanceResponse, err
	}

	path := "/fire/accounts/balance"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttInquiryAccountBalanceResponse); err != nil {
		return ptr_ttInquiryAccountBalanceResponse, err
	}

	return ptr_ttInquiryAccountBalanceResponse, nil
}

//TTInquiryTransaction provides service to Inquiry Transaction that has been submitted before
func (c *Client) TTInquiryTransaction(ctx context.Context, ptr_ttInquiryTransactionRequest *bca.TTInquiryTransactionRequest) (*bca.TTInquiryTransactionResponse, error) {
	var ptr_ttInquiryTransactionResponse *bca.TTInquiryTransactionResponse

	jsonReq, err := json.Marshal(*ptr_ttInquiryTransactionRequest)
	if err != nil {
		return ptr_ttInquiryTransactionResponse, err
	}

	path := "/fire/transactions"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttInquiryTransactionResponse); err != nil {
		return ptr_ttInquiryTransactionResponse, err
	}

	return ptr_ttInquiryTransactionResponse, nil
}

//TTCashTransfer provides service for transaction “Cash Transfer” to Non account holder
func (c *Client) TTCashTransfer(ctx context.Context, ptr_ttCashTransferRequest *bca.TTCashTransferRequest) (*bca.TTCashTransferResponse, error) {
	var ptr_ttCashTransferResponse *bca.TTCashTransferResponse

	jsonReq, err := json.Marshal(*ptr_ttCashTransferRequest)
	if err != nil {
		return ptr_ttCashTransferResponse, err
	}

	path := "/fire/transactions/cash-transfer"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttCashTransferResponse); err != nil {
		return ptr_ttCashTransferResponse, err
	}

	return ptr_ttCashTransferResponse, nil
}

//TTAmendCashTransfer provides service for Amendment “Cash Transfer” to Non account holder
func (c *Client) TTAmendCashTransfer(ctx context.Context, ptr_ttAmendCashTransferRequest *bca.TTAmendCashTransferRequest) (*bca.TTAmendCashTransferResponse, error) {
	var ptr_ttAmendCashTransferResponse *bca.TTAmendCashTransferResponse

	jsonReq, err := json.Marshal(*ptr_ttAmendCashTransferRequest)
	if err != nil {
		return ptr_ttAmendCashTransferResponse, err
	}

	path := "/fire/transactions/cash-transfer/amend"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttAmendCashTransferResponse); err != nil {
		return ptr_ttAmendCashTransferResponse, err
	}

	return ptr_ttAmendCashTransferResponse, nil
}

//TTCancelCashTransfer provides service for Cancellation “Cash Transfer” to Non account holder
func (c *Client) TTCancelCashTransfer(ctx context.Context, ptr_ttCancelCashTransferRequest *bca.TTCancelCashTransferRequest) (*bca.TTCancelCashTransferResponse, error) {
	var ptr_ttCancelCashTransferResponse *bca.TTCancelCashTransferResponse

	jsonReq, err := json.Marshal(*ptr_ttCancelCashTransferRequest)
	if err != nil {
		return ptr_ttCancelCashTransferResponse, err
	}

	path := "/fire/transactions/cash-transfer/cancel"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_ttCancelCashTransferResponse); err != nil {
		return ptr_ttCancelCashTransferResponse, err
	}

	return ptr_ttCancelCashTransferResponse, nil
}
