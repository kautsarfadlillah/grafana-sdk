package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type QueryResponse struct {
	Status    string `json:"status"`
	IsPartial bool   `json:"isPartial"`
	Data      Data   `json:"data"`
	Stats     Stats  `json:"stats"`
}

type Stats struct {
	SeriesFetched string `json:"seriesFetched"`
}

type Data struct {
	Result []Result `json:"result"`
}

type Result struct {
	Metric json.RawMessage `json:"metric"`
	Values []Value         `json:"values"`
}

type Value []interface{}

func (r *Client) Query(ctx context.Context, source int, params url.Values) (QueryResponse, error) {
	var (
		raw  []byte
		code int
		resp QueryResponse
		err  error
	)

	if raw, code, err = r.get(ctx, fmt.Sprintf("api/datasources/proxy/%d/api/v1/query", source), params); err != nil {
		return QueryResponse{}, err
	}

	if code != 200 {
		return QueryResponse{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return QueryResponse{}, err
	}

	return resp, nil
}
