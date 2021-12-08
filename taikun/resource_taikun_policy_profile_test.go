package taikun

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/itera-io/taikungoclient/client/opa_profiles"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/itera-io/taikungoclient/models"
)

func init() {
	resource.AddTestSweepers("taikun_policy_profile", &resource.Sweeper{
		Name:         "taikun_policy_profile",
		Dependencies: []string{"taikun_project"},
		F: func(r string) error {

			meta, err := sharedConfig()
			if err != nil {
				return err
			}
			apiClient := meta.(*apiClient)

			params := opa_profiles.NewOpaProfilesListParams().WithV(ApiVersion)

			var PolicyProfilesList []*models.OpaProfileListDto
			for {
				response, err := apiClient.client.OpaProfiles.OpaProfilesList(params, apiClient)
				if err != nil {
					return err
				}
				PolicyProfilesList = append(PolicyProfilesList, response.GetPayload().Data...)
				if len(PolicyProfilesList) == int(response.GetPayload().TotalCount) {
					break
				}
				offset := int32(len(PolicyProfilesList))
				params = params.WithOffset(&offset)
			}

			for _, e := range PolicyProfilesList {
				if strings.HasPrefix(e.Name, testNamePrefix) {
					params := opa_profiles.NewOpaProfilesDeleteParams().WithV(ApiVersion).WithBody(&models.DeleteOpaProfileCommand{ID: e.ID})
					_, err = apiClient.client.OpaProfiles.OpaProfilesDelete(params, apiClient)
					if err != nil {
						return err
					}
				}
			}

			return nil
		},
	})
}

const testAccResourceTaikunPolicyProfileConfig = `
resource "taikun_policy_profile" "foo" {
  name = "%s"
  lock = %t

  forbid_node_port = %t
  forbid_http_ingress = %t
  require_probe = %t
  unique_ingress = %t
  unique_service_selector = %t

  %s
}
`

func TestAccResourceTaikunPolicyProfile(t *testing.T) {
	firstName := randomTestName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckPrometheus(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckTaikunPolicyProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccResourceTaikunPolicyProfileConfig,
					firstName,
					false,
					false,
					false,
					false,
					false,
					false,
					"",
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckTaikunPolicyProfileExists,
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "name", firstName),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "lock", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_node_port", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_http_ingress", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "require_probe", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_ingress", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_service_selector", "false"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_id"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_name"),
				),
			},
			{
				ResourceName:      "taikun_policy_profile.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceTaikunPolicyProfileLock(t *testing.T) {
	firstName := randomTestName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckPrometheus(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckTaikunPolicyProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccResourceTaikunPolicyProfileConfig,
					firstName,
					false,
					true,
					true,
					true,
					true,
					true,
					"",
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckTaikunPolicyProfileExists,
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "name", firstName),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "lock", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_node_port", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_http_ingress", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "require_probe", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_ingress", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_service_selector", "true"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_id"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_name"),
				),
			},
			{
				Config: fmt.Sprintf(testAccResourceTaikunPolicyProfileConfig,
					firstName,
					true,
					true,
					true,
					true,
					true,
					true,
					"",
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckTaikunPolicyProfileExists,
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "name", firstName),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "lock", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_node_port", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_http_ingress", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "require_probe", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_ingress", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_service_selector", "true"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_id"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_name"),
				),
			},
		},
	})
}

func TestAccResourceTaikunPolicyProfileUpdate(t *testing.T) {
	firstName := randomTestName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckPrometheus(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckTaikunPolicyProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(
					testAccResourceTaikunPolicyProfileConfig,
					firstName,
					false,
					true,
					true,
					true,
					true,
					true,
					"forbidden_tags = [\"tag1\", \"tag2\"]",
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckTaikunPolicyProfileExists,
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "name", firstName),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "lock", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_node_port", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_http_ingress", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "require_probe", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_ingress", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_service_selector", "true"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_id"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_name"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbidden_tags.#", "2"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbidden_tags.0", "tag1"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbidden_tags.1", "tag2"),
				),
			},
			{
				Config: fmt.Sprintf(testAccResourceTaikunPolicyProfileConfig,
					firstName,
					false,
					true,
					false,
					true,
					false,
					true,
					"forbidden_tags = [\"tag3\"]",
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckTaikunPolicyProfileExists,
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "name", firstName),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "lock", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_node_port", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbid_http_ingress", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "require_probe", "true"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_ingress", "false"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "unique_service_selector", "true"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_id"),
					resource.TestCheckResourceAttrSet("taikun_policy_profile.foo", "organization_name"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbidden_tags.#", "1"),
					resource.TestCheckResourceAttr("taikun_policy_profile.foo", "forbidden_tags.0", "tag3"),
				),
			},
		},
	})
}

func testAccCheckTaikunPolicyProfileExists(state *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "taikun_policy_profile" {
			continue
		}

		id, _ := atoi32(rs.Primary.ID)
		params := opa_profiles.NewOpaProfilesListParams().WithV(ApiVersion).WithID(&id)

		response, err := client.client.OpaProfiles.OpaProfilesList(params, client)
		if err != nil || response.Payload.TotalCount != 1 {
			return fmt.Errorf("policy profile doesn't exist (id = %s)", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckTaikunPolicyProfileDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*apiClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "taikun_policy_profile" {
			continue
		}

		retryErr := resource.RetryContext(context.Background(), getReadAfterOpTimeout(false), func() *resource.RetryError {
			id, _ := atoi32(rs.Primary.ID)
			params := opa_profiles.NewOpaProfilesListParams().WithV(ApiVersion).WithID(&id)

			response, err := client.client.OpaProfiles.OpaProfilesList(params, client)
			if err != nil {
				return resource.NonRetryableError(err)
			}
			if response.Payload.TotalCount != 0 {
				return resource.RetryableError(errors.New("policy profile still exists"))
			}
			return nil
		})
		if timedOut(retryErr) {
			return errors.New("policy profile still exists (timed out)")
		}
		if retryErr != nil {
			return retryErr
		}
	}

	return nil
}
