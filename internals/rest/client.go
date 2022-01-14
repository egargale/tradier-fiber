package rest

import (
	"bytes"
	"fmt"
	"time"

	"github.com/goccy/go-json"

	// "tradier-fiber/util"

	// "github.com/spf13/viper"
	// "github.com/davecgh/go-spew/spew"
	// api "github.com/myussufz/fasthttp-api"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const (
	// fasthhtp Maximum number of connections per each host which may be established.
	MaxConnsPerHost = 100
	//  fasthhtp Maximum number of attempts for idempotent calls
	MaxIdemponentCallAttempts = 10
	//
	defaultRetries = 3
	// Header indicating the number of requests remaining.
	rateLimitAvailable = "X-Ratelimit-Available"
	// Header indicating the time at which our rate limit will renew.
	rateLimitExpiry = "X-Ratelimit-Expiry"
	// Error returned by Tradier if we make too big of a request.
	ErrBodyBufferOverflow = "protocol.http.TooBigBody"
	//
)

var (
	// ErrNoAccountSelected is returned if account-specific methods
	// are attempted to be used without selecting an account first.
	ErrNoAccountSelected = errors.New("no account selected")
)

type Position struct {
	CostBasis    float64  `json:"cost_basis"`
	DateAcquired DateTime `json:"date_acquired"`
	Id           int
	Quantity     float64
	Symbol       string
}

type ClientParams struct {
	Endpoint   string
	AuthToken  string
	Client     *fasthttp.Client
	RetryLimit int
	Account    string
}

// DefaultParams returns ClientParams initialized with default values.
func DefaultParams(authToken string) ClientParams {
	return ClientParams{
		Endpoint:   APIEndpoint,
		AuthToken:  authToken,
		Client:     &fasthttp.Client{},
		RetryLimit: defaultRetries,
	}
}

// Client provides methods for making requests to the Tradier API.
type Client struct {
	client     *fasthttp.Client
	endpoint   string
	authHeader string
	retryLimit int
	account    string
}

func NewClient(params ClientParams) *Client {
	return &Client{
		client:     params.Client,
		endpoint:   params.Endpoint,
		authHeader: fmt.Sprintf("Bearer %s", params.AuthToken),
		retryLimit: params.RetryLimit,
		account:    params.Account,
	}
}

func (tc *Client) SelectAccount(account string) {
	tc.account = account
}

// Get the current state of the market (open/closed/etc.)
func (tc *Client) GetMarketState() (MarketStatus, error) {
	url := tc.endpoint + "/v1/markets/clock"
	var result struct {
		Clock MarketStatus
	}
	resp, err := tc.getTradierResponse(url)
	if err != nil {
		fmt.Println("error:", err)
	}
	err1 := json.Unmarshal(resp, &result)
	if err1 != nil {
		fmt.Println("error:", err1)
	}
	fmt.Printf("%+v", result.Clock)
	return result.Clock, err
}

// Get the current account Positions
func (tc *Client) GetAccountPositions() (interface{}, error) {
	if tc.account == "" {
		return nil, ErrNoAccountSelected
	}

	url := tc.endpoint + "/v1/accounts/" + tc.account + "/positions"
	var result struct {
		Positions struct {
			Position []*Position
		}
	}
	var result1 struct {
		Positions struct {
			Position *Position
		}
	}
	resp, err := tc.getTradierResponse(url)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	if err1 := json.Unmarshal(resp, &result); err1 != nil {
		json.Unmarshal(resp, &result1)
		return result1.Positions.Position, err
	} else {
		fmt.Printf("%+v", result.Positions.Position)
		return result.Positions.Position, err
	}
}

func (tc *Client) getTradierResponse(url string) ([]byte, error) {
	tc.client.MaxIdemponentCallAttempts = 10
	fmt.Printf("Client will connect to: %s\n", url)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.Set("Authorization", tc.authHeader)
	req.Header.Set("Accept", "application/json")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Perform the request
	//
	err := fasthttp.DoTimeout(req, resp, 3*time.Second)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return nil, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
	}

	// Verify the content type
	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		fmt.Printf("Expected content type application/json but got %s\n", contentType)
	}

	fmt.Printf("Response body is: %s", resp.Body())
	return resp.Body(), nil
}
