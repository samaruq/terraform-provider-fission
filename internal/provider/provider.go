package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/fission/fission/pkg/controller/client"
	"github.com/fission/fission/pkg/controller/client/rest"
	"github.com/fission/fission/pkg/fission-cli/cmd"
	// "github.com/fission/fission/pkg/fission-cli/cmd/canaryconfig"
	// "github.com/fission/fission/pkg/fission-cli/cmd/environment"
	// "github.com/fission/fission/pkg/fission-cli/cmd/function"
	// "github.com/fission/fission/pkg/fission-cli/cmd/httptrigger"
	// "github.com/fission/fission/pkg/fission-cli/cmd/kubewatch"
	// "github.com/fission/fission/pkg/fission-cli/cmd/mqtrigger"
	// "github.com/fission/fission/pkg/fission-cli/cmd/package"
	// "github.com/fission/fission/pkg/fission-cli/cmd/spec"
	// "github.com/fission/fission/pkg/fission-cli/cmd/support"
	// "github.com/fission/fission/pkg/fission-cli/cmd/timetrigger"
	// "github.com/fission/fission/pkg/fission-cli/cmd/version"
	// "github.com/fission/fission/pkg/fission-cli/console"
	"github.com/fission/fission/pkg/fission-cli/util"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			// DataSourcesMap: map[string]*schema.Resource{
			//         "scaffolding_data_source": dataSourceScaffolding(),
			// },
			ResourcesMap: map[string]*schema.Resource{
				"fission_environment": resourceFissionEnvironment(),
			},
			Schema: map[string]*schema.Schema{},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	cmd.CommandActioner
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		serverUrl, err := util.GetApplicationUrl("application=fission-api", "")
		if err != nil {
			return nil, diag.FromErr(err)
		}

		restClient := rest.NewRESTClient(serverUrl)
		cmd.SetClientset(client.MakeClientset(restClient))

		return &apiClient{}, nil
	}
}
