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

package grammar

import (
	. "github.com/open-component-model/ocm/pkg/regex"

	"github.com/open-component-model/ocm/pkg/contexts/oci/grammar"
)

const (
	ComponentSeparatorChar = grammar.RepositorySeparatorChar
	ComponentSeparator     = grammar.RepositorySeparator

	VersionSeparatorChar = grammar.TagSeparatorChar
	VersionSeparator     = grammar.TagSeparator
)

var (
	// TypeRegexp describes a type name for a repository.
	TypeRegexp = grammar.TypeRegexp

	// AnchoredRepositoryRegexp parses a uniform repository spec.
	AnchoredRepositoryRegexp = Anchored(
		Optional(Capture(TypeRegexp), Literal("::")),
		Capture(grammar.DomainPortRegexp), Optional(grammar.RepositorySeparatorRegexp, Capture(grammar.RepositoryRegexp)),
	)

	// AnchoredGenericRepositoryRegexp describes a CTF reference.
	AnchoredGenericRepositoryRegexp = Anchored(
		Optional(Capture(TypeRegexp), Literal("::")),
		Capture(Match(".*")),
	)

	// ComponentRegexp describes the component name. It cosnsists
	// of a domain ame followed by OCI repository name components.
	ComponentRegexp = Sequence(grammar.DomainRegexp, grammar.RepositorySeparatorRegexp, grammar.RepositoryRegexp)

	// AnchoredComponentVersionRegexp parses a component with an optional version.
	AnchoredComponentVersionRegexp = Anchored(
		Capture(ComponentRegexp),
		Optional(Literal(VersionSeparator), Capture(grammar.TagRegexp)),
	)

	// AnchoredReferenceRegexp parses a complete string representation for default component references including
	// the repository part.
	// It provides 5 captures: type, repository host port, sub path, component and version.
	AnchoredReferenceRegexp = Anchored(
		Optional(Capture(TypeRegexp), Literal("::")),
		Capture(grammar.DomainPortRegexp), Optional(grammar.RepositorySeparatorRegexp, Capture(grammar.RepositoryRegexp)),
		Literal("//"), Capture(ComponentRegexp),
		Optional(Literal(VersionSeparator), Capture(grammar.TagRegexp)),
	)

	// AnchoredGenericReferenceRegexp parses a CTF file based string representation.
	AnchoredGenericReferenceRegexp = Anchored(
		Optional(Capture(TypeRegexp), Literal("::")),
		Capture(Match(".*?")),
		Optional(Literal("//"), Capture(ComponentRegexp),
			Optional(Literal(VersionSeparator), Capture(grammar.TagRegexp))),
	)
)
