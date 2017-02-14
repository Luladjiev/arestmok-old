package storage

import "fmt"

type storageStruct struct {
	routes map[string]interface{}
}

func (s *storageStruct) set(id string, data interface{}) {
	s.routes[id] = data
}

func (s *storageStruct) get(id string) (data interface{}, err error) {
	data, ok := s.routes[id]
	if ok != true {
		return nil, fmt.Errorf("%s not found in storage", id)
	}
	return s.routes[id], nil
}

var storageData = storageStruct{
	routes: make(map[string]interface{}),
}

// Set data in storage
func Set(id string, data interface{}) {
	storageData.set(id, data)
}

// Get data from storage
func Get(id string) (data interface{}, err error) {
	return storageData.get(id)
}
