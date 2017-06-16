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

// NetElement represents a CVP network element returned as part of
// inventory query
type NetElement struct {
	IPAddress            string       `json:"ipAddress"`
	ModelName            string       `json:"modelName"`
	InternalVersion      string       `json:"internalVersion"`
	SystemMacAddress     string       `json:"systemMacAddress"`
	MemTotal             int          `json:"memTotal"`
	BootupTimeStamp      float64      `json:"bootupTimeStamp"`
	MemFree              int          `json:"memFree"`
	Architecture         string       `json:"architecture"`
	InternalBuildID      string       `json:"internalBuildId"`
	HardwareRevision     string       `json:"hardwareRevision"`
	Fqdn                 string       `json:"fqdn"`
	TaskIDList           []CvpTask    `json:"taskIdList"`
	ZtpMode              string       `json:"ztpMode"`
	Version              string       `json:"version"`
	SerialNumber         string       `json:"serialNumber"`
	Key                  string       `json:"key"`
	Type                 string       `json:"type"`
	TempActionList       []TempAction `json:"tempAction"`
	IsDANZEnabled        string       `json:"isDANZEnabled"`
	IsMLAGEnabled        string       `json:"isMLAGEnabled"`
	ComplianceIndication string       `json:"complianceIndication"`
	ComplianceCode       string       `json:"complianceCode"`
	LastSyncUp           int64        `json:"lastSyncUp"`
	UnAuthorized         bool         `json:"unAuthorized"`
	DeviceInfo           string       `json:"deviceInfo"`
	DeviceStatus         string       `json:"deviceStatus"`
	ParentContainerID    string       `json:"parentContainerId"`
}

// TempAction is
type TempAction struct {
	CcID                            string   `json:"ccId"`
	SessionID                       string   `json:"sessionId"`
	ContainerKey                    string   `json:"containerKey"`
	TaskID                          int      `json:"taskId"`
	Info                            string   `json:"info"`
	InfoPreview                     string   `json:"infoPreview"`
	Note                            string   `json:"note"`
	Action                          string   `json:"action"`
	NodeType                        string   `json:"nodeType"`
	NodeID                          string   `json:"nodeId"`
	ToID                            string   `json:"toId"`
	FromID                          string   `json:"fromId"`
	NodeName                        string   `json:"nodeName"`
	ToName                          string   `json:"toName"`
	FromName                        string   `json:"fromName"`
	ChildTasks                      []string `json:"childTasks"`
	ParentTask                      string   `json:"parentTask"`
	OldNodeName                     string   `json:"oldNodeName"`
	ToIDType                        string   `json:"toIdType"`
	ConfigletList                   []string `json:"configletList"`
	IgnoreConfigletList             []string `json:"ignoreConfigletList"`
	ConfigletNamesList              []string `json:"configletNamesList"`
	IgnoreConfigletNamesList        []string `json:"ignoreConfigletNamesList"`
	NodeList                        []string `json:"nodeList"`
	IgnoreNodeList                  []string `json:"ignoreNodeList"`
	NodeNamesList                   []string `json:"nodeNamesList"`
	IgnoreNodeNamesList             []string `json:"ignoreNodeNamesList"`
	NodeIPAddress                   string   `json:"nodeIpAddress"`
	NodeTargetIPAddress             string   `json:"nodeTargetIpAddress"`
	Key                             string   `json:"key"`
	IgnoreNodeID                    string   `json:"ignoreNodeId"`
	IgnoreNodeName                  string   `json:"ignoreNodeName"`
	ImageBundleID                   string   `json:"imageBundleId"`
	Mode                            string   `json:"mode"`
	Timestamp                       int64    `json:"timestamp"`
	ConfigletBuilderList            []string `json:"configletBuilderList"`
	ConfigletBuilderNamesList       []string `json:"configletBuilderNamesList"`
	IgnoreConfigletBuilderList      []string `json:"ignoreConfigletBuilderList"`
	IgnoreConfigletBuilderNamesList []string `json:"ignoreConfigletBuilderNamesList"`
	PageType                        string   `json:"pageType"`
	ViaContainer                    bool     `json:"viaContainer"`
	BestImageContainerID            string   `json:"bestImageContainerId"`
	User                            string   `json:"user"`
	FactoryID                       int      `json:"factoryId"`
	ID                              int      `json:"id"`
}

// CvpInventoryList is a list of NetElements and Containers
type CvpInventoryList struct {
	Total          int               `json:"total"`
	ContainerList  map[string]string `json:"containerList"`
	NetElementList []NetElement      `json:"netElementList"`

	ErrorResponse
}

// GetInventory returns a CvpInventoryList based on a provided query and range.
//
// Failed search returns empty
// {
//   "total": 0,
//   "containerList": {},
//   "netElementList": []
// }
func (c CvpRestAPI) GetInventory(querystr string, start int, end int) (*CvpInventoryList, error) {
	var info CvpInventoryList
	query := &url.Values{
		"queryparam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/inventory/getInventory.do", query)
	if err != nil {
		return nil, fmt.Errorf("GetInventory: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("GetInventory: %s", err)

	}

	if err := info.Error(); err != nil {
		return nil, fmt.Errorf("GetInventory: %s", err)
	}

	return &info, nil
}

// GetAllDevices returns CvpInventoryList of all current inventory
func (c CvpRestAPI) GetAllDevices() ([]NetElement, error) {
	data, err := c.GetInventory("", 0, 0)
	if err != nil {
		return nil, fmt.Errorf("GetAllDevices: %s", err)
	}
	if len(data.NetElementList) > 0 {
		return data.NetElementList, nil
	}
	return nil, nil
}

// GetDeviceByName returns a CvpInventoryList based on device name provided
func (c CvpRestAPI) GetDeviceByName(fqdn string) (*NetElement, error) {
	data, err := c.GetInventory(fqdn, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("GetDeviceByName: %s", err)
	}

	if len(data.NetElementList) > 0 {
		return &data.NetElementList[0], nil
	}
	return nil, nil
}

// GetDevicesInContainer returns a CvpInventoryList based on container name provided
func (c CvpRestAPI) GetDevicesInContainer(name string) ([]NetElement, error) {
	data, err := c.GetInventory(name, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("GetDevicesInContainer: %s", err)
	}

	if len(data.NetElementList) > 0 {
		return data.NetElementList, nil
	}
	return nil, nil
}

// GetUndefinedDevices returns a NetElement list of devices within the Undefined container
func (c CvpRestAPI) GetUndefinedDevices() ([]NetElement, error) {
	var res []NetElement

	data, err := c.GetInventory("undefined", 0, 0)
	if err != nil {
		return nil, fmt.Errorf("GetUndefinedDevices: %s", err)
	}

	numElements := len(data.NetElementList)
	if numElements > 0 {
		var idx int
		res := make([]NetElement, numElements)
		for _, netElement := range data.NetElementList {
			if netElement.ParentContainerID == "undefined_container" {
				res[idx] = netElement
				idx++
			}
		}
	}
	return res, nil
}

// Container is
type Container struct {
	ChildContainerID bool   `json:"childContainerId"`
	FactoryID        int    `json:"factoryId"`
	ID               int    `json:"id"`
	Key              string `json:"key"`
	Name             string `json:"name"`
	ParentID         string `json:"parentId"`
	Type             string `json:"type"`
	UserID           string `json:"userId"`
}

// ContainerList is a list of NetElements and Containers
type ContainerList struct {
	Total         int         `json:"total"`
	ContainerList []Container `json:"data"`

	ErrorResponse
}

// GetContainer returns
// The endpoint searchContainers.do will not return the Undefined_Container in the list
func (c CvpRestAPI) GetContainer(querystr string, start int, end int) (*ContainerList, error) {
	var info ContainerList
	query := &url.Values{
		"queryparam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/inventory/add/searchContainers.do", query)
	if err != nil {
		return nil, fmt.Errorf("GetContainer: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("GetContainer: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, fmt.Errorf("GetContainer: %s", err)
	}
	return &info, nil
}

// GetAllContainers returns all current inventory Containers
func (c CvpRestAPI) GetAllContainers() (*ContainerList, error) {
	return c.GetContainer("", 0, 0)
}

// GetContainerByName returns a Container
func (c CvpRestAPI) GetContainerByName(name string) (*Container, error) {
	data, err := c.GetContainer(name, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("GetContainerByName: %s", err)
	}
	if data.Total > 0 {
		for _, cont := range data.ContainerList {
			if cont.Name == name {
				return &cont, nil
			}
		}
	}
	return nil, nil
}
