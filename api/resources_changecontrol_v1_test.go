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
	"testing"
)

func Test_GetRsc_Error_UnitTest(t *testing.T) {
	respStr := `{"code":5, "message":"resource not found"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetChangeControlRsc("fakeId")
	assert(t, err != nil, "Failed to treat API error code as failure")
}

func Test_GetChangeControlsRsc_UnitTest(t *testing.T) {
	respStr := `{"result":{"value":{"key":{"id":"QM0BiWeRl"},"change":{"name":"Change 20230131_121158","rootStageId":"G0TTuEfr8q","stages":{"values":{"BttT5IDrq":{"name":"Exit BGP Maint","action":{"name":"exitbgpmaintmode","args":{"values":{"DeviceID":"SN-veos1"}}}},"G0TTuEfr8q":{"name":"Change 20230131_121158 Root","rows":{"values":[{"values":["BttT5IDrq"]}]}}}},"notes":"","time":"2023-01-31T12:12:16.054988005Z","user":"cvpadmin"}},"time":"2023-01-31T12:12:16.054988005Z","type":"INITIAL"}}
{"result":{"value":{"key":{"id":"QIKoyUtZjaFd-Y0LH3clY"},"change":{"name":"Change 20230125_151303","rootStageId":"kTUrxEn8coeXMKMVx0J4H","stages":{"values":{"6HvsEQ4tM5AOPPuE9Ftbz":{"name":"Exit BGP Maint","action":{"name":"exitbgpmaintmode","args":{"values":{"DeviceID":"SN-DC1-LEAF1A"}}},"rows":{},"status":"STAGE_STATUS_COMPLETED","error":"Error executing stage 6HvsEQ4tM5AOPPuE9Ftbz : On-boot maintenance check failed: Can not find on boot maintenance status from CLI response, tried path: [maintenanceUnits System onBootMaintenance]"},"kTUrxEn8coeXMKMVx0J4H":{"name":"Change 20230125_151303 Root","rows":{"values":[{"values":["6HvsEQ4tM5AOPPuE9Ftbz"]}]},"status":"STAGE_STATUS_COMPLETED","error":"Error executing stage 6HvsEQ4tM5AOPPuE9Ftbz : On-boot maintenance check failed: Can not find on boot maintenance status from CLI response, tried path: [maintenanceUnits System onBootMaintenance]"}}},"notes":"","time":"2023-01-26T09:29:50.263493510Z","user":"cvpadmin"},"approve":{"value":true,"notes":"Test approval.","time":"2023-02-02T14:38:40.446862005Z","user":"cvpadmin"},"start":{"value":true,"notes":"Started by Scheduled Change Control","time":"2023-02-02T14:40:00.053313982Z","user":"cvpadmin"},"status":"CHANGE_CONTROL_STATUS_COMPLETED","error":"Error executing stage 6HvsEQ4tM5AOPPuE9Ftbz : On-boot maintenance check failed: Can not find on boot maintenance status from CLI response, tried path: [maintenanceUnits System onBootMaintenance]","schedule":{"value":"2023-02-02T14:40:00.043Z","time":"2023-02-02T14:38:20.187702678Z","user":"cvpadmin"}},"time":"2023-02-02T14:40:00.748121056Z","type":"INITIAL"}}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	vals, err := api.GetChangeControlsRsc()
	ok(t, err)
	assert(t, len(vals) == 2, "Expected 2 change controls")
}

func Test_GetChangeControlRsc_UnitTest(t *testing.T) {
	ccId := "vYdaCZu_D798Bpubak4E_"
	respStr := `{"value":{"key":{"id":"vYdaCZu_D798Bpubak4E_"}, "change":{"name":"Change 20230126_113349", "rootStageId":"f9aGfU6iqmKW90-Z9soJV", "stages":{"values":{"f9aGfU6iqmKW90-Z9soJV":{"name":"Change 20230126_113349 Root", "rows":{}}}}, "notes":"", "time":"2023-01-26T12:16:39.819278546Z", "user":"cvpadmin"}, "error":"Reschedule required: Schedule time in the past for ChangeControl ID vYdaCZu_D798Bpubak4E_", "schedule":{"value":"2023-02-02T13:19:45.031602Z", "notes":"Testing sched timer.", "time":"2023-02-01T13:19:45.773491427Z", "user":"cvpadmin"}}, "time":"2023-02-06T12:04:47.521740196Z"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	val, err := api.GetChangeControlRsc(ccId)
	ok(t, err)
	assert(t, val.Change != nil && *val.Change.Name == "Change 20230126_113349", "failed to deserialize change/name")
	assert(t, val.Schedule != nil && val.Schedule.Time != nil, "Failed to handle Rsc API Schedule")
}

func Test_GetChangeControlApprovalsRsc_UnitTest(t *testing.T) {
	respStr := `{"result":{"value":{"key":{"id":"test"},"approve":{"value":true,"notes":"More Notes"},"version":"2023-02-01T15:11:26.624538804Z"},"time":"2023-02-02T11:59:50.507671741Z","type":"INITIAL"}}
{"result":{"value":{"key":{"id":"QIKoyUtZjaFd-Y0LH3clY"},"approve":{"value":true,"notes":"Test approval."},"version":"2023-01-26T09:29:50.263493510Z"},"time":"2023-02-02T14:38:40.446862005Z","type":"INITIAL"}}
{"result":{"value":{"key":{"id":"test2"},"approve":{"value":true,"notes":"More Notes"},"version":"2023-01-31T16:01:16.107005406Z"},"time":"2023-02-01T13:21:19.422609230Z","type":"INITIAL"}}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	vals, err := api.GetChangeControlApprovalsRsc()
	ok(t, err)
	assert(t, len(vals) == 3, "Expected 3 approvals")
}

func Test_GetChangeControlApprovalRsc_UnitTest(t *testing.T) {
	ccId := "vYdaCZu_D798Bpubak4E_"
	respStr := `{"value":{"key":{"id":"vYdaCZu_D798Bpubak4E_"}, "approve":{"value":true, "notes":"More Notes"}, "version":"2023-01-31T16:01:16.107005406Z"}, "time":"2023-02-01T13:21:19.422609230Z"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	val, err := api.GetChangeControlApprovalRsc(ccId)
	ok(t, err)
	assert(t, val.Approve.Value && *val.Approve.Notes == "More Notes", "failed to deserialize approval flag state")
}
