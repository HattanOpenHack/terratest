// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureLoadBalancerExample(t *testing.T) {
	t.Parallel()

	// loadbalancer::tag::1:: Configure Terraform setting up a path to Terraform code.
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/terraform-azure-loadbalancer-example",
	}

	// loadbalancer::tag::4:: At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// loadbalancer::tag::2:: Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// loadbalancer::tag::3:: Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	loadBalancer01Name := terraform.Output(t, terraformOptions, "loadbalancer01_name")
	loadBalancer02Name := terraform.Output(t, terraformOptions, "loadbalancer02_name")

	// loadbalancer::tag::5 Set expected variables for test

	// happy path tests

	// load balancer 01 (with Public IP) exists
	lb01Exists, err1 := azure.GetLoadBalancerE(loadBalancer01Name, resourceGroupName, "")
	assert.NoError(t, err1, "Load Balancer error.")
	assert.True(t, lb01Exists)

	// load balancer 02 (with Private IP on vnet/subnet) exists
	lb02Exists, err2 := azure.GetLoadBalancerE(loadBalancer02Name, resourceGroupName, "")
	assert.NoError(t, err2, "Load Balancer error.")
	assert.True(t, lb02Exists)

}
