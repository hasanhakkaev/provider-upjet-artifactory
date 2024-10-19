// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	user "github.com/upbound/provider-artifactory/internal/controller/anonymous/user"
	key "github.com/upbound/provider-artifactory/internal/controller/api/key"
	customwebhook "github.com/upbound/provider-artifactory/internal/controller/artifactcustomwebhook/customwebhook"
	releasebundlecustomwebhook "github.com/upbound/provider-artifactory/internal/controller/artifactory/releasebundlecustomwebhook"
	releasebundlewebhook "github.com/upbound/provider-artifactory/internal/controller/artifactory/releasebundlewebhook"
	propertycustomwebhook "github.com/upbound/provider-artifactory/internal/controller/artifactpropertycustomwebhook/propertycustomwebhook"
	propertywebhook "github.com/upbound/provider-artifactory/internal/controller/artifactpropertywebhook/propertywebhook"
	webhook "github.com/upbound/provider-artifactory/internal/controller/artifactwebhook/webhook"
	backup "github.com/upbound/provider-artifactory/internal/controller/backup/backup"
	customwebhookbuildcustomwebhook "github.com/upbound/provider-artifactory/internal/controller/buildcustomwebhook/customwebhook"
	webhookbuildwebhook "github.com/upbound/provider-artifactory/internal/controller/buildwebhook/webhook"
	certificate "github.com/upbound/provider-artifactory/internal/controller/certificate/certificate"
	customwebhookdistributioncustomwebhook "github.com/upbound/provider-artifactory/internal/controller/distributioncustomwebhook/customwebhook"
	publickey "github.com/upbound/provider-artifactory/internal/controller/distributionpublickey/publickey"
	webhookdistributionwebhook "github.com/upbound/provider-artifactory/internal/controller/distributionwebhook/webhook"
	customwebhookdockercustomwebhook "github.com/upbound/provider-artifactory/internal/controller/dockercustomwebhook/customwebhook"
	webhookdockerwebhook "github.com/upbound/provider-artifactory/internal/controller/dockerwebhook/webhook"
	alpinerepository "github.com/upbound/provider-artifactory/internal/controller/federatedalpinerepository/alpinerepository"
	bowerrepository "github.com/upbound/provider-artifactory/internal/controller/federatedbowerrepository/bowerrepository"
	cargorepository "github.com/upbound/provider-artifactory/internal/controller/federatedcargorepository/cargorepository"
	chefrepository "github.com/upbound/provider-artifactory/internal/controller/federatedchefrepository/chefrepository"
	cocoapodsrepository "github.com/upbound/provider-artifactory/internal/controller/federatedcocoapodsrepository/cocoapodsrepository"
	composerrepository "github.com/upbound/provider-artifactory/internal/controller/federatedcomposerrepository/composerrepository"
	conanrepository "github.com/upbound/provider-artifactory/internal/controller/federatedconanrepository/conanrepository"
	condarepository "github.com/upbound/provider-artifactory/internal/controller/federatedcondarepository/condarepository"
	cranrepository "github.com/upbound/provider-artifactory/internal/controller/federatedcranrepository/cranrepository"
	debianrepository "github.com/upbound/provider-artifactory/internal/controller/federateddebianrepository/debianrepository"
	dockerv1repository "github.com/upbound/provider-artifactory/internal/controller/federateddockerv1repository/dockerv1repository"
	dockerv2repository "github.com/upbound/provider-artifactory/internal/controller/federateddockerv2repository/dockerv2repository"
	gemsrepository "github.com/upbound/provider-artifactory/internal/controller/federatedgemsrepository/gemsrepository"
	genericrepository "github.com/upbound/provider-artifactory/internal/controller/federatedgenericrepository/genericrepository"
	gitlfsrepository "github.com/upbound/provider-artifactory/internal/controller/federatedgitlfsrepository/gitlfsrepository"
	gorepository "github.com/upbound/provider-artifactory/internal/controller/federatedgorepository/gorepository"
	gradlerepository "github.com/upbound/provider-artifactory/internal/controller/federatedgradlerepository/gradlerepository"
	helmrepository "github.com/upbound/provider-artifactory/internal/controller/federatedhelmrepository/helmrepository"
	ivyrepository "github.com/upbound/provider-artifactory/internal/controller/federatedivyrepository/ivyrepository"
	mavenrepository "github.com/upbound/provider-artifactory/internal/controller/federatedmavenrepository/mavenrepository"
	npmrepository "github.com/upbound/provider-artifactory/internal/controller/federatednpmrepository/npmrepository"
	nugetrepository "github.com/upbound/provider-artifactory/internal/controller/federatednugetrepository/nugetrepository"
	opkgrepository "github.com/upbound/provider-artifactory/internal/controller/federatedopkgrepository/opkgrepository"
	puppetrepository "github.com/upbound/provider-artifactory/internal/controller/federatedpuppetrepository/puppetrepository"
	pypirepository "github.com/upbound/provider-artifactory/internal/controller/federatedpypirepository/pypirepository"
	rpmrepository "github.com/upbound/provider-artifactory/internal/controller/federatedrpmrepository/rpmrepository"
	sbtrepository "github.com/upbound/provider-artifactory/internal/controller/federatedsbtrepository/sbtrepository"
	swiftrepository "github.com/upbound/provider-artifactory/internal/controller/federatedswiftrepository/swiftrepository"
	terraformmodulerepository "github.com/upbound/provider-artifactory/internal/controller/federatedterraformmodulerepository/terraformmodulerepository"
	terraformproviderrepository "github.com/upbound/provider-artifactory/internal/controller/federatedterraformproviderrepository/terraformproviderrepository"
	vagrantrepository "github.com/upbound/provider-artifactory/internal/controller/federatedvagrantrepository/vagrantrepository"
	security "github.com/upbound/provider-artifactory/internal/controller/generalsecurity/security"
	environment "github.com/upbound/provider-artifactory/internal/controller/globalenvironment/environment"
	group "github.com/upbound/provider-artifactory/internal/controller/group/group"
	keypair "github.com/upbound/provider-artifactory/internal/controller/keypair/keypair"
	groupsetting "github.com/upbound/provider-artifactory/internal/controller/ldapgroupsetting/groupsetting"
	groupsettingv2 "github.com/upbound/provider-artifactory/internal/controller/ldapgroupsettingv2/groupsettingv2"
	setting "github.com/upbound/provider-artifactory/internal/controller/ldapsetting/setting"
	settingv2 "github.com/upbound/provider-artifactory/internal/controller/ldapsettingv2/settingv2"
	alpinerepositorylocalalpinerepository "github.com/upbound/provider-artifactory/internal/controller/localalpinerepository/alpinerepository"
	bowerrepositorylocalbowerrepository "github.com/upbound/provider-artifactory/internal/controller/localbowerrepository/bowerrepository"
	cargorepositorylocalcargorepository "github.com/upbound/provider-artifactory/internal/controller/localcargorepository/cargorepository"
	chefrepositorylocalchefrepository "github.com/upbound/provider-artifactory/internal/controller/localchefrepository/chefrepository"
	cocoapodsrepositorylocalcocoapodsrepository "github.com/upbound/provider-artifactory/internal/controller/localcocoapodsrepository/cocoapodsrepository"
	composerrepositorylocalcomposerrepository "github.com/upbound/provider-artifactory/internal/controller/localcomposerrepository/composerrepository"
	conanrepositorylocalconanrepository "github.com/upbound/provider-artifactory/internal/controller/localconanrepository/conanrepository"
	condarepositorylocalcondarepository "github.com/upbound/provider-artifactory/internal/controller/localcondarepository/condarepository"
	cranrepositorylocalcranrepository "github.com/upbound/provider-artifactory/internal/controller/localcranrepository/cranrepository"
	debianrepositorylocaldebianrepository "github.com/upbound/provider-artifactory/internal/controller/localdebianrepository/debianrepository"
	dockerv1repositorylocaldockerv1repository "github.com/upbound/provider-artifactory/internal/controller/localdockerv1repository/dockerv1repository"
	dockerv2repositorylocaldockerv2repository "github.com/upbound/provider-artifactory/internal/controller/localdockerv2repository/dockerv2repository"
	gemsrepositorylocalgemsrepository "github.com/upbound/provider-artifactory/internal/controller/localgemsrepository/gemsrepository"
	genericrepositorylocalgenericrepository "github.com/upbound/provider-artifactory/internal/controller/localgenericrepository/genericrepository"
	gitlfsrepositorylocalgitlfsrepository "github.com/upbound/provider-artifactory/internal/controller/localgitlfsrepository/gitlfsrepository"
	gorepositorylocalgorepository "github.com/upbound/provider-artifactory/internal/controller/localgorepository/gorepository"
	gradlerepositorylocalgradlerepository "github.com/upbound/provider-artifactory/internal/controller/localgradlerepository/gradlerepository"
	helmrepositorylocalhelmrepository "github.com/upbound/provider-artifactory/internal/controller/localhelmrepository/helmrepository"
	huggingfacemlrepository "github.com/upbound/provider-artifactory/internal/controller/localhuggingfacemlrepository/huggingfacemlrepository"
	ivyrepositorylocalivyrepository "github.com/upbound/provider-artifactory/internal/controller/localivyrepository/ivyrepository"
	mavenrepositorylocalmavenrepository "github.com/upbound/provider-artifactory/internal/controller/localmavenrepository/mavenrepository"
	npmrepositorylocalnpmrepository "github.com/upbound/provider-artifactory/internal/controller/localnpmrepository/npmrepository"
	nugetrepositorylocalnugetrepository "github.com/upbound/provider-artifactory/internal/controller/localnugetrepository/nugetrepository"
	opkgrepositorylocalopkgrepository "github.com/upbound/provider-artifactory/internal/controller/localopkgrepository/opkgrepository"
	pubrepository "github.com/upbound/provider-artifactory/internal/controller/localpubrepository/pubrepository"
	puppetrepositorylocalpuppetrepository "github.com/upbound/provider-artifactory/internal/controller/localpuppetrepository/puppetrepository"
	pypirepositorylocalpypirepository "github.com/upbound/provider-artifactory/internal/controller/localpypirepository/pypirepository"
	repositorymultireplication "github.com/upbound/provider-artifactory/internal/controller/localrepositorymultireplication/repositorymultireplication"
	repositorysinglereplication "github.com/upbound/provider-artifactory/internal/controller/localrepositorysinglereplication/repositorysinglereplication"
	rpmrepositorylocalrpmrepository "github.com/upbound/provider-artifactory/internal/controller/localrpmrepository/rpmrepository"
	sbtrepositorylocalsbtrepository "github.com/upbound/provider-artifactory/internal/controller/localsbtrepository/sbtrepository"
	swiftrepositorylocalswiftrepository "github.com/upbound/provider-artifactory/internal/controller/localswiftrepository/swiftrepository"
	terraformbackendrepository "github.com/upbound/provider-artifactory/internal/controller/localterraformbackendrepository/terraformbackendrepository"
	terraformmodulerepositorylocalterraformmodulerepository "github.com/upbound/provider-artifactory/internal/controller/localterraformmodulerepository/terraformmodulerepository"
	terraformproviderrepositorylocalterraformproviderrepository "github.com/upbound/provider-artifactory/internal/controller/localterraformproviderrepository/terraformproviderrepository"
	vagrantrepositorylocalvagrantrepository "github.com/upbound/provider-artifactory/internal/controller/localvagrantrepository/vagrantrepository"
	server "github.com/upbound/provider-artifactory/internal/controller/mailserver/server"
	usermanageduser "github.com/upbound/provider-artifactory/internal/controller/manageduser/user"
	settings "github.com/upbound/provider-artifactory/internal/controller/oauthsettings/settings"
	target "github.com/upbound/provider-artifactory/internal/controller/permission/target"
	set "github.com/upbound/provider-artifactory/internal/controller/propertyset/set"
	providerconfig "github.com/upbound/provider-artifactory/internal/controller/providerconfig"
	proxy "github.com/upbound/provider-artifactory/internal/controller/proxy/proxy"
	replication "github.com/upbound/provider-artifactory/internal/controller/pull/replication"
	replicationpush "github.com/upbound/provider-artifactory/internal/controller/push/replication"
	bundlecustomwebhook "github.com/upbound/provider-artifactory/internal/controller/releasebundlecustomwebhook/bundlecustomwebhook"
	bundlewebhook "github.com/upbound/provider-artifactory/internal/controller/releasebundlewebhook/bundlewebhook"
	repositoryreplication "github.com/upbound/provider-artifactory/internal/controller/remote/repositoryreplication"
	alpinerepositoryremotealpinerepository "github.com/upbound/provider-artifactory/internal/controller/remotealpinerepository/alpinerepository"
	bowerrepositoryremotebowerrepository "github.com/upbound/provider-artifactory/internal/controller/remotebowerrepository/bowerrepository"
	cargorepositoryremotecargorepository "github.com/upbound/provider-artifactory/internal/controller/remotecargorepository/cargorepository"
	chefrepositoryremotechefrepository "github.com/upbound/provider-artifactory/internal/controller/remotechefrepository/chefrepository"
	cocoapodsrepositoryremotecocoapodsrepository "github.com/upbound/provider-artifactory/internal/controller/remotecocoapodsrepository/cocoapodsrepository"
	composerrepositoryremotecomposerrepository "github.com/upbound/provider-artifactory/internal/controller/remotecomposerrepository/composerrepository"
	conanrepositoryremoteconanrepository "github.com/upbound/provider-artifactory/internal/controller/remoteconanrepository/conanrepository"
	condarepositoryremotecondarepository "github.com/upbound/provider-artifactory/internal/controller/remotecondarepository/condarepository"
	cranrepositoryremotecranrepository "github.com/upbound/provider-artifactory/internal/controller/remotecranrepository/cranrepository"
	debianrepositoryremotedebianrepository "github.com/upbound/provider-artifactory/internal/controller/remotedebianrepository/debianrepository"
	dockerrepository "github.com/upbound/provider-artifactory/internal/controller/remotedockerrepository/dockerrepository"
	gemsrepositoryremotegemsrepository "github.com/upbound/provider-artifactory/internal/controller/remotegemsrepository/gemsrepository"
	genericrepositoryremotegenericrepository "github.com/upbound/provider-artifactory/internal/controller/remotegenericrepository/genericrepository"
	gitlfsrepositoryremotegitlfsrepository "github.com/upbound/provider-artifactory/internal/controller/remotegitlfsrepository/gitlfsrepository"
	gorepositoryremotegorepository "github.com/upbound/provider-artifactory/internal/controller/remotegorepository/gorepository"
	gradlerepositoryremotegradlerepository "github.com/upbound/provider-artifactory/internal/controller/remotegradlerepository/gradlerepository"
	helmrepositoryremotehelmrepository "github.com/upbound/provider-artifactory/internal/controller/remotehelmrepository/helmrepository"
	huggingfacemlrepositoryremotehuggingfacemlrepository "github.com/upbound/provider-artifactory/internal/controller/remotehuggingfacemlrepository/huggingfacemlrepository"
	ivyrepositoryremoteivyrepository "github.com/upbound/provider-artifactory/internal/controller/remoteivyrepository/ivyrepository"
	mavenrepositoryremotemavenrepository "github.com/upbound/provider-artifactory/internal/controller/remotemavenrepository/mavenrepository"
	npmrepositoryremotenpmrepository "github.com/upbound/provider-artifactory/internal/controller/remotenpmrepository/npmrepository"
	nugetrepositoryremotenugetrepository "github.com/upbound/provider-artifactory/internal/controller/remotenugetrepository/nugetrepository"
	opkgrepositoryremoteopkgrepository "github.com/upbound/provider-artifactory/internal/controller/remoteopkgrepository/opkgrepository"
	p2repository "github.com/upbound/provider-artifactory/internal/controller/remotep2repository/p2repository"
	pubrepositoryremotepubrepository "github.com/upbound/provider-artifactory/internal/controller/remotepubrepository/pubrepository"
	puppetrepositoryremotepuppetrepository "github.com/upbound/provider-artifactory/internal/controller/remotepuppetrepository/puppetrepository"
	pypirepositoryremotepypirepository "github.com/upbound/provider-artifactory/internal/controller/remotepypirepository/pypirepository"
	rpmrepositoryremoterpmrepository "github.com/upbound/provider-artifactory/internal/controller/remoterpmrepository/rpmrepository"
	sbtrepositoryremotesbtrepository "github.com/upbound/provider-artifactory/internal/controller/remotesbtrepository/sbtrepository"
	swiftrepositoryremoteswiftrepository "github.com/upbound/provider-artifactory/internal/controller/remoteswiftrepository/swiftrepository"
	terraformrepository "github.com/upbound/provider-artifactory/internal/controller/remoteterraformrepository/terraformrepository"
	vcsrepository "github.com/upbound/provider-artifactory/internal/controller/remotevcsrepository/vcsrepository"
	config "github.com/upbound/provider-artifactory/internal/controller/replication/config"
	layout "github.com/upbound/provider-artifactory/internal/controller/repositorylayout/layout"
	settingssaml "github.com/upbound/provider-artifactory/internal/controller/saml/settings"
	token "github.com/upbound/provider-artifactory/internal/controller/scopedtoken/token"
	replicationconfig "github.com/upbound/provider-artifactory/internal/controller/single/replicationconfig"
	userunmanageduser "github.com/upbound/provider-artifactory/internal/controller/unmanageduser/user"
	useruser "github.com/upbound/provider-artifactory/internal/controller/user/user"
	alpinerepositoryvirtualalpinerepository "github.com/upbound/provider-artifactory/internal/controller/virtualalpinerepository/alpinerepository"
	bowerrepositoryvirtualbowerrepository "github.com/upbound/provider-artifactory/internal/controller/virtualbowerrepository/bowerrepository"
	chefrepositoryvirtualchefrepository "github.com/upbound/provider-artifactory/internal/controller/virtualchefrepository/chefrepository"
	composerrepositoryvirtualcomposerrepository "github.com/upbound/provider-artifactory/internal/controller/virtualcomposerrepository/composerrepository"
	conanrepositoryvirtualconanrepository "github.com/upbound/provider-artifactory/internal/controller/virtualconanrepository/conanrepository"
	condarepositoryvirtualcondarepository "github.com/upbound/provider-artifactory/internal/controller/virtualcondarepository/condarepository"
	cranrepositoryvirtualcranrepository "github.com/upbound/provider-artifactory/internal/controller/virtualcranrepository/cranrepository"
	debianrepositoryvirtualdebianrepository "github.com/upbound/provider-artifactory/internal/controller/virtualdebianrepository/debianrepository"
	dockerrepositoryvirtualdockerrepository "github.com/upbound/provider-artifactory/internal/controller/virtualdockerrepository/dockerrepository"
	gemsrepositoryvirtualgemsrepository "github.com/upbound/provider-artifactory/internal/controller/virtualgemsrepository/gemsrepository"
	genericrepositoryvirtualgenericrepository "github.com/upbound/provider-artifactory/internal/controller/virtualgenericrepository/genericrepository"
	gitlfsrepositoryvirtualgitlfsrepository "github.com/upbound/provider-artifactory/internal/controller/virtualgitlfsrepository/gitlfsrepository"
	gorepositoryvirtualgorepository "github.com/upbound/provider-artifactory/internal/controller/virtualgorepository/gorepository"
	gradlerepositoryvirtualgradlerepository "github.com/upbound/provider-artifactory/internal/controller/virtualgradlerepository/gradlerepository"
	helmrepositoryvirtualhelmrepository "github.com/upbound/provider-artifactory/internal/controller/virtualhelmrepository/helmrepository"
	ivyrepositoryvirtualivyrepository "github.com/upbound/provider-artifactory/internal/controller/virtualivyrepository/ivyrepository"
	mavenrepositoryvirtualmavenrepository "github.com/upbound/provider-artifactory/internal/controller/virtualmavenrepository/mavenrepository"
	npmrepositoryvirtualnpmrepository "github.com/upbound/provider-artifactory/internal/controller/virtualnpmrepository/npmrepository"
	nugetrepositoryvirtualnugetrepository "github.com/upbound/provider-artifactory/internal/controller/virtualnugetrepository/nugetrepository"
	p2repositoryvirtualp2repository "github.com/upbound/provider-artifactory/internal/controller/virtualp2repository/p2repository"
	pubrepositoryvirtualpubrepository "github.com/upbound/provider-artifactory/internal/controller/virtualpubrepository/pubrepository"
	puppetrepositoryvirtualpuppetrepository "github.com/upbound/provider-artifactory/internal/controller/virtualpuppetrepository/puppetrepository"
	pypirepositoryvirtualpypirepository "github.com/upbound/provider-artifactory/internal/controller/virtualpypirepository/pypirepository"
	rpmrepositoryvirtualrpmrepository "github.com/upbound/provider-artifactory/internal/controller/virtualrpmrepository/rpmrepository"
	sbtrepositoryvirtualsbtrepository "github.com/upbound/provider-artifactory/internal/controller/virtualsbtrepository/sbtrepository"
	swiftrepositoryvirtualswiftrepository "github.com/upbound/provider-artifactory/internal/controller/virtualswiftrepository/swiftrepository"
	terraformrepositoryvirtualterraformrepository "github.com/upbound/provider-artifactory/internal/controller/virtualterraformrepository/terraformrepository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		user.Setup,
		key.Setup,
		customwebhook.Setup,
		releasebundlecustomwebhook.Setup,
		releasebundlewebhook.Setup,
		propertycustomwebhook.Setup,
		propertywebhook.Setup,
		webhook.Setup,
		backup.Setup,
		customwebhookbuildcustomwebhook.Setup,
		webhookbuildwebhook.Setup,
		certificate.Setup,
		customwebhookdistributioncustomwebhook.Setup,
		publickey.Setup,
		webhookdistributionwebhook.Setup,
		customwebhookdockercustomwebhook.Setup,
		webhookdockerwebhook.Setup,
		alpinerepository.Setup,
		bowerrepository.Setup,
		cargorepository.Setup,
		chefrepository.Setup,
		cocoapodsrepository.Setup,
		composerrepository.Setup,
		conanrepository.Setup,
		condarepository.Setup,
		cranrepository.Setup,
		debianrepository.Setup,
		dockerv1repository.Setup,
		dockerv2repository.Setup,
		gemsrepository.Setup,
		genericrepository.Setup,
		gitlfsrepository.Setup,
		gorepository.Setup,
		gradlerepository.Setup,
		helmrepository.Setup,
		ivyrepository.Setup,
		mavenrepository.Setup,
		npmrepository.Setup,
		nugetrepository.Setup,
		opkgrepository.Setup,
		puppetrepository.Setup,
		pypirepository.Setup,
		rpmrepository.Setup,
		sbtrepository.Setup,
		swiftrepository.Setup,
		terraformmodulerepository.Setup,
		terraformproviderrepository.Setup,
		vagrantrepository.Setup,
		security.Setup,
		environment.Setup,
		group.Setup,
		keypair.Setup,
		groupsetting.Setup,
		groupsettingv2.Setup,
		setting.Setup,
		settingv2.Setup,
		alpinerepositorylocalalpinerepository.Setup,
		bowerrepositorylocalbowerrepository.Setup,
		cargorepositorylocalcargorepository.Setup,
		chefrepositorylocalchefrepository.Setup,
		cocoapodsrepositorylocalcocoapodsrepository.Setup,
		composerrepositorylocalcomposerrepository.Setup,
		conanrepositorylocalconanrepository.Setup,
		condarepositorylocalcondarepository.Setup,
		cranrepositorylocalcranrepository.Setup,
		debianrepositorylocaldebianrepository.Setup,
		dockerv1repositorylocaldockerv1repository.Setup,
		dockerv2repositorylocaldockerv2repository.Setup,
		gemsrepositorylocalgemsrepository.Setup,
		genericrepositorylocalgenericrepository.Setup,
		gitlfsrepositorylocalgitlfsrepository.Setup,
		gorepositorylocalgorepository.Setup,
		gradlerepositorylocalgradlerepository.Setup,
		helmrepositorylocalhelmrepository.Setup,
		huggingfacemlrepository.Setup,
		ivyrepositorylocalivyrepository.Setup,
		mavenrepositorylocalmavenrepository.Setup,
		npmrepositorylocalnpmrepository.Setup,
		nugetrepositorylocalnugetrepository.Setup,
		opkgrepositorylocalopkgrepository.Setup,
		pubrepository.Setup,
		puppetrepositorylocalpuppetrepository.Setup,
		pypirepositorylocalpypirepository.Setup,
		repositorymultireplication.Setup,
		repositorysinglereplication.Setup,
		rpmrepositorylocalrpmrepository.Setup,
		sbtrepositorylocalsbtrepository.Setup,
		swiftrepositorylocalswiftrepository.Setup,
		terraformbackendrepository.Setup,
		terraformmodulerepositorylocalterraformmodulerepository.Setup,
		terraformproviderrepositorylocalterraformproviderrepository.Setup,
		vagrantrepositorylocalvagrantrepository.Setup,
		server.Setup,
		usermanageduser.Setup,
		settings.Setup,
		target.Setup,
		set.Setup,
		providerconfig.Setup,
		proxy.Setup,
		replication.Setup,
		replicationpush.Setup,
		bundlecustomwebhook.Setup,
		bundlewebhook.Setup,
		repositoryreplication.Setup,
		alpinerepositoryremotealpinerepository.Setup,
		bowerrepositoryremotebowerrepository.Setup,
		cargorepositoryremotecargorepository.Setup,
		chefrepositoryremotechefrepository.Setup,
		cocoapodsrepositoryremotecocoapodsrepository.Setup,
		composerrepositoryremotecomposerrepository.Setup,
		conanrepositoryremoteconanrepository.Setup,
		condarepositoryremotecondarepository.Setup,
		cranrepositoryremotecranrepository.Setup,
		debianrepositoryremotedebianrepository.Setup,
		dockerrepository.Setup,
		gemsrepositoryremotegemsrepository.Setup,
		genericrepositoryremotegenericrepository.Setup,
		gitlfsrepositoryremotegitlfsrepository.Setup,
		gorepositoryremotegorepository.Setup,
		gradlerepositoryremotegradlerepository.Setup,
		helmrepositoryremotehelmrepository.Setup,
		huggingfacemlrepositoryremotehuggingfacemlrepository.Setup,
		ivyrepositoryremoteivyrepository.Setup,
		mavenrepositoryremotemavenrepository.Setup,
		npmrepositoryremotenpmrepository.Setup,
		nugetrepositoryremotenugetrepository.Setup,
		opkgrepositoryremoteopkgrepository.Setup,
		p2repository.Setup,
		pubrepositoryremotepubrepository.Setup,
		puppetrepositoryremotepuppetrepository.Setup,
		pypirepositoryremotepypirepository.Setup,
		rpmrepositoryremoterpmrepository.Setup,
		sbtrepositoryremotesbtrepository.Setup,
		swiftrepositoryremoteswiftrepository.Setup,
		terraformrepository.Setup,
		vcsrepository.Setup,
		config.Setup,
		layout.Setup,
		settingssaml.Setup,
		token.Setup,
		replicationconfig.Setup,
		userunmanageduser.Setup,
		useruser.Setup,
		alpinerepositoryvirtualalpinerepository.Setup,
		bowerrepositoryvirtualbowerrepository.Setup,
		chefrepositoryvirtualchefrepository.Setup,
		composerrepositoryvirtualcomposerrepository.Setup,
		conanrepositoryvirtualconanrepository.Setup,
		condarepositoryvirtualcondarepository.Setup,
		cranrepositoryvirtualcranrepository.Setup,
		debianrepositoryvirtualdebianrepository.Setup,
		dockerrepositoryvirtualdockerrepository.Setup,
		gemsrepositoryvirtualgemsrepository.Setup,
		genericrepositoryvirtualgenericrepository.Setup,
		gitlfsrepositoryvirtualgitlfsrepository.Setup,
		gorepositoryvirtualgorepository.Setup,
		gradlerepositoryvirtualgradlerepository.Setup,
		helmrepositoryvirtualhelmrepository.Setup,
		ivyrepositoryvirtualivyrepository.Setup,
		mavenrepositoryvirtualmavenrepository.Setup,
		npmrepositoryvirtualnpmrepository.Setup,
		nugetrepositoryvirtualnugetrepository.Setup,
		p2repositoryvirtualp2repository.Setup,
		pubrepositoryvirtualpubrepository.Setup,
		puppetrepositoryvirtualpuppetrepository.Setup,
		pypirepositoryvirtualpypirepository.Setup,
		rpmrepositoryvirtualrpmrepository.Setup,
		sbtrepositoryvirtualsbtrepository.Setup,
		swiftrepositoryvirtualswiftrepository.Setup,
		terraformrepositoryvirtualterraformrepository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
