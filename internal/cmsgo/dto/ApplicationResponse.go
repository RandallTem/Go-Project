package dto

import "time"

type ApplicationResponse struct {
	UUID                string    `json:"uuid,omitempty"`
	Numauto             string    `json:"numauto,omitempty"`
	ApplicationDateTime time.Time `json:"applicationDateTime,omitempty"`
	IsSuccess           bool      `json:"isSuccess,omitempty"`
	ErrorMessage        string    `json:"errorMessage,omitempty"`
}
