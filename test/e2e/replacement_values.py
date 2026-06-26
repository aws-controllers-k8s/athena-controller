# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.
"""Stores the values used by each of the integration tests for replacing the
Athena-specific test variables.
"""

from acktest.aws import identity

# LAMBDA and HIVE data catalogs require a Lambda function ARN in their
# parameters. Athena does not validate that the function exists at catalog
# creation time, but the ARN must be well-formed, so build it from the current
# test account and region.
LAMBDA_FUNCTION_ARN = (
    f"arn:aws:lambda:{identity.get_region()}:{identity.get_account_id()}"
    ":function:ack-athena-e2e-data-catalog"
)

REPLACEMENT_VALUES = {
    "LAMBDA_FUNCTION_ARN": LAMBDA_FUNCTION_ARN,
}
