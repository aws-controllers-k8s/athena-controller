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

package data_catalog

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackmetrics "github.com/aws-controllers-k8s/runtime/pkg/metrics"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/athena"

	svcapitypes "github.com/aws-controllers-k8s/athena-controller/apis/v1alpha1"
)

// stubAthena returns a resource manager whose Athena client posts to a stub
// that answers every call with an empty document, and the X-Amz-Target of
// each call in order.
func stubAthena(t *testing.T) (*resourceManager, *[]string) {
	targets := &[]string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.ReadAll(r.Body)
		*targets = append(*targets, r.Header.Get("X-Amz-Target"))
		w.Header().Set("Content-Type", "application/x-amz-json-1.1")
		_, _ = io.WriteString(w, "{}")
	}))
	t.Cleanup(srv.Close)
	return &resourceManager{
		sdkapi:  svcsdk.New(svcsdk.Options{Region: "us-west-2", BaseEndpoint: aws.String(srv.URL), Credentials: aws.AnonymousCredentials{}, RetryMaxAttempts: 1}),
		metrics: ackmetrics.NewMetrics("athena"),
	}, targets
}

func testTag(key, value string) *svcapitypes.Tag {
	return &svcapitypes.Tag{Key: aws.String(key), Value: aws.String(value)}
}

func testResource(tags []*svcapitypes.Tag) *resource {
	arn := ackv1alpha1.AWSResourceName("arn:aws:athena:us-west-2:123456789012:DataCatalog/test")
	ko := &svcapitypes.DataCatalog{}
	ko.Spec.Name = aws.String("test")
	ko.Spec.Tags = tags
	ko.Spec.Type = aws.String("LAMBDA")
	ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{ARN: &arn}
	return &resource{ko}
}

func TestSdkUpdate_TagsOnlySkipsUpdateDataCatalog(t *testing.T) {
	rm, targets := stubAthena(t)
	desired := testResource([]*svcapitypes.Tag{testTag("env", "prod")})
	latest := testResource([]*svcapitypes.Tag{testTag("env", "dev")})
	delta := ackcompare.NewDelta()
	delta.Add("Spec.Tags", desired.ko.Spec.Tags, latest.ko.Spec.Tags)

	_, err := rm.sdkUpdate(context.Background(), desired, latest, delta)
	if err != nil {
		t.Fatalf("sdkUpdate: %v", err)
	}
	if want := []string{"AmazonAthena.TagResource"}; !reflect.DeepEqual(*targets, want) {
		t.Errorf("calls = %v, want %v: "+"a tags-only change must not call UpdateDataCatalog", *targets, want)
	}
}

func TestSdkUpdate_TagsAndFieldCallsUpdateDataCatalog(t *testing.T) {
	rm, targets := stubAthena(t)
	desired := testResource([]*svcapitypes.Tag{testTag("env", "prod")})
	desired.ko.Spec.Description = aws.String("new")
	latest := testResource([]*svcapitypes.Tag{testTag("env", "dev")})
	latest.ko.Spec.Description = aws.String("old")
	delta := ackcompare.NewDelta()
	delta.Add("Spec.Tags", desired.ko.Spec.Tags, latest.ko.Spec.Tags)
	delta.Add("Spec.Description", desired.ko.Spec.Description, latest.ko.Spec.Description)

	_, err := rm.sdkUpdate(context.Background(), desired, latest, delta)
	if err != nil {
		t.Fatalf("sdkUpdate: %v", err)
	}
	if want := []string{"AmazonAthena.TagResource", "AmazonAthena.UpdateDataCatalog"}; !reflect.DeepEqual(*targets, want) {
		t.Errorf("calls = %v, want %v", *targets, want)
	}
}
