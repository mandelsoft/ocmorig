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

package comppathopt_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/options/comppathopt"
	v1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/errors"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Common OCM command ustilities for components")
}

var _ = Describe("--path option", func() {
	opts := comppathopt.Option{
		Active: true,
	}

	It("consumes simple name sequence", func() {
		args := []string{"name1", "name2", "name3"}
		rest, err := opts.Complete(args)
		Expect(err).To(Succeed())
		Expect(rest).To(BeNil())

		Expect(opts.Ids).To(Equal([]v1.Identity{
			{
				v1.SystemIdentityName: "name1",
			},
			{
				v1.SystemIdentityName: "name2",
			},
			{
				v1.SystemIdentityName: "name3",
			},
		}))
	})

	It("consumes simple name sequence and stops", func() {
		args := []string{"name1", "name2", ";", "name3"}
		rest, err := opts.Complete(args)
		Expect(err).To(Succeed())
		Expect(rest).To(Equal([]string{"name3"}))

		Expect(opts.Ids).To(Equal([]v1.Identity{
			{
				v1.SystemIdentityName: "name1",
			},
			{
				v1.SystemIdentityName: "name2",
			},
		}))
	})

	It("consumes single complex identity", func() {
		args := []string{"name1", "a=v1", "attr=v2"}
		rest, err := opts.Complete(args)
		Expect(err).To(Succeed())
		Expect(rest).To(BeNil())

		Expect(opts.Ids).To(Equal([]v1.Identity{
			{
				v1.SystemIdentityName: "name1",
				"a":                   "v1",
				"attr":                "v2",
			},
		}))
	})

	It("consumes sequence complex identity", func() {
		args := []string{"name1", "a=v1", "attr=v2", "name2", "attr=v3"}
		rest, err := opts.Complete(args)
		Expect(err).To(Succeed())
		Expect(rest).To(BeNil())

		Expect(opts.Ids).To(Equal([]v1.Identity{
			{
				v1.SystemIdentityName: "name1",
				"a":                   "v1",
				"attr":                "v2",
			},
			{
				v1.SystemIdentityName: "name2",
				"attr":                "v3",
			},
		}))
	})

	It("consumes sequence of complex identities and stops", func() {
		args := []string{"name1", "a=v1", "attr=v2", "name2", "attr=v3", ";", "name3"}
		rest, err := opts.Complete(args)
		Expect(err).To(Succeed())
		Expect(rest).To(Equal([]string{"name3"}))

		Expect(opts.Ids).To(Equal([]v1.Identity{
			{
				v1.SystemIdentityName: "name1",
				"a":                   "v1",
				"attr":                "v2",
			},
			{
				v1.SystemIdentityName: "name2",
				"attr":                "v3",
			},
		}))
	})

	It("consumes sequence of mixed identities", func() {
		args := []string{"name1", "a=v1", "attr=v2", "name2", "name3", "attr=v3"}
		rest, err := opts.Complete(args)
		Expect(err).To(Succeed())
		Expect(rest).To(BeNil())

		Expect(opts.Ids).To(Equal([]v1.Identity{
			{
				v1.SystemIdentityName: "name1",
				"a":                   "v1",
				"attr":                "v2",
			},
			{
				v1.SystemIdentityName: "name2",
			},
			{
				v1.SystemIdentityName: "name3",
				"attr":                "v3",
			},
		}))
	})

	It("fails for initial assignment", func() {
		args := []string{"a=v1", "attr=v2", "name2", "name3", "attr=v3"}
		_, err := opts.Complete(args)
		Expect(err).To(Equal(errors.New("first resource identity argument must be a sole resource name")))
	})

	It("fails for empty key", func() {
		args := []string{"name1", "a=v1", "=v2"}
		_, err := opts.Complete(args)
		Expect(err).To(Equal(errors.New("extra identity key might not be empty in \"=v2\"")))
	})
})
