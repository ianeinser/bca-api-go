package bca

type ForeignExchangeRateRequest struct {
	CurrencyCode string
	RateType     string
}

type RateDetails struct {
	RateType   string  /// Must be among erate, tt, tc, bn.
	Buy        float64 `json:",string"`
	Sell       float64 `json:",string"`
	LastUpdate string
}

type Currency struct {
	CurrencyCode string /// The currency code e.g. IDR, USD, JPY
	RateDetail   []RateDetails
}

type ForeignExchangeRateResponse struct {
	Currencies      []Currency
	InvalidRateType string
	InvalidCurrency string
}
