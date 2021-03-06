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

// DisplayAIAuditSwitch invokes the vod.DisplayAIAuditSwitch API synchronously
// api document: https://help.aliyun.com/api/vod/displayaiauditswitch.html
func (client *Client) DisplayAIAuditSwitch(request *DisplayAIAuditSwitchRequest) (response *DisplayAIAuditSwitchResponse, err error) {
	response = CreateDisplayAIAuditSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// DisplayAIAuditSwitchWithChan invokes the vod.DisplayAIAuditSwitch API asynchronously
// api document: https://help.aliyun.com/api/vod/displayaiauditswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisplayAIAuditSwitchWithChan(request *DisplayAIAuditSwitchRequest) (<-chan *DisplayAIAuditSwitchResponse, <-chan error) {
	responseChan := make(chan *DisplayAIAuditSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisplayAIAuditSwitch(request)
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

// DisplayAIAuditSwitchWithCallback invokes the vod.DisplayAIAuditSwitch API asynchronously
// api document: https://help.aliyun.com/api/vod/displayaiauditswitch.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisplayAIAuditSwitchWithCallback(request *DisplayAIAuditSwitchRequest, callback func(response *DisplayAIAuditSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisplayAIAuditSwitchResponse
		var err error
		defer close(result)
		response, err = client.DisplayAIAuditSwitch(request)
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

// DisplayAIAuditSwitchRequest is the request struct for api DisplayAIAuditSwitch
type DisplayAIAuditSwitchRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DisplayAIAuditSwitchResponse is the response struct for api DisplayAIAuditSwitch
type DisplayAIAuditSwitchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IsDisplay bool   `json:"IsDisplay" xml:"IsDisplay"`
}

// CreateDisplayAIAuditSwitchRequest creates a request to invoke DisplayAIAuditSwitch API
func CreateDisplayAIAuditSwitchRequest() (request *DisplayAIAuditSwitchRequest) {
	request = &DisplayAIAuditSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DisplayAIAuditSwitch", "vod", "openAPI")
	return
}

// CreateDisplayAIAuditSwitchResponse creates a response to parse from DisplayAIAuditSwitch response
func CreateDisplayAIAuditSwitchResponse() (response *DisplayAIAuditSwitchResponse) {
	response = &DisplayAIAuditSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
