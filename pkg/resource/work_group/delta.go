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

package work_group

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

	if ackcompare.HasNilDifference(a.ko.Spec.Configuration, b.ko.Spec.Configuration) {
		delta.Add("Spec.Configuration", a.ko.Spec.Configuration, b.ko.Spec.Configuration)
	} else if a.ko.Spec.Configuration != nil && b.ko.Spec.Configuration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.AdditionalConfiguration, b.ko.Spec.Configuration.AdditionalConfiguration) {
			delta.Add("Spec.Configuration.AdditionalConfiguration", a.ko.Spec.Configuration.AdditionalConfiguration, b.ko.Spec.Configuration.AdditionalConfiguration)
		} else if a.ko.Spec.Configuration.AdditionalConfiguration != nil && b.ko.Spec.Configuration.AdditionalConfiguration != nil {
			if *a.ko.Spec.Configuration.AdditionalConfiguration != *b.ko.Spec.Configuration.AdditionalConfiguration {
				delta.Add("Spec.Configuration.AdditionalConfiguration", a.ko.Spec.Configuration.AdditionalConfiguration, b.ko.Spec.Configuration.AdditionalConfiguration)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.BytesScannedCutoffPerQuery, b.ko.Spec.Configuration.BytesScannedCutoffPerQuery) {
			delta.Add("Spec.Configuration.BytesScannedCutoffPerQuery", a.ko.Spec.Configuration.BytesScannedCutoffPerQuery, b.ko.Spec.Configuration.BytesScannedCutoffPerQuery)
		} else if a.ko.Spec.Configuration.BytesScannedCutoffPerQuery != nil && b.ko.Spec.Configuration.BytesScannedCutoffPerQuery != nil {
			if *a.ko.Spec.Configuration.BytesScannedCutoffPerQuery != *b.ko.Spec.Configuration.BytesScannedCutoffPerQuery {
				delta.Add("Spec.Configuration.BytesScannedCutoffPerQuery", a.ko.Spec.Configuration.BytesScannedCutoffPerQuery, b.ko.Spec.Configuration.BytesScannedCutoffPerQuery)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration, b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration) {
			delta.Add("Spec.Configuration.CustomerContentEncryptionConfiguration", a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration, b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration)
		} else if a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration != nil && b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey, b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey) {
				delta.Add("Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey", a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey, b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey)
			} else if a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey != nil && b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey != nil {
				if *a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey != *b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey {
					delta.Add("Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey", a.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey, b.ko.Spec.Configuration.CustomerContentEncryptionConfiguration.KMSKey)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration, b.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration) {
			delta.Add("Spec.Configuration.EnableMinimumEncryptionConfiguration", a.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration, b.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration)
		} else if a.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration != nil && b.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration != nil {
			if *a.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration != *b.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration {
				delta.Add("Spec.Configuration.EnableMinimumEncryptionConfiguration", a.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration, b.ko.Spec.Configuration.EnableMinimumEncryptionConfiguration)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.EnforceWorkGroupConfiguration, b.ko.Spec.Configuration.EnforceWorkGroupConfiguration) {
			delta.Add("Spec.Configuration.EnforceWorkGroupConfiguration", a.ko.Spec.Configuration.EnforceWorkGroupConfiguration, b.ko.Spec.Configuration.EnforceWorkGroupConfiguration)
		} else if a.ko.Spec.Configuration.EnforceWorkGroupConfiguration != nil && b.ko.Spec.Configuration.EnforceWorkGroupConfiguration != nil {
			if *a.ko.Spec.Configuration.EnforceWorkGroupConfiguration != *b.ko.Spec.Configuration.EnforceWorkGroupConfiguration {
				delta.Add("Spec.Configuration.EnforceWorkGroupConfiguration", a.ko.Spec.Configuration.EnforceWorkGroupConfiguration, b.ko.Spec.Configuration.EnforceWorkGroupConfiguration)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.EngineVersion, b.ko.Spec.Configuration.EngineVersion) {
			delta.Add("Spec.Configuration.EngineVersion", a.ko.Spec.Configuration.EngineVersion, b.ko.Spec.Configuration.EngineVersion)
		} else if a.ko.Spec.Configuration.EngineVersion != nil && b.ko.Spec.Configuration.EngineVersion != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion, b.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion) {
				delta.Add("Spec.Configuration.EngineVersion.EffectiveEngineVersion", a.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion, b.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion)
			} else if a.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion != nil && b.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion != nil {
				if *a.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion != *b.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion {
					delta.Add("Spec.Configuration.EngineVersion.EffectiveEngineVersion", a.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion, b.ko.Spec.Configuration.EngineVersion.EffectiveEngineVersion)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion, b.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion) {
				delta.Add("Spec.Configuration.EngineVersion.SelectedEngineVersion", a.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion, b.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion)
			} else if a.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion != nil && b.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion != nil {
				if *a.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion != *b.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion {
					delta.Add("Spec.Configuration.EngineVersion.SelectedEngineVersion", a.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion, b.ko.Spec.Configuration.EngineVersion.SelectedEngineVersion)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ExecutionRole, b.ko.Spec.Configuration.ExecutionRole) {
			delta.Add("Spec.Configuration.ExecutionRole", a.ko.Spec.Configuration.ExecutionRole, b.ko.Spec.Configuration.ExecutionRole)
		} else if a.ko.Spec.Configuration.ExecutionRole != nil && b.ko.Spec.Configuration.ExecutionRole != nil {
			if *a.ko.Spec.Configuration.ExecutionRole != *b.ko.Spec.Configuration.ExecutionRole {
				delta.Add("Spec.Configuration.ExecutionRole", a.ko.Spec.Configuration.ExecutionRole, b.ko.Spec.Configuration.ExecutionRole)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.IdentityCenterConfiguration, b.ko.Spec.Configuration.IdentityCenterConfiguration) {
			delta.Add("Spec.Configuration.IdentityCenterConfiguration", a.ko.Spec.Configuration.IdentityCenterConfiguration, b.ko.Spec.Configuration.IdentityCenterConfiguration)
		} else if a.ko.Spec.Configuration.IdentityCenterConfiguration != nil && b.ko.Spec.Configuration.IdentityCenterConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter, b.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter) {
				delta.Add("Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter", a.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter, b.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter)
			} else if a.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter != nil && b.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter != nil {
				if *a.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter != *b.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter {
					delta.Add("Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter", a.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter, b.ko.Spec.Configuration.IdentityCenterConfiguration.EnableIdentityCenter)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN, b.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN) {
				delta.Add("Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN", a.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN, b.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN)
			} else if a.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN != nil && b.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN != nil {
				if *a.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN != *b.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN {
					delta.Add("Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN", a.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN, b.ko.Spec.Configuration.IdentityCenterConfiguration.IdentityCenterInstanceARN)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled, b.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled) {
			delta.Add("Spec.Configuration.PublishCloudWatchMetricsEnabled", a.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled, b.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled)
		} else if a.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled != nil && b.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled != nil {
			if *a.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled != *b.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled {
				delta.Add("Spec.Configuration.PublishCloudWatchMetricsEnabled", a.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled, b.ko.Spec.Configuration.PublishCloudWatchMetricsEnabled)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration) {
			delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration)
		} else if a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration != nil && b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType) {
				delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType)
			} else if a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType != nil && b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType != nil {
				if *a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType != *b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType {
					delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.AuthenticationType)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix) {
				delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix)
			} else if a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix != nil && b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix != nil {
				if *a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix != *b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix {
					delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.CreateUserLevelPrefix)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants) {
				delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants)
			} else if a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants != nil && b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants != nil {
				if *a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants != *b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants {
					delta.Add("Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants", a.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants, b.ko.Spec.Configuration.QueryResultsS3AccessGrantsConfiguration.EnableS3AccessGrants)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.RequesterPaysEnabled, b.ko.Spec.Configuration.RequesterPaysEnabled) {
			delta.Add("Spec.Configuration.RequesterPaysEnabled", a.ko.Spec.Configuration.RequesterPaysEnabled, b.ko.Spec.Configuration.RequesterPaysEnabled)
		} else if a.ko.Spec.Configuration.RequesterPaysEnabled != nil && b.ko.Spec.Configuration.RequesterPaysEnabled != nil {
			if *a.ko.Spec.Configuration.RequesterPaysEnabled != *b.ko.Spec.Configuration.RequesterPaysEnabled {
				delta.Add("Spec.Configuration.RequesterPaysEnabled", a.ko.Spec.Configuration.RequesterPaysEnabled, b.ko.Spec.Configuration.RequesterPaysEnabled)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration, b.ko.Spec.Configuration.ResultConfiguration) {
			delta.Add("Spec.Configuration.ResultConfiguration", a.ko.Spec.Configuration.ResultConfiguration, b.ko.Spec.Configuration.ResultConfiguration)
		} else if a.ko.Spec.Configuration.ResultConfiguration != nil && b.ko.Spec.Configuration.ResultConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration, b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration) {
				delta.Add("Spec.Configuration.ResultConfiguration.ACLConfiguration", a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration, b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration)
			} else if a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration != nil && b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption, b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption) {
					delta.Add("Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption", a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption, b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption)
				} else if a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption != nil && b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption != nil {
					if *a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption != *b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption {
						delta.Add("Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption", a.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption, b.ko.Spec.Configuration.ResultConfiguration.ACLConfiguration.S3ACLOption)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration) {
				delta.Add("Spec.Configuration.ResultConfiguration.EncryptionConfiguration", a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration)
			} else if a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration != nil && b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption) {
					delta.Add("Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption", a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption)
				} else if a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption != nil && b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption != nil {
					if *a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption != *b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption {
						delta.Add("Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption", a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.EncryptionOption)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey) {
					delta.Add("Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey", a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey)
				} else if a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey != nil && b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey != nil {
					if *a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey != *b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey {
						delta.Add("Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey", a.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey, b.ko.Spec.Configuration.ResultConfiguration.EncryptionConfiguration.KMSKey)
					}
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner, b.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner) {
				delta.Add("Spec.Configuration.ResultConfiguration.ExpectedBucketOwner", a.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner, b.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner)
			} else if a.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner != nil && b.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner != nil {
				if *a.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner != *b.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner {
					delta.Add("Spec.Configuration.ResultConfiguration.ExpectedBucketOwner", a.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner, b.ko.Spec.Configuration.ResultConfiguration.ExpectedBucketOwner)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Configuration.ResultConfiguration.OutputLocation, b.ko.Spec.Configuration.ResultConfiguration.OutputLocation) {
				delta.Add("Spec.Configuration.ResultConfiguration.OutputLocation", a.ko.Spec.Configuration.ResultConfiguration.OutputLocation, b.ko.Spec.Configuration.ResultConfiguration.OutputLocation)
			} else if a.ko.Spec.Configuration.ResultConfiguration.OutputLocation != nil && b.ko.Spec.Configuration.ResultConfiguration.OutputLocation != nil {
				if *a.ko.Spec.Configuration.ResultConfiguration.OutputLocation != *b.ko.Spec.Configuration.ResultConfiguration.OutputLocation {
					delta.Add("Spec.Configuration.ResultConfiguration.OutputLocation", a.ko.Spec.Configuration.ResultConfiguration.OutputLocation, b.ko.Spec.Configuration.ResultConfiguration.OutputLocation)
				}
			}
		}
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
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
