/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1alpha1 "kubesphere.io/api/application/v1alpha1"
	auditingv1alpha1 "kubesphere.io/api/auditing/v1alpha1"
	clusterv1alpha1 "kubesphere.io/api/cluster/v1alpha1"
	devopsv1alpha1 "kubesphere.io/api/devops/v1alpha1"
	v1alpha3 "kubesphere.io/api/devops/v1alpha3"
	gatewayv1alpha1 "kubesphere.io/api/gateway/v1alpha1"
	v1alpha2 "kubesphere.io/api/iam/v1alpha2"
	networkv1alpha1 "kubesphere.io/api/network/v1alpha1"
	v2beta1 "kubesphere.io/api/notification/v2beta1"
	quotav1alpha2 "kubesphere.io/api/quota/v1alpha2"
	servicemeshv1alpha2 "kubesphere.io/api/servicemesh/v1alpha2"
	storagev1alpha1 "kubesphere.io/api/storage/v1alpha1"
	tenantv1alpha1 "kubesphere.io/api/tenant/v1alpha1"
	tenantv1alpha2 "kubesphere.io/api/tenant/v1alpha2"
	v1beta1 "kubesphere.io/api/types/v1beta1"
	virtualizationv1alpha1 "kubesphere.io/api/virtualization/v1alpha1"
	v1 "kubesphere.io/api/vpc/v1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=application.kubesphere.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("helmapplications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().HelmApplications().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("helmapplicationversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().HelmApplicationVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("helmcategories"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().HelmCategories().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("helmreleases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().HelmReleases().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("helmrepos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().HelmRepos().Informer()}, nil

		// Group=auditing.kubesphere.io, Version=v1alpha1
	case auditingv1alpha1.SchemeGroupVersion.WithResource("rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Auditing().V1alpha1().Rules().Informer()}, nil
	case auditingv1alpha1.SchemeGroupVersion.WithResource("webhooks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Auditing().V1alpha1().Webhooks().Informer()}, nil

		// Group=cluster.kubesphere.io, Version=v1alpha1
	case clusterv1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().Clusters().Informer()}, nil

		// Group=devops.kubesphere.io, Version=v1alpha1
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2ibinaries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iBinaries().Informer()}, nil
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2ibuilders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iBuilders().Informer()}, nil
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2ibuildertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iBuilderTemplates().Informer()}, nil
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2iruns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iRuns().Informer()}, nil

		// Group=devops.kubesphere.io, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("devopsprojects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha3().DevOpsProjects().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("pipelines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha3().Pipelines().Informer()}, nil

		// Group=gateway, Version=v1alpha1
	case gatewayv1alpha1.SchemeGroupVersion.WithResource("gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Gateway().V1alpha1().Gateways().Informer()}, nil

		// Group=iam.kubesphere.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("globalroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().GlobalRoles().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("globalrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().GlobalRoleBindings().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("groups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().Groups().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("groupbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().GroupBindings().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("loginrecords"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().LoginRecords().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("rolebases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().RoleBases().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().Users().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("workspaceroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().WorkspaceRoles().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("workspacerolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().WorkspaceRoleBindings().Informer()}, nil

		// Group=k8s.ovn.org, Version=v1
	case v1.SchemeGroupVersion.WithResource("vpcnetworks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.K8s().V1().VPCNetworks().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("vpcsubnets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.K8s().V1().VPCSubnets().Informer()}, nil

		// Group=network.kubesphere.io, Version=v1alpha1
	case networkv1alpha1.SchemeGroupVersion.WithResource("ipamblocks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().IPAMBlocks().Informer()}, nil
	case networkv1alpha1.SchemeGroupVersion.WithResource("ipamhandles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().IPAMHandles().Informer()}, nil
	case networkv1alpha1.SchemeGroupVersion.WithResource("ippools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().IPPools().Informer()}, nil
	case networkv1alpha1.SchemeGroupVersion.WithResource("namespacenetworkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().NamespaceNetworkPolicies().Informer()}, nil

		// Group=notification.kubesphere.io, Version=v2beta1
	case v2beta1.SchemeGroupVersion.WithResource("configs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Notification().V2beta1().Configs().Informer()}, nil
	case v2beta1.SchemeGroupVersion.WithResource("receivers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Notification().V2beta1().Receivers().Informer()}, nil

		// Group=quota.kubesphere.io, Version=v1alpha2
	case quotav1alpha2.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Quota().V1alpha2().ResourceQuotas().Informer()}, nil

		// Group=servicemesh.kubesphere.io, Version=v1alpha2
	case servicemeshv1alpha2.SchemeGroupVersion.WithResource("servicepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicemesh().V1alpha2().ServicePolicies().Informer()}, nil
	case servicemeshv1alpha2.SchemeGroupVersion.WithResource("strategies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicemesh().V1alpha2().Strategies().Informer()}, nil

		// Group=storage.kubesphere.io, Version=v1alpha1
	case storagev1alpha1.SchemeGroupVersion.WithResource("provisionercapabilities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha1().ProvisionerCapabilities().Informer()}, nil
	case storagev1alpha1.SchemeGroupVersion.WithResource("storageclasscapabilities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha1().StorageClassCapabilities().Informer()}, nil

		// Group=tenant.kubesphere.io, Version=v1alpha1
	case tenantv1alpha1.SchemeGroupVersion.WithResource("workspaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenant().V1alpha1().Workspaces().Informer()}, nil

		// Group=tenant.kubesphere.io, Version=v1alpha2
	case tenantv1alpha2.SchemeGroupVersion.WithResource("workspacetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenant().V1alpha2().WorkspaceTemplates().Informer()}, nil

		// Group=types.kubefed.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("federatedapplications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedApplications().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedclusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedClusterRoles().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedclusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedClusterRoleBindings().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedconfigmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedConfigMaps().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federateddeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedDeployments().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedgroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedGroups().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedgroupbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedGroupBindings().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedIngresses().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedJobs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedlimitranges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedLimitRanges().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatednamespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedNamespaces().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedpersistentvolumeclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedPersistentVolumeClaims().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedSecrets().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedServices().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federatedstatefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Types().V1beta1().FederatedStatefulSets().Informer()}, nil

		// Group=virtualization.ecpaas.io, Version=v1alpha1
	case virtualizationv1alpha1.SchemeGroupVersion.WithResource("virtualmachines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Virtualization().V1alpha1().VirtualMachines().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
