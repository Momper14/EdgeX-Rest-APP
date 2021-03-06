// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetV1DevicesDeviceResourcesResourceURL generates an URL for the get v1 devices device resources resource operation
type GetV1DevicesDeviceResourcesResourceURL struct {
	Device   string
	Resource string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1DevicesDeviceResourcesResourceURL) WithBasePath(bp string) *GetV1DevicesDeviceResourcesResourceURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV1DevicesDeviceResourcesResourceURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetV1DevicesDeviceResourcesResourceURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v1/devices/{device}/resources/{resource}"

	device := o.Device
	if device != "" {
		_path = strings.Replace(_path, "{device}", device, -1)
	} else {
		return nil, errors.New("device is required on GetV1DevicesDeviceResourcesResourceURL")
	}

	resource := o.Resource
	if resource != "" {
		_path = strings.Replace(_path, "{resource}", resource, -1)
	} else {
		return nil, errors.New("resource is required on GetV1DevicesDeviceResourcesResourceURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetV1DevicesDeviceResourcesResourceURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetV1DevicesDeviceResourcesResourceURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetV1DevicesDeviceResourcesResourceURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetV1DevicesDeviceResourcesResourceURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetV1DevicesDeviceResourcesResourceURL")
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
func (o *GetV1DevicesDeviceResourcesResourceURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
