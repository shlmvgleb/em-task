package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RequestWithJSON[T any, R any](ctx context.Context, client *http.Client, url string, body T, headers map[string]string) (R, error) { //nolint:ireturn,lll
	var result R

	bodyData, err := json.Marshal(body)
	if err != nil {
		return result, ErrInvalidRequestBody
	}

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyData))

	req.Header.Set("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("cannot send a request: %w", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return result, fmt.Errorf("cannot read a response from service: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return result, fmt.Errorf("%w: %s", ErrInvalidResponse, resBody)
	}

	if err := json.Unmarshal(resBody, &result); err != nil {
		return result, fmt.Errorf("cannot parse a response from service: %w", err)
	}

	return result, nil
}
