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

package helm

import (
	"os"

	"github.com/open-component-model/ocm/cmds/ocm/clictx"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs"
	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/contexts/oci/ociutils/helm"
	"github.com/open-component-model/ocm/pkg/contexts/oci/ociutils/helm/loader"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

type Handler struct{}

var _ inputs.InputHandler = (*Handler)(nil)

func (h *Handler) Validate(fldPath *field.Path, ctx clictx.Context, input *inputs.BlobInput, inputFilePath string) field.ErrorList {
	allErrs := inputs.ForbidFileInfo(fldPath, input)
	path := fldPath.Child("path")
	if input.Path == "" {
		allErrs = append(allErrs, field.Required(path, "path is required for input"))
	} else {
		inputInfo, filePath, err := inputs.FileInfo(ctx, input.Path, inputFilePath)
		if err != nil {
			allErrs = append(allErrs, field.Invalid(path, filePath, err.Error()))
		}
		if !inputInfo.IsDir() && inputInfo.Mode()&os.ModeType != 0 {
			allErrs = append(allErrs, field.Invalid(path, filePath, "no regular file or directory"))
		}
	}
	return allErrs
}

func (h *Handler) GetBlob(ctx clictx.Context, input *inputs.BlobInput, inputFilePath string) (accessio.TemporaryBlobAccess, string, error) {
	_, inputPath, err := inputs.FileInfo(ctx, input.Path, inputFilePath)
	if err != nil {
		return nil, "", err
	}
	chart, err := loader.Load(inputPath, ctx.FileSystem())
	if err != nil {
		return nil, "", err
	}
	blob, err := helm.SynthesizeArtefactBlob(inputPath, ctx.FileSystem())
	if err != nil {
		return nil, "", err
	}
	return blob, chart.Metadata.Name, err
}

func (h *Handler) Usage() string {
	return `
- <code>helm</code>

  The path must denote an helm chart archive or directory
  relative to the resources file.
  The denoted chart is packed as an OCI artefact set.
  Additional provider info is taken from a file with the same name
  and the suffix <code>.prov</code>.

  If the chart should just be stored as archive, please use the 
  type <code>file</code> or <code>dir</code>.
`
}