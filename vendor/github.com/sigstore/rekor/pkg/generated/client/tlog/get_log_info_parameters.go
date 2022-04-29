// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetLogInfoParams creates a new GetLogInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogInfoParams() *GetLogInfoParams {
	return &GetLogInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogInfoParamsWithTimeout creates a new GetLogInfoParams object
// with the ability to set a timeout on a request.
func NewGetLogInfoParamsWithTimeout(timeout time.Duration) *GetLogInfoParams {
	return &GetLogInfoParams{
		timeout: timeout,
	}
}

// NewGetLogInfoParamsWithContext creates a new GetLogInfoParams object
// with the ability to set a context for a request.
func NewGetLogInfoParamsWithContext(ctx context.Context) *GetLogInfoParams {
	return &GetLogInfoParams{
		Context: ctx,
	}
}

// NewGetLogInfoParamsWithHTTPClient creates a new GetLogInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogInfoParamsWithHTTPClient(client *http.Client) *GetLogInfoParams {
	return &GetLogInfoParams{
		HTTPClient: client,
	}
}

/* GetLogInfoParams contains all the parameters to send to the API endpoint
   for the get log info operation.

   Typically these are written to a http.Request.
*/
type GetLogInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogInfoParams) WithDefaults() *GetLogInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log info params
func (o *GetLogInfoParams) WithTimeout(timeout time.Duration) *GetLogInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log info params
func (o *GetLogInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log info params
func (o *GetLogInfoParams) WithContext(ctx context.Context) *GetLogInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log info params
func (o *GetLogInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log info params
func (o *GetLogInfoParams) WithHTTPClient(client *http.Client) *GetLogInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log info params
func (o *GetLogInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}