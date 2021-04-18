package bca

//Auth represents authentication info used in FIRe transaction
type Auth struct {
	CorporateID string
	AccessCode  string
	BranchCode  string
	UserID      string
	LocalID     string
}

//SenderAccountRequest represents sender details used in TTAccountRequest
type SenderAccountRequest struct {
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

//BeneficiaryAccountRequest represents beneficiary details used in TTAccountRequest
type BeneficiaryAccountRequest struct {
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
type TransactionAccountRequest struct {
	CurrencyID      string
	Amount          float64 `json:",string"`
	PurposeCode     string
	Description1    string
	Description2    string
	DetailOfCharges string
	SourceOfFund    string
	FormNumber      string
}

//TeleTransferAccountRequest is to provides service transaction “Transaction to BCA’s Account” and also “Transfer to Other Bank”
type TeleTransferAccountRequest struct {
	Authentication     Auth
	SenderDetails      SenderAccountRequest
	BeneficiaryDetails BeneficiaryAccountRequest
	TransactionDetails TransactionAccountRequest
}

//BeneficiaryTTAccountResponse represents beneficiary details for response message
type BeneficiaryAccountResponse struct {
	Name                  string
	AccountNumber         string
	ServerBeneAccountName string
}

//TransactionTTAccountResponse represents transaction details for response message
type TransactionAccountResponse struct {
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
type TeleTransferAccountResponse struct {
	Error
	BeneficiaryDetails BeneficiaryAccountResponse
	TransactionDetails TransactionAccountResponse
	StatusTransaction  string
	StatusMessage      string
}

//BeneficiaryInquiryAccountRequest represents beneficiary details to inquire account(s)
type BeneficiaryInquiryAccountRequest struct {
	BankCodeType  string
	BankCodeValue string
	AccountNumber string
}

//InquiryAccountRequest is to provide service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
type InquiryAccountRequest struct {
	Authentication     Auth
	BeneficiaryDetails BeneficiaryInquiryAccountRequest
}

//BeneficiaryInquiryAccountResponse represents response payload to inquire accounts
type BeneficiaryInquiryAccountResponse struct {
	ServerBeneAccountName string
}

//InquiryAccountResponse is to provide service to Inquiry BCA’s Account name or Other Bank Switching’s Account name.
type InquiryAccountResponse struct {
	Error
	BeneficiaryDetails BeneficiaryInquiryAccountResponse
	StatusTransaction  string
	StatusMessage      string
}

//InquiryAccountBalanceRequest represents FID details to inquire account balance
type FIInquiryAccountBalanceRequest struct {
	AccountNumber string
}

//InquiryAccountBalanceRequest is to provide service to Inquiry balance for Vostro’s Account.
type InquiryAccountBalanceRequest struct {
	Authentication Auth
	FIDetails      FIInquiryAccountBalanceRequest
}

//FIInquiryAccountBalanceResponse represents FID details in response message
type FIInquiryAccountBalanceResponse struct {
	CurrencyID     string
	AccountBalance float64 `json:",string"`
}

//TTInquiryAccountBalanceResponse is to provide service to Inquiry balance for Vostro’s Account.
type InquiryAccountBalanceResponse struct {
	Error
	FIDetails         FIInquiryAccountBalanceResponse
	StatusTransaction string
	StatusMessage     string
}

//TransactionTTInquiryTransactionRequest represents transaction details to inquire transaction
type TransactionInquiryTransactionRequest struct {
	InquiryBy    string
	InquiryValue string
}

//InquiryTransactionRequest is to provide service to Inquiry Transaction that has been submitted before
type InquiryTransactionRequest struct {
	Authentication     Auth
	TransactionDetails TransactionInquiryTransactionRequest
}

//SenderInquiryTransactionResponse represents sender details for response message
type SenderInquiryTransactionResponse struct {
	FirstName string
	LastName  string
}

//BeneficiaryInquiryTransactionResponse represents beneficiary details for response message
type BeneficiaryInquiryTransactionResponse struct {
	Name          string
	BankCodeType  string
	BankCodeValue string
	AccountNumber string
}

//TransactionInquiryTransactionResponse represents transaction details for response message
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

//InquiryTransactionResponse is to provide service to Inquiry Transaction that has been submitted before
type InquiryTransactionResponse struct {
	Error
	SenderDetails      SenderInquiryTransactionResponse
	BeneficiaryDetails BeneficiaryInquiryTransactionResponse
	TransactionDetails TransactionTTInquiryTransactionResponse
	StatusTransaction  string
	StatusMessage      string
}

//SenderTeleTransferCashTransferRequest represents sender details used in TTCashTransferRequest
type SenderTeleTransferCashTransferRequest struct {
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

//BeneficiaryTeleTransferCashTransferRequest represents beneficiary details used in TTCashTransferRequest
type BeneficiaryTeleTransferCashTransferRequest struct {
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

//TransactionTeleTransferCashTransferRequest represents transaction details used in TTCashTransferRequest
type TransactionTeleTransferCashTransferRequest struct {
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

//TeleTransferCashTransferRequest is to provide service for transaction “Cash Transfer” to Non account holder.
type TeleTransferCashTransferRequest struct {
	Authentication     Auth
	SenderDetails      SenderTeleTransferCashTransferRequest
	BeneficiaryDetails BeneficiaryTeleTransferCashTransferRequest
	TransactionDetails TransactionTeleTransferCashTransferRequest
}

//BeneficiaryTeleTransferCashTransferResponse represents beneficiary details in response message
type BeneficiaryTeleTransferCashTransferResponse struct {
	Name string
}

//TransactionTeleTransferCashTransferResponse represents transaction details in response message
type TransactionTeleTransferCashTransferResponse struct {
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
type TeleTransferCashTransferResponse struct {
	Error
	BeneficiaryDetails BeneficiaryTeleTransferCashTransferResponse
	TransactionDetails TransactionTeleTransferCashTransferResponse
	StatusTransaction  string
	StatusMessage      string
}

//SenderTeleTransferAmendCashTransfer represents sender details to amend cash-transfer
type SenderTeleTransferAmendCashTransfer struct {
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
type BeneficiaryTeleTransferAmendCashTransfer struct {
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

//Transaction1TeleTransferAmendCashTransfer represents transaction details to amend cash-transfer
type Transaction1TeleTransferAmendCashTransfer struct {
	Description1   string
	Description2   string
	SecretQuestion string
	SecretAnswer   string
}

//AmendmentTeleTransferAmendCashTransfer represents amendment details to amend cash-transfer
type AmendmentTeleTransferAmendCashTransfer struct {
	SenderDetails      SenderTeleTransferAmendCashTransfer
	BeneficiaryDetails BeneficiaryTeleTransferAmendCashTransfer
	TransactionDetails Transaction1TeleTransferAmendCashTransfer
}

//Transaction2TeleTransferAmendCashTransfer represents transaction details to amend cash-transfer
type Transaction2TTAmendCashTransfer struct {
	FormNumber string
}

//TeleTransferAmendCashTransferRequest is to provide service for Amendment “Cash Transfer” to Non account holder
type TeleTransferAmendCashTransferRequest struct {
	Authentication     Auth
	AmendmentDetails   AmendmentTeleTransferAmendCashTransfer
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

//TeleTransferAmendCashTransferResponse is to provide service for Amendment “Cash Transfer” to Non account holder
type TeleTransferAmendCashTransferResponse struct {
	Error
	AmendmentDetails   AmendmentTeleTransferAmendCashTransfer
	TransactionDetails Transaction2TTAmendCashTransfer
	StatusTransaction  string
	StatusMessage      string
}

//TransactionTTCancelCashTransferRequest represents transaction details to cancel cash-transfer
type TransactionTeleTransferCancelCashTransferRequest struct {
	FormNumber string
	Amount     float64 `json:",string"`
	CurrencyID string
}

//TTCancelCashTransferRequest is to provide service for Cancellation “Cash Transfer” to Non account holder
type TeleTransferCancelCashTransferRequest struct {
	Authentication     Auth
	TransactionDetails TransactionTeleTransferCancelCashTransferRequest
}

//TransactionTTCancelCashTransferResponse represents transaction details to cancel cash-transfer
type TransactionTeleTransferCancelCashTransferResponse struct {
	FormNumber      string
	ReleaseDateTime string
}

//TTCancelCashTransferResponse is to provide service for Cancellation “Cash Transfer” to Non account holder
type TeleTransferCancelCashTransferResponse struct {
	Error
	TransactionDetails TransactionTeleTransferCancelCashTransferResponse
	StatusTransaction  string
	StatusMessage      string
}
