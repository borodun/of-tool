package io

type ReceivedFunctions struct {
	Functions []interface{}
}

type FaaSFunctions struct {
	Functions []FaaSFunction
}

type FaaSFunction struct {
	Name string `json:"name"`
}
