---
subcategory: "Container Service (CS)"
layout: "alibabacloudstack"
page_title: "Alibabacloudstack: alibabacloudstack_cs_kubernetes_clusters"
sidebar_current: "docs-alibabacloudstack-datasource-cs-kubernetes-clusters"
description: |-
  Provides a list of Container Service Kubernetes Clusters to be used by the alibabacloudstack_cs_kubernetes_cluster resource.
---

# alibabacloudstack\_cs\_kubernetes\_clusters

This data source provides a list Container Service Kubernetes Clusters on AlibabacloudStack.


## Example Usage

```
# Declare the data source
data "alibabacloudstack_cs_kubernetes_clusters" "k8s_clusters" {
  name_regex  = "my-first-k8s"
}

output "output" {
  value = data.alibabacloudstack_cs_kubernetes_clusters.k8s_clusters.*
}
```

## Argument Reference

The following arguments are supported:

* `ids` - (Optional) Cluster IDs to filter.
* `name_regex` - (Optional) A regex string to filter results by cluster name.
* `state` - (Optional) filter results by cluster state.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `enable_details` - (Optional) Boolean, false by default, only `id` and `name` are exported. Set to true if more details are needed, e.g., `master_disk_category`, `slb_internet_enabled`, `connections`. See full list in attributes.
* `kube_config` - (Optional) Boolean, Set to true to obtain kubeconfig for a cluster and add clusterId in `ids`.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `ids` - A list of matched Kubernetes clusters' ids.
* `names` - A list of matched Kubernetes clusters' names.
* `clusters` - A list of matched Kubernetes clusters. Each element contains the following attributes:
  * `id` - The ID of the container cluster.
  * `name` - The name of the container cluster.
  * `state` - The state of the container cluster.
  * `availability_zone` - The ID of availability zone.
  * `worker_numbers` - The ECS instance node number in the current container cluster.
  * `vswitch_ids` - The ID of VSwitches where the current cluster is located.
  * `vpc_id` - The ID of VPC where the current cluster is located.
  * `slb_internet_enabled` - Whether internet load balancer for API Server is created
  * `security_group_id` - The ID of security group where the current cluster worker node is located.
  * `image_id` - The ID of node image.
  * `nat_gateway_id` - The ID of nat gateway used to launch kubernetes cluster.
  * `master_instance_types` - The instance type of master node.
  * `worker_instance_types` - The instance type of worker node.
  * `master_disk_category` - The system disk category of master node.
  * `master_disk_size` - The system disk size of master node.
  * `worker_disk_category` - The system disk category of worker node.
  * `worker_disk_size` - The system disk size of worker node.
  * `master_nodes` - List of cluster master nodes.
    * `id` - ID of the node.
    * `name` - Node name.
    * `private_ip` - The private IP address of node.
  * `worker_nodes` - List of cluster worker nodes.
    * `id` - ID of the node.
    * `name` - Node name.
    * `private_ip` - The private IP address of node.
  * `connections` - Map of kubernetes cluster connection information.
    * `api_server_internet` - API Server Internet endpoint.
    * `api_server_intranet` - API Server Intranet endpoint.
    * `master_public_ip` - Master node SSH IP address.
    * `service_domain` - Service Access Domain.
  * `node_cidr_mask` - The network mask used on pods for each node.
  * `log_config` - A list of one element containing information about the associated log store. It contains the following attributes:
    * `type` - Type of collecting logs.
    * `project` - Log Service project name.
  * `cluster_network_type` - The network that cluster uses, use flannel or terway.
  * `pod_cidr` - The CIDR block for the pod network when using Flannel.