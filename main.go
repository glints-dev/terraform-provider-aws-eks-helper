package main

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/glints-dev/terraform-provider-aws-eks-helper/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				DataSourcesMap: map[string]*schema.Resource{
					"aws_eks_helper_kube_reserved": dataSourceKubeReserved(),
				},
			}
		},
	})
}

func dataSourceKubeReserved() *schema.Resource {
	return &schema.Resource{
		Description: "`aws_eks_helper_kube_reserved` data source can be used to retrieve resources that should be reserved by Kubernetes for a given instance type.",
		ReadContext: dataSourceKubeReservedRead,
		Schema: map[string]*schema.Schema{
			"instance_type": {
				Description: "The type of instance.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"region": {
				Description: "The region to query instance information from.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cpu": {
				Description: "The amount of CPU (in milli-cores) that should be reserved.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"memory": {
				Description: "The amount of memory (in Mi) that should be reserved.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"storage": {
				Description: "The amount of storage (in Gi) that should be reserved.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"max_pods": {
				Description: "The maximum number of pods that should be scheduled.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func dataSourceKubeReservedRead(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	var diags diag.Diagnostics

	instanceType := d.Get("instance_type").(string)
	region := d.Get("region").(string)

	ec2Service := ec2.New(
		session.Must(session.NewSession()),
		aws.NewConfig().WithRegion(region),
	)

	// Query the AWS EC2 API for instance types.
	response, err := ec2Service.DescribeInstanceTypes(&ec2.DescribeInstanceTypesInput{
		InstanceTypes: []*string{&instanceType},
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if len(response.InstanceTypes) == 0 {
		return diag.Errorf("unknown instance type \"%s\"", instanceType)
	}

	// Get instance type information.
	instanceTypeInfo := provider.NewInstanceTypeInfo(response.InstanceTypes[0])

	if err := d.Set("cpu", instanceTypeInfo.DefaultCPUToReserve()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("memory", instanceTypeInfo.DefaultMemoryToReserve()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("storage", instanceTypeInfo.DefaultStorageToReserve()); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("max_pods", instanceTypeInfo.MaxPodsPerNode); err != nil {
		return diag.FromErr(err)
	}

	// Always run.
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
