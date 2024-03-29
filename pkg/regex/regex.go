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

package regex

import (
	"regexp"
)

var (
	// Alpha defines the alpha atom.
	// This only allows upper and lower case characters.
	Alpha = Match(`[A-Za-z]+`)

	// AlphaNumeric defines the alpha numeric atom, typically a
	// component of names. This only allows upper and lower case characters and digits.
	AlphaNumeric = Match(`[A-Za-z0-9]+`)

	// Identifier is an AlphaNumeric regexp starting with an Alpha regexp.
	Identifier = Sequence(Alpha, Match(`[A-Za-z0-9]`), Optional(Literal("+"), Alpha))
)

// Match compiles the string to a regular expression.
var Match = regexp.MustCompile

// Literal compiles s into a literal regular expression, escaping any regexp
// reserved characters.
func Literal(s string) *regexp.Regexp {
	re := Match(regexp.QuoteMeta(s))

	if _, complete := re.LiteralPrefix(); !complete {
		panic("must be a literal")
	}

	return re
}

// Sequence defines a full expression, where each regular expression must
// follow the previous.
func Sequence(res ...*regexp.Regexp) *regexp.Regexp {
	var s string
	for _, re := range res {
		s += re.String()
	}

	return Match(s)
}

// Optional wraps the expression in a non-capturing group and makes the
// production optional.
func Optional(res ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(Sequence(res...)).String() + `?`)
}

// Repeated wraps the regexp in a non-capturing group to get one or more
// matches.
func Repeated(res ...*regexp.Regexp) *regexp.Regexp {
	return Match(Group(Sequence(res...)).String() + `+`)
}

// Group wraps the regexp in a non-capturing group.
func Group(res ...*regexp.Regexp) *regexp.Regexp {
	return Match(`(?:` + Sequence(res...).String() + `)`)
}

// Or wraps alternative regexps.
func Or(res ...*regexp.Regexp) *regexp.Regexp {
	var s string
	sep := ""
	for _, re := range res {
		s += sep + Group(re).String()
		sep = "|"
	}
	return Match(`(?:` + s + `)`)
}

// Capture wraps the expression in a capturing group.
func Capture(res ...*regexp.Regexp) *regexp.Regexp {
	return Match(`(` + Sequence(res...).String() + `)`)
}

// Anchored anchors the regular expression by adding start and end delimiters.
func Anchored(res ...*regexp.Regexp) *regexp.Regexp {
	return Match(`^` + Sequence(res...).String() + `$`)
}
