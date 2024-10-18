package artifact_property_webhook

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("artifactory_artifact_property_webhook", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "artifactory"
        r.ShortGroup = "artifact_property_webhook"
    })
}
