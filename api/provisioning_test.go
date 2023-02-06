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
	"reflect"
	"testing"
)

func Test_CvpGetConfigletsByDeviceIDRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetConfigletsByDeviceID: GetDeviceConfigletInfo: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetConfigletsByDeviceID("00:50:56:50:a8:af")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpGetConfigletsByDeviceIDJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetConfigletsByDeviceID("00:50:56:50:a8:af"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetConfigletsByDeviceIDNilJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetConfigletsByDeviceID("00:50:56:50:a8:af"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetConfigletsByDeviceIDReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetConfigletsByDeviceID("00:50:56:50:a8:af"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetConfigletsByDeviceIDValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetConfigletsByDeviceID("00:50:56:50:a8:af")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_checkConfigMapping(t *testing.T) {
	type args struct {
		applied       []Configlet
		newconfiglets []Configlet
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   configletAndBuilderKeyNames
		wantErr bool
	}{
		{
			"",
			args{
				applied: []Configlet{
					{
						Reconciled: false,
						Name:       "foo",
						Key:        "foo",
						Type:       "Static",
					},
					{
						Reconciled: true,
						Name:       "foo2",
						Key:        "foo2",
						Type:       "Reconciled",
					},
				},
				newconfiglets: []Configlet{
					{
						Reconciled: false,
						Name:       "foo3",
						Key:        "foo3",
						Type:       "Static",
					},
					{
						Reconciled: false,
						Name:       "foo4",
						Key:        "foo4",
						Type:       "Static",
					},
				},
			},
			true,
			configletAndBuilderKeyNames{
				[]string{"foo", "foo3", "foo4", "foo2"},
				[]string{"foo", "foo3", "foo4", "foo2"},
				nil,
				nil,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, configletAndBuilders, err := checkConfigMapping(tt.args.applied,
				tt.args.newconfiglets)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkConfigMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("checkConfigMapping() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(configletAndBuilders, tt.want1) {
				t.Errorf("checkConfigMapping() got1 = %v, want %v", configletAndBuilders, tt.want1)
			}
		})
	}
}

func Test_checkRemoveConfigMapping(t *testing.T) {
	type args struct {
		applied      []Configlet
		rmConfiglets []Configlet
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		want1   configletAndBuilderKeyNames
		want2   configletAndBuilderKeyNames
		wantErr bool
	}{
		{
			"",
			args{
				applied: []Configlet{
					{
						Reconciled: false,
						Name:       "foo",
						Key:        "foo",
						Type:       "Static",
					},
					{
						Reconciled: false,
						Name:       "foo2",
						Key:        "foo2",
						Type:       "Static",
					},
					{
						Reconciled: true,
						Name:       "foo3",
						Key:        "foo3",
						Type:       "Reconciled",
					},
				},
				rmConfiglets: []Configlet{
					{
						Reconciled: false,
						Name:       "foo2",
						Key:        "foo2",
						Type:       "Static",
					},
				},
			},
			true,
			configletAndBuilderKeyNames{
				[]string{
					"foo",
					"foo3",
				},
				[]string{
					"foo",
					"foo3",
				},
				nil,
				nil,
			},
			configletAndBuilderKeyNames{
				[]string{
					"foo2",
				},
				[]string{"foo2"},
				nil,
				nil,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, configletAndBuilders, rmConfigletAndBuilders, err :=
				checkRemoveConfigMapping(tt.args.applied, tt.args.rmConfiglets)

			if (err != nil) != tt.wantErr {
				t.Errorf("checkRemoveConfigMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("checkRemoveConfigMapping() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(configletAndBuilders, tt.want1) {
				t.Errorf("checkRemoveConfigMapping() configletAndBuilders = %v, want %v",
					configletAndBuilders, tt.want1)
			}

			if !reflect.DeepEqual(rmConfigletAndBuilders, tt.want2) {
				t.Errorf("checkRemoveConfigMapping() rmConfigletAndBuilders = %v, want %v",
					rmConfigletAndBuilders, tt.want2)
			}
		})
	}
}

func Test_changesNeeded(t *testing.T) {
	type args struct {
		applied    []Configlet
		configlets []Configlet
	}

	tests := []struct {
		name    string
		args    args
		want    *configletAndBuilderKeyNames
		want1   *configletAndBuilderKeyNames
		wantErr bool
	}{
		{
			"",
			args{
				applied: []Configlet{
					{
						Reconciled: false,
						Name:       "foo",
						Key:        "foo",
						Type:       "Static",
					},
					{
						Reconciled: false,
						Name:       "foo2",
						Key:        "foo2",
						Type:       "Static",
					},
					{
						Reconciled: true,
						Name:       "foo3",
						Key:        "foo3",
						Type:       "Reconciled",
					},
				},
				configlets: []Configlet{
					{
						Reconciled: false,
						Name:       "foo2",
						Key:        "foo2",
						Type:       "Static",
					},
				},
			},
			&configletAndBuilderKeyNames{
				keys:   []string{"foo2"},
				names:  []string{"foo2"},
				bKeys:  nil,
				bNames: nil,
			},
			&configletAndBuilderKeyNames{
				keys:   []string{"foo", "foo3"},
				names:  []string{"foo", "foo3"},
				bKeys:  nil,
				bNames: nil,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := changesNeeded(tt.args.applied, tt.args.configlets)
			if (err != nil) != tt.wantErr {
				t.Errorf("changesNeeded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("changesNeeded() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("changesNeeded() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
