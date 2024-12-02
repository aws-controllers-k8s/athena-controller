# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Utilities for working with DataCatalog resources"""

import datetime
import time

import boto3
import pytest

DELETE_WAIT_TIME_TIMEOUT_SECONDS = 60 * 10


def wait_until_deleted(
    name: str,
    timeout_seconds: int = DELETE_WAIT_TIME_TIMEOUT_SECONDS,
    interval_seconds: int = 15,
) -> None:
    """Waits until a DataCatalog with a supplied Name is no longer returned from
    the Athena API.

    Raises:
        pytest.fail upon timeout
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while True:
        if datetime.datetime.now() >= timeout:
            pytest.fail(
                "Timed out waiting for DataCatalog to be " "deleted in Athena API"
            )
        time.sleep(interval_seconds)

        latest = get(name)
        if latest is None:
            break


def get(
    name: str,
):
    c = boto3.client("athena")
    try:
        resp = c.get_data_catalog(Name=name)
        return resp["DataCatalog"]
    except c.exceptions.InvalidRequestException:
        return None
