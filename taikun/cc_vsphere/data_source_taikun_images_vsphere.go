package cc_vsphere

import (
	"context"
	tk "github.com/itera-io/taikungoclient"
	"github.com/itera-io/terraform-provider-taikun/taikun/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTaikunImagesVsphere() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve images for a given vSphere cloud credential.",
		ReadContext: dataSourceTaikunImagesVsphereRead,
		Schema: map[string]*schema.Schema{
			"cloud_credential_id": {
				Description:      "vSphere cloud credential ID.",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: utils.StringIsInt,
			},
			"images": {
				Description: "List of retrieved vSphere images.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "vSphere image ID.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "vSphere image name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTaikunImagesVsphereRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudCredentialID, err := utils.Atoi32(d.Get("cloud_credential_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	var offset int32 = 0

	apiClient := meta.(*tk.Client)
	params := apiClient.Client.ImagesAPI.ImagesVsphereImages(ctx, cloudCredentialID)

	var imageList []map[string]interface{}
	for {
		response, res, err := params.Offset(offset).Execute()
		if err != nil {
			return diag.FromErr(tk.CreateError(res, err))
		}
		data := response.GetData()
		imageList = append(imageList, utils.FlattenTaikunImages(data...)...)
		if len(imageList) == int(response.GetTotalCount()) {
			break
		}
		offset = int32(len(imageList))
	}

	if err := d.Set("images", imageList); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(utils.I32toa(cloudCredentialID))
	return nil
}
