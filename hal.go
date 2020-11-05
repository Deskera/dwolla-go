package dwolla

import "fmt"

// ValidationError is a dwolla validation error
type ValidationError struct {
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	Embedded Embedded `json:"_embedded"`
}

// Error implements the error interface
func (v ValidationError) Error() string {
	return fmt.Sprintf("[%s] %s", v.Code, v.Message)
}

// Embedded is a hal embedded resource
type Embedded struct {
	Errors []HALError `json:"errors"`
}

// HALError is a hal error
type HALError struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Path    string    `json:"path"`
	Links   AboutLink `json:"_links"`
}

// HALError implements the error interface
func (e HALError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

type AboutLink struct {
	About Link `json:"about"`
}

// Resource is a hal resource
type Resource struct {
	Links Links `json:"_links,omitempty"`
}

// Link is a hal resource link
type Link struct {
	Href         string `json:"href"`
	ResourceType string `json:"resource-type"`
	Type         string `json:"type"`
}

// Links is a group of resource links
type Links map[string]Link
