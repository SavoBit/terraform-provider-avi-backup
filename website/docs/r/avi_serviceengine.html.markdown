
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

---
layout: "avi"
page_title: "Avi: avi_serviceengine"
sidebar_current: "docs-avi-resource-serviceengine"
description: |-
  Creates and manages Avi ServiceEngine.
---

# avi_serviceengine

The ServiceEngine resource allows the creation and management of Avi ServiceEngine

## Example Usage

```hcl
resource "ServiceEngine" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `availability_zone` - (Optional ) argument_description.
        * `cloud_ref` - (Optional ) argument_description.
        * `container_mode` - (Optional ) argument_description.
        * `container_type` - (Optional ) argument_description.
        * `controller_created` - (Optional ) argument_description.
        * `controller_ip` - (Optional ) argument_description.
        * `data_vnics` - (Optional ) argument_description.
        * `enable_state` - (Optional ) argument_description.
        * `flavor` - (Optional ) argument_description.
        * `host_ref` - (Optional ) argument_description.
        * `hypervisor` - (Optional ) argument_description.
        * `mgmt_vnic` - (Optional ) argument_description.
        * `name` - (Optional ) argument_description.
        * `resources` - (Optional ) argument_description.
        * `se_group_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                    * `uuid` - argument_description.
    
