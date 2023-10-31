/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

import (
	"encoding/json"
)

// MessageImagemapAction
// MessageImagemapAction

type MessageImagemapAction struct {
	ImagemapAction

	/**
	 * Get Area
	 */
	Area *ImagemapArea `json:"area"`

	/**
	 * Get Text
	 */
	Text string `json:"text"`

	/**
	 * Get Label
	 */
	Label string `json:"label,omitempty"`
}

// MarshalJSON customizes the JSON serialization of the MessageImagemapAction struct.
func (r *MessageImagemapAction) MarshalJSON() ([]byte, error) {

	type Alias MessageImagemapAction
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "message",
	})
}