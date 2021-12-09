package aocget

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://adventofcode.com/%d/day/%d/input"

type AoCClient struct {
	h  *http.Client
	sc *http.Cookie
}

//NewClient creates a new http client and saves the session cookie on a struct for reuse
func NewClient(sessionToken string) *AoCClient {
	sc := &http.Cookie{
		Name:  "session",
		Value: sessionToken,
	}
	return &AoCClient{
		h: &http.Client{
			Timeout: 10 * time.Second,
		},
		sc: sc,
	}
}

func (acc *AoCClient) doRequest(ctx context.Context, year, day int) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(baseURL, year, day), nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(acc.sc)

	resp, err := acc.h.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

//DownloadInputString downloads the personalized input for a given year and day combination
// the result is returned as a string
func (acc *AoCClient) DownloadInputString(year, day int) (string, error) {
	b, err := acc.doRequest(context.Background(), year, day)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//DownloadInputStringContext downloads the personalized input for a given year and day combination
// the result is returned as a string
func (acc *AoCClient) DownloadInputStringContext(ctx context.Context, year, day int) (string, error) {
	b, err := acc.doRequest(ctx, year, day)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//DownloadInputBytes downloads the personalized input for a given year and day combination
// the result is returned as a raw []byte for more complicated operations
func (acc *AoCClient) DownloadInputBytes(year, day int) ([]byte, error) {
	return acc.doRequest(context.Background(), year, day)
}

//DownloadInputBytesContext downloads the personalized input for a given year and day combination
// the result is returned as a raw []byte for more complicated operations
func (acc *AoCClient) DownloadInputBytesContext(ctx context.Context, year, day int) ([]byte, error) {
	return acc.doRequest(ctx, year, day)
}
