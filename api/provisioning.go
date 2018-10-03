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

// ActionRequest request structure for saveTopology
type ActionRequest struct {
	Data []Action `json:"data,omitempty"`
}

// Action request structure for saveTopology
type Action struct {
	Action                          string   `json:"action"`
	ConfigletBuilderList            []string `json:"configletBuilderList"`
	ConfigletBuilderNamesList       []string `json:"configletBuilderNamesList"`
	ConfigletList                   []string `json:"configletList"`
	ConfigletNamesList              []string `json:"configletNamesList"`
	FromID                          string   `json:"fromId"`
	FromName                        string   `json:"fromName"`
	IgnoreConfigletBuilderList      []string `json:"ignoreConfigletBuilderList"`
	IgnoreConfigletBuilderNamesList []string `json:"ignoreConfigletBuilderNamesList"`
	IgnoreConfigletList             []string `json:"ignoreConfigletList"`
	IgnoreConfigletNamesList        []string `json:"ignoreConfigletNamesList"`
	IgnoreNodeID                    string   `json:"ignoreNodeId"`
	IgnoreNodeName                  string   `json:"ignoreNodeName"`
	Info                            string   `json:"info"`
	InfoPreview                     string   `json:"infoPreview"`
	NodeID                          string   `json:"nodeId"`
	NodeIPAddress                   string   `json:"nodeIpAddress"`
	NodeName                        string   `json:"nodeName"`
	NodeTargetIPAddress             string   `json:"nodeTargetIpAddress"`
	NodeType                        string   `json:"nodeType"`
	ToID                            string   `json:"toId"`
	ToIDType                        string   `json:"toIdType"`
	ToName                          string   `json:"toName"`

	CCID                 string   `json:"ccId,omitempty"`
	ID                   int      `json:"id,omitempty"`
	Note                 string   `json:"note,omitempty"`
	ChildTasks           []string `json:"childTasks,omitempty"`
	ParentTask           string   `json:"parentTask,omitempty"`
	FactoryID            int      `json:"factoryId,omitempty"`
	BestImageContainerID string   `json:"bestImageContainerId,omitempty"`
	SessionID            string   `json:"sessionId,omitempty"`
	ContainerKey         string   `json:"containerKey,omitempty"`
	TaskID               int      `json:"taskId,omitempty"`
	OldNodeName          string   `json:"oldNodeName,omitempty"`
	NodeList             []string `json:"nodeList,omitempty"`
	IgnoreNodeList       []string `json:"ignoreNodeList,omitempty"`
	NodeNamesList        []string `json:"nodeNamesList,omitempty"`
	IgnoreNodeNamesList  []string `json:"ignoreNodeNamesList,omitempty"`
	UserID               string   `json:"userId,omitempty"`
	Key                  string   `json:"key,omitempty"`
	ImageBundleID        string   `json:"imageBundleId,omitempty"`
	Mode                 string   `json:"mode,omitempty"`
	Timestamp            float64  `json:"timestamp,omitempty"`
	PageType             string   `json:"pageType,omitempty"`
	ViaContainer         bool     `json:"viaContainer,omitempty"`
}

// TopologyResp ..
type TopologyResp struct {
	Topology Topology `json:"topology"`
	Type     string   `json:"type"`

	ErrorResponse
}

// Topology ..
type Topology struct {
	Key                      string        `json:"key"`
	Name                     string        `json:"name"`
	Type                     string        `json:"type"`
	ChildContainerCount      int           `json:"childContainerCount"`
	ChildNetElementCount     int           `json:"childNetElementCount"`
	ParentContainerID        interface{}   `json:"parentContainerId"`
	Mode                     string        `json:"mode"`
	DevStatus                DeviceStatus  `json:"deviceStatus"`
	ChildTaskCount           int           `json:"childTaskCount"`
	ChildContainerList       []Topology    `json:"childContainerList"`
	ChildNetElementList      []NetElement  `json:"childNetElementList"`
	HierarchyNetElementCount int           `json:"hierarchyNetElementCount"`
	TempAction               []interface{} `json:"tempAction"`
	TempEvent                []interface{} `json:"tempEvent"`
}

// DeviceStatus represents status for a device
type DeviceStatus struct {
	Critical     int `json:"critical"`
	Warning      int `json:"warning"`
	Normal       int `json:"normal"`
	ImageUpgrade int `json:"imageUpgrade"`
	Task         int `json:"task"`
	UnAuthorized int `json:"unAuthorized"`
}

// ConfigletMapping represents basic info related to a Configlet
type ConfigletMapping struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ConfigletInfo represents the configlets for a netelemet
type ConfigletInfo struct {
	Total           int                         `json:"total"`
	ConfigletMapper map[string]ConfigletMapping `json:"configletMapper"`
	ConfigletList   []Configlet                 `json:"configletList"`

	ErrorResponse
}

// ContainerData represents a container within CVP
type ContainerData struct {
	Undefined            bool   `json:"undefined"`
	UserID               string `json:"userId"`
	DateTimeInLongFormat int64  `json:"dateTimeInLongFormat"`
	FactoryID            int    `json:"factoryId"`
	Root                 bool   `json:"root"`
	Mode                 string `json:"mode"`
	Name                 string `json:"name"`
	Key                  string `json:"key"`
	ID                   int    `json:"id"`
}

// NetElementContainer maps a specific netelemet to a container
type NetElementContainer struct {
	ContainerKey  string `json:"containerKey"`
	ContainerName string `json:"containerName"`
	NetElementKey string `json:"netElementKey"`
}

// SearchTopologyResp represents the response from SearchTopology request
type SearchTopologyResp struct {
	NetElementContainerList []NetElementContainer `json:"netElementContainerList"`
	Total                   int                   `json:"total"`
	KeywordList             []string              `json:"keywordList"`
	ContainerList           []ContainerData       `json:"containerList"`
	NetElementList          []NetElement          `json:"netElementList"`

	ErrorResponse
}

// ComplianceResp represents a response from a Compliance check
type ComplianceResp struct {
	Architecture         string    `json:"architecture"`
	BootupTimeStamp      float64   `json:"bootupTimeStamp"`
	ComplianceCode       string    `json:"complianceCode"`
	ComplianceIndication string    `json:"complianceIndication"`
	DeviceStatus         string    `json:"deviceStatus"`
	DeviceStatusInfo     string    `json:"deviceStatusInfo"`
	Fqdn                 string    `json:"fqdn"`
	HardwareRevision     string    `json:"hardwareRevision"`
	InternalBuildID      string    `json:"internalBuildId"`
	InternalVersion      string    `json:"internalVersion"`
	IPAddress            string    `json:"ipAddress"`
	IsDANZEnabled        string    `json:"isDANZEnabled"`
	IsMLAGEnabled        string    `json:"isMLAGEnabled"`
	Key                  string    `json:"key"`
	LastSyncUp           int64     `json:"lastSyncUp"`
	MemFree              int       `json:"memFree"`
	MemTotal             int       `json:"memTotal"`
	ModelName            string    `json:"modelName"`
	SerialNumber         string    `json:"serialNumber"`
	SystemMacAddress     string    `json:"systemMacAddress"`
	TaskIDList           []CvpTask `json:"taskIdList"`
	Type                 string    `json:"type"`
	UnAuthorized         bool      `json:"unAuthorized"`
	Version              string    `json:"version"`
	ZtpMode              string    `json:"ztpMode"`
	//tempAction  null `json:"tempAction"`

	ErrorResponse
}

// TaskResp represents a task response
type TaskResp struct {
	Data TaskInfo `json:"data"`
}

// TaskInfo represents task info
type TaskInfo struct {
	TaskIDs []string `json:"taskIds"`
	Status  string   `json:"status"`
}

// GetDeviceByID returns NetElement info related to a device mac.
func (c CvpRestAPI) GetDeviceByID(mac string) (*NetElement, error) {
	var info NetElement

	query := &url.Values{
		"netElementId": {mac},
	}

	resp, err := c.client.Get("/provisioning/getNetElementById.do", query)
	if err != nil {
		return nil, errors.Errorf("GetDeviceByID: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetDeviceByID: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetDeviceByID: %s", err)
	}
	return &info, nil
}

// GetDeviceConfigletInfo returns all configlet info related to a device.
func (c CvpRestAPI) GetDeviceConfigletInfo(mac string) (*ConfigletInfo, error) {
	var info ConfigletInfo
	query := &url.Values{
		"netElementId": {mac},
		"queryParam":   {""},
		"startIndex":   {"0"},
		"endIndex":     {"0"},
	}

	resp, err := c.client.Get("/provisioning/getConfigletsByNetElementId.do", query)
	if err != nil {
		return nil, errors.Errorf("GetDeviceConfigletInfo: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetDeviceConfigletInfo: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetDeviceConfigletInfo: %s", err)
	}
	return &info, nil
}

// GetConfigletsByDeviceID returns the list of configlets applied to a device.
func (c CvpRestAPI) GetConfigletsByDeviceID(mac string) ([]Configlet, error) {
	info, err := c.GetDeviceConfigletInfo(mac)
	if err != nil {
		return nil, errors.Errorf("GetConfigletsByDeviceID: %s", err)
	}
	return info.ConfigletList, nil
}

func (c CvpRestAPI) addTempAction(data *ActionRequest) error {
	var resp ErrorResponse

	query := &url.Values{
		"format":     {"topology"},
		"queryParam": {""},
		"nodeId":     {"root"},
	}

	reqResp, err := c.client.Post("/ztp/addTempAction.do", query, data)
	if err != nil {
		return errors.Errorf("addTempAction: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return errors.Errorf("addTempAction: %s", err)
	}

	if err := resp.Error(); err != nil {
		return errors.Errorf("addTempAction: %s", err)
	}
	return nil
}

// SaveTopology Schedule tasks for many operations like configlet and image bundle
// mapping/removal to/from device or container, addition/deletion of containers,
// deletion of device. Return a list of taskIds created in response to saving
// the topology.
func (c CvpRestAPI) SaveTopology() (*TaskInfo, error) {
	var resp TaskResp

	reqResp, err := c.client.Post("/ztp/v2/saveTopology.do", nil, []string{})
	if err != nil {
		return nil, errors.Errorf("SaveTopology: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("SaveTopology: %s", err)
	}

	return &resp.Data, nil
}

// ApplyConfigletsToDevice apply the configlets to the device.
func (c CvpRestAPI) ApplyConfigletsToDevice(appName string, dev *NetElement, commit bool,
	newConfiglets ...Configlet) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: nil NetElement")
	}

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}

	action, configletMap, builderMap, err := checkConfigMapping(configlets, newConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "ApplyConfigletsToDevice")
	}

	if !action {
		return nil, nil
	}

	// Get a list of the names and keys of the configlets
	ckeys, cnames := keyValueSliceFromMap(configletMap)

	// Get a list of the names and keys of the configlet builders
	cbkeys, cbnames := keyValueSliceFromMap(builderMap)

	info := appName + ": Configlet Assign: to Device " + dev.Fqdn
	infoPreview := "<b>Configlet Assign:</b> to Device" + dev.Fqdn

	data := &ActionRequest{Data: []Action{
		Action{
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            cbkeys,
			ConfigletBuilderNamesList:       cbnames,
			ConfigletList:                   ckeys,
			ConfigletNamesList:              cnames,
			IgnoreConfigletBuilderNamesList: []string{},
			IgnoreConfigletBuilderList:      []string{},
			IgnoreConfigletNamesList:        []string{},
			IgnoreConfigletList:             []string{},
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeIPAddress:                   dev.IPAddress,
			NodeName:                        "",
			NodeTargetIPAddress:             dev.IPAddress,
			FromName:                        "",
			ToName:                          dev.Fqdn,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}
	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ApplyConfigletToDevice apply the configlets to the device.
func (c CvpRestAPI) ApplyConfigletToDevice(appName string, dev *NetElement,
	newConfiglet *Configlet, commit bool) (*TaskInfo, error) {
	var newConfigletList []Configlet
	newConfigletList = append(newConfigletList, *newConfiglet)
	return c.ApplyConfigletsToDevice(appName, dev, commit, newConfigletList...)
}

// RemoveConfigletsFromDevice Remove the configlets from the device.
func (c CvpRestAPI) RemoveConfigletsFromDevice(appName string, dev *NetElement, commit bool,
	remConfiglets ...Configlet) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("RemoveConfigletsFromDevice: nil NetElement")
	}

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("RemoveConfigletsFromDevice: %s", err)
	}

	action, configletMap, builderMap, rmConfigletMap, rmBuilderMap, err :=
		checkRemoveConfigMapping(configlets, remConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "RemoveConfigletsFromDevice")
	}

	if !action {
		return nil, nil
	}

	// Build a list of the configlet names/keys to remove.
	cKeys, cNames := keyValueSliceFromMap(configletMap)

	// Build a list of the configlet names/keys to remove.
	cbKeys, cbNames := keyValueSliceFromMap(builderMap)

	// Build a list of the configlet names/keys to remove.
	rmKeys, rmNames := keyValueSliceFromMap(rmConfigletMap)

	// Build a list of the configlet names/keys to remove.
	rmbKeys, rmbNames := keyValueSliceFromMap(rmBuilderMap)

	info := appName + ": Configlet Remove: from Device " + dev.Fqdn
	infoPreview := "<b>Configlet Remove:</b> from Device" + dev.Fqdn

	data := &ActionRequest{Data: []Action{
		Action{
			ID:                              1,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletList:                   cKeys,
			ConfigletNamesList:              cNames,
			ConfigletBuilderList:            cbKeys,
			ConfigletBuilderNamesList:       cbNames,
			IgnoreConfigletList:             rmKeys,
			IgnoreConfigletNamesList:        rmNames,
			IgnoreConfigletBuilderList:      rmbKeys,
			IgnoreConfigletBuilderNamesList: rmbNames,
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeName:                        "",
			NodeIPAddress:                   dev.IPAddress,
			NodeTargetIPAddress:             dev.IPAddress,
			FromName:                        "",
			ToName:                          dev.Fqdn,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("RemoveConfigletsFromDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ResetDevice Resets/Reboots the device to factory setting.
func (c CvpRestAPI) ResetDevice(appName string, dev *NetElement,
	container *Container, commit bool) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("ResetDevice: nil NetElement ref provided")
	}
	if container == nil {
		return nil, errors.Errorf("ResetDevice: nil Container ref provided")
	}

	info := appName + ": Reset: Device Reset: %s - To be Reset" + dev.Fqdn
	infoPreview := "<b>Device Reset:</b> %s - To be Reset" + dev.Fqdn

	data := &ActionRequest{Data: []Action{
		Action{
			ID:          1,
			Action:      "reset",
			FromID:      dev.ParentContainerID,
			FromName:    container.Name,
			Info:        info,
			InfoPreview: infoPreview,
			NodeID:      dev.SystemMacAddress,
			NodeName:    dev.Fqdn,
			NodeType:    "netelement",
			ToID:        "undefined_container",
			ToIDType:    "container",
			ParentTask:  "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ResetDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// RemoveConfigletFromDevice Remove the configlets from the device.
func (c CvpRestAPI) RemoveConfigletFromDevice(appName string, dev *NetElement,
	remConfiglet *Configlet, commit bool) (*TaskInfo, error) {
	var remConfigletList []Configlet
	remConfigletList = append(remConfigletList, *remConfiglet)
	return c.RemoveConfigletsFromDevice(appName, dev, commit, remConfigletList...)
}

func (c CvpRestAPI) containerOp(containerName, containerKey, parentName,
	parentKey, operation string) (*TaskInfo, error) {

	msg := operation + " container " + containerName + " under container " + parentName

	data := &ActionRequest{Data: []Action{
		Action{
			Info:        msg,
			InfoPreview: msg,
			Action:      operation,
			NodeType:    "container",
			NodeID:      containerKey,
			ToID:        parentKey,
			ToName:      parentName,
			FromID:      "",
			FromName:    "",
			NodeName:    containerName,
		},
	}}
	c.addTempAction(data)
	return c.SaveTopology()
}

// AddContainer adds the container to the specified parent.
func (c CvpRestAPI) AddContainer(containerName, parentName,
	parentKey string) error {
	_, err := c.containerOp(containerName, "New_container1", parentName, parentKey, "add")
	return err
}

// DeleteContainer deletes the container from the specified parent.
func (c CvpRestAPI) DeleteContainer(containerName, containerKey,
	parentName, parentKey string) error {
	_, err := c.containerOp(containerName, containerKey, parentName, parentKey, "delete")
	return err
}

// SearchTopologyWithRange searches the topology for items matching the query parameter
// and returning those within the specified range.
//
// If query yields no hits, then result is (SearchTopologyResp{})
func (c CvpRestAPI) SearchTopologyWithRange(querystr string, start int,
	end int) (*SearchTopologyResp, error) {
	var resp SearchTopologyResp
	query := &url.Values{
		"queryParam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/provisioning/searchTopology.do", query)
	if err != nil {
		return nil, errors.Errorf("SearchTopologyWithRange: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("SearchTopologyWithRange: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("SearchTopologyWithRange: %s", err)
	}
	return &resp, nil
}

// SearchTopology searches the topology for items matching the query parameter.
func (c CvpRestAPI) SearchTopology(query string) (*SearchTopologyResp, error) {
	return c.SearchTopologyWithRange(query, 0, 0)
}

// CheckCompliance Check that a device is in compliance, that is the configlets
// applied to the device match the devices running configuration.
//
// Supported only for NetElements
//
func (c CvpRestAPI) CheckCompliance(nodeKey string, nodeType string) (*ComplianceResp, error) {
	var info ComplianceResp
	data := map[string]string{
		"nodeId":   nodeKey,
		"nodeType": nodeType,
	}

	resp, err := c.client.Post("/provisioning/checkCompliance.do", nil, data)
	if err != nil {
		return nil, errors.Errorf("CheckCompliance: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("CheckCompliance: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("CheckCompliance: %s", err)
	}

	return &info, nil
}

// GetParentContainerForDevice returns the Container for specified deviceMAC
func (c CvpRestAPI) GetParentContainerForDevice(deviceMAC string) (*Container, error) {
	results, err := c.SearchTopologyWithRange(deviceMAC, 0, 0)
	if err != nil {
		return nil, errors.Errorf("GetParentContainerForDevice: %s", err)
	}
	for _, netContainerInfo := range results.NetElementContainerList {
		if netContainerInfo.NetElementKey == deviceMAC {
			return c.GetContainerByName(netContainerInfo.ContainerName)
		}
	}
	return nil, nil
}

// MoveDeviceToContainer moves a specified netelement to a container.
func (c CvpRestAPI) MoveDeviceToContainer(device *NetElement, container *Container,
	commit bool) (*TaskInfo, error) {
	if device == nil {
		return nil, errors.Errorf("MoveDeviceToContainer: nil NetElement")
	}
	if container == nil {
		return nil, errors.Errorf("MoveDeviceToContainer: nil Container")
	}

	var fromID string
	if device.ParentContainerID != "" {
		fromID = device.ParentContainerID
	} else {
		parentCont, err := c.GetParentContainerForDevice(device.SystemMacAddress)
		if err != nil {
			return nil, errors.Errorf("MoveDeviceToContainer: %s", err)
		}
		if parentCont == nil {
			return nil, errors.Errorf("MoveDeviceToContainer: No parent container found for "+
				"device [%s]", device.SystemMacAddress)
		}
		fromID = parentCont.Key
	}

	msg := "Moving device " + device.Fqdn + " from container " + fromID +
		" to container " + container.Name

	data := &ActionRequest{Data: []Action{
		Action{
			Info:        msg,
			InfoPreview: msg,
			Action:      "update",
			NodeType:    "netelement",
			NodeID:      device.Key,
			ToID:        container.Key,
			ToName:      container.Name,
			ToIDType:    "container",
			FromID:      fromID,
			NodeName:    device.Fqdn,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("MoveDeviceToContainer: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ImageInfo represents information related to an Image within CVP
type ImageInfo struct {
	AppliedContainersCount   int    `json:"appliedContainersCount"`
	AppliedDevicesCount      int    `json:"appliedDevicesCount"`
	FactoryID                int    `json:"factoryId"`
	ID                       int    `json:"id"`
	ImageID                  string `json:"imageId"`
	ImageFile                string `json:"imageFile"`
	ImageFileName            string `json:"imageFileName"`
	ImageSize                string `json:"imageSize"`
	IsHotFix                 string `json:"isHotFix"`
	IsRebootRequired         string `json:"isRebootRequired"`
	Key                      string `json:"key"`
	MD5                      string `json:"md5"`
	SHA512                   string `json:"sha512"`
	Name                     string `json:"name"`
	SwiMaxHwepoch            string `json:"swiMaxHwepoch"`
	SwiVarient               string `json:"swiVarient"`
	UploadedDateinLongFormat int64  `json:"uploadedDateinLongFormat"`
	User                     string `json:"user"`
	Version                  string `json:"version"`
}

// ImageResp response from Image request
type ImageResp struct {
	Total int         `json:"total"`
	Data  []ImageInfo `json:"data"`

	ErrorResponse
}

// ImageBundleInfo represents ImageBundle object within CVP
type ImageBundleInfo struct {
	AppliedContainersCount   int         `json:"appliedContainersCount"`
	AppliedDevicesCount      int         `json:"appliedDevicesCount"`
	FactoryID                int         `json:"factoryId"`
	ID                       int         `json:"id"`
	IsCertifiedImageBundle   string      `json:"isCertifiedImageBundle"`
	ImageIds                 []string    `json:"imageIds"`
	Images                   []ImageInfo `json:"images,omitempty"`
	Key                      string      `json:"key"`
	Name                     string      `json:"name"`
	Note                     string      `json:"note"`
	UploadedBy               string      `json:"uploadedBy,omitempty"`
	UploadedDateinLongFormat int64       `json:"uploadedDateinLongFormat"`
	User                     string      `json:"user"`

	ErrorResponse
}

// ImageBundleResp response data/payload for ImageBundle query
type ImageBundleResp struct {
	Total                 int                   `json:"total"`
	Data                  []ImageBundleInfo     `json:"data"`
	ImageBundleMapper     map[string]*ImageInfo `json:"imageBundleMapper"`
	AssignedImageBundleID string                `json:"assignedImageBundleId"`

	ErrorResponse
}

// GetImages returns a list of Images based on a specific query string and range
func (c CvpRestAPI) GetImages(querystr string, start int, end int) ([]ImageInfo, error) {
	var resp ImageResp
	query := &url.Values{
		"queryParam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/image/getImages.do", query)
	if err != nil {
		return nil, errors.Errorf("GetImages: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetImages: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetImages: %s", err)
	}
	return resp.Data, nil
}

// GetImageByName returns an ImageInfo object based on name provided
func (c CvpRestAPI) GetImageByName(name string) (*ImageInfo, error) {
	resp, err := c.GetImages(name, 0, 0)
	if err != nil {
		return nil, errors.Errorf("GetImageByName: %s", err)
	}

	for _, image := range resp {
		if image.Name == name {
			return &image, nil
		}
	}
	return nil, nil
}

// GetImageBundles returns a list of ImageBundles based on a specific query string and range
func (c CvpRestAPI) GetImageBundles(querystr string, start, end int) ([]ImageBundleInfo, error) {
	var resp ImageBundleResp
	query := &url.Values{
		"queryParam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/image/getImageBundles.do", query)
	if err != nil {
		return nil, errors.Errorf("GetImageBundles: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetImageBundles: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetImageBundles: %s", err)
	}
	return resp.Data, nil

}

// GetAllImageBundles gets all ImageBundles
func (c CvpRestAPI) GetAllImageBundles() ([]ImageBundleInfo, error) {
	return c.GetImageBundles("", 0, 0)
}

// GetImageBundleByName gets ImageBundle by specified name
func (c CvpRestAPI) GetImageBundleByName(name string) (*ImageBundleInfo, error) {
	// Hack around string returned for ID
	type tmp struct {
		AppliedContainersCount   int         `json:"appliedContainersCount"`
		AppliedDevicesCount      int         `json:"appliedDevicesCount"`
		FactoryID                int         `json:"factoryId"`
		ID                       string      `json:"id"`
		IsCertifiedImageBundle   string      `json:"isCertifiedImageBundle"`
		ImageIds                 []string    `json:"imageIds"`
		Images                   []ImageInfo `json:"images,omitempty"`
		Key                      string      `json:"key"`
		Name                     string      `json:"name"`
		Note                     string      `json:"note"`
		UploadedBy               string      `json:"uploadedBy,omitempty"`
		UploadedDateinLongFormat int64       `json:"uploadedDateinLongFormat"`
		User                     string      `json:"user"`
		ErrorResponse
	}
	var resp tmp

	query := &url.Values{
		"name": {name},
	}

	reqResp, err := c.client.Get("/image/getImageBundleByName.do", query)
	if err != nil {
		return nil, errors.Errorf("GetImageBundleByName: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetImageBundleByName: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetImageBundleByName: %s", err)
	}
	ret := &ImageBundleInfo{
		AppliedContainersCount:   resp.AppliedContainersCount,
		AppliedDevicesCount:      resp.AppliedDevicesCount,
		FactoryID:                resp.FactoryID,
		ID:                       1,
		IsCertifiedImageBundle:   resp.IsCertifiedImageBundle,
		ImageIds:                 resp.ImageIds,
		Images:                   resp.Images,
		Key:                      resp.Key,
		Name:                     resp.Name,
		Note:                     resp.Note,
		UploadedBy:               resp.UploadedBy,
		UploadedDateinLongFormat: resp.UploadedDateinLongFormat,
		User:                     resp.User,
	}
	return ret, nil

}

// ApplyImageToDevice Applies image bundle to device
func (c CvpRestAPI) ApplyImageToDevice(imageInfo *ImageBundleInfo, netElement *NetElement,
	commit bool) (*TaskInfo, error) {
	if imageInfo == nil {
		return nil, errors.Errorf("ApplyImageToDevice: nil ImageBundleInfo")
	}
	if netElement == nil {
		return nil, errors.Errorf("ApplyImageToDevice: nil NetElement")
	}

	msg := "Apply image " + imageInfo.Name + " to NetElement " + netElement.Fqdn

	data := &ActionRequest{Data: []Action{
		Action{
			Info:        msg,
			InfoPreview: msg,
			Note:        "",
			Action:      "associate",
			NodeType:    "imagebundle",
			NodeID:      strconv.Itoa(imageInfo.ID),
			ToID:        netElement.Key,
			ToIDType:    "netelement",
			FromID:      "",
			NodeName:    imageInfo.Name,
			FromName:    "",
			ToName:      netElement.Fqdn,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyImageToDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ApplyImageToContainer Applies image bundle to container
func (c CvpRestAPI) ApplyImageToContainer(imageInfo *ImageBundleInfo, container *Container,
	commit bool) (*TaskInfo, error) {
	if imageInfo == nil {
		return nil, errors.Errorf("ApplyImageToContainer: nil ImageBundleInfo")
	}
	if container == nil {
		return nil, errors.Errorf("ApplyImageToContainer: nil Container")
	}

	msg := "Apply image " + imageInfo.Name + " to Container " + container.Name
	data := &ActionRequest{Data: []Action{
		Action{
			ID:          1,
			Info:        msg,
			InfoPreview: msg,
			Note:        "",
			Action:      "associate",
			NodeType:    "imagebundle",
			NodeID:      strconv.Itoa(imageInfo.ID),
			ToID:        container.Key,
			ToIDType:    "container",
			FromID:      "",
			NodeName:    imageInfo.Name,
			FromName:    "",
			ToName:      container.Name,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyImageToContainer: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// RemoveImageFromContainer removes image bundle from container
func (c CvpRestAPI) RemoveImageFromContainer(imageInfo *ImageBundleInfo,
	container *Container) (*TaskInfo, error) {
	if imageInfo == nil {
		return nil, errors.Errorf("RemoveImageFromContainer: nil ImageBundleInfo")
	}
	if container == nil {
		return nil, errors.Errorf("RemoveImageFromContainer: nil Container")
	}

	msg := "Remove image " + imageInfo.Name + " from Container " + container.Name

	data := &ActionRequest{Data: []Action{
		Action{
			ID:             1,
			Info:           msg,
			InfoPreview:    msg,
			Note:           "",
			Action:         "associate",
			NodeType:       "imagebundle",
			NodeID:         "",
			ToID:           container.Key,
			ToIDType:       "container",
			FromID:         "",
			NodeName:       "",
			FromName:       "",
			ToName:         container.Name,
			IgnoreNodeID:   strconv.Itoa(imageInfo.ID),
			IgnoreNodeName: imageInfo.Name,
		},
	}}
	c.addTempAction(data)
	return c.SaveTopology()
}

// DeployDevice Move a device from the undefined container to a target container.
// Optionally, apply device-specific configlets to the device.
func (c CvpRestAPI) DeployDevice(netElement *NetElement, container *Container,
	configlets ...Configlet) (*TaskInfo, error) {
	return c.DeployDeviceWithImage(netElement, container, "", configlets...)
}

// DeployDeviceWithImage Move a device from the undefined container to a target container
// and apply image. Optionally, apply device-specific configlets to the device.
func (c CvpRestAPI) DeployDeviceWithImage(netElement *NetElement, container *Container,
	image string, configlets ...Configlet) (*TaskInfo, error) {
	if _, err := c.MoveDeviceToContainer(netElement, container, false); err != nil {
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	conf, err := c.GetTempConfigByNetElementID(netElement.SystemMacAddress)
	applyConfiglets := conf.ProposedConfiglets
	if configlets != nil {
		applyConfiglets = append(applyConfiglets, configlets...)
	}

	if _, err = c.ApplyConfigletsToDevice("DeployDevice", netElement, false,
		applyConfiglets...); err != nil {
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	if image != "" {
		imageBundle, err := c.GetImageBundleByName(image)
		if err != nil {
			return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
		}
		if _, err = c.ApplyImageToDevice(imageBundle, netElement, false); err != nil {
			return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
		}
	}
	return c.SaveTopology()
}

// TempConfig ...
type TempConfig struct {
	ExistingConfiglets        []string    `json:"existingConfiglets"`
	IgnoredConfiglets         []string    `json:"ignoredConfiglets"`
	AssignedConfiglets        []string    `json:"assignedConfiglets"`
	ProposedConfiglets        []Configlet `json:"proposedConfiglets"`
	DeviceConfigletBuilders   []string    `json:"deviceConfigletBuilders"`
	AssignedConfigletBuilders []string    `json:"assignedConfigletBuilders"`
	IgnoredConfigletBuilders  []string    `json:"ignoredConfigletBuilders"`
	DeviceConfiglets          []string    `json:"deviceConfiglets"`
	ExistingConfigletBuilders []string    `json:"existingConfigletBuilders"`

	ErrorResponse
}

// GetTempConfigByNetElementID gets the current temporary config for the supplied netElement
func (c CvpRestAPI) GetTempConfigByNetElementID(netElementID string) (*TempConfig, error) {
	var resp TempConfig
	query := &url.Values{
		"netElementId": {netElementID},
	}

	reqResp, err := c.client.Get("/provisioning/getTempConfigsByNetElementId.do", query)
	if err != nil {
		return nil, errors.Errorf("GetTempConfigByNetElementID: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetTempConfigByNetElementID: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetTempConfigByNetElementID: %s", err)
	}
	return &resp, nil

}

// GetAllTempActions gets the list of current actions outstanding
func (c CvpRestAPI) GetAllTempActions(start, end int) ([]Action, error) {
	var resp struct {
		Total int
		Data  []Action
		ErrorResponse
	}

	query := &url.Values{
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/provisioning/getAllTempActions.do", query)
	if err != nil {
		return nil, errors.Errorf("GetTempActions: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetTempActions: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetTempActions: %s", err)
	}
	return resp.Data, nil

}

// GetTempAction returns the first outstanding action
func (c CvpRestAPI) GetTempAction() (*Action, error) {
	results, err := c.GetAllTempActions(0, 1)
	if err != nil {
		return nil, errors.Errorf("GetTempAction: %s", err)
	}
	if len(results) > 0 {
		return &results[0], nil
	}
	return nil, nil
}

// checkConfigMapping Checks whether the new configlets to be applied are
// already applied or not. Returns actionReqd ( bool ), configletMap,
// and builderMap.
func checkConfigMapping(applied []Configlet, newconfiglets []Configlet) (bool,
	map[string]string, map[string]string, error) {
	builderMap := make(map[string]string)
	configletMap := make(map[string]string)
	for _, configlet := range applied {
		switch configlet.Type {
		case "Static":
			fallthrough
		case "Generated":
			fallthrough
		case "Reconciled":
			configletMap[configlet.Key] = configlet.Name
		case "Builder":
			builderMap[configlet.Key] = configlet.Name
		default:
			return false, nil, nil, errors.Errorf("Configlet [%s] Invalid Type [%s]",
				configlet.Name, configlet.Type)
		}
	}

	var actionReqd bool
	for _, configlet := range newconfiglets {
		if _, found := configletMap[configlet.Key]; found {
			continue
		}
		if _, found := builderMap[configlet.Key]; found {
			continue
		}
		// didn't find this configlet in either maps, so it's new
		actionReqd = true

		switch configlet.Type {
		case "Static":
			fallthrough
		case "Generated":
			fallthrough
		case "Reconciled":
			configletMap[configlet.Key] = configlet.Name
		case "Builder":
			builderMap[configlet.Key] = configlet.Name
		default:
			return false, nil, nil, errors.Errorf("Configlet [%s] Invalid Type [%s]",
				configlet.Name, configlet.Type)
		}
	}
	return actionReqd, configletMap, builderMap, nil
}

// checkRemoveConfigMapping Creates map of configlets that needs to be there after removal of
// specific configlets
func checkRemoveConfigMapping(applied []Configlet, rmConfiglets []Configlet) (bool,
	map[string]string, map[string]string, map[string]string, map[string]string, error) {
	rmBuilderMap := make(map[string]string)
	rmConfigletMap := make(map[string]string)

	for _, configlet := range rmConfiglets {
		switch configlet.Type {
		case "Static":
			fallthrough
		case "Generated":
			fallthrough
		case "Reconciled":
			rmConfigletMap[configlet.Key] = configlet.Name
		case "Builder":
			rmBuilderMap[configlet.Key] = configlet.Name
		default:
			return false, nil, nil, nil, nil,
				errors.Errorf("Invalid Configlet Type [%s]", configlet.Type)
		}
	}

	var actionReqd bool
	configletMap := make(map[string]string)
	builderMap := make(map[string]string)
	for _, configlet := range applied {
		if _, found := rmConfigletMap[configlet.Key]; found {
			continue
		}
		if _, found := rmBuilderMap[configlet.Key]; found {
			continue
		}
		// didn't find this configlet in either maps, so it's new
		actionReqd = true

		switch configlet.Type {
		case "Static":
			fallthrough
		case "Generated":
			fallthrough
		case "Reconciled":
			configletMap[configlet.Key] = configlet.Name
		case "Builder":
			builderMap[configlet.Key] = configlet.Name
		default:
			return false, nil, nil, nil, nil,
				errors.Errorf("Invalid Configlet Type [%s]", configlet.Type)
		}
	}

	return actionReqd, configletMap, builderMap, rmConfigletMap, rmBuilderMap, nil
}

// keyValueSliceFromMap returns two string slices, one for keys in the map provided,
// and one for the values.
func keyValueSliceFromMap(m map[string]string) ([]string, []string) {
	keys := []string{}
	values := []string{}
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}
