package local_pub_repository

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("artifactory_local_pub_repository", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "artifactory"
        r.ShortGroup = "local_pub_repository"
    })
}
