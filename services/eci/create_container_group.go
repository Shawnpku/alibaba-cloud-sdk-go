//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package eci

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateContainerGroup invokes the eci.CreateContainerGroup API synchronously
// api document: https://help.aliyun.com/api/eci/createcontainergroup.html
func (client *Client) CreateContainerGroup(request *CreateContainerGroupRequest) (response *CreateContainerGroupResponse, err error) {
	response = CreateCreateContainerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateContainerGroupWithChan invokes the eci.CreateContainerGroup API asynchronously
// api document: https://help.aliyun.com/api/eci/createcontainergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateContainerGroupWithChan(request *CreateContainerGroupRequest) (<-chan *CreateContainerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateContainerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateContainerGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateContainerGroupWithCallback invokes the eci.CreateContainerGroup API asynchronously
// api document: https://help.aliyun.com/api/eci/createcontainergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateContainerGroupWithCallback(request *CreateContainerGroupRequest, callback func(response *CreateContainerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateContainerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateContainerGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateContainerGroupRequest is the request struct for api CreateContainerGroup
type CreateContainerGroupRequest struct {
	*requests.RpcRequest
	Action                  string                                          `position:"Query" name:"Action"`
	OwnerId                 requests.Integer                                `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount    string                                          `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId         requests.Integer                                `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount            string                                          `position:"Query" name:"OwnerAccount"`
	RegionId                string                                          `position:"Query" name:"RegionId"`
	ZoneId                  string                                          `position:"Query" name:"ZoneId"`
	SecurityGroupId         string                                          `position:"Query" name:"SecurityGroupId"`
	VSwitchId               string                                          `position:"Query" name:"VSwitchId"`
	ContainerGroupName      string                                          `position:"Query" name:"ContainerGroupName"`
	RestartPolicy           string                                          `position:"Query" name:"RestartPolicy"`
	Tag                     *[]CreateContainerGroup_Tag                     `position:"Query" name:"Tag"`
	ImageRegistryCredential *[]CreateContainerGroup_ImageRegistryCredential `position:"Query" name:"ImageRegistryCredential"`
	Container               *[]CreateContainerGroup_Container               `position:"Query" name:"Container"`
	Volume                  *[]CreateContainerGroup_Volume                  `position:"Query" name:"Volume"`
	EipInstanceId           string                                          `position:"Query" name:"EipInstanceId"`
	InitContainer           *[]CreateContainerGroup_InitContainer           `position:"Query" name:"InitContainer"`
	Cpu                     requests.Float                                  `position:"Query" name:"Cpu"`
	Memory                  requests.Float                                  `position:"Query" name:"Memory"`
	DnsConfig               CreateContainerGroup_DnsConfig                  `position:"Query" name:"DnsConfig"`
}

type CreateContainerGroup_Tag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type CreateContainerGroup_ImageRegistryCredential struct {
	Server   string `name:"Server"`
	UserName string `name:"UserName"`
	Password string `name:"Password"`
}

type CreateContainerGroup_Container struct {
	Image           string                                 `name:"Image"`
	Name            string                                 `name:"Name"`
	Cpu             requests.Float                         `name:"Cpu"`
	Memory          requests.Float                         `name:"Memory"`
	WorkingDir      string                                 `name:"WorkingDir"`
	ImagePullPolicy string                                 `name:"ImagePullPolicy"`
	Command         []string                               `name:"Command"`
	Arg             []string                               `name:"Arg"`
	VolumeMount     *[]CreateContainerGroup_VolumeMount    `name:"VolumeMount"`
	Port            *[]CreateContainerGroup_Port           `name:"Port"`
	EnvironmentVar  *[]CreateContainerGroup_EnvironmentVar `name:"EnvironmentVar"`
	ReadinessProbe  CreateContainerGroup_ReadinessProbe    `name:"ReadinessProbe"`
	LivenessProbe   CreateContainerGroup_LivenessProbe     `name:"LivenessProbe"`
	SecurityContext CreateContainerGroup_SecurityContext   `name:"SecurityContext"`
}

type CreateContainerGroup_Volume struct {
	Name             string                                `name:"Name"`
	Type             string                                `name:"Type"`
	NFSVolume        CreateContainerGroup_NFSVolume        `name:"NFSVolume"`
	ConfigFileVolume CreateContainerGroup_ConfigFileVolume `name:"ConfigFileVolume"`
}

type CreateContainerGroup_InitContainer struct {
	Name            string                                 `name:"Name"`
	Image           string                                 `name:"Image"`
	Cpu             requests.Float                         `name:"Cpu"`
	Memory          requests.Float                         `name:"Memory"`
	WorkingDir      string                                 `name:"WorkingDir"`
	ImagePullPolicy string                                 `name:"ImagePullPolicy"`
	Command         []string                               `name:"Command"`
	Arg             []string                               `name:"Arg"`
	VolumeMount     *[]CreateContainerGroup_VolumeMount    `name:"VolumeMount"`
	Port            *[]CreateContainerGroup_Port           `name:"Port"`
	EnvironmentVar  *[]CreateContainerGroup_EnvironmentVar `name:"EnvironmentVar"`
	SecurityContext CreateContainerGroup_SecurityContext   `name:"SecurityContext"`
}

type CreateContainerGroup_DnsConfig struct {
	NameServer []string                       `name:"NameServer"`
	Search     []string                       `name:"Search"`
	Option     *[]CreateContainerGroup_Option `name:"Option"`
}

type CreateContainerGroup_VolumeMount struct {
	MountPath string           `name:"MountPath"`
	ReadOnly  requests.Boolean `name:"ReadOnly"`
	Name      string           `name:"Name"`
}

type CreateContainerGroup_Port struct {
	Protocol string           `name:"Protocol"`
	Port     requests.Integer `name:"Port"`
}

type CreateContainerGroup_EnvironmentVar struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type CreateContainerGroup_ReadinessProbe struct {
	InitialDelaySeconds requests.Integer               `name:"InitialDelaySeconds"`
	PeriodSeconds       requests.Integer               `name:"PeriodSeconds"`
	SuccessThreshold    requests.Integer               `name:"SuccessThreshold"`
	FailureThreshold    requests.Integer               `name:"FailureThreshold"`
	TimeoutSeconds      requests.Integer               `name:"TimeoutSeconds"`
	HttpGet             CreateContainerGroup_HttpGet   `name:"HttpGet"`
	Exec                CreateContainerGroup_Exec      `name:"Exec"`
	TcpSocket           CreateContainerGroup_TcpSocket `name:"TcpSocket"`
}

type CreateContainerGroup_HttpGet struct {
	Path   string           `name:"Path"`
	Port   requests.Integer `name:"Port"`
	Scheme string           `name:"Scheme"`
}

type CreateContainerGroup_Exec struct {
	Command []string `name:"Command"`
}

type CreateContainerGroup_TcpSocket struct {
	Port requests.Integer `name:"Port"`
}

type CreateContainerGroup_LivenessProbe struct {
	InitialDelaySeconds requests.Integer               `name:"InitialDelaySeconds"`
	PeriodSeconds       requests.Integer               `name:"PeriodSeconds"`
	SuccessThreshold    requests.Integer               `name:"SuccessThreshold"`
	FailureThreshold    requests.Integer               `name:"FailureThreshold"`
	TimeoutSeconds      requests.Integer               `name:"TimeoutSeconds"`
	HttpGet             CreateContainerGroup_HttpGet   `name:"HttpGet"`
	Exec                CreateContainerGroup_Exec      `name:"Exec"`
	TcpSocket           CreateContainerGroup_TcpSocket `name:"TcpSocket"`
}

type CreateContainerGroup_SecurityContext struct {
	ReadOnlyRootFilesystem requests.Boolean                `name:"ReadOnlyRootFilesystem"`
	RunAsUser              requests.Integer                `name:"RunAsUser"`
	Capability             CreateContainerGroup_Capability `name:"Capability"`
}

type CreateContainerGroup_Capability struct {
	Add []string `name:"Add"`
}

type CreateContainerGroup_NFSVolume struct {
	Server   string           `name:"Server"`
	Path     string           `name:"Path"`
	ReadOnly requests.Boolean `name:"ReadOnly"`
}

type CreateContainerGroup_ConfigFileVolume struct {
	ConfigFileToPath *[]CreateContainerGroup_ConfigFileToPath `name:"ConfigFileToPath"`
}

type CreateContainerGroup_ConfigFileToPath struct {
	Content string `name:"Content"`
	Path    string `name:"Path"`
}

type CreateContainerGroup_Option struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

// CreateContainerGroupResponse is the response struct for api CreateContainerGroup
type CreateContainerGroupResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ContainerGroupId string `json:"ContainerGroupId" xml:"ContainerGroupId"`
}

// CreateCreateContainerGroupRequest creates a request to invoke CreateContainerGroup API
func CreateCreateContainerGroupRequest() (request *CreateContainerGroupRequest) {
	request = &CreateContainerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eci", "2018-08-08", "CreateContainerGroup", "eci", "openAPI")
	return
}

// CreateCreateContainerGroupResponse creates a response to parse from CreateContainerGroup response
func CreateCreateContainerGroupResponse() (response *CreateContainerGroupResponse) {
	response = &CreateContainerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
