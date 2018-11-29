package uis

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateUisNode invokes the uis.CreateUisNode API synchronously
// api document: https://help.aliyun.com/api/uis/createuisnode.html
func (client *Client) CreateUisNode(request *CreateUisNodeRequest) (response *CreateUisNodeResponse, err error) {
	response = CreateCreateUisNodeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUisNodeWithChan invokes the uis.CreateUisNode API asynchronously
// api document: https://help.aliyun.com/api/uis/createuisnode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUisNodeWithChan(request *CreateUisNodeRequest) (<-chan *CreateUisNodeResponse, <-chan error) {
	responseChan := make(chan *CreateUisNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUisNode(request)
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

// CreateUisNodeWithCallback invokes the uis.CreateUisNode API asynchronously
// api document: https://help.aliyun.com/api/uis/createuisnode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUisNodeWithCallback(request *CreateUisNodeRequest, callback func(response *CreateUisNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUisNodeResponse
		var err error
		defer close(result)
		response, err = client.CreateUisNode(request)
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

// CreateUisNodeRequest is the request struct for api CreateUisNode
type CreateUisNodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UisNodeBandwidth     requests.Integer `position:"Query" name:"UisNodeBandwidth"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	UisNodeAreaId        string           `position:"Query" name:"UisNodeAreaId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UisId                string           `position:"Query" name:"UisId"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
	IpAddrsNum           requests.Integer `position:"Query" name:"IpAddrsNum"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateUisNodeResponse is the response struct for api CreateUisNode
type CreateUisNodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	UisNodeId string `json:"UisNodeId" xml:"UisNodeId"`
}

// CreateCreateUisNodeRequest creates a request to invoke CreateUisNode API
func CreateCreateUisNodeRequest() (request *CreateUisNodeRequest) {
	request = &CreateUisNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "CreateUisNode", "uis", "openAPI")
	return
}

// CreateCreateUisNodeResponse creates a response to parse from CreateUisNode response
func CreateCreateUisNodeResponse() (response *CreateUisNodeResponse) {
	response = &CreateUisNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
