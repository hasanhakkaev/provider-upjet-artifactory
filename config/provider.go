/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/upbound/provider-artifactory/config/artifactcustomwebhook"
	"github.com/upbound/provider-artifactory/config/artifactpropertycustomwebhook"
	"github.com/upbound/provider-artifactory/config/artifactpropertywebhook"
	"github.com/upbound/provider-artifactory/config/artifactwebhook"
	"github.com/upbound/provider-artifactory/config/backup"
	"github.com/upbound/provider-artifactory/config/buildcustomwebhook"
	"github.com/upbound/provider-artifactory/config/buildwebhook"
	"github.com/upbound/provider-artifactory/config/certificate"
	"github.com/upbound/provider-artifactory/config/distributioncustomwebhook"
	"github.com/upbound/provider-artifactory/config/distributionpublickey"
	"github.com/upbound/provider-artifactory/config/distributionwebhook"
	"github.com/upbound/provider-artifactory/config/dockercustomwebhook"
	"github.com/upbound/provider-artifactory/config/dockerwebhook"
	"github.com/upbound/provider-artifactory/config/federatedalpinerepository"
	"github.com/upbound/provider-artifactory/config/federatedbowerrepository"
	"github.com/upbound/provider-artifactory/config/federatedcargorepository"
	"github.com/upbound/provider-artifactory/config/federatedchefrepository"
	"github.com/upbound/provider-artifactory/config/federatedcocoapodsrepository"
	"github.com/upbound/provider-artifactory/config/federatedcomposerrepository"
	"github.com/upbound/provider-artifactory/config/federatedconanrepository"
	"github.com/upbound/provider-artifactory/config/federatedcondarepository"
	"github.com/upbound/provider-artifactory/config/federatedcranrepository"
	"github.com/upbound/provider-artifactory/config/federateddebianrepository"
	"github.com/upbound/provider-artifactory/config/federateddockerrepository"
	"github.com/upbound/provider-artifactory/config/federateddockerv1repository"
	"github.com/upbound/provider-artifactory/config/federateddockerv2repository"
	"github.com/upbound/provider-artifactory/config/federatedgemsrepository"
	"github.com/upbound/provider-artifactory/config/federatedgenericrepository"
	"github.com/upbound/provider-artifactory/config/federatedgitlfsrepository"
	"github.com/upbound/provider-artifactory/config/federatedgorepository"
	"github.com/upbound/provider-artifactory/config/federatedgradlerepository"
	"github.com/upbound/provider-artifactory/config/federatedhelmrepository"
	"github.com/upbound/provider-artifactory/config/federatedivyrepository"
	"github.com/upbound/provider-artifactory/config/federatedmavenrepository"
	"github.com/upbound/provider-artifactory/config/federatednpmrepository"
	"github.com/upbound/provider-artifactory/config/federatednugetrepository"
	"github.com/upbound/provider-artifactory/config/federatedopkgrepository"
	"github.com/upbound/provider-artifactory/config/federatedpuppetrepository"
	"github.com/upbound/provider-artifactory/config/federatedpypirepository"
	"github.com/upbound/provider-artifactory/config/federatedrpmrepository"
	"github.com/upbound/provider-artifactory/config/federatedsbtrepository"
	"github.com/upbound/provider-artifactory/config/federatedswiftrepository"
	"github.com/upbound/provider-artifactory/config/federatedterraformmodulerepository"
	"github.com/upbound/provider-artifactory/config/federatedterraformproviderrepository"
	"github.com/upbound/provider-artifactory/config/federatedvagrantrepository"
	"github.com/upbound/provider-artifactory/config/generalsecurity"
	"github.com/upbound/provider-artifactory/config/globalenvironment"
	"github.com/upbound/provider-artifactory/config/group"
	"github.com/upbound/provider-artifactory/config/keypair"
	"github.com/upbound/provider-artifactory/config/ldapgroupsetting"
	"github.com/upbound/provider-artifactory/config/ldapgroupsettingv2"
	"github.com/upbound/provider-artifactory/config/ldapsetting"
	"github.com/upbound/provider-artifactory/config/ldapsettingv2"
	"github.com/upbound/provider-artifactory/config/localalpinerepository"
	"github.com/upbound/provider-artifactory/config/localbowerrepository"
	"github.com/upbound/provider-artifactory/config/localcargorepository"
	"github.com/upbound/provider-artifactory/config/localchefrepository"
	"github.com/upbound/provider-artifactory/config/localcocoapodsrepository"
	"github.com/upbound/provider-artifactory/config/localcomposerrepository"
	"github.com/upbound/provider-artifactory/config/localconanrepository"
	"github.com/upbound/provider-artifactory/config/localcondarepository"
	"github.com/upbound/provider-artifactory/config/localcranrepository"
	"github.com/upbound/provider-artifactory/config/localdebianrepository"
	"github.com/upbound/provider-artifactory/config/localdockerv1repository"
	"github.com/upbound/provider-artifactory/config/localdockerv2repository"
	"github.com/upbound/provider-artifactory/config/localgemsrepository"
	"github.com/upbound/provider-artifactory/config/localgenericrepository"
	"github.com/upbound/provider-artifactory/config/localgitlfsrepository"
	"github.com/upbound/provider-artifactory/config/localgorepository"
	"github.com/upbound/provider-artifactory/config/localgradlerepository"
	"github.com/upbound/provider-artifactory/config/localhelmrepository"
	"github.com/upbound/provider-artifactory/config/localhuggingfacemlrepository"
	"github.com/upbound/provider-artifactory/config/localivyrepository"
	"github.com/upbound/provider-artifactory/config/localmavenrepository"
	"github.com/upbound/provider-artifactory/config/localnpmrepository"
	"github.com/upbound/provider-artifactory/config/localnugetrepository"
	"github.com/upbound/provider-artifactory/config/localopkgrepository"
	"github.com/upbound/provider-artifactory/config/localpubrepository"
	"github.com/upbound/provider-artifactory/config/localpuppetrepository"
	"github.com/upbound/provider-artifactory/config/localpypirepository"
	"github.com/upbound/provider-artifactory/config/localrepositorymultireplication"
	"github.com/upbound/provider-artifactory/config/localrepositorysinglereplication"
	"github.com/upbound/provider-artifactory/config/localrpmrepository"
	"github.com/upbound/provider-artifactory/config/localsbtrepository"
	"github.com/upbound/provider-artifactory/config/localswiftrepository"
	"github.com/upbound/provider-artifactory/config/localterraformbackendrepository"
	"github.com/upbound/provider-artifactory/config/localterraformmodulerepository"
	"github.com/upbound/provider-artifactory/config/localterraformproviderrepository"
	"github.com/upbound/provider-artifactory/config/localvagrantrepository"
	"github.com/upbound/provider-artifactory/config/mailserver"
	"github.com/upbound/provider-artifactory/config/manageduser"
	"github.com/upbound/provider-artifactory/config/oauthsettings"
	"github.com/upbound/provider-artifactory/config/propertyset"
	"github.com/upbound/provider-artifactory/config/proxy"
	"github.com/upbound/provider-artifactory/config/releasebundlecustomwebhook"
	"github.com/upbound/provider-artifactory/config/releasebundlewebhook"
	"github.com/upbound/provider-artifactory/config/remotealpinerepository"
	"github.com/upbound/provider-artifactory/config/remotebowerrepository"
	"github.com/upbound/provider-artifactory/config/remotecargorepository"
	"github.com/upbound/provider-artifactory/config/remotechefrepository"
	"github.com/upbound/provider-artifactory/config/remotecocoapodsrepository"
	"github.com/upbound/provider-artifactory/config/remotecomposerrepository"
	"github.com/upbound/provider-artifactory/config/remoteconanrepository"
	"github.com/upbound/provider-artifactory/config/remotecondarepository"
	"github.com/upbound/provider-artifactory/config/remotecranrepository"
	"github.com/upbound/provider-artifactory/config/remotedebianrepository"
	"github.com/upbound/provider-artifactory/config/remotedockerrepository"
	"github.com/upbound/provider-artifactory/config/remotegemsrepository"
	"github.com/upbound/provider-artifactory/config/remotegenericrepository"
	"github.com/upbound/provider-artifactory/config/remotegitlfsrepository"
	"github.com/upbound/provider-artifactory/config/remotegorepository"
	"github.com/upbound/provider-artifactory/config/remotegradlerepository"
	"github.com/upbound/provider-artifactory/config/remotehelmrepository"
	"github.com/upbound/provider-artifactory/config/remotehuggingfacemlrepository"
	"github.com/upbound/provider-artifactory/config/remoteivyrepository"
	"github.com/upbound/provider-artifactory/config/remotemavenrepository"
	"github.com/upbound/provider-artifactory/config/remotenpmrepository"
	"github.com/upbound/provider-artifactory/config/remotenugetrepository"
	"github.com/upbound/provider-artifactory/config/remoteopkgrepository"
	"github.com/upbound/provider-artifactory/config/remotep2repository"
	"github.com/upbound/provider-artifactory/config/remotepubrepository"
	"github.com/upbound/provider-artifactory/config/remotepuppetrepository"
	"github.com/upbound/provider-artifactory/config/remotepypirepository"
	"github.com/upbound/provider-artifactory/config/remoterpmrepository"
	"github.com/upbound/provider-artifactory/config/remotesbtrepository"
	"github.com/upbound/provider-artifactory/config/remoteswiftrepository"
	"github.com/upbound/provider-artifactory/config/remoteterraformrepository"
	"github.com/upbound/provider-artifactory/config/remotevcsrepository"
	"github.com/upbound/provider-artifactory/config/repositorylayout"
	"github.com/upbound/provider-artifactory/config/scopedtoken"
	"github.com/upbound/provider-artifactory/config/unmanageduser"
	"github.com/upbound/provider-artifactory/config/user"
	"github.com/upbound/provider-artifactory/config/virtualalpinerepository"
	"github.com/upbound/provider-artifactory/config/virtualbowerrepository"
	"github.com/upbound/provider-artifactory/config/virtualchefrepository"
	"github.com/upbound/provider-artifactory/config/virtualcomposerrepository"
	"github.com/upbound/provider-artifactory/config/virtualconanrepository"
	"github.com/upbound/provider-artifactory/config/virtualcondarepository"
	"github.com/upbound/provider-artifactory/config/virtualcranrepository"
	"github.com/upbound/provider-artifactory/config/virtualdebianrepository"
	"github.com/upbound/provider-artifactory/config/virtualdockerrepository"
	"github.com/upbound/provider-artifactory/config/virtualgemsrepository"
	"github.com/upbound/provider-artifactory/config/virtualgenericrepository"
	"github.com/upbound/provider-artifactory/config/virtualgitlfsrepository"
	"github.com/upbound/provider-artifactory/config/virtualgorepository"
	"github.com/upbound/provider-artifactory/config/virtualgradlerepository"
	"github.com/upbound/provider-artifactory/config/virtualhelmrepository"
	"github.com/upbound/provider-artifactory/config/virtualivyrepository"
	"github.com/upbound/provider-artifactory/config/virtualmavenrepository"
	"github.com/upbound/provider-artifactory/config/virtualnpmrepository"
	"github.com/upbound/provider-artifactory/config/virtualnugetrepository"
	"github.com/upbound/provider-artifactory/config/virtualp2repository"
	"github.com/upbound/provider-artifactory/config/virtualpubrepository"
	"github.com/upbound/provider-artifactory/config/virtualpuppetrepository"
	"github.com/upbound/provider-artifactory/config/virtualpypirepository"
	"github.com/upbound/provider-artifactory/config/virtualrpmrepository"
	"github.com/upbound/provider-artifactory/config/virtualsbtrepository"
	"github.com/upbound/provider-artifactory/config/virtualswiftrepository"
	"github.com/upbound/provider-artifactory/config/virtualterraformrepository"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "artifactory"
	modulePath     = "github.com/upbound/provider-artifactory"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		backup.Configure,
		generalsecurity.Configure,
		ldapgroupsetting.Configure,
		ldapgroupsettingv2.Configure,
		ldapsetting.Configure,
		ldapsettingv2.Configure,
		mailserver.Configure,
		oauthsettings.Configure,
		propertyset.Configure,
		proxy.Configure,
		repositorylayout.Configure,
		federatedalpinerepository.Configure,
		federatedbowerrepository.Configure,
		federatedcargorepository.Configure,
		federatedchefrepository.Configure,
		federatedcocoapodsrepository.Configure,
		federatedcomposerrepository.Configure,
		federatedconanrepository.Configure,
		federatedcondarepository.Configure,
		federatedcranrepository.Configure,
		federateddebianrepository.Configure,
		federateddockerrepository.Configure,
		federateddockerv1repository.Configure,
		federateddockerv2repository.Configure,
		federatedgemsrepository.Configure,
		federatedgenericrepository.Configure,
		federatedgitlfsrepository.Configure,
		federatedgorepository.Configure,
		federatedgradlerepository.Configure,
		federatedhelmrepository.Configure,
		federatedivyrepository.Configure,
		federatedmavenrepository.Configure,
		federatednpmrepository.Configure,
		federatednugetrepository.Configure,
		federatedopkgrepository.Configure,
		federatedpuppetrepository.Configure,
		federatedpypirepository.Configure,
		federatedrpmrepository.Configure,
		federatedsbtrepository.Configure,
		federatedswiftrepository.Configure,
		federatedterraformmodulerepository.Configure,
		federatedterraformproviderrepository.Configure,
		federatedvagrantrepository.Configure,
		localalpinerepository.Configure,
		localbowerrepository.Configure,
		localcargorepository.Configure,
		localchefrepository.Configure,
		localcocoapodsrepository.Configure,
		localcomposerrepository.Configure,
		localconanrepository.Configure,
		localcondarepository.Configure,
		localcranrepository.Configure,
		localdebianrepository.Configure,
		localdockerv1repository.Configure,
		localdockerv2repository.Configure,
		localgemsrepository.Configure,
		localgenericrepository.Configure,
		localgitlfsrepository.Configure,
		localgorepository.Configure,
		localgradlerepository.Configure,
		localhelmrepository.Configure,
		localhuggingfacemlrepository.Configure,
		localivyrepository.Configure,
		localmavenrepository.Configure,
		localnpmrepository.Configure,
		localnugetrepository.Configure,
		localopkgrepository.Configure,
		localpubrepository.Configure,
		localpuppetrepository.Configure,
		localpypirepository.Configure,
		localrpmrepository.Configure,
		localsbtrepository.Configure,
		localswiftrepository.Configure,
		localterraformmodulerepository.Configure,
		localterraformproviderrepository.Configure,
		localterraformbackendrepository.Configure,
		localvagrantrepository.Configure,
		remotealpinerepository.Configure,
		remotebowerrepository.Configure,
		remotecargorepository.Configure,
		remotechefrepository.Configure,
		remotecocoapodsrepository.Configure,
		remotecomposerrepository.Configure,
		remoteconanrepository.Configure,
		remotecondarepository.Configure,
		remotecranrepository.Configure,
		remotedebianrepository.Configure,
		remotedockerrepository.Configure,
		remotegemsrepository.Configure,
		remotegenericrepository.Configure,
		remotegitlfsrepository.Configure,
		remotegorepository.Configure,
		remotegradlerepository.Configure,
		remotehelmrepository.Configure,
		remotehuggingfacemlrepository.Configure,
		remoteivyrepository.Configure,
		remotemavenrepository.Configure,
		remotenpmrepository.Configure,
		remotenugetrepository.Configure,
		remoteopkgrepository.Configure,
		remotep2repository.Configure,
		remotepubrepository.Configure,
		remotepuppetrepository.Configure,
		remotepypirepository.Configure,
		remoterpmrepository.Configure,
		remotesbtrepository.Configure,
		remoteswiftrepository.Configure,
		remoteterraformrepository.Configure,
		remotevcsrepository.Configure,
		localrepositorymultireplication.Configure,
		localrepositorysinglereplication.Configure,
		certificate.Configure,
		distributionpublickey.Configure,
		globalenvironment.Configure,
		group.Configure,
		keypair.Configure,
		scopedtoken.Configure,
		manageduser.Configure,
		unmanageduser.Configure,
		user.Configure,
		virtualalpinerepository.Configure,
		virtualbowerrepository.Configure,
		virtualchefrepository.Configure,
		virtualcomposerrepository.Configure,
		virtualconanrepository.Configure,
		virtualcondarepository.Configure,
		virtualcranrepository.Configure,
		virtualdebianrepository.Configure,
		virtualdockerrepository.Configure,
		virtualgemsrepository.Configure,
		virtualgenericrepository.Configure,
		virtualgitlfsrepository.Configure,
		virtualgorepository.Configure,
		virtualgradlerepository.Configure,
		virtualhelmrepository.Configure,
		virtualivyrepository.Configure,
		virtualmavenrepository.Configure,
		virtualnpmrepository.Configure,
		virtualnugetrepository.Configure,
		virtualp2repository.Configure,
		virtualpubrepository.Configure,
		virtualpuppetrepository.Configure,
		virtualpypirepository.Configure,
		virtualrpmrepository.Configure,
		virtualsbtrepository.Configure,
		virtualswiftrepository.Configure,
		virtualterraformrepository.Configure,
		artifactcustomwebhook.Configure,
		artifactpropertycustomwebhook.Configure,
		artifactpropertywebhook.Configure,
		artifactwebhook.Configure,
		releasebundlecustomwebhook.Configure,
		releasebundlewebhook.Configure,
		buildcustomwebhook.Configure,
		buildwebhook.Configure,
		distributioncustomwebhook.Configure,
		distributionwebhook.Configure,
		dockercustomwebhook.Configure,
		dockerwebhook.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
