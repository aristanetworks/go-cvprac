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
	"errors"
	"testing"
)

func Test_CvpGetAllUsersRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetAllUsers: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllUsers(0, 0)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetAllUsersJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllUsers(0, 0); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllUsersEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllUsers(0, 0); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllUsersReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllUsers(0, 0); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetAllUsersValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllUsers(0, 0)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_CvpGetAllUsersValid2_UnitTest(t *testing.T) {
	data := `{
  "total": 2,
  "roles": {
    "cvpadmin": [
      "network-admin"
    ],
    "user1": []
  },
  "users": [
    {
      "userId": "cvpadmin",
      "firstName": "",
      "lastName": "",
      "email": "cwomble@arista.com",
      "lastAccessed": 1550588080210,
      "userStatus": "Enabled",
      "currentStatus": "Online",
      "contactNumber": null,
      "userType": "Local",
      "factoryId": 1,
      "id": 28
    },
    {
      "userId": "user1",
      "firstName": "",
      "lastName": "",
      "email": "user1@arista.com",
      "lastAccessed": 0,
      "userStatus": "Enabled",
      "currentStatus": "",
      "contactNumber": null,
      "userType": "Local",
      "factoryId": 1,
      "id": 28
    }
  ]
}`
	client := NewMockClient(data, nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllUsers(0, 0)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_CvpGetUserRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetUser: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetUser("user")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetUserJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetUser("user"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetUserEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetUser("user"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetUserReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetUser("user"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetUserValid_UnitTest(t *testing.T) {
	data := `{
  "roles": [
    "network-admin"
  ],
  "user": {
    "userId": "cvpadmin",
    "firstName": "",
    "lastName": "",
    "email": "cwomble@arista.com",
    "lastAccessed": 1550677918078,
    "userStatus": "Enabled",
    "currentStatus": "Online",
    "contactNumber": null,
    "userType": "Local",
    "factoryId": 1,
    "id": 28
  }
}`
	client := NewMockClient(data, nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetUser("cvpadmin")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}

}
