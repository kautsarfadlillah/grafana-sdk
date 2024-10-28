package sdk

import (
	"context"
	"encoding/json"
	"fmt"
)

type TimeRequest struct {
	Query string `json:"query"`
	Time  int64  `json:"time"`
}

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

func (r *Client) Query(ctx context.Context, source int, request TimeRequest) (QueryResponse, error) {
	var (
		raw  []byte
		resp QueryResponse
		err  error
	)

	if raw, err = json.Marshal(request); err != nil {
		return QueryResponse{}, err
	}

	if raw, _, err = r.post(ctx, fmt.Sprintf("api/datasources/proxy/%d/api/v1/query", source), nil, nil); err != nil {
		return QueryResponse{}, err
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return QueryResponse{}, err
	}

	return resp, nil
}
