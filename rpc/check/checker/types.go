// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package checker

import "errors"

var errJsonConvert = errors.New("json convert error")

type (
	CheckRequest struct {
		Book string `json:"book,omitempty"`
	}

	CheckResponse struct {
		Found bool  `json:"found,omitempty"`
		Price int64 `json:"price,omitempty"`
	}
)