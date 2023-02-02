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

// FIXME: document structs and factor out common properties.
// FIXME: break out shared utilities (resultRsc) to "resources.go"?

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

type StageMap struct {
	Values map[string]StageRsc `json:"values"`
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
	Name        *string    `json:"name,omitempty"`
	RootStageId *string    `json:"rootStageId,omitempty"`
	Stages      *StageMap  `json:"stages,omitempty"`
	Notes       *string    `json:"notes,omitempty"`
	User        *string    `json:"user,omitempty"`
	Time        *time.Time `json:"time,omitempty"`
}

type FlagRsc struct {
	Value bool       `json:"value"`
	Notes string     `json:"notes"`
	Time  *time.Time `json:"time,omitempty"`
	User  string     `json:"user,omitempty"`
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

type ApproveConfigRsc struct {
	Key     ChangeControlKeyRsc `json:"key"`
	Approve FlagConfigRsc       `json:"approve"`
	Version *time.Time          `json:"version,omitempty"`
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

func deserChangeControl(data []byte) (*ChangeControlRsc, error) {
	result := ValueRsc{}
	cc := ChangeControlRsc{}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Errorf("ChangeControlRsc [Result]: %s Payload:\n%s",
			err, data)
	}

	if err := json.Unmarshal(result.Value, &cc); err != nil {
		return nil, errors.Errorf("ChangeControlRsc [Inner]: %s Payload:\n%s",
			err, data)
	}

	return &cc, nil
}

func deserApproval(data []byte) (*ApproveConfigRsc, error) {
	result := ValueRsc{}
	cc := ApproveConfigRsc{}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Errorf("ApproveConfigRsc [Result]: %s Payload:\n%s",
			err, data)
	}

	if err := json.Unmarshal(result.Value, &cc); err != nil {
		return nil, errors.Errorf("ApproveConfigRsc [Inner]: %s Payload:\n%s",
			err, data)
	}

	return &cc, nil
}

// GetChangeControlsRsc returns a list of `ChangeControlRsc`s via the
// ChangeControl resource API.
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
// `ccId` via the ChangeControl resource API.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) GetChangeControlRsc(ccId string) (*ChangeControlRsc, error) {
	query := &url.Values{
		"key.id": {ccId},
	}

	resp, err := c.client.Get("/api/resources/changecontrol/v1/ChangeControl", query)

	if err != nil {
		return nil, errors.Errorf("GetChangeControlRsc: %s",
			err)
	}

	return deserChangeControl(resp)
}

// DeleteChangeControlRsc removes a single `ChangeControlRsc` matching the given
// `ccId` via the ChangeControl resource API.
//
// This method will fail if a ChangeControl has finished running.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) DeleteChangeControlRsc(ccId string) error {
	query := &url.Values{
		"key.id": {ccId},
	}

	_, err := c.client.Delete("/api/resources/changecontrol/v1/ChangeControlConfig", query, nil)

	if err != nil {
		return errors.Errorf("DeleteChangeControlRsc: %s",
			err)
	}

	return nil
}

// CreateChangeControlRsc creates a new `ChangeControlRsc` with the given uuid `ccId`,
// `name`, and task list (in serial or parallel).
//
// The user must ensure that any referenced tasks exist beforehand.
//
// This endpoint is available on CVP 2021.1.0 or newer.
func (c CvpRestAPI) CreateChangeControlRsc(ccId, name string, tasks []string, sequential bool) (*ChangeControlRsc, error) {
	var rootStages [][]StageRsc

	for i, task := range tasks {
		stageName := fmt.Sprintf("stage%d", i)
		action := ActionRsc{Name: "task"}
		action.Args.Values = make(map[string]string)
		action.Args.Values["TaskID"] = task
		stage := StageRsc{
			Name:   stageName,
			Action: &action,
		}

		if sequential || i == 0 {
			el := []StageRsc{stage}
			rootStages = append(rootStages, el)
			// change.Stages.Values["root"].Rows.Values
		} else {
			rootStages[0] = append(rootStages[0], stage)
		}
	}

	return c.CreateChangeControlWithActionsRsc(ccId, name, rootStages)
}

// CreateChangeControlWithActionsRsc creates a new `ChangeControlRsc` with the given uuid `ccId`,
// `name`, executing the given `rootStages` attached to a new root node. `extraStages` can be optionally
// included as needed.
//
// `rootStages` consists of rows of 'StageRsc's which are visited serially. Elements in each row are
// run in parallel.
//
// This endpoint is available on CVP 2021.1.0 or newer.
func (c CvpRestAPI) CreateChangeControlWithActionsRsc(ccId, name string, rootStages [][]StageRsc, extraStages ...StageRsc) (*ChangeControlRsc, error) {
	root := "root"
	stages := StageMap{Values: make(map[string]StageRsc)}
	change := ChangeRsc{
		Name:        &name,
		RootStageId: &root,
		Stages:      &stages,
	}

	cfg := ChangeControlRsc{Change: &change}
	cfg.Key.Id = ccId

	rootStage := StageRsc{
		Name: "root",
	}

	var rows []StringList

	for _, parallelStages := range rootStages {
		var rowToAdd []string
		for _, stage := range parallelStages {
			rowToAdd = append(rowToAdd, stage.Name)
			change.Stages.Values[stage.Name] = stage
		}
		rows = append(rows, StringList{Values: rowToAdd})
	}

	for _, stage := range extraStages {
		change.Stages.Values[stage.Name] = stage
	}

	matrix := StringMatrix{Values: rows}
	rootStage.Rows = &matrix

	change.Stages.Values["root"] = rootStage

	resp, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("CreateChangeControlRsc: %s", err)
		return nil, err
	}

	return deserChangeControl(resp)
}

// AddNoteToChangeControlRsc adds a note to the `ChangeControlRsc` matching
// `ccId` via the ChangeControl resource API.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) AddNoteToChangeControlRsc(ccId, notes string) error {
	change := ChangeRsc{
		Notes: &notes,
	}

	cfg := ChangeControlRsc{Change: &change}
	cfg.Key.Id = ccId

	rsp, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("AddNoteToChangeControlRsc: %s", err)
	}

	fmt.Printf("R: %+v\n", string(rsp))

	return err
}

// StartChangeControlRsc starts the the `ChangeControlRsc` matching the uuid
// `ccId` via the ChangeControl resource API, with an expository `note`.
//
// Will return an error if the matching change control has not yet been approved.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) StartChangeControlRsc(ccId, notes string) error {
	start := FlagRsc{
		Notes: notes,
		Value: true,
	}

	cfg := ChangeControlRsc{Start: &start}
	cfg.Key.Id = ccId

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("StartChangeControlRsc: %s", err)
	}

	return err
}

// StopChangeControlRsc stops the the `ChangeControlRsc` matching the uuid
// `ccId` via the ChangeControl resource API, with an expository `note`.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) StopChangeControlRsc(ccId, notes string) error {
	start := FlagRsc{
		Notes: notes,
		Value: false,
	}

	cfg := ChangeControlRsc{Start: &start}
	cfg.Key.Id = ccId

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("StopChangeControlRsc: %s", err)
	}

	return err
}

// ScheduleChangeControlRsc schedules the Change Control given by `ccId` to occur
// at `schedTime`, with optional `notes`.
//
// This endpoint is available on CVP 2022.1.0 or newer.
func (c CvpRestAPI) ScheduleChangeControlRsc(ccId string, schedTime time.Time, notes string) error {
	cfg := ChangeControlRsc{}
	cfg.Key.Id = ccId

	// Note: The API seems to disallow setting the user or time on cvprac v2022.1.0 -- not tested others.
	sched := TimestampFlagRsc{}
	sched.Value = &schedTime
	sched.Notes = notes

	cfg.Schedule = &sched

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("ScheduleChangeControlRsc: %s", err)
	}

	// FIXME: check error handling of the top-level struct received.
	// Scheduling unapproved jobs seems to have no issue, scheduling unapproved
	// causes the change control to become unapproved again, counter to what docs say...
	//
	// Error string can appear at several levels in the hierarchy

	return err
}

// DescheduleChangeControlRsc removes any schedule data from a Change Control given by `ccId`.
//
// This endpoint is available on CVP 2022.1.0 or newer.
func (c CvpRestAPI) DescheduleChangeControlRsc(ccId, notes string) error {
	cfg := ChangeControlRsc{}
	cfg.Key.Id = ccId

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

// GetChangeControlApprovalsRsc returns a list of `ApproveConfigRsc`s via the
// ChangeControl resource API.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) GetChangeControlApprovalsRsc() ([]ApproveConfigRsc, error) {
	ccs := []ApproveConfigRsc{}
	resp, err := c.client.Get("/api/resources/changecontrol/v1/ApproveConfig/all", nil)

	if err != nil {
		return nil, errors.Errorf("GetChangeControlApprovalsRsc: %s",
			err)
	}

	results, err := resultList(resp)

	if err != nil {
		return nil, errors.Errorf("GetChangeControlApprovalsRsc: %s",
			err)
	}

	for _, result := range results {
		var cc ApproveConfigRsc
		if err = json.Unmarshal(result.Result.Value, &cc); err != nil {
			return nil, errors.Errorf("GetChangeControlApprovalsRsc: %s Payload:\n%s",
				err, resp)
		}
		ccs = append(ccs, cc)
	}

	return ccs, nil
}

// GetChangeControlRsc returns a single `ChangeControlRsc` matching the given
// `ccId` via the ChangeControl resource API.
//
// This endpoint is available on CVP 2021.2.0 or newer.
func (c CvpRestAPI) GetChangeControlApprovalRsc(ccId string) (*ApproveConfigRsc, error) {
	query := &url.Values{
		"key.id": {ccId},
	}

	resp, err := c.client.Get("/api/resources/changecontrol/v1/ApproveConfig", query)

	if err != nil {
		return nil, errors.Errorf("GetChangeControlApprovalRsc: %s",
			err)
	}

	return deserApproval(resp)
}

// ApproveChangeControlRsc sets the approval state for a single `ChangeControlRsc` matching the given
// `ccId` via the ChangeControl resource API.
//
// This requires a `version` matching the last modified timestamp of the intended change control
// (i.e., `<ChangeControlRsc>.Change.Time`). This call will error if another modification has
// taken place, causing a `version` mismatch.
//
// This endpoint is available on CVP 2021.2.0 or newer.
// explain version.
func (c CvpRestAPI) ApproveChangeControlRsc(ccId string, approve bool, version time.Time, notes string) error {
	flag := FlagConfigRsc{
		Value: approve,
		Notes: notes,
	}
	approval := ApproveConfigRsc{
		Approve: flag,
		Version: &version,
	}
	approval.Key.Id = ccId

	_, err := c.client.Post("/api/resources/changecontrol/v1/ApproveConfig", nil, approval)

	if err != nil {
		err = errors.Errorf("ApproveChangeControlRsc: %s", err)
	}

	return err
}
