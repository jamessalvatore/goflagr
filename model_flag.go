/*
 * Flagr
 *
 * Flagr is a feature flagging, A/B testing and dynamic configuration microservice. The base path for all the APIs is \"/api/v1\".
 *
 * API version: 1.1.16
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goflagr

import (
	"time"
)

type Flag struct {
	Id int64 `json:"id,omitempty"`
	// unique key representation of the flag
	Key         string    `json:"key,omitempty"`
	Description string    `json:"description"`
	Enabled     bool      `json:"enabled"`
	Tags        []Tag     `json:"tags,omitempty"`
	Segments    []Segment `json:"segments,omitempty"`
	Variants    []Variant `json:"variants,omitempty"`
	// enabled data records will get data logging in the metrics pipeline, for example, kafka.
	DataRecordsEnabled bool `json:"dataRecordsEnabled"`
	// it will override the entityType in the evaluation logs if it's not empty
	EntityType string `json:"entityType,omitempty"`
	// flag usage details in markdown format
	Notes     string    `json:"notes,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
