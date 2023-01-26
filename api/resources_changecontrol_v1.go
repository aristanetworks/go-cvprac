//
// Copyright (c) 2016-2017, Arista Networks, Inc. All rights reserved.
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
	"fmt"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type ResultRsc struct {
	Result ValueRsc
}

type ValueRsc struct {
	Value json.RawMessage
	Time  time.Time
	Type  string
}

type ChangeControlKeyRsc struct {
	Id string
}

type StringMap struct {
	Values map[string]string
}

type StringList struct {
	Values []string
}

type ActionRsc struct {
	Name    string
	Timeout uint32
	Args    StringMap
}

type StageRsc struct {
	Name   string
	Action *ActionRsc
	Rows   struct {
		Values []StringList
	}
	Status *string
	Error  *string
}

type ChangeRsc struct {
	Name        string
	RootStageId string
	Stages      struct {
		Values map[string]StageRsc
	}
	Notes string
	User  string
	Time  time.Time
}

type FlagRsc struct {
	Value bool
	Notes string
	Time  time.Time
	User  string
}

type FlagConfigRsc struct {
	Value bool
	Notes string
}

type TimestampFlagRsc struct {
	Value time.Time
	Notes string
	Time  time.Time
	User  string
}

type TimestampFlagConfigRsc struct {
	Value bool
	Notes string
}

type ChangeControlRsc struct {
	Key      ChangeControlKeyRsc
	Change   ChangeRsc
	Flag     *FlagRsc
	Start    *FlagRsc
	Status   *string
	Schedule *TimestampFlagRsc

	Error *string
}

func resultList(data []byte) ([]ResultRsc, error) {
	out := []ResultRsc{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	for decoder.More() {
		el := ResultRsc{}

		if err := decoder.Decode(&el); err != nil {
			return nil, err
		}

		out = append(out, el)
	}

	return out, nil
}

// GetChangeControlsRsc returns a list of `ChangeControlRsc`s via the
// ChangeControl resource API availablre on CVP 2021.2.0 or newer.
func (c CvpRestAPI) GetChangeControlsRsc() ([]ChangeControlRsc, error) {
	ccs := []ChangeControlRsc{}
	resp, err := c.client.Get("/api/resources/changecontrol/v1/ChangeControl/all", nil)

	if err != nil {
		return nil, errors.Errorf("GetChangeControlsRsc: %s",
			err)
	}

	results, err := resultList(resp)

	if err != nil {
		return nil, errors.Errorf("GetChangeControlsRsc: %s",
			err)
	}

	for _, result := range results {
		var cc ChangeControlRsc
		if err = json.Unmarshal(result.Result.Value, &cc); err != nil {
			return nil, errors.Errorf("GetChangeControlsRsc: %s Payload:\n%s",
				err, resp)
		}
		ccs = append(ccs, cc)
	}

	return ccs, nil
}

func (c CvpRestAPI) GetChangeControlRsc(key string) (ChangeControlRsc, error) {
	result := ValueRsc{}
	cc := ChangeControlRsc{}

	query := &url.Values{
		"key.id": {key},
	}

	resp, err := c.client.Get("/api/resources/changecontrol/v1/ChangeControl", query)

	if err != nil {
		return cc, errors.Errorf("GetChangeControlRsc: %s",
			err)
	}

	if err = json.Unmarshal(resp, &result); err != nil {
		return cc, errors.Errorf("GetChangeControlRsc [Result]: %s Payload:\n%s",
			err, resp)
	}

	if err = json.Unmarshal(result.Value, &cc); err != nil {
		return cc, errors.Errorf("GetChangeControlRsc [Inner]: %s Payload:\n%s",
			err, resp)
	}

	return cc, nil
}

func (c CvpRestAPI) ScheduleChangeControlRsc(key string, sched_time time.Time, notes string) error {
	data := map[string]interface{}{
		"key":      map[string]interface{}{"id": key},
		"schedule": map[string]interface{}{"value": sched_time, "notes": notes},
		// "time":     time.Now(),
	}

	js, err := json.Marshal(data)

	// if err != nil {
	// 	err = errors.Errorf("ScheduleChangeControlRsc: %s", err)
	// }

	resp, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, data)

	if err != nil {
		err = errors.Errorf("ScheduleChangeControlRsc: %s", err)
	}

	fmt.Printf("Sched: got back %v from %v\n", string(resp), string(js))

	return err
}
