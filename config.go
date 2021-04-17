package bca

//Config represents configuration that needed by BCA API
type Config struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	APISecret    string
	URL          string
	CorporateID  string
	OriginHost   string

	ChannelID    string
	CredentialID string
	CompanyCode  string

	FIReCorporateID string
	AccessCode      string
	BranchCode      string
	UserID          string
	LocalID         string

	LogLevel int
	LogPath  string
}
