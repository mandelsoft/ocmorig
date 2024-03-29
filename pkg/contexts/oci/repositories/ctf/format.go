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

package ctf

import (
	"sort"
	"strings"
	"sync"

	"github.com/mandelsoft/vfs/pkg/vfs"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/common/accessobj"
	"github.com/open-component-model/ocm/pkg/contexts/oci/cpi"
	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/ctf/format"
	"github.com/open-component-model/ocm/pkg/errors"
)

const (
	ArtefactIndexFileName = format.ArtefactIndexFileName
	BlobsDirectoryName    = format.BlobsDirectoryName
)

var accessObjectInfo = &accessobj.AccessObjectInfo{
	DescriptorFileName:       ArtefactIndexFileName,
	ObjectTypeName:           "repository",
	ElementDirectoryName:     BlobsDirectoryName,
	ElementTypeName:          "blob",
	DescriptorHandlerFactory: NewStateHandler,
}

type Object = Repository

type FormatHandler interface {
	accessio.Option

	Format() accessio.FileFormat

	Open(ctx cpi.Context, acc accessobj.AccessMode, path string, opts accessio.Options) (*Object, error)
	Create(ctx cpi.Context, path string, opts accessio.Options, mode vfs.FileMode) (*Object, error)
	Write(obj *Object, path string, opts accessio.Options, mode vfs.FileMode) error
}

type formatHandler struct {
	accessobj.FormatHandler
}

var (
	FormatDirectory = RegisterFormat(accessobj.FormatDirectory)
	FormatTAR       = RegisterFormat(accessobj.FormatTAR)
	FormatTGZ       = RegisterFormat(accessobj.FormatTGZ)
)

////////////////////////////////////////////////////////////////////////////////

var (
	fileFormats = map[accessio.FileFormat]FormatHandler{}
	lock        sync.RWMutex
)

func RegisterFormat(f accessobj.FormatHandler) FormatHandler {
	lock.Lock()
	defer lock.Unlock()
	h := &formatHandler{f}
	fileFormats[f.Format()] = h
	return h
}

func GetFormat(name accessio.FileFormat) FormatHandler {
	lock.RLock()
	defer lock.RUnlock()
	return fileFormats[name]
}

func SupportedFormats() []accessio.FileFormat {
	lock.RLock()
	defer lock.RUnlock()
	result := make([]accessio.FileFormat, 0, len(fileFormats))
	for f := range fileFormats {
		result = append(result, f)
	}
	sort.Slice(result, func(i, j int) bool { return strings.Compare(string(result[i]), string(result[j])) < 0 })
	return result
}

////////////////////////////////////////////////////////////////////////////////

func OpenFromBlob(ctx cpi.Context, acc accessobj.AccessMode, blob accessio.BlobAccess, opts ...accessio.Option) (*Object, error) {
	o := accessio.AccessOptions(opts...)
	if o.File != nil || o.Reader != nil {
		return nil, errors.ErrInvalid("file or reader option nor possible for blob access")
	}
	reader, err := blob.Reader()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	o.Reader = reader
	fmt := accessio.FormatTar
	mime := blob.MimeType()
	if strings.HasSuffix(mime, "+gzip") {
		fmt = accessio.FormatTGZ
	}
	o.FileFormat = &fmt
	return Open(ctx, acc&accessobj.ACC_READONLY, "", 0, o)
}

func Open(ctx cpi.Context, acc accessobj.AccessMode, path string, mode vfs.FileMode, opts ...accessio.Option) (*Object, error) {
	o, create, err := accessobj.HandleAccessMode(acc, path, opts...)
	if err != nil {
		return nil, err
	}
	h, ok := fileFormats[*o.FileFormat]
	if !ok {
		return nil, errors.ErrUnknown(accessobj.KIND_FILEFORMAT, o.FileFormat.String())
	}
	if create {
		return h.Create(ctx, path, o, mode)
	}
	return h.Open(ctx, acc, path, o)
}

func Create(ctx cpi.Context, acc accessobj.AccessMode, path string, mode vfs.FileMode, opts ...accessio.Option) (*Object, error) {
	o := accessio.AccessOptions(opts...).DefaultFormat(accessio.FormatDirectory)
	h, ok := fileFormats[*o.FileFormat]
	if !ok {
		return nil, errors.ErrUnknown(accessobj.KIND_FILEFORMAT, o.FileFormat.String())
	}
	return h.Create(ctx, path, o, mode)
}

func (h *formatHandler) Open(ctx cpi.Context, acc accessobj.AccessMode, path string, opts accessio.Options) (*Object, error) {
	obj, err := h.FormatHandler.Open(accessObjectInfo, acc, path, opts)
	return _Wrap(ctx, NewRepositorySpec(acc, path, opts), obj, err)
}

func (h *formatHandler) Create(ctx cpi.Context, path string, opts accessio.Options, mode vfs.FileMode) (*Object, error) {
	obj, err := h.FormatHandler.Create(accessObjectInfo, path, opts, mode)
	return _Wrap(ctx, NewRepositorySpec(accessobj.ACC_CREATE, path, opts), obj, err)
}

// WriteToFilesystem writes the current object to a filesystem.
func (h *formatHandler) Write(obj *Object, path string, opts accessio.Options, mode vfs.FileMode) error {
	return h.FormatHandler.Write(obj.base.Access(), path, opts, mode)
}
