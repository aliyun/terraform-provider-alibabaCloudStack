package alibabacloudstack

import (
	"github.com/aliyun/terraform-provider-alibabacloudstack/alibabacloudstack/errmsgs"
	"errors"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cr_ee"
)

func (c *CrService) ListCrEEInstances(pageNo int, pageSize int) (*cr_ee.ListInstanceResponse, error) {
	response := &cr_ee.ListInstanceResponse{}
	request := cr_ee.CreateListInstanceRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.PageNo = requests.NewInteger(pageNo)
	request.PageSize = requests.NewInteger(pageSize)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.ListInstance(request)
	})
	response, ok := raw.(*cr_ee.ListInstanceResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, "ListInstance", action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.ListInstanceIsSuccess {
		return response, errmsgs.WrapErrorf(errors.New(response.Code), errmsgs.DataDefaultErrorMsg, "ListInstance", action, errmsgs.AlibabacloudStackSdkGoERROR)
	}
	return response, nil
}

func (c *CrService) DescribeCrEEInstance(instanceId string) (*cr_ee.GetInstanceResponse, error) {
	request := cr_ee.CreateGetInstanceRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	resource := instanceId
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.GetInstance(request)
	})
	response, ok := raw.(*cr_ee.GetInstanceResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"INSTANCE_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.GetInstanceIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}
	return response, nil
}

func (c *CrService) GetCrEEInstanceUsage(instanceId string) (*cr_ee.GetInstanceUsageResponse, error) {
	response := &cr_ee.GetInstanceUsageResponse{}
	request := cr_ee.CreateGetInstanceUsageRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	resource := instanceId
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.GetInstanceUsage(request)
	})
	response, ok := raw.(*cr_ee.GetInstanceUsageResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"INSTANCE_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.GetInstanceUsageIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}
	return response, nil
}

func (c *CrService) ListCrEEInstanceEndpoint(instanceId string) (*cr_ee.ListInstanceEndpointResponse, error) {
	response := &cr_ee.ListInstanceEndpointResponse{}
	request := cr_ee.CreateListInstanceEndpointRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	resource := instanceId
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.ListInstanceEndpoint(request)
	})
	response, ok := raw.(*cr_ee.ListInstanceEndpointResponse)
	if err != nil{
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"INSTANCE_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	response, _ = raw.(*cr_ee.ListInstanceEndpointResponse)
	if !response.ListInstanceEndpointIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}
	return response, nil
}

func (c *CrService) ListCrEENamespaces(instanceId string, pageNo int, pageSize int) (*cr_ee.ListNamespaceResponse, error) {
	response := &cr_ee.ListNamespaceResponse{}
	request := cr_ee.CreateListNamespaceRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.PageNo = requests.NewInteger(pageNo)
	request.PageSize = requests.NewInteger(pageSize)
	resource := c.GenResourceId(instanceId)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.ListNamespace(request)
	})
	response, ok := raw.(*cr_ee.ListNamespaceResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.ListNamespaceIsSuccess {
		return response, errmsgs.WrapErrorf(errors.New(response.Code), errmsgs.DataDefaultErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR)
	}
	return response, nil
}

func (c *CrService) DescribeCrEENamespace(id string) (*cr_ee.GetNamespaceResponse, error) {
	strRet := c.ParseResourceId(id)
	instanceId := strRet[0]
	namespaceName := strRet[1]
	response := &cr_ee.GetNamespaceResponse{}
	request := cr_ee.CreateGetNamespaceRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.NamespaceName = namespaceName
	resource := c.GenResourceId(instanceId, namespaceName)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.GetNamespace(request)
	})
	response, ok := raw.(*cr_ee.GetNamespaceResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"NAMESPACE_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.GetNamespaceIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}
	return response, nil
}

func (c *CrService) DeleteCrEENamespace(instanceId string, namespaceName string) (*cr_ee.DeleteNamespaceResponse, error) {
	response := &cr_ee.DeleteNamespaceResponse{}
	request := cr_ee.CreateDeleteNamespaceRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.NamespaceName = namespaceName
	resource := c.GenResourceId(instanceId, namespaceName)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.DeleteNamespace(request)
	})
	response, ok := raw.(*cr_ee.DeleteNamespaceResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"NAMESPACE_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	response, _ = raw.(*cr_ee.DeleteNamespaceResponse)
	if !response.DeleteNamespaceIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}
	return response, nil
}

func (c *CrService) WaitForCrEENamespace(instanceId string, namespaceName string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	resource := c.GenResourceId(instanceId, namespaceName)

	for {
		resp, err := c.DescribeCrEENamespace(c.GenResourceId(instanceId, namespaceName))
		if err != nil {
			if errmsgs.NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return errmsgs.WrapError(err)
			}
		}

		if resp.NamespaceName == namespaceName && status != Deleted {
			return nil
		}
		if time.Now().After(deadline) {
			return errmsgs.WrapErrorf(err, errmsgs.WaitTimeoutMsg, resource, GetFunc(1), timeout, resp.NamespaceName, namespaceName, errmsgs.ProviderERROR)
		}
		time.Sleep(3 * time.Second)
	}
}

func (c *CrService) ListCrEERepos(instanceId string, namespace string, pageNo int, pageSize int) (*cr_ee.ListRepositoryResponse, error) {
	response := &cr_ee.ListRepositoryResponse{}
	request := cr_ee.CreateListRepositoryRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.RepoNamespaceName = namespace
	request.RepoStatus = "ALL"
	request.PageNo = requests.NewInteger(pageNo)
	request.PageSize = requests.NewInteger(pageSize)
	resource := c.GenResourceId(instanceId, namespace)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.ListRepository(request)
	})
	response, ok := raw.(*cr_ee.ListRepositoryResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	response, _ = raw.(*cr_ee.ListRepositoryResponse)
	if !response.ListRepositoryIsSuccess {
		return response, errmsgs.WrapErrorf(errors.New(response.Code), errmsgs.DataDefaultErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR)
	}
	return response, nil
}

func (c *CrService) DescribeCrEERepo(id string) (*cr_ee.GetRepositoryResponse, error) {
	strRet := c.ParseResourceId(id)
	instanceId := strRet[0]
	namespace := strRet[1]
	repo := strRet[2]
	response := &cr_ee.GetRepositoryResponse{}
	request := cr_ee.CreateGetRepositoryRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.RepoNamespaceName = namespace
	request.RepoName = repo
	resource := c.GenResourceId(instanceId, namespace, repo)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.GetRepository(request)
	})
	response, ok := raw.(*cr_ee.GetRepositoryResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"REPO_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	response, _ = raw.(*cr_ee.GetRepositoryResponse)
	if !response.GetRepositoryIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}

	return response, nil
}

func (c *CrService) DeleteCrEERepo(instanceId, namespace, repo, repoId string) (*cr_ee.DeleteRepositoryResponse, error) {
	response := &cr_ee.DeleteRepositoryResponse{}
	request := cr_ee.CreateDeleteRepositoryRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.RepoId = repoId
	resource := c.GenResourceId(instanceId, namespace, repo)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.DeleteRepository(request)
	})
	response, ok := raw.(*cr_ee.DeleteRepositoryResponse)
	if err != nil {
		errmsg := ""
		if ok {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		if errmsgs.IsExpectedErrors(err, []string{"REPO_NOT_EXIST"}) {
			return response, errmsgs.WrapErrorf(err, errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.DeleteRepositoryIsSuccess {
		return response, c.wrapCrServiceError(resource, action, response.Code)
	}
	return response, nil
}

func (c *CrService) WaitForCrEERepo(instanceId string, namespace string, repo string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	resource := c.GenResourceId(instanceId, namespace, repo)

	for {
		resp, err := c.DescribeCrEERepo(c.GenResourceId(instanceId, namespace, repo))
		if err != nil {
			if errmsgs.NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return errmsgs.WrapError(err)
			}
		}
		if resp.RepoName == repo && status != Deleted {
			return nil
		}
		if time.Now().After(deadline) {
			return errmsgs.WrapErrorf(err, errmsgs.WaitTimeoutMsg, resource, GetFunc(1), timeout, resp.RepoName, repo, errmsgs.ProviderERROR)
		}
		time.Sleep(3 * time.Second)
	}
}

func (c *CrService) ListCrEERepoTags(instanceId string, repoId string, pageNo int, pageSize int) (*cr_ee.ListRepoTagResponse, error) {
	response := &cr_ee.ListRepoTagResponse{}
	request := cr_ee.CreateListRepoTagRequest()
	c.client.InitRpcRequest(*request.RpcRequest)
	request.InstanceId = instanceId
	request.RepoId = repoId
	request.PageNo = requests.NewInteger(pageNo)
	request.PageSize = requests.NewInteger(pageSize)
	resource := c.GenResourceId(instanceId, repoId)
	action := request.GetActionName()

	raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
		return creeClient.ListRepoTag(request)
	})
	response, ok := raw.(*cr_ee.ListRepoTagResponse)
	if err != nil {
		errmsg := ""
		if ok  {
			errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
		}
		return response, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
	}
	addDebug(action, raw, request.RpcRequest, request)

	if !response.ListRepoTagIsSuccess {
		return response, errmsgs.WrapErrorf(errors.New(response.Code), errmsgs.DataDefaultErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR)
	}
	return response, nil
}

func (c *CrService) DescribeCrEESyncRule(id string) (*cr_ee.SyncRulesItem, error) {
	strRet := c.ParseResourceId(id)
	instanceId := strRet[0]
	namespace := strRet[1]
	syncRuleId := strRet[2]

	pageNo := 1
	for {
		response := &cr_ee.ListRepoSyncRuleResponse{}
		request := cr_ee.CreateListRepoSyncRuleRequest()
		c.client.InitRpcRequest(*request.RpcRequest)
		request.InstanceId = instanceId
		request.NamespaceName = namespace
		request.PageNo = requests.NewInteger(pageNo)
		request.PageSize = requests.NewInteger(PageSizeLarge)
		raw, err := c.client.WithCrEEClient(func(creeClient *cr_ee.Client) (interface{}, error) {
			return creeClient.ListRepoSyncRule(request)
		})
		response, ok := raw.(*cr_ee.ListRepoSyncRuleResponse)
		if err != nil {
			errmsg := ""
			if ok {
				errmsg = errmsgs.GetBaseResponseErrorMessage(response.BaseResponse)
			}
			return nil, errmsgs.WrapErrorf(err, errmsgs.RequestV1ErrorMsg, id, request.GetActionName(), errmsgs.AlibabacloudStackSdkGoERROR, errmsg)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)

		if !response.ListRepoSyncRuleIsSuccess {
			return nil, c.wrapCrServiceError(id, request.GetActionName(), response.Code)
		}

		for _, rule := range response.SyncRules {
			if rule.SyncRuleId == syncRuleId && rule.LocalInstanceId == instanceId {
				return &rule, nil
			}
		}

		if len(response.SyncRules) < PageSizeLarge {
			return nil, errmsgs.WrapErrorf(errors.New("sync rule not found"), errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
		}

		pageNo++
	}
}

func (c *CrService) wrapCrServiceError(resource string, action string, code string) error {
	switch code {
	case "INSTANCE_NOT_EXIST", "NAMESPACE_NOT_EXIST", "REPO_NOT_EXIST":
		return errmsgs.WrapErrorf(errors.New(code), errmsgs.NotFoundMsg, errmsgs.AlibabacloudStackSdkGoERROR)
	default:
		return errmsgs.WrapErrorf(errors.New(code), errmsgs.DefaultErrorMsg, resource, action, errmsgs.AlibabacloudStackSdkGoERROR)
	}
}

func (c *CrService) GenResourceId(args ...string) string {
	return strings.Join(args, COLON_SEPARATED)
}

func (c *CrService) ParseResourceId(id string) []string {
	return strings.Split(id, COLON_SEPARATED)
}
