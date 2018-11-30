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

package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	cvprac "gopkg.in/aristanetworks/go-cvprac.v2"
	"gopkg.in/aristanetworks/go-cvprac.v2/api"

	"net/http"

	resty "gopkg.in/resty.v1"
)

// NumRetryRequests specifies the number for retries to attempt
const NumRetryRequests = 3

// UNDEFPORT undefined port
const UNDEFPORT = -1

type authInfo struct {
	Username string
	Password string
}

// CvpClient represents a CVP client api connection
type CvpClient struct {
	cvpapi.ClientInterface
	Hosts    []string
	HostPool *HostIterator
	Port     int
	Protocol string
	authInfo *authInfo
	Timeout  time.Duration
	Client   *resty.Client
	Headers  map[string]string
	SessID   string
	url      string
	API      *cvpapi.CvpRestAPI
	Debug    bool
}

// Option is a Client Option...function that sets a value and returns
type Option func(*CvpClient) error

// Hosts sets the available host ip/names to use for this
// Client
func Hosts(hosts ...string) Option {
	return func(c *CvpClient) error {
		var err error
		if len(hosts) <= 0 {
			return errors.New("Must define at least one host")
		}
		c.Hosts = hosts
		c.HostPool, err = NewHostIterator(c.Hosts)
		if err != nil {
			return err
		}
		return nil
	}
}

// Port sets the port option for this Client
func Port(port int) Option {
	return func(c *CvpClient) error {
		if port <= 0 {
			return errors.Errorf("Invalid port number [%d]", port)
		}
		c.Port = port
		return nil
	}
}

// Protocol sets the protocol option for this Client
func Protocol(proto string) Option {
	return func(c *CvpClient) error {
		switch proto {
		case "http":
		case "https":
		default:
			return errors.Errorf("Invalid protocol [%s]", proto)
		}
		c.Protocol = proto
		return nil
	}
}

// ConnectTimeout sets the connection timeout for this Client
func ConnectTimeout(timeout int) Option {
	return func(c *CvpClient) error {
		if timeout < 0 {
			return errors.New("Timeout (seconds) must be >= 0")
		}
		c.Timeout = time.Duration(timeout) * time.Second
		return nil
	}
}

// Debug sets the debug option for this Client
func Debug(enable bool) Option {
	return func(c *CvpClient) error {
		c.Debug = enable
		return nil
	}
}

// SetOption takes one or more option function and applies them in order
func (c *CvpClient) SetOption(options ...Option) error {
	for _, opt := range options {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// SetHosts sets the hosts
func (c *CvpClient) SetHosts(host ...string) error {
	return c.SetOption(Hosts(host...))
}

// SetPort sets the port for this connection
func (c *CvpClient) SetPort(port int) error {
	return c.SetOption(Port(port))
}

// SetProtocol sets the protocol (i.e. http or https) associated with this
// connection
func (c *CvpClient) SetProtocol(proto string) error {
	return c.SetOption(Protocol(proto))
}

// SetConnectTimeout sets the connection timeout associated with this
// connection
func (c *CvpClient) SetConnectTimeout(timeout int) error {
	return c.SetOption(ConnectTimeout(timeout))
}

// SetDebug enables or disables debugging.
func (c *CvpClient) SetDebug(enable bool) error {
	return c.SetOption(Debug(enable))
}

// NewCvpClient creates a new CVP RESTful Client
func NewCvpClient(options ...Option) (*CvpClient, error) {
	var err error
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   "Golang cvprac/" + cvprac.Version,
	}
	c := &CvpClient{
		Headers:  headers,
		Port:     UNDEFPORT,
		Protocol: "https",
		Timeout:  time.Duration(60 * time.Second),
		Hosts:    []string{"localhost"},
	}

	if err := c.SetOption(options...); err != nil {
		return nil, err
	}
	c.HostPool, err = NewHostIterator(c.Hosts)
	if err != nil {
		return nil, err
	}

	c.Client = resty.New().SetHeaders(headers)

	c.API = cvpapi.NewCvpRestAPI(c)

	return c, nil
}

// GetSessionID returns the current Session ID
func (c *CvpClient) GetSessionID() string {
	return c.SessID
}

func (c *CvpClient) initSession(host string) error {
	var port int

	c.Client.SetHeaders(c.Headers)

	if c.Protocol == "https" {
		port = 443
		c.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	} else {
		port = 80
	}

	if c.Port != UNDEFPORT {
		port = c.Port
	}

	// formulate and set our base url
	c.url = fmt.Sprintf("%s://%s:%d/web", c.Protocol, host, port)
	c.Client.SetHostURL(c.url)
	c.Client.SetTimeout(c.Timeout)
	c.Client.SetDebug(c.Debug)
	return nil
}

// Connect Login to CVP and get a session ID and cookie.
func (c *CvpClient) Connect(username string, password string) error {
	c.authInfo = &authInfo{username, password}

	if err := c.createSession(true); err != nil {
		return err
	}
	return nil
}

func (c *CvpClient) createSession(allNodes bool) error {
	var err error
	var errorMsg []string

	numNodes := len(c.Hosts)
	if !allNodes && numNodes > 1 {
		numNodes--
	}
	for nodeIter := 0; nodeIter < numNodes; nodeIter++ {
		host := c.HostPool.Cycle()

		c.initSession(host)
		if err = c.login(); err != nil {
			tmpMsg := fmt.Sprintf("createSession: Host %s Error: %s", host, err.Error())
			errorMsg = append(errorMsg, tmpMsg)
		} else {
			return nil
		}
	}
	return errors.New(strings.Join(errorMsg, "\n"))
}

func (c *CvpClient) login() error {
	var loginResp cvpapi.LoginResp

	c.SessID = ""
	request := c.Client.R()

	auth := "{\"userId\":\"" + c.authInfo.Username +
		"\", \"password\":\"" + c.authInfo.Password + "\"}"

	resp, err := request.SetBody(auth).Post("/login/authenticate.do")
	if err != nil {
		return errors.Wrapf(err, "login")
	}

	if err = checkResponse(resp); err != nil {
		return errors.Wrapf(err, "checkResponse failed")
	}

	if err = json.Unmarshal(resp.Body(), &loginResp); err != nil {
		return errors.Wrapf(err, "unmarshal failed")
	}
	c.SessID = loginResp.SessionID

	return nil
}

func (c *CvpClient) makeRequest(reqType string, url string, params *url.Values,
	data interface{}) ([]byte, error) {
	var err error
	var resp *resty.Response
	var formattedParams map[string]string

	retryCnt := NumRetryRequests

	if params != nil {
		formattedParams, err = parseURLValues(params)
		if err != nil {
			return nil, err
		}
	}

	nodeCnt := len(c.Hosts)
	for nodeCnt > 0 {

		request := c.Client.R()
		request.SetQueryParams(formattedParams)

		// If we've seen an error
		if err != nil {
			// Decrement count as another node will be tried, if there
			// are no more nodes then return error.
			nodeCnt--
			if nodeCnt == 0 {
				return nil, err
			}
			// Not the first time through the loop. Retrying request so
			// create a session to another CVP node...but exclude this one.
			if err := c.createSession(false); err != nil {
				return nil, err
			}
			retryCnt = NumRetryRequests
			err = nil
		}

		// Check reqType
		switch reqType {
		case "GET":
			resp, err = request.Get(url)
		case "POST":
			resp, err = request.SetBody(data).Post(url)
		default:
			return nil, errors.Errorf("Invalid. Request type [%s] not implemented", reqType)
		}

		if err != nil {
			return nil, err
		}
		// Underlying request issue. Could be getsockopt error (like network not reachable)
		if resp.RawResponse == nil {
			// retry another session
			err = errors.New("RawResponse error")
			continue
		}

		status := resp.StatusCode()

		if status == 301 {
			// retry another session
			err = errors.Errorf("Status: %d", status)
			continue
		}
		// From 2018.2.0 onwards, a '401' response is returned for
		// an 'Unauthorized' user, unlike previous releases where a
		// CvpError with code '112498' was returned. This might happen
		// if the session expires, which is after 12 hours of inactivity.
		// In this case, the session must be refreshed. Retry same host.
		if status == 401 {
			retryCnt--
			if retryCnt <= 0 {
				err = errors.Errorf("Status: %d", status)
			}
			continue
		}
		// client error
		if status != http.StatusOK {
			// retry another session
			err = errors.Errorf("Status: %d", status)
			continue
		}

		var info cvpapi.ErrorResponse
		if err = json.Unmarshal(resp.Body(), &info); err == nil {
			// check and see if we have a CVP error payload
			if err = info.Error(); err != nil {
				return nil, err
			}
		}
		break
	}
	return resp.Body(), nil
}

// Get implemented as part of cvprac api client interface
func (c *CvpClient) Get(url string, params *url.Values) ([]byte, error) {
	return c.makeRequest("GET", url, params, nil)
}

// Post implemented as part of cvprac api client interface
func (c *CvpClient) Post(url string, params *url.Values, data interface{}) ([]byte, error) {
	return c.makeRequest("POST", url, params, data)
}

func parseURLValues(params *url.Values) (map[string]string, error) {
	newMap := make(map[string]string)
	for k, v := range *params {
		if len(v) > 1 {
			return nil, errors.Errorf("Parsing URL Values: Multiple values for "+
				"param %s. Values: %v", k, v)
		}
		newMap[k] = params.Get(k)
	}
	return newMap, nil
}

func checkResponse(resp *resty.Response) error {
	// Underlying request issue. Could be getsockopt error (like network not reachable)
	if resp.RawResponse == nil {
		// retry another session
		return errors.New("RawResponse error")
	}

	status := resp.StatusCode()

	if status != http.StatusOK {
		return errors.Errorf("checkResponse Status: %s, StatusCode: %d", resp.Status(), status)
	}

	var info cvpapi.ErrorResponse
	if err := json.Unmarshal(resp.Body(), &info); err != nil {
		return errors.Wrapf(err, "checkResponse unmarshal error")
	}

	if err := info.Error(); err != nil {
		return errors.Wrapf(err, "checkResponse Request error")
	}
	return nil
}
