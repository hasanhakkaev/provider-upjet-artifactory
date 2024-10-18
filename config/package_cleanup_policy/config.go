package package_cleanup_policy

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("artifactory_package_cleanup_policy", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "artifactory"
        r.ShortGroup = "package_cleanup_policy"
    })
}
