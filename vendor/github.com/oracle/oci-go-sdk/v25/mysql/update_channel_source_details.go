// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v25/common"
)

// UpdateChannelSourceDetails Parameters detailing how to provision the source for the given Channel.
type UpdateChannelSourceDetails interface {
}

type updatechannelsourcedetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *updatechannelsourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatechannelsourcedetails updatechannelsourcedetails
	s := struct {
		Model Unmarshalerupdatechannelsourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatechannelsourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "MYSQL":
		mm := UpdateChannelSourceFromMysqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updatechannelsourcedetails) String() string {
	return common.PointerString(m)
}
