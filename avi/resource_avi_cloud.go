/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi


import (
    "log"
    "strings"
    "github.com/avinetworks/sdk/go/clients"
    "github.com/hashicorp/terraform/helper/schema"
)
func ResourceCloudSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
             "apic_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceAPICConfigurationSchema(),
                                                        },
             "apic_mode" :{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                         Default: false,
                                                                                                                                                                        },
             "autoscale_polling_interval" :{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                         Default: 60,
                                                                                                                                                                        },
             "aws_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceAwsConfigurationSchema(),
                                                        },
             "azure_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceAzureConfigurationSchema(),
                                                        },
             "cloudstack_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceCloudStackConfigurationSchema(),
                                                        },
             "custom_tags" :{
                             Type: schema.TypeList, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceCustomTagSchema(),
                                                        },
             "dhcp_enabled" :{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                         Default: false,
                                                                                                                                                                        },
             "dns_provider_ref" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            Computed:  true,
                                                                                    },
             "docker_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceDockerConfigurationSchema(),
                                                        },
             "east_west_dns_provider_ref" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            Computed:  true,
                                                                                    },
             "east_west_ipam_provider_ref" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            Computed:  true,
                                                                                    },
             "enable_vip_static_routes" :{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                         Default: false,
                                                                                                                                                                        },
             "gcp_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceGCPConfigurationSchema(),
                                                        },
             "ip6_autocfg_enabled" :{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                         Default: false,
                                                                                                                                                                        },
             "ipam_provider_ref" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            Computed:  true,
                                                                                    },
             "license_tier" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                                                                                    },
             "license_type" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                                                                                    },
             "linuxserver_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceLinuxServerConfigurationSchema(),
                                                        },
             "mtu" :{
                             Type: schema.TypeInt, 
                             Optional: true,
                                                         Default: 1500,
                                                                                                                                                                        },
             "name" :{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
             "nsx_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceNsxConfigurationSchema(),
                                                        },
             "obj_name_prefix" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                                                                                    },
             "openstack_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceOpenStackConfigurationSchema(),
                                                        },
             "oshiftk8s_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceOShiftK8SConfigurationSchema(),
                                                        },
             "prefer_static_routes" :{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                         Default: false,
                                                                                                                                                                        },
             "proxy_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceProxyConfigurationSchema(),
                                                        },
             "rancher_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourceRancherConfigurationSchema(),
                                                        },
             "se_group_template_ref" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            Computed:  true,
                                                                                    },
             "state_based_dns_registration" :{
                             Type: schema.TypeBool, 
                             Optional: true,
                                                         Default: true,
                                                                                                                                                                        },
             "tenant_ref" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                                                                            Computed:  true,
                                                                                    },
             "uuid" :{
                             Type: schema.TypeString, 
                             Optional: true,
                                                                                    Computed:  true,
                                                                                    },
             "vca_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourcevCloudAirConfigurationSchema(),
                                                        },
             "vcenter_configuration" :{
                             Type: schema.TypeSet, 
                             Optional: true,
                                                                                                                                                                         Elem: ResourcevCenterConfigurationSchema(),
                                                        },
             "vtype" :{
                             Type: schema.TypeString, 
                             Required: true,
                                                                                                                },
        }
}


func resourceAviCloud() *schema.Resource {
    return &schema.Resource{
        Create: resourceAviCloudCreate,
        Read:   ResourceAviCloudRead,
        Update: resourceAviCloudUpdate,
        Delete: resourceAviCloudDelete,
        Schema: ResourceCloudSchema(),
        Importer: &schema.ResourceImporter{
                State: ResourceCloudImporter,
        },
    }
}

func ResourceCloudImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
    s := ResourceCloudSchema()
    return ResourceImporter(d, m, "cloud", s)
}

func ResourceAviCloudRead(d *schema.ResourceData, meta interface{}) error {
    s := ResourceCloudSchema()
    err := ApiRead(d, meta, "cloud", s)
    if err != nil {
        log.Printf("[ERROR] in reading object %v\n", err)
    }
    return err
}

func resourceAviCloudCreate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceCloudSchema()
    err := ApiCreateOrUpdate(d, meta, "cloud", s)
        if err == nil {
        err = ResourceAviCloudRead(d, meta)
    }
    return err
}

func resourceAviCloudUpdate(d *schema.ResourceData, meta interface{}) error {
    s := ResourceCloudSchema()
    var err error
    err = ApiCreateOrUpdate(d, meta, "cloud", s)
    if err == nil {
        err = ResourceAviCloudRead(d, meta)
    }
    return err
}

func resourceAviCloudDelete(d *schema.ResourceData, meta interface{}) error {
    objType := "cloud"
    if ApiDeleteSystemDefaultCheck(d) {
        return nil
    }
    client := meta.(*clients.AviClient)
    uuid := d.Get("uuid").(string)
    sess_tenant := client.AviSession.GetTenant()
    if uuid != "" {
        path := "api/" + objType + "/" + uuid
        err := client.AviSession.Delete(path, sess_tenant)
        if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
            log.Println("[INFO] resourceAviCloudDelete not found")
            return err
        }
        d.SetId("")
    }
    return nil
}


