package bca

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/juju/errors"
)

//API is an interface for making call to BCA API
type API interface {
	//Call(method, path string, body io.Reader, v interface{}) error
	Call(method, path string, body []byte, v interface{}) error
	CallRaw(method, path string, headers http.Header, body io.Reader, v interface{}) error
}

//APIImplementation represents config that used for HTTP client needs
type APIImplementation struct {
	APIKey     string
	APISecret  string
	OriginHost string
	URL        string
	HTTPClient *http.Client
	LogLevel   int
	Logger     *log.Logger
}

//NewAPI is used to initialize new APIImplementation
func NewAPI(cfg Config) APIImplementation {
	return APIImplementation{
		APIKey:     cfg.APIKey,
		APISecret:  cfg.APISecret,
		URL:        cfg.URL,
		OriginHost: cfg.OriginHost,
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// 0: no logging
		// 1: errors only
		// 2: errors + informational (default)
		// 3: errors + informational + debug
		LogLevel: cfg.LogLevel,
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
	}
}

//Call is the implementation for invoking BCA API with its authentication
//func (c *APIImplementation) Call(method, path, accessToken string, body io.Reader, v interface{}) error {
func (c *APIImplementation) Call(method, path, accessToken string, additionalHeader map[string]string, body []byte, v interface{}) error {

	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+accessToken)
	headers.Add("Origin", c.OriginHost)
	headers.Add("X-BCA-Key", c.APIKey)

	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z07:00")
	headers.Add("X-BCA-Timestamp", timestamp)

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(body)
	//signature := generateSignature(c.APISecret, method, path, accessToken, buf.String(), timestamp)
	signature := generateSignature(c.APISecret, method, path, accessToken, string(body), timestamp)
	headers.Add("X-BCA-Signature", signature)

	// Add additional headers as by FundTransferDomestic
	for key, val := range additionalHeader {
		headers.Add(key, val)
	}

	//return c.CallRaw(method, path, "application/json", headers, body, v)
	return c.CallRaw(method, path, "application/json", headers, bytes.NewBuffer(body), v)

}

//CallRaw is the implementation for invoking API without any wrapper
func (c *APIImplementation) CallRaw(method, path, contentType string, headers http.Header, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, contentType, headers, body)

	if err != nil {
		return err
	}

	return c.Do(req, v)
}

//NewRequest is used to create new HTTP request of BCA API
func (c *APIImplementation) NewRequest(method, path, contentType string, headers http.Header, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = c.URL + path

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		if c.LogLevel > 0 {
			c.Logger.Printf("Cannot create Stripe request: %v\n", err)
		}
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if headers != nil {
		for k, v := range headers {
			for _, line := range v {
				req.Header.Add(k, line)
			}
		}
	}

	return req, nil
}

//Do is used by Call to execute BCA HTTP request and parse the response
func (c *APIImplementation) Do(req *http.Request, v interface{}) error {
	logLevel := c.LogLevel
	logger := c.Logger

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return err
	}

	if logLevel > 2 {
		logger.Println("BCA response: ", string(resBody))
	}

	if v != nil {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}
	}

	return nil
}

func canonicalize(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func sortQueryParam(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", errors.Trace(err)
	}

	u.RawQuery = u.Query().Encode()
	return u.String(), nil
}

func generateSignature(apiSecret, method, path, accessToken, requestBody, timestamp string) string {
	canonicalReqBody := canonicalize(requestBody)
	h := sha256.New()
	//h.Write([]byte(requestBody))
	if _, err := h.Write([]byte(canonicalReqBody)); err != nil {
		//return "", "", errors.Trace(err)
		panic(err)
	}
	sortedURL, err := sortQueryParam(path)
	if err != nil {
		//return "", "", errors.Trace(err)
		panic(err)
	}
	//strToSign := method + ":" + path + ":" + accessToken + ":" + hex.EncodeToString(h.Sum(nil)) + ":" + timestamp

	strToSign := method + ":" +
		sortedURL + ":" +
		accessToken + ":" +
		strings.ToLower(hex.EncodeToString(h.Sum(nil))) + ":" +
		timestamp

	mac := hmac.New(sha256.New, []byte(apiSecret))
	//mac.Write([]byte(strToSign))
	if _, err = mac.Write([]byte(strToSign)); err != nil {
		//return "", strToSign, errors.Trace(err)
		panic(err)
	}
	return hex.EncodeToString(mac.Sum(nil))
}
