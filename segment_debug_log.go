/*
 * Flagr
 *
 * Flagr is a feature flagging, A/B testing and dynamic configuration microservice
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goflagr

type SegmentDebugLog struct {

	SegmentID int64 `json:"segmentID,omitempty"`

	Msg string `json:"msg,omitempty"`
}
