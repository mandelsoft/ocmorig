// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package core

import (
	"context"
	"reflect"

	"github.com/open-component-model/ocm/pkg/contexts/credentials"
	datacontext2 "github.com/open-component-model/ocm/pkg/contexts/datacontext"
	"github.com/open-component-model/ocm/pkg/runtime"
)

const CONTEXT_TYPE = "oci.context.gardener.cloud"

const CommonTransportFormat = "CommonTransportFormat"

type Context interface {
	datacontext2.Context

	AttributesContext() datacontext2.AttributesContext
	CredentialsContext() credentials.Context

	RepositorySpecHandlers() RepositorySpecHandlers
	MapUniformRepositorySpec(u *UniformRepositorySpec, aliases Aliases) (RepositorySpec, error)

	RepositoryTypes() RepositoryTypeScheme

	RepositoryForSpec(spec RepositorySpec, creds ...credentials.CredentialsSource) (Repository, error)
	RepositoryForConfig(data []byte, unmarshaler runtime.Unmarshaler, creds ...credentials.CredentialsSource) (Repository, error)
}

var key = reflect.TypeOf(_context{})

// DefaultContext is the default context initialized by init functions
var DefaultContext = Builder{}.New()

// ForContext returns the Context to use for context.Context.
// This is either an explicit context or the default context.
func ForContext(ctx context.Context) Context {
	return datacontext2.ForContextByKey(ctx, key, DefaultContext).(Context)
}

////////////////////////////////////////////////////////////////////////////////

type _context struct {
	datacontext2.Context

	sharedattributes datacontext2.AttributesContext
	credentials      credentials.Context

	knownRepositoryTypes RepositoryTypeScheme
	specHandlers         RepositorySpecHandlers
}

func newContext(shared datacontext2.AttributesContext, creds credentials.Context, reposcheme RepositoryTypeScheme, specHandlers RepositorySpecHandlers) Context {
	if shared == nil {
		shared = creds.ConfigContext()
	}
	c := &_context{
		sharedattributes:     shared,
		credentials:          creds,
		knownRepositoryTypes: reposcheme,
		specHandlers:         specHandlers,
	}
	c.Context = datacontext2.NewContextBase(c, CONTEXT_TYPE, key, shared.GetAttributes())
	return c
}

var _ Context = &_context{}

func (c *_context) AttributesContext() datacontext2.AttributesContext {
	return c.sharedattributes
}

func (c *_context) CredentialsContext() credentials.Context {
	return c.credentials
}

func (c *_context) RepositoryTypes() RepositoryTypeScheme {
	return c.knownRepositoryTypes
}

func (c *_context) RepositorySpecHandlers() RepositorySpecHandlers {
	return c.specHandlers
}

func (c *_context) MapUniformRepositorySpec(u *UniformRepositorySpec, aliases Aliases) (RepositorySpec, error) {
	return c.specHandlers.MapUniformRepositorySpec(c, u, aliases)
}

func (c *_context) RepositorySpecForConfig(data []byte, unmarshaler runtime.Unmarshaler) (RepositorySpec, error) {
	return c.knownRepositoryTypes.DecodeRepositorySpec(data, unmarshaler)
}

func (c *_context) RepositoryForSpec(spec RepositorySpec, creds ...credentials.CredentialsSource) (Repository, error) {
	cred, err := credentials.CredentialsChain(creds).Credentials(c.CredentialsContext())
	if err != nil {
		return nil, err
	}
	return spec.Repository(c, cred)
}

func (c *_context) RepositoryForConfig(data []byte, unmarshaler runtime.Unmarshaler, creds ...credentials.CredentialsSource) (Repository, error) {
	spec, err := c.knownRepositoryTypes.DecodeRepositorySpec(data, unmarshaler)
	if err != nil {
		return nil, err
	}
	return c.RepositoryForSpec(spec, creds...)
}