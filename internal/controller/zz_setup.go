// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	user "github.com/upbound/provider-artifactory/internal/controller/anonymous/user"
	customwebhook "github.com/upbound/provider-artifactory/internal/controller/artifact_custom_webhook/customwebhook"
	propertycustomwebhook "github.com/upbound/provider-artifactory/internal/controller/artifact_property_custom_webhook/propertycustomwebhook"
	propertywebhook "github.com/upbound/provider-artifactory/internal/controller/artifact_property_webhook/propertywebhook"
	webhook "github.com/upbound/provider-artifactory/internal/controller/artifact_webhook/webhook"
	releasebundlecustomwebhook "github.com/upbound/provider-artifactory/internal/controller/artifactory/releasebundlecustomwebhook"
	releasebundlewebhook "github.com/upbound/provider-artifactory/internal/controller/artifactory/releasebundlewebhook"
	backup "github.com/upbound/provider-artifactory/internal/controller/backup/backup"
	customwebhookbuild_custom_webhook "github.com/upbound/provider-artifactory/internal/controller/build_custom_webhook/customwebhook"
	webhookbuild_webhook "github.com/upbound/provider-artifactory/internal/controller/build_webhook/webhook"
	certificate "github.com/upbound/provider-artifactory/internal/controller/certificate/certificate"
	customwebhookdistribution_custom_webhook "github.com/upbound/provider-artifactory/internal/controller/distribution_custom_webhook/customwebhook"
	publickey "github.com/upbound/provider-artifactory/internal/controller/distribution_public_key/publickey"
	webhookdistribution_webhook "github.com/upbound/provider-artifactory/internal/controller/distribution_webhook/webhook"
	customwebhookdocker_custom_webhook "github.com/upbound/provider-artifactory/internal/controller/docker_custom_webhook/customwebhook"
	webhookdocker_webhook "github.com/upbound/provider-artifactory/internal/controller/docker_webhook/webhook"
	alpinerepository "github.com/upbound/provider-artifactory/internal/controller/federated_alpine_repository/alpinerepository"
	bowerrepository "github.com/upbound/provider-artifactory/internal/controller/federated_bower_repository/bowerrepository"
	cargorepository "github.com/upbound/provider-artifactory/internal/controller/federated_cargo_repository/cargorepository"
	chefrepository "github.com/upbound/provider-artifactory/internal/controller/federated_chef_repository/chefrepository"
	cocoapodsrepository "github.com/upbound/provider-artifactory/internal/controller/federated_cocoapods_repository/cocoapodsrepository"
	composerrepository "github.com/upbound/provider-artifactory/internal/controller/federated_composer_repository/composerrepository"
	conanrepository "github.com/upbound/provider-artifactory/internal/controller/federated_conan_repository/conanrepository"
	condarepository "github.com/upbound/provider-artifactory/internal/controller/federated_conda_repository/condarepository"
	cranrepository "github.com/upbound/provider-artifactory/internal/controller/federated_cran_repository/cranrepository"
	debianrepository "github.com/upbound/provider-artifactory/internal/controller/federated_debian_repository/debianrepository"
	dockerrepository "github.com/upbound/provider-artifactory/internal/controller/federated_docker_repository/dockerrepository"
	dockerv1repository "github.com/upbound/provider-artifactory/internal/controller/federated_docker_v1_repository/dockerv1repository"
	dockerv2repository "github.com/upbound/provider-artifactory/internal/controller/federated_docker_v2_repository/dockerv2repository"
	gemsrepository "github.com/upbound/provider-artifactory/internal/controller/federated_gems_repository/gemsrepository"
	genericrepository "github.com/upbound/provider-artifactory/internal/controller/federated_generic_repository/genericrepository"
	gitlfsrepository "github.com/upbound/provider-artifactory/internal/controller/federated_gitlfs_repository/gitlfsrepository"
	gorepository "github.com/upbound/provider-artifactory/internal/controller/federated_go_repository/gorepository"
	gradlerepository "github.com/upbound/provider-artifactory/internal/controller/federated_gradle_repository/gradlerepository"
	helmrepository "github.com/upbound/provider-artifactory/internal/controller/federated_helm_repository/helmrepository"
	ivyrepository "github.com/upbound/provider-artifactory/internal/controller/federated_ivy_repository/ivyrepository"
	mavenrepository "github.com/upbound/provider-artifactory/internal/controller/federated_maven_repository/mavenrepository"
	npmrepository "github.com/upbound/provider-artifactory/internal/controller/federated_npm_repository/npmrepository"
	nugetrepository "github.com/upbound/provider-artifactory/internal/controller/federated_nuget_repository/nugetrepository"
	opkgrepository "github.com/upbound/provider-artifactory/internal/controller/federated_opkg_repository/opkgrepository"
	puppetrepository "github.com/upbound/provider-artifactory/internal/controller/federated_puppet_repository/puppetrepository"
	pypirepository "github.com/upbound/provider-artifactory/internal/controller/federated_pypi_repository/pypirepository"
	rpmrepository "github.com/upbound/provider-artifactory/internal/controller/federated_rpm_repository/rpmrepository"
	sbtrepository "github.com/upbound/provider-artifactory/internal/controller/federated_sbt_repository/sbtrepository"
	swiftrepository "github.com/upbound/provider-artifactory/internal/controller/federated_swift_repository/swiftrepository"
	terraformmodulerepository "github.com/upbound/provider-artifactory/internal/controller/federated_terraform_module_repository/terraformmodulerepository"
	terraformproviderrepository "github.com/upbound/provider-artifactory/internal/controller/federated_terraform_provider_repository/terraformproviderrepository"
	vagrantrepository "github.com/upbound/provider-artifactory/internal/controller/federated_vagrant_repository/vagrantrepository"
	security "github.com/upbound/provider-artifactory/internal/controller/general_security/security"
	environment "github.com/upbound/provider-artifactory/internal/controller/global_environment/environment"
	group "github.com/upbound/provider-artifactory/internal/controller/group/group"
	keypair "github.com/upbound/provider-artifactory/internal/controller/keypair/keypair"
	groupsetting "github.com/upbound/provider-artifactory/internal/controller/ldap_group_setting/groupsetting"
	groupsettingv2 "github.com/upbound/provider-artifactory/internal/controller/ldap_group_setting_v2/groupsettingv2"
	setting "github.com/upbound/provider-artifactory/internal/controller/ldap_setting/setting"
	settingv2 "github.com/upbound/provider-artifactory/internal/controller/ldap_setting_v2/settingv2"
	alpinerepositorylocal_alpine_repository "github.com/upbound/provider-artifactory/internal/controller/local_alpine_repository/alpinerepository"
	bowerrepositorylocal_bower_repository "github.com/upbound/provider-artifactory/internal/controller/local_bower_repository/bowerrepository"
	cargorepositorylocal_cargo_repository "github.com/upbound/provider-artifactory/internal/controller/local_cargo_repository/cargorepository"
	chefrepositorylocal_chef_repository "github.com/upbound/provider-artifactory/internal/controller/local_chef_repository/chefrepository"
	cocoapodsrepositorylocal_cocoapods_repository "github.com/upbound/provider-artifactory/internal/controller/local_cocoapods_repository/cocoapodsrepository"
	composerrepositorylocal_composer_repository "github.com/upbound/provider-artifactory/internal/controller/local_composer_repository/composerrepository"
	conanrepositorylocal_conan_repository "github.com/upbound/provider-artifactory/internal/controller/local_conan_repository/conanrepository"
	condarepositorylocal_conda_repository "github.com/upbound/provider-artifactory/internal/controller/local_conda_repository/condarepository"
	cranrepositorylocal_cran_repository "github.com/upbound/provider-artifactory/internal/controller/local_cran_repository/cranrepository"
	debianrepositorylocal_debian_repository "github.com/upbound/provider-artifactory/internal/controller/local_debian_repository/debianrepository"
	dockerv1repositorylocal_docker_v1_repository "github.com/upbound/provider-artifactory/internal/controller/local_docker_v1_repository/dockerv1repository"
	dockerv2repositorylocal_docker_v2_repository "github.com/upbound/provider-artifactory/internal/controller/local_docker_v2_repository/dockerv2repository"
	gemsrepositorylocal_gems_repository "github.com/upbound/provider-artifactory/internal/controller/local_gems_repository/gemsrepository"
	genericrepositorylocal_generic_repository "github.com/upbound/provider-artifactory/internal/controller/local_generic_repository/genericrepository"
	gitlfsrepositorylocal_gitlfs_repository "github.com/upbound/provider-artifactory/internal/controller/local_gitlfs_repository/gitlfsrepository"
	gorepositorylocal_go_repository "github.com/upbound/provider-artifactory/internal/controller/local_go_repository/gorepository"
	gradlerepositorylocal_gradle_repository "github.com/upbound/provider-artifactory/internal/controller/local_gradle_repository/gradlerepository"
	helmrepositorylocal_helm_repository "github.com/upbound/provider-artifactory/internal/controller/local_helm_repository/helmrepository"
	huggingfacemlrepository "github.com/upbound/provider-artifactory/internal/controller/local_huggingfaceml_repository/huggingfacemlrepository"
	ivyrepositorylocal_ivy_repository "github.com/upbound/provider-artifactory/internal/controller/local_ivy_repository/ivyrepository"
	mavenrepositorylocal_maven_repository "github.com/upbound/provider-artifactory/internal/controller/local_maven_repository/mavenrepository"
	npmrepositorylocal_npm_repository "github.com/upbound/provider-artifactory/internal/controller/local_npm_repository/npmrepository"
	nugetrepositorylocal_nuget_repository "github.com/upbound/provider-artifactory/internal/controller/local_nuget_repository/nugetrepository"
	opkgrepositorylocal_opkg_repository "github.com/upbound/provider-artifactory/internal/controller/local_opkg_repository/opkgrepository"
	pubrepository "github.com/upbound/provider-artifactory/internal/controller/local_pub_repository/pubrepository"
	puppetrepositorylocal_puppet_repository "github.com/upbound/provider-artifactory/internal/controller/local_puppet_repository/puppetrepository"
	pypirepositorylocal_pypi_repository "github.com/upbound/provider-artifactory/internal/controller/local_pypi_repository/pypirepository"
	repositorymultireplication "github.com/upbound/provider-artifactory/internal/controller/local_repository_multi_replication/repositorymultireplication"
	repositorysinglereplication "github.com/upbound/provider-artifactory/internal/controller/local_repository_single_replication/repositorysinglereplication"
	rpmrepositorylocal_rpm_repository "github.com/upbound/provider-artifactory/internal/controller/local_rpm_repository/rpmrepository"
	sbtrepositorylocal_sbt_repository "github.com/upbound/provider-artifactory/internal/controller/local_sbt_repository/sbtrepository"
	swiftrepositorylocal_swift_repository "github.com/upbound/provider-artifactory/internal/controller/local_swift_repository/swiftrepository"
	terraformmodulerepositorylocal_terraform_module_repository "github.com/upbound/provider-artifactory/internal/controller/local_terraform_module_repository/terraformmodulerepository"
	terraformproviderrepositorylocal_terraform_provider_repository "github.com/upbound/provider-artifactory/internal/controller/local_terraform_provider_repository/terraformproviderrepository"
	terraformbackendrepository "github.com/upbound/provider-artifactory/internal/controller/local_terraformbackend_repository/terraformbackendrepository"
	vagrantrepositorylocal_vagrant_repository "github.com/upbound/provider-artifactory/internal/controller/local_vagrant_repository/vagrantrepository"
	server "github.com/upbound/provider-artifactory/internal/controller/mail_server/server"
	usermanaged_user "github.com/upbound/provider-artifactory/internal/controller/managed_user/user"
	settings "github.com/upbound/provider-artifactory/internal/controller/oauth_settings/settings"
	set "github.com/upbound/provider-artifactory/internal/controller/property_set/set"
	providerconfig "github.com/upbound/provider-artifactory/internal/controller/providerconfig"
	proxy "github.com/upbound/provider-artifactory/internal/controller/proxy/proxy"
	alpinerepositoryremote_alpine_repository "github.com/upbound/provider-artifactory/internal/controller/remote_alpine_repository/alpinerepository"
	bowerrepositoryremote_bower_repository "github.com/upbound/provider-artifactory/internal/controller/remote_bower_repository/bowerrepository"
	cargorepositoryremote_cargo_repository "github.com/upbound/provider-artifactory/internal/controller/remote_cargo_repository/cargorepository"
	chefrepositoryremote_chef_repository "github.com/upbound/provider-artifactory/internal/controller/remote_chef_repository/chefrepository"
	cocoapodsrepositoryremote_cocoapods_repository "github.com/upbound/provider-artifactory/internal/controller/remote_cocoapods_repository/cocoapodsrepository"
	composerrepositoryremote_composer_repository "github.com/upbound/provider-artifactory/internal/controller/remote_composer_repository/composerrepository"
	conanrepositoryremote_conan_repository "github.com/upbound/provider-artifactory/internal/controller/remote_conan_repository/conanrepository"
	condarepositoryremote_conda_repository "github.com/upbound/provider-artifactory/internal/controller/remote_conda_repository/condarepository"
	cranrepositoryremote_cran_repository "github.com/upbound/provider-artifactory/internal/controller/remote_cran_repository/cranrepository"
	debianrepositoryremote_debian_repository "github.com/upbound/provider-artifactory/internal/controller/remote_debian_repository/debianrepository"
	dockerrepositoryremote_docker_repository "github.com/upbound/provider-artifactory/internal/controller/remote_docker_repository/dockerrepository"
	gemsrepositoryremote_gems_repository "github.com/upbound/provider-artifactory/internal/controller/remote_gems_repository/gemsrepository"
	genericrepositoryremote_generic_repository "github.com/upbound/provider-artifactory/internal/controller/remote_generic_repository/genericrepository"
	gitlfsrepositoryremote_gitlfs_repository "github.com/upbound/provider-artifactory/internal/controller/remote_gitlfs_repository/gitlfsrepository"
	gorepositoryremote_go_repository "github.com/upbound/provider-artifactory/internal/controller/remote_go_repository/gorepository"
	gradlerepositoryremote_gradle_repository "github.com/upbound/provider-artifactory/internal/controller/remote_gradle_repository/gradlerepository"
	helmrepositoryremote_helm_repository "github.com/upbound/provider-artifactory/internal/controller/remote_helm_repository/helmrepository"
	huggingfacemlrepositoryremote_huggingfaceml_repository "github.com/upbound/provider-artifactory/internal/controller/remote_huggingfaceml_repository/huggingfacemlrepository"
	ivyrepositoryremote_ivy_repository "github.com/upbound/provider-artifactory/internal/controller/remote_ivy_repository/ivyrepository"
	mavenrepositoryremote_maven_repository "github.com/upbound/provider-artifactory/internal/controller/remote_maven_repository/mavenrepository"
	npmrepositoryremote_npm_repository "github.com/upbound/provider-artifactory/internal/controller/remote_npm_repository/npmrepository"
	nugetrepositoryremote_nuget_repository "github.com/upbound/provider-artifactory/internal/controller/remote_nuget_repository/nugetrepository"
	opkgrepositoryremote_opkg_repository "github.com/upbound/provider-artifactory/internal/controller/remote_opkg_repository/opkgrepository"
	p2repository "github.com/upbound/provider-artifactory/internal/controller/remote_p2_repository/p2repository"
	pubrepositoryremote_pub_repository "github.com/upbound/provider-artifactory/internal/controller/remote_pub_repository/pubrepository"
	puppetrepositoryremote_puppet_repository "github.com/upbound/provider-artifactory/internal/controller/remote_puppet_repository/puppetrepository"
	pypirepositoryremote_pypi_repository "github.com/upbound/provider-artifactory/internal/controller/remote_pypi_repository/pypirepository"
	repositoryreplication "github.com/upbound/provider-artifactory/internal/controller/remote_repository_replication/repositoryreplication"
	rpmrepositoryremote_rpm_repository "github.com/upbound/provider-artifactory/internal/controller/remote_rpm_repository/rpmrepository"
	sbtrepositoryremote_sbt_repository "github.com/upbound/provider-artifactory/internal/controller/remote_sbt_repository/sbtrepository"
	swiftrepositoryremote_swift_repository "github.com/upbound/provider-artifactory/internal/controller/remote_swift_repository/swiftrepository"
	terraformrepository "github.com/upbound/provider-artifactory/internal/controller/remote_terraform_repository/terraformrepository"
	vcsrepository "github.com/upbound/provider-artifactory/internal/controller/remote_vcs_repository/vcsrepository"
	layout "github.com/upbound/provider-artifactory/internal/controller/repository_layout/layout"
	token "github.com/upbound/provider-artifactory/internal/controller/scoped_token/token"
	userunmanaged_user "github.com/upbound/provider-artifactory/internal/controller/unmanaged_user/user"
	useruser "github.com/upbound/provider-artifactory/internal/controller/user/user"
	alpinerepositoryvirtual_alpine_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_alpine_repository/alpinerepository"
	bowerrepositoryvirtual_bower_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_bower_repository/bowerrepository"
	chefrepositoryvirtual_chef_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_chef_repository/chefrepository"
	composerrepositoryvirtual_composer_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_composer_repository/composerrepository"
	conanrepositoryvirtual_conan_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_conan_repository/conanrepository"
	condarepositoryvirtual_conda_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_conda_repository/condarepository"
	cranrepositoryvirtual_cran_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_cran_repository/cranrepository"
	debianrepositoryvirtual_debian_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_debian_repository/debianrepository"
	dockerrepositoryvirtual_docker_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_docker_repository/dockerrepository"
	gemsrepositoryvirtual_gems_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_gems_repository/gemsrepository"
	genericrepositoryvirtual_generic_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_generic_repository/genericrepository"
	gitlfsrepositoryvirtual_gitlfs_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_gitlfs_repository/gitlfsrepository"
	gorepositoryvirtual_go_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_go_repository/gorepository"
	gradlerepositoryvirtual_gradle_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_gradle_repository/gradlerepository"
	helmrepositoryvirtual_helm_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_helm_repository/helmrepository"
	ivyrepositoryvirtual_ivy_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_ivy_repository/ivyrepository"
	mavenrepositoryvirtual_maven_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_maven_repository/mavenrepository"
	npmrepositoryvirtual_npm_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_npm_repository/npmrepository"
	nugetrepositoryvirtual_nuget_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_nuget_repository/nugetrepository"
	p2repositoryvirtual_p2_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_p2_repository/p2repository"
	pubrepositoryvirtual_pub_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_pub_repository/pubrepository"
	puppetrepositoryvirtual_puppet_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_puppet_repository/puppetrepository"
	pypirepositoryvirtual_pypi_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_pypi_repository/pypirepository"
	rpmrepositoryvirtual_rpm_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_rpm_repository/rpmrepository"
	sbtrepositoryvirtual_sbt_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_sbt_repository/sbtrepository"
	swiftrepositoryvirtual_swift_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_swift_repository/swiftrepository"
	terraformrepositoryvirtual_terraform_repository "github.com/upbound/provider-artifactory/internal/controller/virtual_terraform_repository/terraformrepository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		user.Setup,
		customwebhook.Setup,
		propertycustomwebhook.Setup,
		propertywebhook.Setup,
		webhook.Setup,
		releasebundlecustomwebhook.Setup,
		releasebundlewebhook.Setup,
		backup.Setup,
		customwebhookbuild_custom_webhook.Setup,
		webhookbuild_webhook.Setup,
		certificate.Setup,
		customwebhookdistribution_custom_webhook.Setup,
		publickey.Setup,
		webhookdistribution_webhook.Setup,
		customwebhookdocker_custom_webhook.Setup,
		webhookdocker_webhook.Setup,
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
		dockerrepository.Setup,
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
		alpinerepositorylocal_alpine_repository.Setup,
		bowerrepositorylocal_bower_repository.Setup,
		cargorepositorylocal_cargo_repository.Setup,
		chefrepositorylocal_chef_repository.Setup,
		cocoapodsrepositorylocal_cocoapods_repository.Setup,
		composerrepositorylocal_composer_repository.Setup,
		conanrepositorylocal_conan_repository.Setup,
		condarepositorylocal_conda_repository.Setup,
		cranrepositorylocal_cran_repository.Setup,
		debianrepositorylocal_debian_repository.Setup,
		dockerv1repositorylocal_docker_v1_repository.Setup,
		dockerv2repositorylocal_docker_v2_repository.Setup,
		gemsrepositorylocal_gems_repository.Setup,
		genericrepositorylocal_generic_repository.Setup,
		gitlfsrepositorylocal_gitlfs_repository.Setup,
		gorepositorylocal_go_repository.Setup,
		gradlerepositorylocal_gradle_repository.Setup,
		helmrepositorylocal_helm_repository.Setup,
		huggingfacemlrepository.Setup,
		ivyrepositorylocal_ivy_repository.Setup,
		mavenrepositorylocal_maven_repository.Setup,
		npmrepositorylocal_npm_repository.Setup,
		nugetrepositorylocal_nuget_repository.Setup,
		opkgrepositorylocal_opkg_repository.Setup,
		pubrepository.Setup,
		puppetrepositorylocal_puppet_repository.Setup,
		pypirepositorylocal_pypi_repository.Setup,
		repositorymultireplication.Setup,
		repositorysinglereplication.Setup,
		rpmrepositorylocal_rpm_repository.Setup,
		sbtrepositorylocal_sbt_repository.Setup,
		swiftrepositorylocal_swift_repository.Setup,
		terraformmodulerepositorylocal_terraform_module_repository.Setup,
		terraformproviderrepositorylocal_terraform_provider_repository.Setup,
		terraformbackendrepository.Setup,
		vagrantrepositorylocal_vagrant_repository.Setup,
		server.Setup,
		usermanaged_user.Setup,
		settings.Setup,
		set.Setup,
		providerconfig.Setup,
		proxy.Setup,
		alpinerepositoryremote_alpine_repository.Setup,
		bowerrepositoryremote_bower_repository.Setup,
		cargorepositoryremote_cargo_repository.Setup,
		chefrepositoryremote_chef_repository.Setup,
		cocoapodsrepositoryremote_cocoapods_repository.Setup,
		composerrepositoryremote_composer_repository.Setup,
		conanrepositoryremote_conan_repository.Setup,
		condarepositoryremote_conda_repository.Setup,
		cranrepositoryremote_cran_repository.Setup,
		debianrepositoryremote_debian_repository.Setup,
		dockerrepositoryremote_docker_repository.Setup,
		gemsrepositoryremote_gems_repository.Setup,
		genericrepositoryremote_generic_repository.Setup,
		gitlfsrepositoryremote_gitlfs_repository.Setup,
		gorepositoryremote_go_repository.Setup,
		gradlerepositoryremote_gradle_repository.Setup,
		helmrepositoryremote_helm_repository.Setup,
		huggingfacemlrepositoryremote_huggingfaceml_repository.Setup,
		ivyrepositoryremote_ivy_repository.Setup,
		mavenrepositoryremote_maven_repository.Setup,
		npmrepositoryremote_npm_repository.Setup,
		nugetrepositoryremote_nuget_repository.Setup,
		opkgrepositoryremote_opkg_repository.Setup,
		p2repository.Setup,
		pubrepositoryremote_pub_repository.Setup,
		puppetrepositoryremote_puppet_repository.Setup,
		pypirepositoryremote_pypi_repository.Setup,
		repositoryreplication.Setup,
		rpmrepositoryremote_rpm_repository.Setup,
		sbtrepositoryremote_sbt_repository.Setup,
		swiftrepositoryremote_swift_repository.Setup,
		terraformrepository.Setup,
		vcsrepository.Setup,
		layout.Setup,
		token.Setup,
		userunmanaged_user.Setup,
		useruser.Setup,
		alpinerepositoryvirtual_alpine_repository.Setup,
		bowerrepositoryvirtual_bower_repository.Setup,
		chefrepositoryvirtual_chef_repository.Setup,
		composerrepositoryvirtual_composer_repository.Setup,
		conanrepositoryvirtual_conan_repository.Setup,
		condarepositoryvirtual_conda_repository.Setup,
		cranrepositoryvirtual_cran_repository.Setup,
		debianrepositoryvirtual_debian_repository.Setup,
		dockerrepositoryvirtual_docker_repository.Setup,
		gemsrepositoryvirtual_gems_repository.Setup,
		genericrepositoryvirtual_generic_repository.Setup,
		gitlfsrepositoryvirtual_gitlfs_repository.Setup,
		gorepositoryvirtual_go_repository.Setup,
		gradlerepositoryvirtual_gradle_repository.Setup,
		helmrepositoryvirtual_helm_repository.Setup,
		ivyrepositoryvirtual_ivy_repository.Setup,
		mavenrepositoryvirtual_maven_repository.Setup,
		npmrepositoryvirtual_npm_repository.Setup,
		nugetrepositoryvirtual_nuget_repository.Setup,
		p2repositoryvirtual_p2_repository.Setup,
		pubrepositoryvirtual_pub_repository.Setup,
		puppetrepositoryvirtual_puppet_repository.Setup,
		pypirepositoryvirtual_pypi_repository.Setup,
		rpmrepositoryvirtual_rpm_repository.Setup,
		sbtrepositoryvirtual_sbt_repository.Setup,
		swiftrepositoryvirtual_swift_repository.Setup,
		terraformrepositoryvirtual_terraform_repository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
