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

package common

import (
	"fmt"
	"strings"

	"github.com/open-component-model/ocm/cmds/ocm/pkg/options"
	"github.com/open-component-model/ocm/pkg/contexts/clictx"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc"
	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/errors"
)

func ConsumeIdentities(pattern bool, args []string, stop ...string) ([]metav1.Identity, []string, error) {
	result := []metav1.Identity{}
	for i, a := range args {
		for _, s := range stop {
			if s == a {
				return result, args[i+1:], nil
			}
		}
		i := strings.Index(a, "=")
		if i < 0 {
			result = append(result, metav1.Identity{compdesc.SystemIdentityName: a})
		} else {
			if len(result) == 0 {
				if !pattern {
					return nil, nil, fmt.Errorf("first resource identity argument must be a sole resource name")
				}
				result = append(result, metav1.Identity{a[:i]: a[i+1:]})
			} else {
				result[len(result)-1][a[:i]] = a[i+1:]
			}
			if i == 0 {
				return nil, nil, fmt.Errorf("extra identity key might not be empty in %q", a)
			}
		}
	}
	return result, nil, nil
}

func MapArgsToIdentities(args ...string) ([]metav1.Identity, error) {
	result, _, err := ConsumeIdentities(false, args)
	return result, err
}

func MapArgsToIdentityPattern(args ...string) (metav1.Identity, error) {
	result, _, err := ConsumeIdentities(true, args)
	if err == nil {
		switch len(result) {
		case 0:
			return nil, nil
		case 1:
			return result[0], err
		default:
			if len(result) > 1 {
				return nil, errors.Newf("only one identity pattern possible (sole name in between)")
			}
		}
	}
	return nil, err
}

////////////////////////////////////////////////////////////////////////////////

type OptionWithSessionCompleter interface {
	CompleteWithSession(ctx clictx.OCM, session ocm.Session) error
}

func CompleteOptionsWithSession(ctx clictx.Context, session ocm.Session) options.OptionsProcessor {
	return func(opt options.Options) error {
		if c, ok := opt.(OptionWithSessionCompleter); ok {
			return c.CompleteWithSession(ctx.OCM(), session)
		}
		return nil
	}
}
