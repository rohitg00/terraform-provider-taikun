package taikun

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccDataSourceTaikunBillingCredentialsConfig = `
resource "taikun_billing_credential" "foo" {
  name = "%s"

  prometheus_password = "%s"
  prometheus_url      = "%s"
  prometheus_username = "%s"
}

data "taikun_billing_credentials" "all" {
   depends_on = [
    taikun_billing_credential.foo
  ]
}`

func TestAccDataSourceTaikunBillingCredentials(t *testing.T) {
	billingCredentialName := randomTestName()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckPrometheus(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataSourceTaikunBillingCredentialsConfig,
					billingCredentialName,
					os.Getenv("PROMETHEUS_PASSWORD"),
					os.Getenv("PROMETHEUS_URL"),
					os.Getenv("PROMETHEUS_USERNAME"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.taikun_billing_credentials.all", "id", "all"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.#"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.created_by"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.id"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.is_locked"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.is_default"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.name"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.organization_id"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.organization_name"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.prometheus_password"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.prometheus_url"),
					resource.TestCheckResourceAttrSet("data.taikun_billing_credentials.all", "billing_credentials.0.prometheus_username"),
				),
			},
		},
	})
}
