package business

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	bca "github.com/ianeinser/go-bca"
)

var (
	httpHeaderChannelID    string = "ChannelID"
	httpHeaderCredentialID string = "CredentialID"
)

//Client is used to invoke BCA Business Banking API
type Client struct {
	Client       bca.APIImplementation
	CorporateID  string
	AccessToken  string
	ChannelID    string
	CredentialID string
}

//NewClient is used to initialize new business.Client
func NewClient(config bca.Config) Client {
	return Client{
		CorporateID:  config.CorporateID,
		ChannelID:    config.ChannelID,
		CredentialID: config.CredentialID,
		Client:       bca.NewAPI(config),
	}
}

//BBBalanceInformation is used to Get your KlikBCA Bisnis account balance information with maximum of 20 accounts in a request
func (c Client) BBBalanceInformation(ctx context.Context, request bca.BBBalanceInformationRequest) (bca.BBBalanceInformationResponse, error) {
	var response bca.BBBalanceInformationResponse
	path := fmt.Sprintf("/banking/v3/corporates/%s/accounts/%s", request.CorporateID, request.AccountNumber)

	if err := c.Client.Call("GET", path, c.AccessToken, nil, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

//BBAccountStatement is used to get your KlikBCA Bisnis account statement for a period up to 31 days
func (c Client) BBAccountStatement(ctx context.Context, request bca.BBAccountStatementRequest) (bca.BBAccountStatementResponse, error) {
	var response bca.BBAccountStatementResponse
	path := fmt.Sprintf("/banking/v3/corporates/%s/accounts/%s/statements", request.CorporateID, request.AccountNumber)

	startDate := request.StartDate
	endDate := request.EndDate

	v := url.Values{}
	v.Add("StartDate", startDate.Format("2006-01-02"))
	v.Add("EndDate", endDate.Format("2006-01-02"))
	path += "?" + v.Encode()

	if err := c.Client.Call("GET", path, c.AccessToken, nil, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

//BBFundTransfer is used to send fund transfer instructions to BCA using this service. The source of fund transfer must be from corporate’s own deposit account. The recipient may be any deposit account within BCA
func (c Client) BBFundTransfer(ctx context.Context, request bca.BBFundTransferRequest) (bca.BBFundTransferResponse, error) {
	var response bca.BBFundTransferResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/banking/corporates/transfers"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, &response); err != nil {
		return response, err
	}
	return response, nil
}

//BBDomesticFundTransfer is used to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate's own deposit account. The recipient may be any deposit account within domestic bank except BCA
func (c Client) BBDomesticFundTransfer(ctx context.Context, request bca.BBDomesticFundTransferRequest) (bca.BBDomesticFundTransferResponse, error) {
	var response bca.BBDomesticFundTransferResponse

	jsonReq, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	path := "/banking/corporates/transfers/domestic"

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("POST", path, c.AccessToken, headers, jsonReq, &response); err != nil {
		return response, err
	}
	return response, nil
}

//BBAccountStatementOffline is used to get your bulk statement in form of file for a period up to 7 days
func (c Client) BBAccountStatementOffline(ctx context.Context, request bca.BBAccountStatementOfflineRequest) (bca.BBAccountStatementOfflineResponse, error) {
	var response bca.BBAccountStatementOfflineResponse
	path := fmt.Sprintf("/banking/offline/corporates/accounts/%s/filestatements", request.AccountNumber)

	startDate := request.StartDate
	endDate := request.EndDate

	v := url.Values{}
	v.Add("StartDate", startDate.Format("2006-01-02"))
	v.Add("EndDate", endDate.Format("2006-01-02"))
	path += "?" + v.Encode()

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("GET", path, c.AccessToken, headers, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

//BBInquiryTransferStatus is used to get fund transfer status
func (c Client) BBInquiryTransferStatus(ctx context.Context, request bca.BBInquiryTransferStatusRequest) (bca.BBInquiryTransferStatusResponse, error) {
	var response bca.BBInquiryTransferStatusResponse
	path := fmt.Sprintf("/banking/corporates/transfers/status/%s", request.TransactionID)

	transactionDate := request.TransactionDate

	v := url.Values{}
	v.Add("TransactionDate", transactionDate.Format("2006-01-02"))
	v.Add("TransferType", request.TransferType)
	path += "?" + v.Encode()

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("GET", path, c.AccessToken, headers, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

//BBInquiryDomesticAccount is used to get beneficiary account information including beneficiary account name
func (c Client) BBInquiryDomesticAccount(ctx context.Context, request bca.BBInquiryDomesticAccountRequest) (bca.BBInquiryDomesticAccountResponse, error) {
	var response bca.BBInquiryDomesticAccountResponse
	path := fmt.Sprintf("/banking/corporates/transfers/v2/domestic/beneficiaries/banks/%s/accounts/%s", request.BeneficiaryBankCode, request.BeneficiaryAccountNumber)

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("GET", path, c.AccessToken, headers, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}
