package uikit

import (
	"encoding/json"
	"strings"
)

// JSONString encodes v and return v as string.
// If error return empty string and an error.
func JSONString(v map[string]interface{}) (string, error) {

	if v == nil {
		return "", nil
	}

	sb := &strings.Builder{}
	err := json.NewEncoder(sb).Encode(v)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
