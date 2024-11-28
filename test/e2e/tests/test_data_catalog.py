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

"""Integration tests for the Athena DataCatalog resource"""

import time
import pytest

from acktest.k8s import condition
from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import service_marker, CRD_GROUP, CRD_VERSION, load_athena_resource
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e import data_catalog

DATA_CATALOG_RESOURCE_PLURAL = "datacatalogs"

CREATE_WAIT_SECONDS = 10
MODIFY_WAIT_SECONDS = 10
DELETE_WAIT_SECONDS = 10


@pytest.fixture(scope="module")
def simple_data_catalog():
    resources = get_bootstrap_resources()
    global functionARN
    functionARN = resources.LambdaFn.arn
    
    data_catalog_name = random_suffix_name("my-simple-dc", 24)
    
    replacements = REPLACEMENT_VALUES.copy()
    replacements["DATA_CATALOG_NAME"] = data_catalog_name
    replacements["FUNCTION_ARN"] = functionARN

    resource_data = load_athena_resource(
        "data_catalog_simple",
        additional_replacements=replacements,
    )

    ref = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        DATA_CATALOG_RESOURCE_PLURAL,
        data_catalog_name,
        namespace="default",
    )
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref)

    assert cr is not None
    assert k8s.get_resource_exists(ref)

    yield (ref, cr)

    try:
        _, deleted = k8s.delete_custom_resource(ref, DELETE_WAIT_SECONDS)
        assert deleted
    except:
        pass


@service_marker
@pytest.mark.canary
class TestDataCatalog:
    def test_crud(self, simple_data_catalog):
        ref, _ = simple_data_catalog

        time.sleep(CREATE_WAIT_SECONDS)
        condition.assert_synced(ref)

        cr = k8s.get_resource(ref)

        assert "spec" in cr
        assert "name" in cr["spec"]
        data_catalog_name = cr["spec"]["name"]

        latest = data_catalog.get(data_catalog_name)
        assert latest is not None
        
        assert "Description" in latest
        description = latest["Description"]
        assert description == "initial description"
        
        assert "Type" in latest
        dcType = latest["Type"]
        assert dcType == "HIVE"
        
        assert "Parameters" in latest
        assert len(latest["Parameters"]) == 2
        params = latest["Parameters"]
        assert "metadata-function" in params
        assert params["metadata-function"] == functionARN
        
        # update the CR
        updates = {
            "spec": {
                "description": "updated description",
                "type": "LAMBDA",
                "parameters": {
                    "metadata-function": functionARN,
                    "record-function": functionARN,
                },
            },
        }
        k8s.patch_custom_resource(ref, updates)
        time.sleep(MODIFY_WAIT_SECONDS)

        latest = data_catalog.get(data_catalog_name)
        assert latest is not None

        assert "Description" in latest
        description = latest["Description"]
        assert description == "updated description"
        
        assert "Type" in latest
        dcType = latest["Type"]
        assert dcType == "LAMBDA"
        
        assert "Parameters" in latest
        assert len(latest["Parameters"]) == 3
        params = latest["Parameters"]
        assert "metadata-function" in params
        assert "record-function" in params
        assert params["metadata-function"] == functionARN
        assert params["record-function"] == functionARN

        # delete the CR
        _, deleted = k8s.delete_custom_resource(ref, DELETE_WAIT_SECONDS)
        assert deleted
        data_catalog.wait_until_deleted(data_catalog_name)
