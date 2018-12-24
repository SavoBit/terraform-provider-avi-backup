/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"errors"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourcePoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ignore_servers": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"analytics_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcePoolAnalyticsPolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"analytics_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"apic_epg_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"application_persistence_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"autoscale_launch_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"autoscale_networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"autoscale_policy_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"capacity_estimation": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"capacity_estimation_ttfb_thresh": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"conn_pool_properties": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceConnPoolPropertiesSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"connection_ramp_duration": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"default_server_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  80,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"domain_name": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"east_west": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"external_autoscale_groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"fail_action": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceFailActionSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"fewest_tasks_feedback_delay": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"graceful_disable_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"health_monitor_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"host_check_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"inline_health_monitor": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ipaddrgroup_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"lb_algorithm": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_LEAST_CONNECTIONS",
		},
		"lb_algorithm_consistent_hash_hdr": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"lb_algorithm_core_nonaffinity": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"lb_algorithm_hash": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS",
		},
		"lookup_server_by_name": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"max_concurrent_connections_per_server": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"max_conn_rate_per_server": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceRateProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"min_health_monitors_up": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"min_servers_up": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNetworkFilterSchema(),
		},
		"nsx_securitygroup": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"pki_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"placement_networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePlacementNetworkSchema(),
		},
		"request_queue_depth": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  128,
		},
		"request_queue_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"rewrite_host_header_to_server_name": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"rewrite_host_header_to_sni": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"server_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"server_reselect": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPServerReselectSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"server_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServerSchema(),
		},
		"service_metadata": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"sni_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ssl_key_and_certificate_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssl_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_service_port": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vrf_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolCreate,
		Read:   ResourceAviPoolRead,
		Update: resourceAviPoolUpdate,
		Delete: resourceAviPoolDelete,
		Schema: ResourcePoolSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolImporter,
		},
	}
}

func ResourcePoolImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolSchema()
	return ResourceImporter(d, m, "pool", s)
}

func ResourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	err := ApiRead(d, meta, "pool", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	if ignore_servers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignore_servers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err := errors.New("Error Invalid Plan. cannot set ignore_servers and servers together.")
			return err
		}
	}
	err := ApiCreateOrUpdate(d, meta, "pool", s)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}
	return err
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	var err error
	set_ignore_servers := false
	s := ResourcePoolSchema()

	if ignore_servers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignore_servers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err = errors.New("Error Invalid Plan. cannot set ignore_servers and servers together.")
			return err
		}
		client := meta.(*clients.AviClient)
		pUUID := UUIDFromID(d.Id())
		path := "api/pool" + "/" + pUUID
		log.Printf("[DEBUG] resourceAviPoolUpdate reading object with id %v\n", pUUID)
		var obj interface{}
		var apiResponse interface{}
		err = client.AviSession.Get(path, &obj)
		if err == nil {
			// found pool so unpack it
			td := make(map[string]interface{})
			if localData, err := SchemaToAviData(d, s); err == nil {
				apiResponse, err = SetDefaultsInAPIRes(obj, localData, s)
			} else {
				log.Printf("[ERROR] resourceAviPool in SchemaToAviData: %v\n", err)
			}
			if _, err := ApiDataToSchema(apiResponse, td, s); err == nil {
				log.Printf("[DEBUG] read servers %v from existing object ", td["servers"])
				d.Set("servers", td["servers"])
			}
		} else {
			log.Printf("[ERROR] ApiRead object with uuid %v not found err %v\n", pUUID, err)
		}
		set_ignore_servers = true
	}
	err = ApiCreateOrUpdate(d, meta, "pool", s)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}
	d.Set("ignore_servers", set_ignore_servers)
	return err
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	if ignore_servers, ok := d.GetOk("ignore_servers"); ok {
		if servers, ok := d.GetOk("servers"); ok && ignore_servers.(bool) && servers != nil {
			log.Printf("[ERROR] cannot set ignore_servers and servers together.")
			err := errors.New("Error Invalid Plan. cannot set ignore_servers and servers together.")
			return err
		}
	}
	objType := "pool"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[ERROR] resourceAviPoolDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
