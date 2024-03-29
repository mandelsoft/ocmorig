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

package cpi

// This is the Context Provider Interface for credential providers

import (
	"github.com/open-component-model/ocm/pkg/common"
	"github.com/open-component-model/ocm/pkg/contexts/credentials/core"
	"github.com/open-component-model/ocm/pkg/contexts/datacontext"
)

const (
	KIND_CREDENTIALS = core.KIND_CREDENTIALS
	KIND_REPOSITORY  = core.KIND_REPOSITORY
)

const CONTEXT_TYPE = core.CONTEXT_TYPE

type (
	Context                = core.Context
	Repository             = core.Repository
	RepositoryType         = core.RepositoryType
	Credentials            = core.Credentials
	CredentialsSource      = core.CredentialsSource
	CredentialsChain       = core.CredentialsChain
	CredentialsSpec        = core.CredentialsSpec
	RepositorySpec         = core.RepositorySpec
	GenericRepositorySpec  = core.GenericRepositorySpec
	GenericCredentialsSpec = core.GenericCredentialsSpec
)

type (
	ConsumerIdentity = core.ConsumerIdentity
	IdentityMatcher  = core.IdentityMatcher
)

var DefaultContext = core.DefaultContext

func New(m ...datacontext.BuilderMode) Context {
	return core.Builder{}.New(m...)
}

func NewGenericCredentialsSpec(name string, repospec *GenericRepositorySpec) *GenericCredentialsSpec {
	return core.NewGenericCredentialsSpec(name, repospec)
}

func NewCredentialsSpec(name string, repospec RepositorySpec) CredentialsSpec {
	return core.NewCredentialsSpec(name, repospec)
}

func ToGenericCredentialsSpec(spec CredentialsSpec) (*GenericCredentialsSpec, error) {
	return core.ToGenericCredentialsSpec(spec)
}

func ToGenericRepositorySpec(spec RepositorySpec) (*GenericRepositorySpec, error) {
	return core.ToGenericRepositorySpec(spec)
}

func RegisterRepositoryType(name string, atype RepositoryType) {
	core.DefaultRepositoryTypeScheme.Register(name, atype)
}

func RegisterStandardIdentityMatcher(typ string, matcher IdentityMatcher, desc string) {
	core.StandardIdentityMatchers.Register(typ, matcher, desc)
}

func NewCredentials(props common.Properties) Credentials {
	return core.NewCredentials(props)
}

func ErrUnknownCredentials(name string) error {
	return core.ErrUnknownCredentials(name)
}

func ErrUnknownRepository(kind, name string) error {
	return core.ErrUnknownRepository(kind, name)
}

var (
	CompleteMatch = core.CompleteMatch
	NoMatch       = core.NoMatch
	PartialMatch  = core.PartialMatch
)
