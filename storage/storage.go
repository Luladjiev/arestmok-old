package storage

import "fmt"

// RouteStruct is a struct for specific route
type RouteStruct struct {
	Config    interface{}
	Structure interface{}
	Data      []interface{}
}

type storageStruct struct {
	Routes map[string]RouteStruct
}

func (s *storageStruct) set(id string, route RouteStruct) {
	s.Routes[id] = route
}

func (s *storageStruct) get(id string) (route RouteStruct, err error) {
	route, ok := s.Routes[id]
	if ok != true {
		return route, fmt.Errorf("%s not found in storage", id)
	}
	return s.Routes[id], nil
}

var storageData = storageStruct{
	Routes: make(map[string]RouteStruct),
}

// Set data in storage
func Set(id string, data RouteStruct) {
	storageData.set(id, data)
}

// Get data from storage
func Get(id string) (route RouteStruct, err error) {
	return storageData.get(id)
}
