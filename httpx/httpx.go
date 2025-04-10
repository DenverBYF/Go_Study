package httpx

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Ret struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// httGet http get
func HttpGet(ctx context.Context, hc http.Client, url string, cookies []*http.Cookie) (*Ret, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest failed")
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}

	return httpCombine(ctx, hc, req)
}

// httpGetWithHeader ...
func HttpGetWithHeader(ctx context.Context, hc http.Client, url string, header http.Header) (*Ret, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest failed")
	}
	req.Header = header
	return httpCombine(ctx, hc, req)
}

// httpPost ...
func HttpPost(ctx context.Context, hc http.Client, url string, arg string, header map[string]string) (*Ret, error) {
	payload := strings.NewReader(arg)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest failed")
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	return httpCombine(ctx, hc, req)
}

// httpCombine ...
func httpCombine(ctx context.Context, hc http.Client, req *http.Request) (*Ret, error) {
	res, err := hc.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient Do failed")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read http response body failed")
	}

	return &Ret{
		Body:       body,
		StatusCode: res.StatusCode,
		Headers:    res.Header.Clone(),
	}, nil
}
