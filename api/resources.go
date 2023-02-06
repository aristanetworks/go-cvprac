//
// Copyright (c) 2016-2023, Arista Networks, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright notice,
//   this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//
//   * Neither the name of Arista Networks nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL ARISTA NETWORKS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
// BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
// OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN
// IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package cvpapi

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Result type for a `Get<x>Rsc` or `Get<x>sRsc` request.
type resultRsc struct {
	// Response body.
	Result *valueRsc `json:"result"`
	// The time that this resource was last modified.
	Time *time.Time `json:"time,omitempty"`

	resourceGetError
}

// The response body for a `Get<x>Rsc` or `Get<x>sRsc` request.
type valueRsc struct {
	// Generic resource type.
	Value json.RawMessage `json:"value"`
	// Optional time at which this resource was last modified.
	Time time.Time `json:"time"`
	// Optional type data included with some resources (e.g.,
	// change control approvals).
	Type string `json:"type"`

	resourceGetError
}

// Failure state associated with resource API requests.
type resourceGetError struct {
	// Error code supplied if a request fails.
	Code *int `json:"code,omitempty"`
	// Optional explanation for a request failure.
	Message *string `json:"message,omitempty"`
}

// Errors associated with the state of an item from the resource API.
type ResourceError struct {
	// An error attached to this item.
	//
	// The presence of this string does not necessarily mean that *this*
	// request failed, e.g., a scheduled execution might expire before its change
	// is approved.
	Error *string `json:"error,omitempty"`
}

// Wrapper type for a list of strings in Resource API requests
// and responses.
type StringList struct {
	Values []string `json:"values"`
}

// Wrapper type for a 2D string array in Resource API requests
// and responses.
type StringMatrix struct {
	Values []StringList `json:"values"`
}

// Wrapper type for a string map in Resource API requests
// and responses.
type StringMap struct {
	Values map[string]string `json:"values"`
}

// A boolean flag.
type FlagRsc struct {
	Value bool `json:"value"`

	ModifiedBy
}

// A timestamp flag.
type TimestampFlagRsc struct {
	Value *time.Time `json:"value,omitempty"`

	ModifiedBy
}

// Modification data for fields of a Resource API value.
type ModifiedBy struct {
	// The time this field was updated at.
	Time *time.Time `json:"time,omitempty"`
	// The user who last updated this field.
	User *string `json:"user,omitempty"`
	// A comment explaining the last change made to this field.
	Notes *string `json:"notes,omitempty"`
}

func (v *resourceGetError) GetError() error {
	if v.Code != nil {
		err := errors.Errorf("resource API lookup failed [code %d]", v.Code)
		if v.Message != nil {
			err = errors.Errorf("%s: %s", err, *v.Message)
		}
		return err
	}

	return nil
}

// Resource APIs hand us back newline-delimited Json on /all endpoints.
// This converts a call into a list of usable Json objects.
func resultList(data []byte) ([]resultRsc, error) {
	out := []resultRsc{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	for decoder.More() {
		el := resultRsc{}

		if err := decoder.Decode(&el); err != nil {
			return nil, err
		}

		if err := el.GetError(); err != nil {
			return nil, err
		}

		if el.Result == nil {
			return nil, errors.Errorf("missing 'Value' field from ndJson response")
		}

		out = append(out, el)
	}

	return out, nil
}
