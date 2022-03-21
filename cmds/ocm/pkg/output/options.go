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

package output

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type OtherOptions interface {
	AddFlags(fs *pflag.FlagSet)
	Complete() error
}

type Options struct {
	output string

	Output *string
	Sort   []string
	Others OtherOptions
}

func (o *Options) AddFlags(fs *pflag.FlagSet, outputs Outputs) {
	s := ""
	sep := ""
	for o := range outputs {
		if o != "" {
			s = fmt.Sprintf("%s%s%s", s, sep, o)
			sep = ", "
		}
	}
	fs.StringVarP(&o.output, "output", "o", "", fmt.Sprintf("output mode (%s)", s))
	fs.StringArrayVarP(&o.Sort, "sort", "s", nil, "sort fields")

	if o.Others != nil {
		o.Others.AddFlags(fs)
	}
}

func (o *Options) Complete() error {
	if o.output != "" {
		o.Output = &o.output
	}
	var fields []string

	for _, f := range o.Sort {
		split := strings.Split(f, ",")
		for _, s := range split {
			s = strings.TrimSpace(s)
			if s != "" {
				fields = append(fields, s)
			}
		}
	}
	o.Sort = fields
	if o.Others != nil {
		return o.Others.Complete()
	}
	return nil
}