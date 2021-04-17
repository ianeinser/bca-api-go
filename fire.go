package bca

//Auth represents authentication info used in FIRe transaction
type Auth struct {
	CorporateID string
	AccessCode  string
	BranchCode  string
	UserID      string
	LocalID     string
}

//SenderTTAccountRequest represents sender details used in TTAccountRequest
type SenderTTAccountRequest struct {
	FirstName            string
	LastName             string
	DateOfBirth          string
	Address1             string
	Address2             string
	City                 string
	StateID              string
	PostalCode           string
	CountryID            string
	Mobile               string
	IdentificationType   string
	IdentificationNumber string
	AccountNumber        string
}

//BeneficiaryTTAccountRequest represents beneficiary details used in TTAccountRequest
type BeneficiaryTTAccountRequest struct {
	Name                 string
	DateOfBirth          string
	Address1             string
	Address2             string
	City                 string
	StateID              string
	PostalCode           string
	CountryID            string
	Mobile               string
	IdentificationType   string
	IdentificationNumber string
	NationalityID        string
	Occupation           string
	BankCodeType         string
	BankCodeValue        string
	BankCountryID        string
	BankAddress          string
	BankCity             string
	AccountNumber        string
}

//TransactionTTAccountRequest represents transaction details used in TTAccountRequest
type TransactionTTAccountRequest struct {
	CurrencyID      string
	Amount          float64 `json:",string"`
	PurposeCode     string
	Description1    string
	Description2    string
	DetailOfCharges string
	SourceOfFund    string
	FormNumber      string
}

//TTAccountRequest is to provides service transaction “Transaction to BCA’s Account” and also “Transfer to Other Bank”
type TTAccountRequest struct {
	Authentication     Auth
	SenderDetails      SenderTTAccountRequest
	BeneficiaryDetails BeneficiaryTTAccountRequest
	TransactionDetails TransactionTTAccountRequest
}

//BeneficiaryTTAccountResponse represents beneficiary details for response message
type BeneficiaryTTAccountResponse struct {
	Name                  string
	AccountNumber         string
	ServerBeneAccountName string
}

//TransactionTTAccountResponse represents transaction details for response message
type TransactionTTAccountResponse struct {
	CurrencyID        string
	Amount            float64 `json:",string"`
	Description1      string
	Description2      string
	FormNumber        string
	ReferenceNumber   string
	ReleaseDateTime   string
	StatusTransaction string `json:",omitempty"`
	StatusMessage     string `json:",omitempty"`
}

//TTAccountResponse is to provide service transaction “Transaction to BCA’s Account” and also “Transfer to Other Bank”
type TTAccountResponse struct {
	Error
	BeneficiaryDetails BeneficiaryTTAccountResponse
	TransactionDetails TransactionTTAccountResponse
	StatusTransaction  string
	StatusMessage      string
}

//BeneficiaryTTInquiryAccountRequest represents beneficiary details to inquire account(s)
type BeneficiaryTTInquiryAccountRequest struct {
	BankCodeType  string
	BankCodeValue string
	AccountNumber string
}

//TTInquiryAccountRequest is to provide service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
type TTInquiryAccountRequest struct {
	Authentication     Auth
	BeneficiaryDetails BeneficiaryTTInquiryAccountRequest
}

//BeneficiaryTTInquiryAccountResponse represents response payload to inquire accounts
type BeneficiaryTTInquiryAccountResponse struct {
	ServerBeneAccountName string
}

//TTInquiryAccountResponse is to provide service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
type TTInquiryAccountResponse struct {
	Error
	BeneficiaryDetails BeneficiaryTTInquiryAccountResponse
	StatusTransaction  string
	StatusMessage      string
}

//FITTInquiryAccountBalanceRequest represents FID details to inquire account balance
type FITTInquiryAccountBalanceRequest struct {
	AccountNumber string
}

//TTInquiryAccountBalanceRequest is to provide service to Inquiry balance for Vostro’s Account.
type TTInquiryAccountBalanceRequest struct {
	Authentication Auth
	FIDetails      FITTInquiryAccountBalanceRequest
}

//FITTInquiryAccountBalanceResponse represents FID details in response message
type FITTInquiryAccountBalanceResponse struct {
	CurrencyID     string
	AccountBalance float64 `json:",string"`
}

//TTInquiryAccountBalanceResponse is to provide service to Inquiry balance for Vostro’s Account.
type TTInquiryAccountBalanceResponse struct {
	Error
	FIDetails         FITTInquiryAccountBalanceResponse
	StatusTransaction string
	StatusMessage     string
}

//TransactionTTInquiryTransactionRequest represents transaction details to inquire transaction
type TransactionTTInquiryTransactionRequest struct {
	InquiryBy    string
	InquiryValue string
}

//TTInquiryTransactionRequest is to provide service to Inquiry Transaction that has been submitted before
type TTInquiryTransactionRequest struct {
	Authentication     Auth
	TransactionDetails TransactionTTInquiryTransactionRequest
}

//SenderTTInquiryTransactionResponse represents sender details for response message
type SenderTTInquiryTransactionResponse struct {
	FirstName string
	LastName  string
}

//BeneficiaryTTInquiryTransactionResponse represents beneficiary details for response message
type BeneficiaryTTInquiryTransactionResponse struct {
	Name          string
	BankCodeType  string
	BankCodeValue string
	AccountNumber string
}

//TransactionTTInquiryTransactionResponse represents transaction details for response message
type TransactionTTInquiryTransactionResponse struct {
	AmountPaid      float64 `json:",string"`
	CurrencyID      string
	ReleaseDateTime string
	LocalID         string
	FormNumber      string
	ReferenceNumber string
	PIN             string
	Description1    string
	Description2    string
}

//TTInquiryTransactionResponse is to provide service to Inquiry Transaction that has been submitted before
type TTInquiryTransactionResponse struct {
	Error
	SenderDetails      SenderTTInquiryTransactionResponse
	BeneficiaryDetails BeneficiaryTTInquiryTransactionResponse
	TransactionDetails TransactionTTInquiryTransactionResponse
	StatusTransaction  string
	StatusMessage      string
}

//SenderTTCashTransferRequest represents sender details used in TTCashTransferRequest
type SenderTTCashTransferRequest struct {
	FirstName            string
	LastName             string
	DateOfBirth          string
	Address1             string
	Address2             string
	City                 string
	StateID              string
	PostalCode           string
	CountryID            string
	Mobile               string
	IdentificationType   string
	IdentificationNumber string
}

//BeneficiaryTTCashTransferRequest represents beneficiary details used in TTCashTransferRequest
type BeneficiaryTTCashTransferRequest struct {
	Name                 string
	DateOfBirth          string
	Address1             string
	Address2             string
	City                 string
	StateID              string
	PostalCode           string
	CountryID            string
	Mobile               string
	IdentificationType   string
	IdentificationNumber string
	NationalityID        string
	Occupation           string
}

//TransactionTTCashTransferRequest represents transaction details used in TTCashTransferRequest
type TransactionTTCashTransferRequest struct {
	PIN             string
	SecretQuestion  string
	SecretAnswer    string
	CurrencyID      string
	Amount          float64 `json:",string"`
	PurposeCode     string
	Description1    string
	Description2    string
	DetailOfCharges string
	SourceOfFund    string
	FormNumber      string
}

//TTCashTransferRequest is to provide service for transaction “Cash Transfer” to Non account holder.
type TTCashTransferRequest struct {
	Authentication     Auth
	SenderDetails      SenderTTCashTransferRequest
	BeneficiaryDetails BeneficiaryTTCashTransferRequest
	TransactionDetails TransactionTTCashTransferRequest
}

//BeneficiaryTTCashTransferResponse represents beneficiary details in response message
type BeneficiaryTTCashTransferResponse struct {
	Name string
}

//TransactionTTCashTransferResponse represents transaction details in response message
type TransactionTTCashTransferResponse struct {
	PIN             string
	CurrencyID      string
	Amount          float64 `json:",string"`
	Description1    string
	Description2    string
	FormNumber      string
	ReferenceNumber string
	ReleaseDateTime string
}

//TTCashTransferResponse is to provide service for transaction “Cash Transfer” to Non account holder.
type TTCashTransferResponse struct {
	Error
	BeneficiaryDetails BeneficiaryTTCashTransferResponse
	TransactionDetails TransactionTTCashTransferResponse
	StatusTransaction  string
	StatusMessage      string
}

//SenderTTAmendCashTransfer represents sender details to amend cash-transfer
type SenderTTAmendCashTransfer struct {
	FirstName            string
	LastName             string
	DateOfBirth          string
	Address1             string
	Address2             string
	City                 string
	StateID              string
	PostalCode           string
	CountryID            string
	Mobile               string
	IdentificationType   string
	IdentificationNumber string
}

//BeneficiaryTTAmendCashTransfer represents beneficiary details to amend cash-transfer
type BeneficiaryTTAmendCashTransfer struct {
	Name                 string
	DateOfBirth          string
	Address1             string
	Address2             string
	City                 string
	StateID              string
	PostalCode           string
	CountryID            string
	Mobile               string
	IdentificationType   string
	IdentificationNumber string
	NationalityID        string
	Occupation           string
}

//Transaction1TTAmendCashTransfer represents transaction details to amend cash-transfer
type Transaction1TTAmendCashTransfer struct {
	Description1   string
	Description2   string
	SecretQuestion string
	SecretAnswer   string
}

//AmendmentTTAmendCashTransfer represents amendment details to amend cash-transfer
type AmendmentTTAmendCashTransfer struct {
	SenderDetails      SenderTTAmendCashTransfer
	BeneficiaryDetails BeneficiaryTTAmendCashTransfer
	TransactionDetails Transaction1TTAmendCashTransfer
}

//Transaction2TTAmendCashTransfer represents transaction details to amend cash-transfer
type Transaction2TTAmendCashTransfer struct {
	FormNumber string
}

//TTAmendCashTransferRequest is to provide service for Amendment “Cash Transfer” to Non account holder
type TTAmendCashTransferRequest struct {
	Authentication     Auth
	AmendmentDetails   AmendmentTTAmendCashTransfer
	TransactionDetails Transaction2TTAmendCashTransfer
}

//TransactionTTAmendCashTransferResponse represents transaction details to amend cash-transfer
type TransactionTTAmendCashTransferResponse struct {
	Description1   string
	Description2   string
	SecretQuestion string
	SecretAnswer   string
	FormNumber     string
}

//TTAmendCashTransferResponse is to provide service for Amendment “Cash Transfer” to Non account holder
type TTAmendCashTransferResponse struct {
	Error
	AmendmentDetails   AmendmentTTAmendCashTransfer
	TransactionDetails Transaction2TTAmendCashTransfer
	StatusTransaction  string
	StatusMessage      string
}

//TransactionTTCancelCashTransferRequest represents transaction details to cancel cash-transfer
type TransactionTTCancelCashTransferRequest struct {
	FormNumber string
	Amount     float64 `json:",string"`
	CurrencyID string
}

//TTCancelCashTransferRequest is to provide service for Cancellation “Cash Transfer” to Non account holder
type TTCancelCashTransferRequest struct {
	Authentication     Auth
	TransactionDetails TransactionTTCancelCashTransferRequest
}

//TransactionTTCancelCashTransferResponse represents transaction details to cancel cash-transfer
type TransactionTTCancelCashTransferResponse struct {
	FormNumber      string
	ReleaseDateTime string
}

//TTCancelCashTransferResponse is to provide service for Cancellation “Cash Transfer” to Non account holder
type TTCancelCashTransferResponse struct {
	Error
	TransactionDetails TransactionTTCancelCashTransferResponse
	StatusTransaction  string
	StatusMessage      string
}
