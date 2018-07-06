// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerPolicies_basic(t *testing.T) {
	providers := testAccProviders
	config := legacyTestProviderConfig() + `
	data "oci_load_balancer_policies" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	resourceName := "data.oci_load_balancer_policies.t"

	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.#"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.name"),
				),
			},
		},
	})
}
