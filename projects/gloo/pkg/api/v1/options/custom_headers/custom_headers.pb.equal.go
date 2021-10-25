// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/custom_headers/custom_headers.proto

package custom_headers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *HeaderManipulation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderManipulation)
	if !ok {
		that2, ok := that.(HeaderManipulation)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetRequestHeadersToAdd()) != len(target.GetRequestHeadersToAdd()) {
		return false
	}
	for idx, v := range m.GetRequestHeadersToAdd() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestHeadersToAdd()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRequestHeadersToAdd()[idx]) {
				return false
			}
		}

	}

	if len(m.GetRequestHeadersToRemove()) != len(target.GetRequestHeadersToRemove()) {
		return false
	}
	for idx, v := range m.GetRequestHeadersToRemove() {

		if strings.Compare(v, target.GetRequestHeadersToRemove()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetResponseHeadersToAdd()) != len(target.GetResponseHeadersToAdd()) {
		return false
	}
	for idx, v := range m.GetResponseHeadersToAdd() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResponseHeadersToAdd()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResponseHeadersToAdd()[idx]) {
				return false
			}
		}

	}

	if len(m.GetResponseHeadersToRemove()) != len(target.GetResponseHeadersToRemove()) {
		return false
	}
	for idx, v := range m.GetResponseHeadersToRemove() {

		if strings.Compare(v, target.GetResponseHeadersToRemove()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *HeaderValueOption) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderValueOption)
	if !ok {
		that2, ok := that.(HeaderValueOption)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetHeader()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeader()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeader(), target.GetHeader()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAppend()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAppend()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAppend(), target.GetAppend()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HeaderValue) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderValue)
	if !ok {
		that2, ok := that.(HeaderValue)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetValue(), target.GetValue()) != 0 {
		return false
	}

	return true
}
