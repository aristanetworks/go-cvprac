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
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Key type for identifying individual change controls.
type ChangeControlKeyRsc struct {
	Id string `json:"id"`
}

// Wrapper type for a collection of named stages in the change
// control Resource API.
type StageMap struct {
	Values map[string]StageRsc `json:"values"`
}

// An individual action within a change control.
type ActionRsc struct {
	// The name of the action to execute.
	//
	// Refer to CVP documentation for valid choices and their
	// corresponding arguments.
	Name string `json:"name"`
	// Timeout of this action in seconds.
	Timeout uint32 `json:"timeout"`
	// Key-value pairs of strings used as arguments for
	// this action.
	Args StringMap `json:"args"`
}

// A stage within a change control.
type StageRsc struct {
	// The name of this stage.
	//
	// This is used when this stage is referred to by either the
	// root stage ID, or as a row within another stage.
	Name string `json:"name"`
	// The action attached to this stage,
	Action *ActionRsc `json:"action,omitempty"`
	// A matrix containing the names of other sub-stages to execute.
	//
	// Entries in each row are executed in parallel, while
	// rows are run serially.
	Rows *StringMatrix `json:"rows,omitempty"`
	// Execution state of this stage, filled in once a stage has
	// begun (e.g., `STAGE_STATUS_{UNSPECIIFED,RUNNING,COMPLETED}`).
	Status *string `json:"status,omitempty"`

	ResourceError
}

// The main body of a change control.
type ChangeRsc struct {
	// The user-friendly name of this change.
	Name *string `json:"name,omitempty"`
	// The name of the root stage to be executed. This should match
	// an entry in `Stages`.
	RootStageId *string `json:"rootStageId,omitempty"`
	// All stages associated with this change.
	Stages *StageMap `json:"stages,omitempty"`

	ModifiedBy
}

// A change control paired with any scheduling information.
type ChangeControlRsc struct {
	// A unique key or UUID identifying this change control.
	Key ChangeControlKeyRsc `json:"key"`
	// The main body and actions of this change.
	Change *ChangeRsc `json:"change,omitempty"`
	// Flag which starts or stops the execution of this change control.
	Start *FlagRsc `json:"start,omitempty"`
	// An optional time at which that this change control will automatically
	// execute if it has been approved.
	Schedule *TimestampFlagRsc `json:"schedule,omitempty"`

	ResourceError
}

// Approval state for a change control.
type ApproveConfigRsc struct {
	// The unique key or UUID identifying the target change control.
	Key ChangeControlKeyRsc `json:"key"`
	// Approval state and metadata.
	Approve FlagRsc `json:"approve"`
	// The timestamp matching the last update time of the target change control.
	//
	// If `Version` does not match the change control, then this approval will
	// fail or be cancelled.
	Version *time.Time `json:"version,omitempty"`
}

func deserChangeControl(data []byte) (*ChangeControlRsc, error) {
	result := valueRsc{}
	cc := ChangeControlRsc{}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Errorf("ChangeControlRsc [Result]: %s Payload:\n%s",
			err, data)
	}

	if err := result.GetError(); err != nil {
		return nil, errors.Errorf("ChangeControlRsc [API]: %s", err)
	}

	if err := json.Unmarshal(result.Value, &cc); err != nil {
		return nil, errors.Errorf("ChangeControlRsc [Inner]: %s Payload:\n%s",
			err, data)
	}

	return &cc, nil
}

func deserApproval(data []byte) (*ApproveConfigRsc, error) {
	result := valueRsc{}
	cc := ApproveConfigRsc{}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Errorf("ApproveConfigRsc [Result]: %s Payload:\n%s",
			err, data)
	}

	if err := result.GetError(); err != nil {
		return nil, errors.Errorf("ApproveConfigRsc [API]: %s", err)
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
	change := ChangeRsc{}
	change.Notes = &notes

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
	start := FlagRsc{Value: true}
	start.Notes = &notes

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
	start := FlagRsc{Value: false}
	start.Notes = &notes

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
	sched := TimestampFlagRsc{Value: &schedTime}
	sched.Notes = &notes

	cfg.Schedule = &sched

	_, err := c.client.Post("/api/resources/changecontrol/v1/ChangeControlConfig", nil, cfg)

	if err != nil {
		err = errors.Errorf("ScheduleChangeControlRsc: %s", err)
	}

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
	sched.Notes = &notes

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
	flag := FlagRsc{Value: approve}
	flag.Notes = &notes

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
