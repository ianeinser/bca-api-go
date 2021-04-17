package bca

import "time"

//BBBalanceInformationRequest is to get your KlikBCA Bisnis account balance information with maximum of 20 accounts in a request
type BBBalanceInformationRequest struct {
	CorporateID   string
	AccountNumber string
}

//BalanceSuccessBBBalanceInformationResponse represents account balance details when it is successfully retrieved
type BalanceSuccessBBBalanceInformationResponse struct {
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
type BalanceFailedBBBalanceInformationResponse struct {
	Indonesian    string
	English       string
	AccountNumber string
}

//BBBalanceInformationResponse is to get your KlikBCA Bisnis account balance information with maximum of 20 accounts in a request
type BBBalanceInformationResponse struct {
	Error
	AccountDetailDataSuccess []BalanceSuccessBBBalanceInformationResponse `json:",omitempty"`
	AccountDetailDataFailed  []BalanceFailedBBBalanceInformationResponse  `json:",omitempty"`
}

//BBAccountStatementRequest is to get your KlikBCA Bisnis account statement for a period up to 31 days
type BBAccountStatementRequest struct {
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

//BBAccountStatementResponse is to get your KlikBCA Bisnis account statement for a period up to 31 days
type BBAccountStatementResponse struct {
	Error
	Currency     string
	StartBalance float64 `json:",string"`
	StartDate    string
	EndDate      string
	Data         []AccountStatement
}

//BBFundTransferRequest is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate’s own deposit account. The recipient may be any deposit account within BCA
type BBFundTransferRequest struct {
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

//BBFundTransferResponse is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate’s own deposit account. The recipient may be any deposit account within BCA
type BBFundTransferResponse struct {
	Error
	TransactionID   string
	TransactionDate string
	ReferenceID     string
	Status          string
}

//BBDomesticFundTransferRequest is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate's own deposit account. The recipient may be any deposit account within domestic bank except BCA.
type BBDomesticFundTransferRequest struct {
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

//BBDomesticFundTransferResponse is to send fund transfer instructions to BCA using this service. The source of fund transfer must be from your corporate's own deposit account. The recipient may be any deposit account within domestic bank except BCA.
type BBDomesticFundTransferResponse struct {
	Error
	TransactionID   string
	TransactionDate string
	ReferenceID     string
	PPUNumber       string
}

//BBAccountStatementOfflineRequest is to get your bulk statement in form of file for a period up to 7 days
type BBAccountStatementOfflineRequest struct {
	AccountNumber string
	StartDate     time.Time
	EndDate       time.Time
}

//BBAccountStatementOfflineResponse is to get your bulk statement in form of file for a period up to 7 days
type BBAccountStatementOfflineResponse struct {
	RequestID  string
	ResponseWS string
}

//BBInquiryTransferStatusRequest is to get fund transfer status
type BBInquiryTransferStatusRequest struct {
	TransactionID   string
	TransactionDate time.Time
	TransferType    string
}

//ReasonBBInquiryTransferStatusResponse represents
type ReasonBBInquiryTransferStatusResponse struct {
	English    string
	Indonesian string
}

//BBInquiryTransferStatusResponse is to get fund transfer status
type BBInquiryTransferStatusResponse struct {
	Error
	TransactionID            string
	TransactionDate          time.Time
	TransferType             string
	SourceAccountNumber      string
	BeneficiaryAccountNumber string
	CurrencyCode             string
	Amount                   float64 `json:",string"`
	StatusCode               string
	Reason                   ReasonBBInquiryTransferStatusResponse
}

//BBInquiryDomesticAccountRequest is to get beneficiary account information including beneficiary account name
type BBInquiryDomesticAccountRequest struct {
	BeneficiaryAccountNumber string
	BeneficiaryBankCode      string
}

//BBInquiryDomesticAccountResponse is to get beneficiary account information including beneficiary account name
type BBInquiryDomesticAccountResponse struct {
	BeneficiaryBankCode      string
	BeneficiaryAccountNumber string
	BeneficiaryAccountName   string
}
