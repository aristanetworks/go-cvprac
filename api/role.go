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

// RoleList list of roles
type RoleList struct {
	Total int    `json:"total"`
	Roles []Role `json:"Roles"`

	ErrorResponse
}

// SingleRole represents one role
type SingleRole struct {
	RoleDate Role `json:"role"`

	ErrorResponse
}

// GetAllRoles returns all the existing roles in CVP
func (c CvpRestAPI) GetAllRoles(start, end int) (*RoleList, error) {
	var roles RoleList

	query := &url.Values{
		"queryparam": {""},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/role/getRoles.do", query)
	if err != nil {
		return nil, errors.Errorf("GetAllRoles: %s", err)
	}

	if err = json.Unmarshal(resp, &roles); err != nil {
		return nil, errors.Errorf("GetAllRoles: %s Payload:\n%s", err, resp)
	}

	if err := roles.Error(); err != nil {
		// Entity does not exist
		if roles.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetAllRoles: %s", err)
	}
	return &roles, nil
}

// GetRole returns the role with ID 'roleID'
func (c CvpRestAPI) GetRole(roleID string) (*SingleRole, error) {
	var info SingleRole

	query := &url.Values{"userId": {roleID}}

	resp, err := c.client.Get("/role/getRole.do", query)
	if err != nil {
		return nil, errors.Errorf("GetRole: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetRole: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetRole: %s", err)
	}
	return &info, nil
}
