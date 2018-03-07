/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceUserAccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"old_password": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"local": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"full_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"email": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceAviUserAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserAccountCreate,
		Read:   ResourceAviUserAccountRead,
		Update: resourceAviUserAccountUpdate,
		Delete: resourceAviUserAccountDelete,
		Schema: ResourceUserAccountSchema(),
	}
}

func ResourceAviUserAccountRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAviUserAccountCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountSchema()
	var err error
	client := meta.(*clients.AviClient)
	var robj interface{}
	obj := d
	if data, err := SchemaToAviData(obj, s); err == nil {
		path := "api/useraccount"
		err = client.AviSession.Put(path, data, &robj)
	}
	return err
}

func resourceAviUserAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAviUserAccountDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
