package request

import "github.com/gorilla/mux"
import "github.com/humpback/humpback-agent/models"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
ResolveClusterGroupRequest
Method:  GET
Route:   /v1/cluster/groups/{groupid}
GroupID: cluster groupid
*/
type ClusterGroupRequest struct {
	GroupID string `json:"groupid"`
}

func ResolveClusterGroupRequest(r *http.Request) (*ClusterGroupRequest, error) {

	vars := mux.Vars(r)
	groupid := strings.TrimSpace(vars["groupid"])
	if len(strings.TrimSpace(groupid)) == 0 {
		return nil, fmt.Errorf("groupid invalid, can not be empty")
	}

	return &ClusterGroupRequest{
		GroupID: groupid,
	}, nil
}

/*
ResolveClusterGroupEnginesRequest
Method:  GET
Route:   /v1/cluster/groups/{groupid}/engines
GroupID: cluster groupid
*/
type ClusterGroupEnginesRequest struct {
	GroupID string `json:"groupid"`
}

func ResolveClusterGroupEnginesRequest(r *http.Request) (*ClusterGroupEnginesRequest, error) {

	vars := mux.Vars(r)
	groupid := strings.TrimSpace(vars["groupid"])
	if len(strings.TrimSpace(groupid)) == 0 {
		return nil, fmt.Errorf("groupid invalid, can not be empty")
	}

	return &ClusterGroupEnginesRequest{
		GroupID: groupid,
	}, nil
}

/*
ResolveClusterEngineRequest
Method:  GET
Route:   /v1/cluster/engines/{server}
Server: cluster engine host address
*/
type ClusterEngineRequest struct {
	Server string `json:"server"`
}

func ResolveClusterEngineRequest(r *http.Request) (*ClusterEngineRequest, error) {

	vars := mux.Vars(r)
	server := strings.TrimSpace(vars["server"])
	if len(strings.TrimSpace(server)) == 0 {
		return nil, fmt.Errorf("engine server invalid, can not be empty")
	}

	return &ClusterEngineRequest{
		Server: server,
	}, nil
}

const (
	GROUP_CREATE_EVENT = "create"
	GROUP_REMOVE_EVENT = "remove"
	GROUP_CHANGE_EVENT = "change"
)

type ClusterGroupEventRequest struct {
	GroupID string `json:"groupid"`
	Event   string `json:"event"`
}

/*
ResolveClusterGroupEventRequest
Method:  POST
Route:   /v1/cluster/groups/event
GroupID: cluster groupid
Event:   event type
*/
func ResolveClusterGroupEventRequest(r *http.Request) (*ClusterGroupEventRequest, error) {

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	request := &ClusterGroupEventRequest{}
	if err := json.NewDecoder(bytes.NewReader(buf)).Decode(request); err != nil {
		return nil, err
	}

	request.Event = strings.ToLower(request.Event)
	request.Event = strings.TrimSpace(request.Event)
	if request.Event != GROUP_CREATE_EVENT &&
		request.Event != GROUP_REMOVE_EVENT &&
		request.Event != GROUP_CHANGE_EVENT {
		return nil, fmt.Errorf("event type invalid.")
	}
	return request, nil
}

/*
ResolveClusterCreateContainerRequest
Method:  POST
Route:   /v1/cluster/containers
GroupID:   cluster groupid
Instances: create container instance count
Config:    container config
*/
type ClusterCreateContainerRequest struct {
	GroupID   string           `json:"groupid"`
	Instances int              `json:"instances"`
	Config    models.Container `json:"config"`
}

func ResolveClusterCreateContainerRequest(r *http.Request) (*ClusterCreateContainerRequest, error) {

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	request := &ClusterCreateContainerRequest{}
	if err := json.NewDecoder(bytes.NewReader(buf)).Decode(request); err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(request.GroupID)) == 0 {
		return nil, fmt.Errorf("create container groupid invalid, can not be empty")
	}

	if request.Instances < 0 {
		return nil, fmt.Errorf("create container instances invalid, should be larger than 0")
	}

	if len(strings.TrimSpace(request.Config.Name)) == 0 {
		return nil, fmt.Errorf("create container name can not be empty")
	}

	return request, nil
}
