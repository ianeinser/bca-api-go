package fire

import (
	"context"
	"encoding/json"

	bca "github.com/ianeinser/go-bca"
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
func (c Client) TTAccount(ctx context.Context, request bca.TTAccountRequest) (bca.TTAccountResponse, error) {
	var response bca.TTAccountResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/transactions/to-account"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}

//TTInquiryAccount provides service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
func (c Client) TTInquiryAccount(ctx context.Context, request bca.TTInquiryAccountRequest) (bca.TTInquiryAccountResponse, error) {
	var response bca.TTInquiryAccountResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/accounts"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}

//TTInquiryAccountBalance provides service to Inquiry balance for Vostro’s Account
func (c Client) TTInquiryAccountBalance(ctx context.Context, request bca.TTInquiryAccountBalanceRequest) (bca.TTInquiryAccountBalanceResponse, error) {
	var response bca.TTInquiryAccountBalanceResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/accounts/balance"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}

//TTInquiryTransaction provides service to Inquiry Transaction that has been submitted before
func (c Client) TTInquiryTransaction(ctx context.Context, request bca.TTInquiryTransactionRequest) (bca.TTInquiryTransactionResponse, error) {
	var response bca.TTInquiryTransactionResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/transactions"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}

//TTCashTransfer provides service for transaction “Cash Transfer” to Non account holder
func (c Client) TTCashTransfer(ctx context.Context, request bca.TTCashTransferRequest) (bca.TTCashTransferResponse, error) {
	var response bca.TTCashTransferResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/transactions/cash-transfer"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}

//TTAmendCashTransfer provides service for Amendment “Cash Transfer” to Non account holder
func (c Client) TTAmendCashTransfer(ctx context.Context, request bca.TTAmendCashTransferRequest) (bca.TTAmendCashTransferResponse, error) {
	var response bca.TTAmendCashTransferResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/transactions/cash-transfer/amend"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}

//TTCancelCashTransfer provides service for Cancellation “Cash Transfer” to Non account holder
func (c Client) TTCancelCashTransfer(ctx context.Context, request bca.TTCancelCashTransferRequest) (bca.TTCancelCashTransferResponse, error) {
	var response bca.TTCancelCashTransferResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/fire/transactions/cash-transfer/cancel"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}

	return response, nil
}
