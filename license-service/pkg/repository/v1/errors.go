// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 
//
package v1

import "errors"

var (
	// ErrNoData is a comman error when we are not able to find any data in db we should give this
	ErrNoData = errors.New("No Data Found")
	// ErrNodeNotFound is returned when the node we are looking for does not exist in database
	ErrNodeNotFound = errors.New("Node does not exist")
)
