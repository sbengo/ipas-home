// Copyright 2015 Vadim Kravcenko
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package gojenkins

import "errors"

// Nodes

type Computers struct {
	BusyExecutors  int            `json:"busyExecutors"`
	Computers      []nodeResponse `json:"computer"`
	DisplayName    string         `json:"displayName"`
	TotalExecutors int            `json:"totalExecutors"`
}

type Node struct {
	Raw     *nodeResponse
	Jenkins *Jenkins
	Base    string
}

type nodeResponse struct {
	Actions             []interface{} `json:"actions"`
	DisplayName         string        `json:"displayName"`
	Executors           []struct{}    `json:"executors"`
	Icon                string        `json:"icon"`
	IconClassName       string        `json:"iconClassName"`
	Idle                bool          `json:"idle"`
	JnlpAgent           bool          `json:"jnlpAgent"`
	LaunchSupported     bool          `json:"launchSupported"`
	LoadStatistics      struct{}      `json:"loadStatistics"`
	ManualLaunchAllowed bool          `json:"manualLaunchAllowed"`
	MonitorData         struct {
		Hudson_NodeMonitors_ArchitectureMonitor interface{} `json:"hudson.node_monitors.ArchitectureMonitor"`
		Hudson_NodeMonitors_ClockMonitor        interface{} `json:"hudson.node_monitors.ClockMonitor"`
		Hudson_NodeMonitors_DiskSpaceMonitor    interface{} `json:"hudson.node_monitors.DiskSpaceMonitor"`
		Hudson_NodeMonitors_ResponseTimeMonitor struct {
			Average int64 `json:"average"`
		} `json:"hudson.node_monitors.ResponseTimeMonitor"`
		Hudson_NodeMonitors_SwapSpaceMonitor      interface{} `json:"hudson.node_monitors.SwapSpaceMonitor"`
		Hudson_NodeMonitors_TemporarySpaceMonitor interface{} `json:"hudson.node_monitors.TemporarySpaceMonitor"`
	} `json:"monitorData"`
	NumExecutors       int64         `json:"numExecutors"`
	Offline            bool          `json:"offline"`
	OfflineCause       struct{}      `json:"offlineCause"`
	OfflineCauseReason string        `json:"offlineCauseReason"`
	OneOffExecutors    []interface{} `json:"oneOffExecutors"`
	TemporarilyOffline bool          `json:"temporarilyOffline"`
}

func (n *Node) Info() (*nodeResponse, error) {
	_, err := n.Poll()
	if err != nil {
		return nil, err
	}
	return n.Raw, nil
}

func (n *Node) GetName() string {
	return n.Raw.DisplayName
}

func (n *Node) Delete() (bool, error) {
	resp, err := n.Jenkins.Requester.Post(n.Base+"/doDelete", nil, nil, nil)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == 200, nil
}

func (n *Node) IsOnline() (bool, error) {
	_, err := n.Poll()
	if err != nil {
		return false, err
	}
	return !n.Raw.Offline, nil
}

func (n *Node) IsTemporarilyOffline() (bool, error) {
	_, err := n.Poll()
	if err != nil {
		return false, err
	}
	return n.Raw.TemporarilyOffline, nil
}

func (n *Node) IsIdle() (bool, error) {
	_, err := n.Poll()
	if err != nil {
		return false, err
	}
	return n.Raw.Idle, nil
}

func (n *Node) IsJnlpAgent() (bool, error) {
	_, err := n.Poll()
	if err != nil {
		return false, err
	}
	return n.Raw.JnlpAgent, nil
}

func (n *Node) SetOnline() (bool, error) {
	_, err := n.Poll()

	if err != nil {
		return false, err
	}

	if n.Raw.Offline && !n.Raw.TemporarilyOffline {
		return false, errors.New("Node is Permanently offline, can't bring it up")
	}

	if n.Raw.Offline && n.Raw.TemporarilyOffline {
		return n.ToggleTemporarilyOffline()
	}

	return false, nil
}

func (n *Node) SetOffline() (bool, error) {
	if !n.Raw.Offline {
		return n.ToggleTemporarilyOffline()
	}
	return false, errors.New("Node already Offline")
}

func (n *Node) ToggleTemporarilyOffline(options ...interface{}) (bool, error) {
	state_before, err := n.IsTemporarilyOffline()
	if err != nil {
		return false, err
	}
	qr := map[string]string{"offlineMessage": "requested from gojenkins"}
	if len(options) > 0 {
		qr["offlineMessage"] = options[0].(string)
	}
	_, err = n.Jenkins.Requester.GetJSON(n.Base+"/toggleOffline", nil, qr)
	if err != nil {
		return false, err
	}
	new_state, err := n.IsTemporarilyOffline()
	if err != nil {
		return false, err
	}
	if state_before == new_state {
		return false, errors.New("Node state not changed")
	}
	return false, nil
}

func (n *Node) Poll() (int, error) {
	_, err := n.Jenkins.Requester.GetJSON(n.Base, n.Raw, nil)
	if err != nil {
		return 0, err
	}
	return n.Jenkins.Requester.LastResponse.StatusCode, nil
}
