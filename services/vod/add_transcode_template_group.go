package vod

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

// AddTranscodeTemplateGroup invokes the vod.AddTranscodeTemplateGroup API synchronously
// api document: https://help.aliyun.com/api/vod/addtranscodetemplategroup.html
func (client *Client) AddTranscodeTemplateGroup(request *AddTranscodeTemplateGroupRequest) (response *AddTranscodeTemplateGroupResponse, err error) {
	response = CreateAddTranscodeTemplateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddTranscodeTemplateGroupWithChan invokes the vod.AddTranscodeTemplateGroup API asynchronously
// api document: https://help.aliyun.com/api/vod/addtranscodetemplategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTranscodeTemplateGroupWithChan(request *AddTranscodeTemplateGroupRequest) (<-chan *AddTranscodeTemplateGroupResponse, <-chan error) {
	responseChan := make(chan *AddTranscodeTemplateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTranscodeTemplateGroup(request)
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

// AddTranscodeTemplateGroupWithCallback invokes the vod.AddTranscodeTemplateGroup API asynchronously
// api document: https://help.aliyun.com/api/vod/addtranscodetemplategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTranscodeTemplateGroupWithCallback(request *AddTranscodeTemplateGroupRequest, callback func(response *AddTranscodeTemplateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTranscodeTemplateGroupResponse
		var err error
		defer close(result)
		response, err = client.AddTranscodeTemplateGroup(request)
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

// AddTranscodeTemplateGroupRequest is the request struct for api AddTranscodeTemplateGroup
type AddTranscodeTemplateGroupRequest struct {
	*requests.RpcRequest
	TranscodeTemplateList    string           `position:"Query" name:"TranscodeTemplateList"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	Name                     string           `position:"Query" name:"Name"`
	ResourceRealOwnerId      requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	TranscodeTemplateGroupId string           `position:"Query" name:"TranscodeTemplateGroupId"`
}

// AddTranscodeTemplateGroupResponse is the response struct for api AddTranscodeTemplateGroup
type AddTranscodeTemplateGroupResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	TranscodeTemplateGroupId string `json:"TranscodeTemplateGroupId" xml:"TranscodeTemplateGroupId"`
}

// CreateAddTranscodeTemplateGroupRequest creates a request to invoke AddTranscodeTemplateGroup API
func CreateAddTranscodeTemplateGroupRequest() (request *AddTranscodeTemplateGroupRequest) {
	request = &AddTranscodeTemplateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "AddTranscodeTemplateGroup", "vod", "openAPI")
	return
}

// CreateAddTranscodeTemplateGroupResponse creates a response to parse from AddTranscodeTemplateGroup response
func CreateAddTranscodeTemplateGroupResponse() (response *AddTranscodeTemplateGroupResponse) {
	response = &AddTranscodeTemplateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
