// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// RemoveTeamToTestURL generates an URL for the remove team to test operation
type RemoveTeamToTestURL struct {
	Teamname string
	Testid   int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RemoveTeamToTestURL) WithBasePath(bp string) *RemoveTeamToTestURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RemoveTeamToTestURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RemoveTeamToTestURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/tests/{testid}/teams/{teamname}"

	teamname := o.Teamname
	if teamname != "" {
		_path = strings.Replace(_path, "{teamname}", teamname, -1)
	} else {
		return nil, errors.New("teamname is required on RemoveTeamToTestURL")
	}

	testid := swag.FormatInt64(o.Testid)
	if testid != "" {
		_path = strings.Replace(_path, "{testid}", testid, -1)
	} else {
		return nil, errors.New("testid is required on RemoveTeamToTestURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/DevTest"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RemoveTeamToTestURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RemoveTeamToTestURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RemoveTeamToTestURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RemoveTeamToTestURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RemoveTeamToTestURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *RemoveTeamToTestURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
