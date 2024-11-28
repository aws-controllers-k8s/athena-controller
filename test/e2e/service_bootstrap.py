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
"""Bootstraps the resources required to run the Athena integration tests.
"""
import logging

from acktest.bootstrapping import Resources, BootstrapFailureException

from e2e import bootstrap_directory
from e2e.bootstrap_resources import BootstrapResources
from acktest.bootstrapping.function import Function
from acktest.aws.identity import get_region, get_account_id

def service_bootstrap() -> Resources:
    logging.getLogger().setLevel(logging.INFO)

    aws_region = get_region()
    account_id = get_account_id()
    code_uri=f"{account_id}.dkr.ecr.{aws_region}.amazonaws.com/ack-e2e-testing-athena-controller:v1"
    
    resources = BootstrapResources(
        LambdaFn=Function(
            name_prefix="aws-athena-data-catalog-",
            code_uri=code_uri,
            service="athena",
        ),
    )

    try:
        resources.bootstrap()
    except BootstrapFailureException as ex:
        exit(254)

    return resources

if __name__ == "__main__":
    config = service_bootstrap()
    # Write config to current directory by default
    config.serialize(bootstrap_directory)
