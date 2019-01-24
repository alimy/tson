// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package tson

import (
	"io"
)

const (
	tagName = "tson"
)

// Template json template
type Template struct {
	// TODO
}

// Execute merge json template and data
func (t *Template) Execute(w io.Writer, data interface{}) error {
	// TODO
	return nil
}

// ExecuteTemplate merge json template and data
func (t *Template) ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	// TODO
	return nil
}

// Parse parse json template from a defined struct object with 'tson' tag
func (t *Template) Parse(source interface{}) (*Template, error) {
	// TODO
	return &Template{}, nil
}

// Parse parse json template from text string
func (t *Template) ParseFrom(text string) (*Template, error) {
	// TODO
	return &Template{}, nil
}
