package repository

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudscale_server", func(r *config.Resource) {
		r.ShortGroup = ""

		r.References["interfaces.network_uuid"] = config.Reference{
			Type: "github.com/splattner/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Network",
		}

		r.References["interfaces.addresses.subnet_uuid"] = config.Reference{
			Type: "github.com/splattner/provider-cloudscale/apis/cluster/cloudscale/v1alpha1.Subnet",
		}

	})
}
