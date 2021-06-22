package provider

import (
	"context"
	"fmt"

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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
	// load arguments
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

	// build environment spec
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

	// create environment
	_, err := client.Client().V1().Environment().Create(envSpec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", namespace, name))

	return diags
}

func resourceFissionEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	exists, err := resourceFissionEnvironmentExists(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	if !exists {
		d.SetId("")
		return diag.Diagnostics{}
	}

	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	// build environment spec
	envMeta := metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
	}

	// get environment object
	env, err := client.Client().V1().Environment().Get(&envMeta)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("environment", env)

	return diag.Diagnostics{}
}

func resourceFissionEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		name      string
		namespace string
		image     string
		diags     diag.Diagnostics
		data      interface{}
		ok        bool
	)
	// load arguments
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

	// build environment spec
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

	// delete environment
	_, err := client.Client().V1().Environment().Update(envSpec)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildID(envSpec.ObjectMeta))

	return diag.diagnostics{}
}

func resourceFissionEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	namespace, name, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	// build environment spec
	envMeta := metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
	}

	// delete environment
	err := client.Client().V1().Environment().Delete(&envMeta)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diag.diagnostics{}
}

func resourceFissionEnvironmentExists() {}
