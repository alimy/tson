package tson

import (
	"io"
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

// Parse parse json template content
func (t *Template) Parse(text string) (*Template, error) {
	// TODO
	return &Template{}, nil
}
