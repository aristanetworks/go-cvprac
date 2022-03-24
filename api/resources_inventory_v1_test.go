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
	"testing"
)

func Test_CvpOnboardDeviceError_UnitTest(t *testing.T) {
	expectedResp := OnboardDeviceResponse{}

	respStr := `{"value": {"key":{"requestId":"123456789"},"hostnameOrIp":"1.2.3.4","device_type":"eos"}}`

	json.Unmarshal([]byte(respStr), expectedResp)

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	_, err := api.OnboardDevice("1.2.3.4", "eos")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_CvpDecomDeviceError_UnitTest(t *testing.T) {
	expectedResp := DecomDeviceResponse{}

	respStr := `{"value": {"key":{"request_id":"123456789"},"device_id": "123456789"}}`

	json.Unmarshal([]byte(respStr), expectedResp)

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	_, err := api.DecomDevice("123456789")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}
