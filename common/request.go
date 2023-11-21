package common

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetWithHeader(ctx context.Context, url string, header map[string]string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, response); err != nil {
		return err
	}
	return nil
}

func GetWithBear(ctx context.Context, url, bear string, response interface{}) error {
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", bear)
	return GetWithHeader(ctx, url, header, response)
}
