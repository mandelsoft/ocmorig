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

package oci

import (
	"context"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/contexts/oci/core"
)

const (
	KIND_OCIARTEFACT = core.KIND_OCIARTEFACT
	KIND_MEDIATYPE   = accessio.KIND_MEDIATYPE
	KIND_BLOB        = accessio.KIND_BLOB
)

const CONTEXT_TYPE = core.CONTEXT_TYPE

const CommonTransportFormat = core.CommonTransportFormat

type (
	Context                          = core.Context
	Repository                       = core.Repository
	RepositorySpecHandlers           = core.RepositorySpecHandlers
	RepositorySpecHandler            = core.RepositorySpecHandler
	UniformRepositorySpec            = core.UniformRepositorySpec
	RepositoryTypeScheme             = core.RepositoryTypeScheme
	RepositorySpec                   = core.RepositorySpec
	IntermediateRepositorySpecAspect = core.IntermediateRepositorySpecAspect
	GenericRepositorySpec            = core.GenericRepositorySpec
	ArtefactAccess                   = core.ArtefactAccess
	NamespaceLister                  = core.NamespaceLister
	NamespaceAccess                  = core.NamespaceAccess
	ManifestAccess                   = core.ManifestAccess
	IndexAccess                      = core.IndexAccess
	BlobAccess                       = core.BlobAccess
	DataAccess                       = core.DataAccess
)

func DefaultContext() core.Context {
	return core.DefaultContext
}

func ForContext(ctx context.Context) Context {
	return core.ForContext(ctx)
}

func DefinedForContext(ctx context.Context) (Context, bool) {
	return core.DefinedForContext(ctx)
}

func IsErrBlobNotFound(err error) bool {
	return accessio.IsErrBlobNotFound(err)
}

func ToGenericRepositorySpec(spec RepositorySpec) (*GenericRepositorySpec, error) {
	return core.ToGenericRepositorySpec(spec)
}
