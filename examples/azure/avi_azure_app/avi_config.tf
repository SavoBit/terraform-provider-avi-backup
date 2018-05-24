provider "azurerm" {
  subscription_id = "${var.subscription_id}"
  client_id 		  = "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		  = "${var.tenant_id}"
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.azurerm_network_interface.controller_nic.private_ip_address}"
  avi_tenant     = "admin"
  avi_version    = "${var.avi_version}"
}

data "avi_applicationprofile" "system_http_profile" {
  name = "System-HTTP"
}

data "avi_applicationprofile" "system_https_profile" {
  name = "System-Secure-HTTP"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_cloud" "azure_cloud_cfg" {
  name = "AZURE"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"
}

data "avi_healthmonitor" "system_http_healthmonitor" {
  name = "System-HTTP"
}

data "avi_networkprofile" "system_tcp_profile" {
  name = "System-TCP-Proxy"
}

data "avi_analyticsprofile" "system_analytics_profile" {
  name = "System-Analytics-Profile"
}

data "avi_sslkeyandcertificate" "system_default_cert" {
  name = "System-Default-Cert"
}

data "avi_sslprofile" "system_standard_sslprofile" {
  name = "System-Standard"
}

data "avi_serviceenginegroup" "se_group" {
  name      = "Default-Group"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"
}


resource "avi_pool" "azure-pool-v1" {
  name                      = "azure_poolv1"
  health_monitor_refs       = ["${data.avi_healthmonitor.system_http_healthmonitor.id}"]
  server_count              = 0
  tenant_ref                = "${data.avi_tenant.default_tenant.id}"
  vrf_ref                   = "${data.avi_vrfcontext.terraform_vrf.id}"
  cloud_ref                 = "${data.avi_cloud.azure_cloud_cfg.id}"
  //external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.resource_group_name}"]
  external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.project_name}-terraform-resource-group"]
  fail_action = {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "azure-pool-v2" {
  name                      = "azure_poolv2"
  health_monitor_refs       = ["${data.avi_healthmonitor.system_http_healthmonitor.id}"]
  server_count              = 0
  tenant_ref                = "${data.avi_tenant.default_tenant.id}"
  vrf_ref                   = "${data.avi_vrfcontext.terraform_vrf.id}"
  cloud_ref                 = "${data.avi_cloud.azure_cloud_cfg.id}"
  //external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}@${var.resource_group_name}"]
  external_autoscale_groups = ["${azurerm_virtual_machine_scale_set.terraform_scale_set_v2.name}@${var.project_name}-terraform-resource-group"]
  fail_action = {
    type = "FAIL_ACTION_CLOSE_CONN"
  }

}

resource "azurerm_virtual_machine_scale_set" "terraform_scale_set_v1" {
  name                = "${var.project_name}_scale_set_v1"
  location            = "${var.location}"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
  upgrade_policy_mode = "Manual"

  sku {
    name     = "Standard_A0"
    tier     = "Standard"
    capacity = 1
  }

  storage_profile_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_profile_os_disk {
    name              = ""
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_profile_data_disk {
    lun           = 0
    caching       = "ReadWrite"
    create_option = "Empty"
    disk_size_gb  = 10
  }

  os_profile {
    computer_name_prefix = "${var.project_name}testvm"
    admin_username       = "myadmin"
    admin_password       = "${var.azure_vm_password}"
  }

  network_profile {
    name    = "${var.project_name}terraformnetworkprofile"
    primary = true
    ip_configuration {
      name      = "${var.project_name}TestIPConfiguration"
      subnet_id = "${data.azurerm_subnet.terraform_subnet.id}"
    }
  }

  os_profile_linux_config {
    disable_password_authentication = false
    ssh_keys {
      path     = "/home/myadmin/.ssh/authorized_keys"
      key_data = "${file("~/.ssh/id_rsa.pub")}"
    }
  }

  extension {
    name                 = "vmssextension"
    publisher            = "Microsoft.OSTCExtensions"
    type                 = "CustomScriptForLinux"
    type_handler_version = "1.2"
    settings = <<SETTINGS
    {
        "commandToExecute": "apt-get -y install nginx"
    }
    SETTINGS
  }

}

output "azure_scale_set1" {
  value = "${azurerm_virtual_machine_scale_set.terraform_scale_set_v1.name}"
}

resource "azurerm_virtual_machine_scale_set" "terraform_scale_set_v2" {
  name                = "${var.project_name}_scale_set_v2"
  location            = "${var.location}"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
  upgrade_policy_mode = "Manual"

  sku {
    name     = "Standard_A0"
    tier     = "Standard"
    capacity = 1
  }

  storage_profile_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  storage_profile_os_disk {
    name              = ""
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_profile_data_disk {
    lun           = 0
    caching       = "ReadWrite"
    create_option = "Empty"
    disk_size_gb  = 10
  }

  os_profile {
    computer_name_prefix = "${var.project_name}testvm"
    admin_username       = "myadmin"
    admin_password       = "${var.azure_vm_password}"
  }

  network_profile {
    name    = "${var.project_name}terraformnetworkprofile"
    primary = true
    ip_configuration {
      name      = "${var.project_name}TestIPConfiguration"
      subnet_id = "${data.azurerm_subnet.terraform_subnet.id}"
    }
  }

  os_profile_linux_config {
    disable_password_authentication = false
    ssh_keys {
      path     = "/home/myadmin/.ssh/authorized_keys"
      key_data = "${file("~/.ssh/id_rsa.pub")}"
    }
  }

  extension {
    name      = "vmssextension"
    publisher = "Microsoft.OSTCExtensions"
    type      = "CustomScriptForLinux"
    type_handler_version = "1.2"
    settings = <<SETTINGS
    {
        "commandToExecute": "apt-get -y install nginx"
    }
    SETTINGS
  }

}

resource "avi_poolgroup" "azure-poolgroup" {
  name       = "azure_poolgroup"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"

  members = {
    pool_ref = "${avi_pool.azure-pool-v1.id}"
    ratio    = 100
  }

  members = {
    pool_ref = "${avi_pool.azure-pool-v2.id}"
    ratio    = 100
  }
}

resource "avi_virtualservice" "azure-virtualservice" {
  name                         = "azure_vs"
  pool_group_ref               = "${avi_poolgroup.azure-poolgroup.id}"
  tenant_ref                   = "${data.avi_tenant.default_tenant.id}"
  cloud_type                   = "CLOUD_AZURE"
  application_profile_ref      = "${data.avi_applicationprofile.system_https_profile.id}"
  network_profile_ref          = "${data.avi_networkprofile.system_tcp_profile.id}"
  cloud_ref                    = "${data.avi_cloud.azure_cloud_cfg.id}"
  analytics_profile_ref        = "${data.avi_analyticsprofile.system_analytics_profile.id}"
  ssl_key_and_certificate_refs = ["${data.avi_sslkeyandcertificate.system_default_cert.id}"]
  ssl_profile_ref              = "${data.avi_sslprofile.system_standard_sslprofile.id}"
  se_group_ref                 = "${data.avi_serviceenginegroup.se_group.id}"
  vrf_context_ref              = "${data.avi_vrfcontext.terraform_vrf.id}"
  scaleout_ecmp                = true
  enabled                      = true
  // vsvip_ref = "${avi_vsvip.azure-vs-vsvip.id}"
  vip {
    auto_allocate_ip  = true
    avi_allocated_vip = true
    avi_allocated_fip = false
    # auto_allocate_floating_ip = true
    subnet_uuid       = "${data.azurerm_subnet.terraform_subnet.name}"
    enabled = true
    subnet = {
      ip_addr = {
        addr = "${var.azure_vip_subnet_ip}"
        type = "V4"
      }
      mask = "${var.azure_vip_subnet_mask}"
    }
  }

  services {
    port           = 80
    enable_ssl     = true
    port_range_end = 80
  }
  services {
    port           = 443
    enable_ssl     = true
    port_range_end = 443
  }
  analytics_policy {
    enabled = true
    metrics_realtime_update = {
      enabled  = true
      duration = 0
    }
  }
}
