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

type resultRsc struct {
	Result ValueRsc   `json:"result"`
	Time   *time.Time `json:"time,omitempty"`
}

type ValueRsc struct {
	Value json.RawMessage `json:"value"`
	Time  time.Time       `json:"time"`
	Type  string          `json:"type"`
}

type ChangeControlKeyRsc struct {
	Id string `json:"id"`
}

type StringMap struct {
	Values map[string]string `json:"values"`
}

type StringList struct {
	Values []string `json:"values"`
}

type StringMatrix struct {
	Values []StringList `json:"values"`
}

type ActionRsc struct {
	Name    string    `json:"name"`
	Timeout uint32    `json:"timeout"`
	Args    StringMap `json:"args"`
}

type StageRsc struct {
	Name   string        `json:"name"`
	Action *ActionRsc    `json:"action,omitempty"`
	Rows   *StringMatrix `json:"rows,omitempty"`
	Status *string       `json:"status,omitempty"`
	Error  *string       `json:"error,omitempty"`
}

type ChangeRsc struct {
	Name        string `json:"name"`
	RootStageId string `json:"rootStageId"`
	Stages      struct {
		Values map[string]StageRsc `json:"values"`
	} `json:"stages"`
	Notes *string    `json:"notes,omitempty"`
	User  *string    `json:"user,omitempty"`
	Time  *time.Time `json:"time,omitempty"`
}

type FlagRsc struct {
	Value bool      `json:"value"`
	Notes string    `json:"notes"`
	Time  time.Time `json:"time"`
	User  string    `json:"user"`
}

type FlagConfigRsc struct {
	Value bool   `json:"value"`
	Notes string `json:"notes"`
}

type TimestampFlagRsc struct {
	Value *time.Time `json:"value,omitempty"`
	Notes string     `json:"notes"`
	Time  *time.Time `json:"time,omitempty"`
	User  *string    `json:"user,omitempty"`
}

type TimestampFlagConfigRsc struct {
	Value bool   `json:"value"`
	Notes string `json:"notes"`
}

type ChangeControlRsc struct {
	Key      ChangeControlKeyRsc `json:"key"`
	Change   *ChangeRsc          `json:"change,omitempty"`
	Flag     *FlagRsc            `json:"flag,omitempty"`
	Start    *FlagRsc            `json:"start,omitempty"`
	Status   *string             `json:"status,omitempty"`
	Schedule *TimestampFlagRsc   `json:"schedule,omitempty"`

	Error *string `json:"error,omitempty"`
}

// Resource APIs hand us back ndJson on /all endpoints.
// This converts a call into a list of usable Json objects.
func resultList(data []byte) ([]resultRsc, error) {
	out := []resultRsc{}

	decoder := json.NewDecoder(bytes.NewReader(data))

	for decoder.More() {
		el := resultRsc{}

		if err := decoder.Decode(&el); err != nil {
			return nil, err
		}

		out = append(out, el)
	}

	return out, nil
}

// GetChangeControlsRsc returns a list of `ChangeControlRsc`s via the
// ChangeControl resource API available.
//
// This endpoint is available on CVP 2021.2.0 or newer.
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

// GetChangeControlRsc returns a single `ChangeControlRsc` matching the given
// `key` via the ChangeControl resource API available.
//
// This endpoint is available on CVP 2021.2.0 or newer.
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

// CreateChangeControlRsc creates a new `ChangeControlRsc` with the given uuid `key`,
// `name`, and task list (in serial or parallel).
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) CreateChangeControlRsc(key, name string, tasks []string, sequential bool) error {
	change := ChangeRsc{
		Name:        name,
		RootStageId: "root",
	}
	change.Stages.Values = make(map[string]StageRsc)

	cfg := ChangeControlRsc{Change: &change}
	cfg.Key.Id = key

	rootStage := StageRsc{
		Name: "root",
	}

	var rows []StringList

	for i, task := range tasks {
		localName := fmt.Sprintf("stage%d", i)
		localAction := ActionRsc{Name: "task"}
		localAction.Args.Values = make(map[string]string)
		localAction.Args.Values["TaskID"] = task
		localStage := StageRsc{
			Name:   localName,
			Action: &localAction,
		}

		change.Stages.Values[localName] = localStage

		if sequential || i == 0 {
			el := StringList{Values: []string{localName}}
			rows = append(rows, el)
			// change.Stages.Values["root"].Rows.Values
		} else {
			rows[0].Values = append(rows[0].Values, localName)
		}
	}

	matrix := StringMatrix{Values: rows}
	rootStage.Rows = &matrix

	change.Stages.Values["root"] = rootStage

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("CreateChangeControlRsc: %s", err)
	}

	return err
}

// ScheduleChangeControlRsc schedules the Change Control given by `key` to occur
// at `sched_time`, with optional `notes`.
//
// This endpoint is available on CVP 2022.1.0 or newer.
func (c CvpRestAPI) ScheduleChangeControlRsc(key string, sched_time time.Time, notes string) error {
	cfg := ChangeControlRsc{}
	cfg.Key.Id = key

	// Note: The API seems to disallow setting the user or time on cvprac v2022.1.0 -- not tested others.
	sched := TimestampFlagRsc{}
	sched.Value = &sched_time
	sched.Notes = notes

	cfg.Schedule = &sched

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("ScheduleChangeControlRsc: %s", err)
	}

	// FIXME: check error handling of the top-level struct received.
	// Scheduling unapproved jobs seems to have no issue, scheduling unapproved
	// causes the change control to become unapproved again, counter to what docs say...

	return err
}

// DescheduleChangeControlRsc removes any schedule data from a Change Control given by `key`.
//
// This endpoint is available on CVP 2022.1.0 or newer.
func (c CvpRestAPI) DescheduleChangeControlRsc(key string, notes string) error {
	cfg := ChangeControlRsc{}
	cfg.Key.Id = key

	// Note: The API seems to disallow setting the user or time on cvprac v2022.1.0 -- not tested others.
	sched := TimestampFlagRsc{}
	sched.Notes = notes

	cfg.Schedule = &sched

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("DescheduleChangeControlRsc: %s", err)
	}

	return err
}
