// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CloudEvent struct {
	ID         string                 `json:"id"`
	Source     string                 `json:"source"`
	Type       string                 `json:"type"`
	Subject    *string                `json:"subject"`
	Attributes map[string]interface{} `json:"attributes"`
	Data       map[string]interface{} `json:"data"`
}

type CloudEventInput struct {
	Source     string                 `json:"source"`
	Type       string                 `json:"type"`
	Subject    *string                `json:"subject"`
	Attributes map[string]interface{} `json:"attributes"`
	Data       map[string]interface{} `json:"data"`
}

type ReceiveRequest struct {
	Type    string  `json:"type"`
	Subject *string `json:"subject"`
}