package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func resourceFissionEnvironment() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceFissionEnvironmentCreate,
		ReadContext:   resourceFissionEnvironmentRead,
		UpdateContext: resourceFissionEnvironmentUpdate,
		DeleteContext: resourceFissionEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"namespace": {
				Description: "Namespace of the environment.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
			},
			"image": {
				Description: "Container image of the environment.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceFissionEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		name      string
		namespace string
		image     string
		diags     diag.Diagnostics
		data      interface{}
		ok        bool
	)
	if data, ok = d.GetOk("name"); ok {
		name = data.(string)
	}
	if data, ok = d.GetOk("namespace"); ok {
		namespace = data.(string)
	}
	if data, ok = d.GetOk("image"); ok {
		image = data.(string)
	}

	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	envSpec := &fv1.Environment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: fv1.EnvironmentSpec{
			Version: 1,
			Runtime: fv1.Runtime{
				Image: image,
			},
			Resources: v1.ResourceRequirements{},
		},
	}

	client.Client().V1().Environment().Create(envSpec)

	return diags
}

func resourceFissionEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceFissionEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceFissionEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
