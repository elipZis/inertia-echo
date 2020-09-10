package model

import (
	"encoding/json"
	"strings"
)

//
type Filter struct {
	Search  string `json:"search"`
	Deleted string `json:"deleted"`
	Role    string `json:"role"`
}

//
func (this Filter) UnmarshalJSON(data []byte) error {
	type filter Filter // prevent recursion
	res := new(filter)

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	this = Filter{
		strings.ToLower(res.Search),
		res.Deleted,
		res.Role,
	}

	return nil
}
