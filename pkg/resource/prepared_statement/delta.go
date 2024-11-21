// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package prepared_statement

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.QueryStatement, b.ko.Spec.QueryStatement) {
		delta.Add("Spec.QueryStatement", a.ko.Spec.QueryStatement, b.ko.Spec.QueryStatement)
	} else if a.ko.Spec.QueryStatement != nil && b.ko.Spec.QueryStatement != nil {
		if *a.ko.Spec.QueryStatement != *b.ko.Spec.QueryStatement {
			delta.Add("Spec.QueryStatement", a.ko.Spec.QueryStatement, b.ko.Spec.QueryStatement)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.WorkGroup, b.ko.Spec.WorkGroup) {
		delta.Add("Spec.WorkGroup", a.ko.Spec.WorkGroup, b.ko.Spec.WorkGroup)
	} else if a.ko.Spec.WorkGroup != nil && b.ko.Spec.WorkGroup != nil {
		if *a.ko.Spec.WorkGroup != *b.ko.Spec.WorkGroup {
			delta.Add("Spec.WorkGroup", a.ko.Spec.WorkGroup, b.ko.Spec.WorkGroup)
		}
	}

	return delta
}