package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"telegram_bot/clients/lib/e"
)

const (
	getUpdatesMethod = "getUpdates"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(basePath string) string {
	return "bot" + basePath
}

func (c *Client) Updates(offset, limit int) ([]Updates, error) {
	query := url.Values{}
	query.Set("offset", strconv.Itoa(offset))
	query.Set("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}
	var res UpdateResp

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	const errMsg = "can't make request"
	defer func() { err = e.WrapErrorf("can't make request", err) }()
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func (c *Client) SendMessage(message string) {}
