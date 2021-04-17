package bca

import "time"

//BalanceInformationRequest is to get your KlikBCA Bisnis account balance information with maximum of 20 accounts in a request
type BalanceInformationRequest struct {
	CorporateID   string
	AccountNumber string
}

//BalanceSuccessBBBalanceInformationResponse represents account balance details when it is successfully retrieved
type BalanceSuccessBalanceInformationResponse struct {
	AccountNumber    string
	Currency         string
	Balance          float64 `json:",string"`
	AvailableBalance float64 `json:",string"`
	FloatAmount      float64 `json:",string"`
	HoldAmount       float64 `json:",string"`
	Plafon           float64 `json:",string"`
	Indonesian       string
	English          string
}

//BalanceFailedBBBalanceInformationResponse represents account balance details when it cannot be retrieved
type BalanceFailedBalanceInformationResponse struct {
	Indonesian    string
	English       string
	AccountNumber string
}

//BalanceInformationResponse is to get your KlikBCA Bisnis account balance information with maximum of 20 accounts in a request
type BalanceInformationResponse struct {
	Error
	AccountDetailDataSuccess []BalanceSuccessBalanceInformationResponse `json:",omitempty"`
	AccountDetailDataFailed  []BalanceFailedBalanceInformationResponse  `json:",omitempty"`
}

//AccountStatementRequest is to get your KlikBCA Bisnis account statement for a period up to 31 days
type AccountStatementRequest struct {
	CorporateID   string
	AccountNumber string
	StartDate     time.Time
	EndDate       time.Time
}

//AccountStatement represents account statement information
type AccountStatement struct {
	TransactionDate   string
	BranchCode        string
	TransactionType   string
	TransactionAmount float64 `json:",string"`
	TransactionName   string
	Trailer           string
}

//AccountStatementResponse is to get your KlikBCA Bisnis account statement for a period up to 31 days
type AccountStatementResponse struct {
	Error
	Currency     string
	StartBalance float64 `json:",string"`
	StartDate    string
	EndDate      string
	Data         []AccountStatement
}

//FundTransferRequest is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate’s own deposit account. The recipient may be any deposit account within BCA
type FundTransferRequest struct {
	CorporateID              string
	SourceAccountNumber      string
	TransactionID            string
	TransactionDate          string
	ReferenceID              string
	CurrencyCode             string
	Amount                   float64 `json:",string"`
	BeneficiaryAccountNumber string
	Remark1                  string
	Remark2                  string
}

//FundTransferResponse is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate’s own deposit account. The recipient may be any deposit account within BCA
type FundTransferResponse struct {
	Error
	TransactionID   string
	TransactionDate string
	ReferenceID     string
	Status          string
}

//DomesticFundTransferRequest is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate's own deposit account. The recipient may be any deposit account within domestic bank except BCA.
type DomesticFundTransferRequest struct {
	TransactionID            string
	TransactionDate          string
	ReferenceID              string
	SourceAccountNumber      string
	BeneficiaryAccountNumber string
	BeneficiaryBankCode      string
	BeneficiaryName          string
	Amount                   float64 `json:",string"`
	TransferType             string
	BeneficiaryCustType      string
	BeneficiaryCustResidence string
	CurrencyCode             string
	Remark1                  string
	Remark2                  string
	BeneficiaryEmail         string
}

//DomesticFundTransferResponse is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate's own deposit account. The recipient may be any deposit account within domestic bank except BCA.
type DomesticFundTransferResponse struct {
	Error
	TransactionID   string
	TransactionDate string
	ReferenceID     string
	PPUNumber       string
}

//AccountStatementOfflineRequest is to get your bulk statement in form of file for a period up to 7 days
type AccountStatementOfflineRequest struct {
	AccountNumber string
	StartDate     time.Time
	EndDate       time.Time
}

//AccountStatementOfflineResponse is to get your bulk statement in form of file for a period up to 7 days
type AccountStatementOfflineResponse struct {
	RequestID  string
	ResponseWS string
}

//InquiryTransferStatusRequest is to get fund transfer status
type InquiryTransferStatusRequest struct {
	TransactionID   string
	TransactionDate time.Time
	TransferType    string
}

//ReasonInquiryTransferStatusResponse represents
type ReasonInquiryTransferStatusResponse struct {
	English    string
	Indonesian string
}

//BBInquiryTransferStatusResponse is to get fund transfer status
type InquiryTransferStatusResponse struct {
	Error
	TransactionID            string
	TransactionDate          time.Time
	TransferType             string
	SourceAccountNumber      string
	BeneficiaryAccountNumber string
	CurrencyCode             string
	Amount                   float64 `json:",string"`
	StatusCode               string
	Reason                   ReasonInquiryTransferStatusResponse
}

//InquiryDomesticAccountRequest is to get beneficiary account information including beneficiary account name
type InquiryDomesticAccountRequest struct {
	BeneficiaryAccountNumber string
	BeneficiaryBankCode      string
}

//InquiryDomesticAccountResponse is to get beneficiary account information including beneficiary account name
type InquiryDomesticAccountResponse struct {
	BeneficiaryBankCode      string
	BeneficiaryAccountNumber string
	BeneficiaryAccountName   string
}
