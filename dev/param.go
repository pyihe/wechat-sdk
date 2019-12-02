package dev

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
)

type Params map[string]interface{}

func NewPayParam() Params {
	return make(Params)
}

func (m Params) Get(key string) interface{} {
	if m == nil {
		return nil
	}
	return m[key]
}

func (m Params) Add(key string, value interface{}) {
	if m == nil {
		m = make(Params)
	}
	m[key] = value
}

func (m Params) MarshalXML() (reader io.Reader, err error) {
	buffer := bytes.NewBuffer(make([]byte, 0))

	if _, err = io.WriteString(buffer, "<xml>"); err != nil {
		return
	}

	for k, v := range m {
		if _, err = io.WriteString(buffer, "<"+k+">"); err != nil {
			return
		}
		if err = xml.EscapeText(buffer, []byte(fmt.Sprintf("%v", v))); err != nil {
			return
		}
		if _, err = io.WriteString(buffer, "</"+k+">"); err != nil {
			return
		}
	}

	if _, err = io.WriteString(buffer, "</xml>"); err != nil {
		return
	}
	return buffer, nil
}

func (m Params) SortKey() (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
