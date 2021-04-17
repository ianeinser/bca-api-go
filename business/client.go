package business

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	bca "github.com/ianeinser/bca-api-go"
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

//BalanceInformation is used to Get your KlikBCA Bisnis account balance information with maximum of 20 accounts in a request
func (c *Client) BalanceInformation(ctx context.Context, ptr_balanceInformationRequest *bca.BalanceInformationRequest) (*bca.BalanceInformationResponse, error) {
	var ptr_balanceInformationResponse *bca.BalanceInformationResponse
	path := fmt.Sprintf("/banking/v3/corporates/%s/accounts/%s", (*ptr_balanceInformationRequest).CorporateID, (*ptr_balanceInformationRequest).AccountNumber)

	if err := c.Client.Call("GET", path, c.AccessToken, nil, nil, ptr_balanceInformationResponse); err != nil {
		return ptr_balanceInformationResponse, err
	}
	return ptr_balanceInformationResponse, nil
}

//AccountStatement is used to get your KlikBCA Bisnis account statement for a period up to 31 days
func (c *Client) AccountStatement(ctx context.Context, ptr_accountStatementRequest *bca.AccountStatementRequest) (*bca.AccountStatementResponse, error) {
	var ptr_accountStatementResponse *bca.AccountStatementResponse
	path := fmt.Sprintf("/banking/v3/corporates/%s/accounts/%s/statements", (*ptr_accountStatementRequest).CorporateID, (*ptr_accountStatementRequest).AccountNumber)

	startDate := (*ptr_accountStatementRequest).StartDate
	endDate := (*ptr_accountStatementRequest).EndDate

	v := url.Values{}
	v.Add("StartDate", startDate.Format("2006-01-02"))
	v.Add("EndDate", endDate.Format("2006-01-02"))
	path += "?" + v.Encode()

	if err := c.Client.Call("GET", path, c.AccessToken, nil, nil, ptr_accountStatementResponse); err != nil {
		return ptr_accountStatementResponse, err
	}
	return ptr_accountStatementResponse, nil
}

//FundTransfer is used to send fund transfer instructions to BCA using this service. The source of fund transfer must be from corporate’s own deposit account. The recipient may be any deposit account within BCA
func (c *Client) FundTransfer(ctx context.Context, ptr_fundTransferRequest *bca.FundTransferRequest) (*bca.FundTransferResponse, error) {
	var ptr_fundTransferResponse *bca.FundTransferResponse

	jsonReq, err := json.Marshal(*ptr_fundTransferRequest)
	if err != nil {
		return ptr_fundTransferResponse, err
	}

	path := "/banking/corporates/transfers"

	if err := c.Client.Call("POST", path, c.AccessToken, nil, jsonReq, ptr_fundTransferResponse); err != nil {
		return ptr_fundTransferResponse, err
	}
	return ptr_fundTransferResponse, nil
}

//DomesticFundTransfer is used to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate's own deposit account. The recipient may be any deposit account within domestic bank except BCA
func (c *Client) DomesticFundTransfer(ctx context.Context, ptr_domesticFundTransferRequest *bca.DomesticFundTransferRequest) (*bca.DomesticFundTransferResponse, error) {
	var ptr_domesticFundTransferResponse *bca.DomesticFundTransferResponse

	jsonReq, err := json.Marshal(*ptr_domesticFundTransferRequest)
	if err != nil {
		return ptr_domesticFundTransferResponse, err
	}

	path := "/banking/corporates/transfers/domestic"

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("POST", path, c.AccessToken, headers, jsonReq, ptr_domesticFundTransferResponse); err != nil {
		return ptr_domesticFundTransferResponse, err
	}
	return ptr_domesticFundTransferResponse, nil
}

//AccountStatementOffline is used to get your bulk statement in form of file for a period up to 7 days
func (c *Client) AccountStatementOffline(ctx context.Context, ptr_accountStatementOfflineRequest *bca.AccountStatementOfflineRequest) (*bca.AccountStatementOfflineResponse, error) {
	var ptr_accountStatementOfflineResponse *bca.AccountStatementOfflineResponse
	path := fmt.Sprintf("/banking/offline/corporates/accounts/%s/filestatements", (*ptr_accountStatementOfflineRequest).AccountNumber)

	startDate := (*ptr_accountStatementOfflineRequest).StartDate
	endDate := (*ptr_accountStatementOfflineRequest).EndDate

	v := url.Values{}
	v.Add("StartDate", startDate.Format("2006-01-02"))
	v.Add("EndDate", endDate.Format("2006-01-02"))
	path += "?" + v.Encode()

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("GET", path, c.AccessToken, headers, nil, ptr_accountStatementOfflineResponse); err != nil {
		return ptr_accountStatementOfflineResponse, err
	}
	return ptr_accountStatementOfflineResponse, nil
}

//InquiryTransferStatus is used to get fund transfer status
func (c *Client) InquiryTransferStatus(ctx context.Context, ptr_inquiryTransferStatusRequest *bca.InquiryTransferStatusRequest) (*bca.InquiryTransferStatusResponse, error) {
	var ptr_inquiryTransferStatusResponse *bca.InquiryTransferStatusResponse
	path := fmt.Sprintf("/banking/corporates/transfers/status/%s", (*ptr_inquiryTransferStatusRequest).TransactionID)

	transactionDate := (*ptr_inquiryTransferStatusRequest).TransactionDate

	v := url.Values{}
	v.Add("TransactionDate", transactionDate.Format("2006-01-02"))
	v.Add("TransferType", (*ptr_inquiryTransferStatusRequest).TransferType)
	path += "?" + v.Encode()

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("GET", path, c.AccessToken, headers, nil, ptr_inquiryTransferStatusResponse); err != nil {
		return ptr_inquiryTransferStatusResponse, err
	}
	return ptr_inquiryTransferStatusResponse, nil
}

//InquiryDomesticAccount is used to get beneficiary account information including beneficiary account name
func (c *Client) InquiryDomesticAccount(ctx context.Context, ptr_inquiryDomesticAccountRequest *bca.InquiryDomesticAccountRequest) (*bca.InquiryDomesticAccountResponse, error) {
	var ptr_inquiryDomesticAccountResponse *bca.InquiryDomesticAccountResponse
	path := fmt.Sprintf("/banking/corporates/transfers/v2/domestic/beneficiaries/banks/%s/accounts/%s", (*ptr_inquiryDomesticAccountRequest).BeneficiaryBankCode, (*ptr_inquiryDomesticAccountRequest).BeneficiaryAccountNumber)

	headers := map[string]string{
		httpHeaderChannelID:    c.ChannelID,
		httpHeaderCredentialID: c.CredentialID,
	}

	if err := c.Client.Call("GET", path, c.AccessToken, headers, nil, ptr_inquiryDomesticAccountResponse); err != nil {
		return ptr_inquiryDomesticAccountResponse, err
	}
	return ptr_inquiryDomesticAccountResponse, nil
}
