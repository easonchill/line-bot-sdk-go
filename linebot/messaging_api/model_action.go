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
	"fmt"
)

type ActionInterface interface {
	GetType() string
}

func (e Action) GetType() string {
	return e.Type
}

type UnknownAction struct {
	ActionInterface
	Type string
	Raw  map[string]json.RawMessage
}

func (e UnknownAction) GetType() string {
	return e.Type
}

func setDiscriminatorPropertyAction(r ActionInterface) ActionInterface {
	switch v := r.(type) {
	case *CameraAction:
		if v.Type == "" {
			v.Type = "camera"
		}
		return v
	case CameraAction:
		if v.Type == "" {
			v.Type = "camera"
		}
		return v
	case *CameraRollAction:
		if v.Type == "" {
			v.Type = "cameraRoll"
		}
		return v
	case CameraRollAction:
		if v.Type == "" {
			v.Type = "cameraRoll"
		}
		return v
	case *DatetimePickerAction:
		if v.Type == "" {
			v.Type = "datetimepicker"
		}
		return v
	case DatetimePickerAction:
		if v.Type == "" {
			v.Type = "datetimepicker"
		}
		return v
	case *LocationAction:
		if v.Type == "" {
			v.Type = "location"
		}
		return v
	case LocationAction:
		if v.Type == "" {
			v.Type = "location"
		}
		return v
	case *MessageAction:
		if v.Type == "" {
			v.Type = "message"
		}
		return v
	case MessageAction:
		if v.Type == "" {
			v.Type = "message"
		}
		return v
	case *PostbackAction:
		if v.Type == "" {
			v.Type = "postback"
		}
		return v
	case PostbackAction:
		if v.Type == "" {
			v.Type = "postback"
		}
		return v
	case *RichMenuSwitchAction:
		if v.Type == "" {
			v.Type = "richmenuswitch"
		}
		return v
	case RichMenuSwitchAction:
		if v.Type == "" {
			v.Type = "richmenuswitch"
		}
		return v
	case *UriAction:
		if v.Type == "" {
			v.Type = "uri"
		}
		return v
	case UriAction:
		if v.Type == "" {
			v.Type = "uri"
		}
		return v

	default:
		return v
	}
}

// Action
// Action

// https://developers.line.biz/en/reference/messaging-api/#action-objects

type Action struct {
	// Type of action

	Type string `json:"type,omitempty"`
	// Label for the action.

	Label string `json:"label,omitempty"`
}

func UnmarshalAction(data []byte) (ActionInterface, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalAction: %w", err)
	}

	var discriminator string
	err = json.Unmarshal(raw["type"], &discriminator)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalAction: Cannot read type: %w", err)
	}

	switch discriminator {
	case "camera":
		var camera CameraAction
		if err := json.Unmarshal(data, &camera); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read camera: %w", err)
		}
		return camera, nil
	case "cameraRoll":
		var cameraRoll CameraRollAction
		if err := json.Unmarshal(data, &cameraRoll); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read cameraRoll: %w", err)
		}
		return cameraRoll, nil
	case "datetimepicker":
		var datetimepicker DatetimePickerAction
		if err := json.Unmarshal(data, &datetimepicker); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read datetimepicker: %w", err)
		}
		return datetimepicker, nil
	case "location":
		var location LocationAction
		if err := json.Unmarshal(data, &location); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read location: %w", err)
		}
		return location, nil
	case "message":
		var message MessageAction
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read message: %w", err)
		}
		return message, nil
	case "postback":
		var postback PostbackAction
		if err := json.Unmarshal(data, &postback); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read postback: %w", err)
		}
		return postback, nil
	case "richmenuswitch":
		var richmenuswitch RichMenuSwitchAction
		if err := json.Unmarshal(data, &richmenuswitch); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read richmenuswitch: %w", err)
		}
		return richmenuswitch, nil
	case "uri":
		var uri UriAction
		if err := json.Unmarshal(data, &uri); err != nil {
			return nil, fmt.Errorf("UnmarshalAction: Cannot read uri: %w", err)
		}
		return uri, nil

	default:
		var unknown UnknownAction
		unknown.Type = discriminator
		unknown.Raw = raw
		return unknown, nil
	}
}
