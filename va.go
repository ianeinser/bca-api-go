package bca

//InquiryStatusPaymentRequest represents Virtual Account payment status request message
type InquiryStatusPaymentRequest struct {
	CompanyCode    string
	CustomerNumber string
	RequestID      string
}

//DetailsBill represents bills details to retrieve Virtual Account payment status
type DetailsBill struct {
	BillReference string
	BillNumber    string
}

//VAInquiryStatusPaymentResponse represents Virtual Account transaction info
type VAInquiryStatusPaymentResponse struct {
	DetailBills       []DetailsBill
	RequestID         string
	TransactionDate   string
	PaymentFlagStatus string
	Reference         string
	TotalAmount       float64 `json:",string"`
	PaidAmount        float64 `json:",string"`
}

//VAInquiryStatusPaymentResponse represents Virtual Account payment status response message
type InquiryStatusPaymentResponse struct {
	Error
	TransactionData []InquiryStatusPaymentResponse
}
