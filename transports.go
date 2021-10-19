package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(uppercaseRequest)
		if ok {
			v, err := svc.Uppercase(req.S)
			if err != nil {
				return uppercaseResponse{v, err.Error()}, nil
			}
			return uppercaseResponse{v, ""}, nil
		}
		return uppercaseResponse{"", "the argument that was supplied is not a string "}, nil
	}
}

func makeCountEndpoint(svc stringService) endpoint.Endpoint {
	return func (_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(countRequest)
		if ok {
			v := svc.Count(req.S)
			return countResponse{v}, nil
		}
		return countResponse{0}, errors.New("the argument that was supplied wass not a string")
	}
}
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil	{
		return nil, err
	}
	return request, nil
}
func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil	{
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
