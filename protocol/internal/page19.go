// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package internal

import (
	"encoding/json"
	"errors"
	"fmt"
)

// PageResourceType Resource type as it was perceived by the rendering engine.
type PageResourceType int

// Valid returns true if enum is set.
func (e PageResourceType) Valid() bool {
	return e >= 1 && e <= 13
}

func (e PageResourceType) String() string {
	switch e {
	case 0:
		return "PageResourceTypeNotSet"
	case 1:
		return "Document"
	case 2:
		return "Stylesheet"
	case 3:
		return "Image"
	case 4:
		return "Media"
	case 5:
		return "Font"
	case 6:
		return "Script"
	case 7:
		return "TextTrack"
	case 8:
		return "XHR"
	case 9:
		return "Fetch"
	case 10:
		return "EventSource"
	case 11:
		return "WebSocket"
	case 12:
		return "Manifest"
	case 13:
		return "Other"
	}
	return fmt.Sprintf("PageResourceType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e PageResourceType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("internal.PageResourceType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *PageResourceType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"Document\"":
		*e = 1
	case "\"Stylesheet\"":
		*e = 2
	case "\"Image\"":
		*e = 3
	case "\"Media\"":
		*e = 4
	case "\"Font\"":
		*e = 5
	case "\"Script\"":
		*e = 6
	case "\"TextTrack\"":
		*e = 7
	case "\"XHR\"":
		*e = 8
	case "\"Fetch\"":
		*e = 9
	case "\"EventSource\"":
		*e = 10
	case "\"WebSocket\"":
		*e = 11
	case "\"Manifest\"":
		*e = 12
	case "\"Other\"":
		*e = 13
	default:
		return fmt.Errorf("internal.PageResourceType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// PageFrameID Unique frame identifier.
type PageFrameID string
