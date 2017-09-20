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

	ID         int      `json:"id,omitempty"`
	Note       string   `json:"note,omitempty"`
	ChildTasks []string `json:"childTasks,omitempty"`
	ParentTask string   `json:"parentTask,omitempty"`
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

// ConfigletInfo represents basic info related to a Configlet
type ConfigletInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// DeviceConfiglets represents the configlets for a netelemet
type DeviceConfiglets struct {
	Total           int                      `json:"total"`
	ConfigletMapper map[string]ConfigletInfo `json:"configletMapper"`
	ConfigletList   []Configlet              `json:"configletList"`

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

// GetConfigletsByDeviceID returns the list of configlets applied to a device.
func (c CvpRestAPI) GetConfigletsByDeviceID(mac string) ([]Configlet, error) {
	var info DeviceConfiglets
	query := &url.Values{
		"netElementId": {mac},
		"queryParam":   {""},
		"startIndex":   {"0"},
		"endIndex":     {"0"},
	}

	resp, err := c.client.Get("/provisioning/getConfigletsByNetElementId.do", query)
	if err != nil {
		return nil, fmt.Errorf("GetConfigletsByDeviceID: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("GetConfigletsByDeviceID: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, fmt.Errorf("GetConfigletsByDeviceID: %s", err)
	}

	return info.ConfigletList, nil
}

func (c CvpRestAPI) addTempAction(data ActionRequest) error {
	var resp ErrorResponse

	query := &url.Values{
		"format":     {"topology"},
		"queryParam": {""},
		"nodeId":     {"root"},
	}

	reqResp, err := c.client.Post("/ztp/addTempAction.do", query, data)
	if err != nil {
		return fmt.Errorf("addTempAction: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return fmt.Errorf("addTempAction: %s", err)
	}

	if err := resp.Error(); err != nil {
		return fmt.Errorf("addTempAction: %s", err)
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
		return nil, fmt.Errorf("SaveTopology: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, fmt.Errorf("SaveTopology: %s", err)
	}

	return &resp.Data, nil
}

// ApplyConfigletsToDevice apply the configlets to the device.
func (c CvpRestAPI) ApplyConfigletsToDevice(appName string, dev *NetElement, commit bool,
	newConfiglets ...Configlet) (*TaskInfo, error) {

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, fmt.Errorf("ApplyConfigletsToDevice: %s", err)
	}

	// Get a list of the names and keys of the configlets
	var cnames []string
	var ckeys []string

	for _, configlet := range configlets {
		cnames = append(cnames, configlet.Name)
		ckeys = append(ckeys, configlet.Key)
	}

	// Add the new configlets to the end of the arrays
	for _, entry := range newConfiglets {
		cnames = append(cnames, entry.Name)
		ckeys = append(ckeys, entry.Key)
	}

	info := appName + ": Configlet Assign: to Device " + dev.Fqdn
	infoPreview := "<b>Configlet Assign:</b> to Device" + dev.Fqdn

	data := ActionRequest{Data: []Action{
		Action{
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            []string{},
			ConfigletBuilderNamesList:       []string{},
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
		return nil, fmt.Errorf("ApplyConfigletsToDevice: %s", err)
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

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, fmt.Errorf("RemoveConfigletsFromDevice: %s", err)
	}

	// Get a list of the names/keys of configlets to keep
	// Do not add configlets from remConfiglets
	var keepKeys []string
	var keepNames []string
	var found bool
	// for each configlet applied to device
	for _, configlet := range configlets {
		found = false
		// look at our list of configlets to remove
		for _, rconfiglet := range remConfiglets {
			// if the configlet is part of remove list
			// skip it.
			if configlet.Key == rconfiglet.Key {
				found = true
				break
			}
		}
		if !found {
			keepNames = append(keepNames, configlet.Name)
			keepKeys = append(keepKeys, configlet.Key)
		}
	}

	// Build a list of the configlet names/keys to remove.
	var delKeys []string
	var delNames []string

	// Add the new configlets to the end of the arrays
	for _, entry := range remConfiglets {
		delNames = append(delNames, entry.Name)
		delKeys = append(delKeys, entry.Key)
	}

	info := appName + ": Configlet Remove: from Device " + dev.Fqdn
	infoPreview := "<b>Configlet Remove:</b> from Device" + dev.Fqdn

	data := ActionRequest{Data: []Action{
		Action{
			ID:                              1,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletList:                   keepKeys,
			ConfigletNamesList:              keepNames,
			ConfigletBuilderList:            []string{},
			ConfigletBuilderNamesList:       []string{},
			IgnoreConfigletList:             delKeys,
			IgnoreConfigletNamesList:        delNames,
			IgnoreConfigletBuilderList:      []string{},
			IgnoreConfigletBuilderNamesList: []string{},
			ToID:                dev.SystemMacAddress,
			ToIDType:            "netelement",
			FromID:              "",
			NodeName:            "",
			NodeIPAddress:       dev.IPAddress,
			NodeTargetIPAddress: dev.IPAddress,
			FromName:            "",
			ToName:              dev.Fqdn,
			ChildTasks:          []string{},
			ParentTask:          "",
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, fmt.Errorf("RemoveConfigletsFromDevice: %s", err)
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

	data := ActionRequest{Data: []Action{
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
		return nil, fmt.Errorf("SearchTopologyWithRange: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, fmt.Errorf("SearchTopologyWithRange: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, fmt.Errorf("SearchTopologyWithRange: %s", err)
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
		return nil, fmt.Errorf("CheckCompliance: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, fmt.Errorf("CheckCompliance: %s", err)
	}

	if err := info.Error(); err != nil {
		return nil, fmt.Errorf("CheckCompliance: %s", err)
	}

	return &info, nil
}

// GetParentContainerForDevice returns the Container for specified deviceMAC
func (c CvpRestAPI) GetParentContainerForDevice(deviceMAC string) (*Container, error) {
	results, err := c.SearchTopologyWithRange(deviceMAC, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("GetParentContainerForDevice: %s", err)
	}
	if results.Total > 0 {
		name := results.NetElementContainerList[0].ContainerName
		return c.GetContainerByName(name)
	}
	return nil, nil
}

// MoveDeviceToContainer moves a specified netelement to a container.
func (c CvpRestAPI) MoveDeviceToContainer(device *NetElement, container *Container,
	commit bool) (*TaskInfo, error) {
	var fromID string
	if device.ParentContainerID != "" {
		fromID = device.ParentContainerID
	} else {
		parentCont, err := c.GetParentContainerForDevice(device.SystemMacAddress)
		if err != nil {
			return nil, fmt.Errorf("MoveDeviceToContainer: %s", err)
		}
		fromID = parentCont.Key
	}

	msg := "Moving device " + device.Fqdn + " from container " + fromID +
		" to container " + container.Name

	data := ActionRequest{Data: []Action{
		Action{
			ID:          1,
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
		return nil, fmt.Errorf("MoveDeviceToContainer: %s", err)
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
		return nil, fmt.Errorf("GetImages: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, fmt.Errorf("GetImages: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, fmt.Errorf("GetImages: %s", err)
	}
	return resp.Data, nil
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
		return nil, fmt.Errorf("GetImageBundles: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, fmt.Errorf("GetImageBundles: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, fmt.Errorf("GetImageBundles: %s", err)
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
		return nil, fmt.Errorf("GetImageBundleByName: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, fmt.Errorf("GetImageBundleByName: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, fmt.Errorf("GetImageBundleByName: %s", err)
	}
	ret := &ImageBundleInfo{
		AppliedContainersCount: resp.AppliedContainersCount,
		AppliedDevicesCount:    resp.AppliedDevicesCount,
		FactoryID:              resp.FactoryID,
		ID:                     1,
		IsCertifiedImageBundle:   resp.IsCertifiedImageBundle,
		ImageIds:                 resp.ImageIds,
		Images:                   resp.Images,
		Key:                      resp.Key,
		Name:                     resp.Name,
		Note:                     resp.Note,
		UploadedBy:               resp.UploadedBy,
		UploadedDateinLongFormat: resp.UploadedDateinLongFormat,
		User: resp.User,
	}
	return ret, nil

}

// ApplyImageToDevice Applies image bundle to device
func (c CvpRestAPI) ApplyImageToDevice(imageInfo *ImageBundleInfo, netElement *NetElement,
	commit bool) (*TaskInfo, error) {
	msg := "Apply image " + imageInfo.Name + " to NetElement " + netElement.Fqdn

	data := ActionRequest{Data: []Action{
		Action{
			ID:          1,
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
		return nil, fmt.Errorf("ApplyImageToDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ApplyImageToContainer Applies image bundle to container
func (c CvpRestAPI) ApplyImageToContainer(imageInfo *ImageBundleInfo, container *Container,
	commit bool) (*TaskInfo, error) {
	msg := "Apply image " + imageInfo.Name + " to Container " + container.Name
	data := ActionRequest{Data: []Action{
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
		return nil, fmt.Errorf("ApplyImageToContainer: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// RemoveImageFromContainer removes image bundle from container
func (c CvpRestAPI) RemoveImageFromContainer(imageInfo *ImageBundleInfo,
	container *Container) (*TaskInfo, error) {
	msg := "Remove image " + imageInfo.Name + " from Container " + container.Name

	data := ActionRequest{Data: []Action{
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
		return nil, fmt.Errorf("DeployDeviceWithImage: %s", err)
	}

	conf, err := c.GetTempConfigByNetElementID(netElement.SystemMacAddress)
	applyConfiglets := conf.ProposedConfiglets
	if configlets != nil {
		applyConfiglets = append(applyConfiglets, configlets...)
	}

	if _, err = c.ApplyConfigletsToDevice("DeployDevice", netElement, false,
		applyConfiglets...); err != nil {
		return nil, fmt.Errorf("DeployDeviceWithImage: %s", err)
	}

	if image != "" {
		imageBundle, err := c.GetImageBundleByName(image)
		if err != nil {
			return nil, fmt.Errorf("DeployDeviceWithImage: %s", err)
		}
		if _, err = c.ApplyImageToDevice(imageBundle, netElement, false); err != nil {
			return nil, fmt.Errorf("DeployDeviceWithImage: %s", err)
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
		return nil, fmt.Errorf("GetTempConfigByNetElementID: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, fmt.Errorf("GetTempConfigByNetElementID: %s", err)
	}

	if err := resp.Error(); err != nil {
		return nil, fmt.Errorf("GetTempConfigByNetElementID: %s", err)
	}
	return &resp, nil

}
