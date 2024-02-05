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

// ClipboardAction
// ClipboardAction
// https://developers.line.biz/en/reference/messaging-api/#clipboard-action
type ClipboardAction struct {
	Action

	/**
	 * Label for the action.
	 */
	Label string `json:"label,omitempty"`

	/**
	 * Text that is copied to the clipboard. Max character limit: 1000  (Required)
	 */
	ClipboardText string `json:"clipboardText"`
}

// MarshalJSON customizes the JSON serialization of the ClipboardAction struct.
func (r *ClipboardAction) MarshalJSON() ([]byte, error) {

	type Alias ClipboardAction
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "clipboard",
	})
}
