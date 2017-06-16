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
	"fmt"
	"net/url"
	"strconv"
)

// Configlet represents a Configlet
type Configlet struct {
	IsDefault            string `json:"isDefault"`
	DateTimeInLongFormat int    `json:"dateTimeInLongFormat"`
	ContainerCount       int    `json:"containerCount"`
	NetElementCount      int    `json:"netElementCount"`
	IsAutoBuilder        string `json:"isAutoBuilder"`
	Reconciled           bool   `json:"reconciled"`
	FactoryID            int    `json:"factoryId"`
	Config               string `json:"config"`
	User                 string `json:"user"`
	Note                 string `json:"note"`
	Name                 string `json:"name"`
	Key                  string `json:"key"`
	ID                   int    `json:"id"`
	Type                 string `json:"type"`

	ErrorResponse
}

func (c Configlet) String() string {
	return c.Name
}

// ConfigletHistoryEntry represents a configlet history entry
type ConfigletHistoryEntry struct {
	ConfigletID                 string `json:"configletId"`
	OldUserID                   string `json:"oldUserId"`
	NewUserID                   string `json:"newUserId"`
	OldConfig                   string `json:"oldConfig"`
	NewConfig                   string `json:"newConfig"`
	OldDate                     string `json:"oldDate"`
	NewDate                     string `json:"newDate"`
	OldDateTimeInLongFormat     int    `json:"oldDateTimeInLongFormat"`
	UpdatedDateTimeInLongFormat int    `json:"updatedDateTimeInLongFormat"`
	FactoryID                   int    `json:"factoryId"`
	Key                         string `json:"key"`
	ID                          int    `json:"id"`
}

// ConfigletHistoryList represents a list of ConfigletHistoryEntry's
type ConfigletHistoryList struct {
	Total       int                     `json:"total"`
	HistoryList []ConfigletHistoryEntry `json:"configletHistory"`

	ErrorResponse
}

// ConfigletOpReturn represents the
type ConfigletOpReturn struct {
	Data Configlet `json:"data"`

	ErrorResponse
}

// GetConfigletByName returns the configlet with the specified name
func (c CvpRestAPI) GetConfigletByName(name string) (*Configlet, error) {
	var info Configlet

	query := &url.Values{"name": {name}}

	resp, err := c.client.Get("/configlet/getConfigletByName.do", query)
	if err != nil {
		return nil, fmt.Errorf("GetConfigletByName: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("GetConfigletByName: %s", err)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, fmt.Errorf("GetConfigletByName: %s", err)
	}
	return &info, nil
}

// GetConfigletHistory returns the history for a configlet provided the key, and a range.
func (c CvpRestAPI) GetConfigletHistory(key string, start int,
	end int) (*ConfigletHistoryList, error) {
	var info ConfigletHistoryList

	//queryparam := url.Values{"name": {key},}

	query := &url.Values{
		"configletId": {key},
		"queryparam":  {""},
		"startIndex":  {strconv.Itoa(start)},
		"endIndex":    {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/configlet/getConfigletHistory.do", query)
	if err != nil {
		return nil, fmt.Errorf("GetConfigletHistory: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("GetConfigletHistory: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, fmt.Errorf("GetConfigletHistory: %s", err)
	}

	return &info, nil
}

// GetAllConfigletHistory returns all the history for a given configlet
func (c CvpRestAPI) GetAllConfigletHistory(key string) (*ConfigletHistoryList, error) {
	return c.GetConfigletHistory(key, 0, 0)
}

// AddConfiglet creates/adds a configlet
func (c CvpRestAPI) AddConfiglet(name string, config string) (*Configlet, error) {
	var info ConfigletOpReturn

	data := map[string]string{
		"name":   name,
		"config": config,
	}

	resp, err := c.client.Post("/configlet/addConfiglet.do", nil, data)
	if err != nil {
		return nil, fmt.Errorf("AddConfiglet: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("AddConfiglet: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, fmt.Errorf("AddConfiglet: %s", err)
	}

	return &info.Data, nil
}

// DeleteConfiglet deletes a configlet.
func (c CvpRestAPI) DeleteConfiglet(name string, key string) error {
	var info ErrorResponse

	data := []map[string]string{
		map[string]string{
			"name": name,
			"key":  key,
		},
	}
	resp, err := c.client.Post("/configlet/deleteConfiglet.do", nil, data)
	if err != nil {
		return fmt.Errorf("DeleteConfiglet: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return fmt.Errorf("DeleteConfiglet: %s", err)
	}

	if err := info.Error(); err != nil {
		return fmt.Errorf("DeleteConfiglet: %s", err)
	}

	return nil
}

// UpdateConfiglet updates a configlet.
func (c CvpRestAPI) UpdateConfiglet(config string, name string, key string) error {
	var info ErrorResponse

	data := map[string]string{
		"config": config,
		"key":    key,
		"name":   name,
	}

	resp, err := c.client.Post("/configlet/updateConfiglet.do", nil, data)
	if err != nil {
		return fmt.Errorf("UpdateConfiglet: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return fmt.Errorf("UpdateConfiglet: %s", err)
	}

	if err := info.Error(); err != nil {
		return fmt.Errorf("UpdateConfiglet: %s", err)
	}

	return nil
}
