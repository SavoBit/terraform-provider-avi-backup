
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
page_title: "AVI: avi_gslb"
sidebar_current: "docs-avi-datasource-gslb"
description: |-
  Get information of Avi Gslb.
---

# avi_gslb

This data source is used to to get avi_gslb objects.

## Example Usage

```hcl
data "Gslb" "foo_Gslb" {
    uuid = "Gslb-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Gslb by name.
* `uuid` - (Optional) Search Gslb by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `clear_on_max_retries` - Max retries after which the remote site is treated as a fresh start.
* `client_ip_addr_group` - Group to specify if the client ip addresses are public or private.
* `description` - General description.
* `dns_configs` - Sub domain configuration for the gslb.
* `is_federated` - This field indicates that this object is replicated across gslb federation.
* `leader_cluster_uuid` - Mark this site as leader of gslb configuration.
* `maintenance_mode` - This field disables the configuration operations on the leader for all federated objects.
* `name` - Name for the gslb object.
* `send_interval` - Frequency with which group members communicate.
* `send_interval_prior_to_maintenance_mode` - The user can specify a send-interval while entering maintenance mode.
* `sites` - Select avi site member belonging to this gslb.
* `tenant_ref` - It is a reference to an object of type tenant.
* `third_party_sites` - Third party site member belonging to this gslb.
* `uuid` - Uuid of the gslb object.
* `view_id` - The view-id is used in change-leader mode to differentiate partitioned groups while they have the same gslb namespace.

