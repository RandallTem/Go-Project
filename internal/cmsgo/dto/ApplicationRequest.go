package dto

import "time"

type ApplicationRequest struct {
	UUID                string    `json:"uuid,omitempty"`
	Numauto             string    `json:"numauto,omitempty"`
	ApplicationDateTime time.Time `json:"applicationDateTime,omitempty"`
}
