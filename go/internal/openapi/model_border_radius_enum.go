/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// BorderRadiusEnum An enumeration.
type BorderRadiusEnum string

// List of BorderRadiusEnum
const (
	BORDERRADIUSENUM_NONE BorderRadiusEnum = "none"
	BORDERRADIUSENUM_LG BorderRadiusEnum = "lg"
	BORDERRADIUSENUM_MD BorderRadiusEnum = "md"
	BORDERRADIUSENUM_SM BorderRadiusEnum = "sm"
	BORDERRADIUSENUM_FULL BorderRadiusEnum = "full"
)

var allowedBorderRadiusEnumEnumValues = []BorderRadiusEnum{
	"none",
	"lg",
	"md",
	"sm",
	"full",
}

func (v *BorderRadiusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BorderRadiusEnum(value)
	for _, existing := range allowedBorderRadiusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BorderRadiusEnum", value)
}

// NewBorderRadiusEnumFromValue returns a pointer to a valid BorderRadiusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBorderRadiusEnumFromValue(v string) (*BorderRadiusEnum, error) {
	ev := BorderRadiusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BorderRadiusEnum: valid values are %v", v, allowedBorderRadiusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BorderRadiusEnum) IsValid() bool {
	for _, existing := range allowedBorderRadiusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BorderRadiusEnum value
func (v BorderRadiusEnum) Ptr() *BorderRadiusEnum {
	return &v
}

type NullableBorderRadiusEnum struct {
	value *BorderRadiusEnum
	isSet bool
}

func (v NullableBorderRadiusEnum) Get() *BorderRadiusEnum {
	return v.value
}

func (v *NullableBorderRadiusEnum) Set(val *BorderRadiusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBorderRadiusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBorderRadiusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBorderRadiusEnum(val *BorderRadiusEnum) *NullableBorderRadiusEnum {
	return &NullableBorderRadiusEnum{value: val, isSet: true}
}

func (v NullableBorderRadiusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBorderRadiusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
