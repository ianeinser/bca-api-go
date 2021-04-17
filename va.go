package bca

//VAInquiryStatusPaymentRequest represents Virtual Account payment status request message
type VAInquiryStatusPaymentRequest struct {
	CompanyCode    string
	CustomerNumber string
	RequestID      string
}

//BillVAInquiryStatusPaymentResponse represents bills details to retrieve Virtual Account payment status
type BillVAInquiryStatusPaymentResponse struct {
	BillReference string
	BillNumber    string
}

//TransactionVAInquiryStatusPaymentResponse represents Virtual Account transaction info
type TransactionVAInquiryStatusPaymentResponse struct {
	DetailBills       []BillVAInquiryStatusPaymentResponse
	RequestID         string
	TransactionDate   string
	PaymentFlagStatus string
	Reference         string
	TotalAmount       float64 `json:",string"`
	PaidAmount        float64 `json:",string"`
}

//VAInquiryStatusPaymentResponse represents Virtual Account payment status response message
type VAInquiryStatusPaymentResponse struct {
	Error
	TransactionData []TransactionVAInquiryStatusPaymentResponse
}
