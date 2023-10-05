/*
 * Flagr
 *
 * Flagr is a feature flagging, A/B testing and dynamic configuration microservice. The base path for all the APIs is \"/api/v1\".
 *
 * API version: 1.1.16
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goflagr

type CreateFlagRequest struct {
	Description string `json:"description"`
	// unique key representation of the flag
	Key string `json:"key,omitempty"`
	// template for flag creation
	Template string `json:"template,omitempty"`
}
