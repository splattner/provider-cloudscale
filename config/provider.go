package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	networkCluster "github.com/splattner/provider-cloudscale/config/cluster/network"
	objectuserCluster "github.com/splattner/provider-cloudscale/config/cluster/objectsuser"
	serverCluster "github.com/splattner/provider-cloudscale/config/cluster/server"
	subnetCluster "github.com/splattner/provider-cloudscale/config/cluster/subnet"
	networkNamespaced "github.com/splattner/provider-cloudscale/config/namespaced/network"
	objectuserNamespaced "github.com/splattner/provider-cloudscale/config/namespaced/objectsuser"
	serverNamespaced "github.com/splattner/provider-cloudscale/config/namespaced/server"
	subnetNamespaced "github.com/splattner/provider-cloudscale/config/namespaced/subnet"
)

const (
	resourcePrefix = "cloudscale"
	modulePath     = "github.com/splattner/provider-cloudscale"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudscale.ch"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectuserCluster.Configure,
		networkCluster.Configure,
		serverCluster.Configure,
		subnetCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("m.cloudscale.ch"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectuserNamespaced.Configure,
		networkNamespaced.Configure,
		serverNamespaced.Configure,
		subnetNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
