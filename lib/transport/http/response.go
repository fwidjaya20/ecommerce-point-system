package http

import (
	"context"
)

type ResponseStructure struct {
	Data     interface{}            `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}

func Response(ctx context.Context, data interface{}, metadata map[string]interface{}) interface{} {
	meta := make(map[string]interface{})

	for k, v := range metadata {
		meta[k] = v
	}

	return ResponseStructure{
		Data:     data,
		Metadata: meta,
	}
}