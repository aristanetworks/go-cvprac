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
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// CvpLabelList is a list of Labels
type CvpLabelList struct {
	Total     int     `json:"total"`
	LabelList []Label `json:"labels"`

	ErrorResponse
}

// Label represents a label
type Label struct {
	Key                  string `json:"key"`
	Name                 string `json:"name"`
	Note                 string `json:"note"`
	Type                 string `json:"type"`
	User                 string `json:"user"`
	DateTimeInLongFormat int64  `json:"dateTimeInLongFormat"`
	NetElementCount      int    `json:"netElementCount"`
	FactoryID            int    `json:"factoryId"`
	ID                   int    `json:"id"`

	ErrorResponse
}

// GetLabels returns the labels for
func (c CvpRestAPI) GetLabels(module, labelType, searchStr string,
	start int, end int) ([]Label, error) {
	var info CvpLabelList

	query := &url.Values{
		"module":     {module},
		"type":       {labelType},
		"queryparam": {searchStr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/label/getLabels.do", query)
	if err != nil {
		return nil, errors.Errorf("GetLabels: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetLabels: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("SearchConfiglets: %s", err)
	}
	return info.LabelList, nil
}

// GetLabelInfo returns the label info for the specified labelID
func (c CvpRestAPI) GetLabelInfo(labelID string) (*Label, error) {
	var info Label

	query := &url.Values{"labelId": {labelID}}

	resp, err := c.client.Get("/label/getLabelInfo.do", query)
	if err != nil {
		return nil, errors.Errorf("GetLabelInfo: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetLabelInfo: %s", err)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetLabelInfo: %s", err)
	}
	return &info, nil
}

// AddLabel adds a label
func (c CvpRestAPI) AddLabel(name string, note string, labeltype string) (*Label, error) {
	var info Label

	data := map[string]string{
		"name": name,
		"note": note,
		"type": labeltype,
	}

	resp, err := c.client.Post("/label/addLabel.do", nil, data)
	if err != nil {
		return nil, errors.Errorf("AddLabel: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("AddLabel: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("AddLabel: %s", err)
	}

	return &info, nil
}

// DeleteLabels deletes a list of Labels.
func (c CvpRestAPI) DeleteLabels(keys []string) error {
	var info ErrorResponse

	data := map[string][]string{
		"data": keys,
	}

	resp, err := c.client.Post("/label/deleteLabel.do", nil, data)
	if err != nil {
		return errors.Errorf("DeleteLabel: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("DeleteLabel: %s", err)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("DeleteLabel: %s", err)
	}
	return nil
}

// DeleteLabel deletes a Label.
func (c CvpRestAPI) DeleteLabel(key string) error {
	return c.DeleteLabels([]string{key})
}

// UpdateLabel updates a configlet.
func (c CvpRestAPI) UpdateLabel(name, key, note, labelType string) error {
	var info ErrorResponse

	data := map[string]string{
		"key":  key,
		"name": name,
		"note": note,
		"type": labelType,
	}

	resp, err := c.client.Post("/label/updateLabel.do", nil, data)
	if err != nil {
		return errors.Errorf("UpdateLabel: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("UpdateLabel: %s", err)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("UpdateLabel: %s", err)
	}
	return nil
}

// UpdateLabelNote updates a label note.
func (c CvpRestAPI) UpdateLabelNote(key, note string) error {
	var info ErrorResponse

	data := map[string]string{
		"key":  key,
		"note": note,
	}

	resp, err := c.client.Post("/label/updateNotesToLabel.do", nil, data)
	if err != nil {
		return errors.Errorf("UpdateLabelNote: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("UpdateLabelNote: %s", err)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("UpdateLabelNote: %s", err)
	}
	return nil
}
