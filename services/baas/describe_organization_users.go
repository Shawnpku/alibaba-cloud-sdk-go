package baas

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

// DescribeOrganizationUsers invokes the baas.DescribeOrganizationUsers API synchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationusers.html
func (client *Client) DescribeOrganizationUsers(request *DescribeOrganizationUsersRequest) (response *DescribeOrganizationUsersResponse, err error) {
	response = CreateDescribeOrganizationUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOrganizationUsersWithChan invokes the baas.DescribeOrganizationUsers API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationUsersWithChan(request *DescribeOrganizationUsersRequest) (<-chan *DescribeOrganizationUsersResponse, <-chan error) {
	responseChan := make(chan *DescribeOrganizationUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOrganizationUsers(request)
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

// DescribeOrganizationUsersWithCallback invokes the baas.DescribeOrganizationUsers API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationUsersWithCallback(request *DescribeOrganizationUsersRequest, callback func(response *DescribeOrganizationUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOrganizationUsersResponse
		var err error
		defer close(result)
		response, err = client.DescribeOrganizationUsers(request)
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

// DescribeOrganizationUsersRequest is the request struct for api DescribeOrganizationUsers
type DescribeOrganizationUsersRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Query" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
}

// DescribeOrganizationUsersResponse is the response struct for api DescribeOrganizationUsers
type DescribeOrganizationUsersResponse struct {
	*responses.BaseResponse
	RequestId string                                  `json:"RequestId" xml:"RequestId"`
	Success   bool                                    `json:"Success" xml:"Success"`
	ErrorCode int                                     `json:"ErrorCode" xml:"ErrorCode"`
	Result    []ResultItemInDescribeOrganizationUsers `json:"Result" xml:"Result"`
}

// CreateDescribeOrganizationUsersRequest creates a request to invoke DescribeOrganizationUsers API
func CreateDescribeOrganizationUsersRequest() (request *DescribeOrganizationUsersRequest) {
	request = &DescribeOrganizationUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "DescribeOrganizationUsers", "", "")
	return
}

// CreateDescribeOrganizationUsersResponse creates a response to parse from DescribeOrganizationUsers response
func CreateDescribeOrganizationUsersResponse() (response *DescribeOrganizationUsersResponse) {
	response = &DescribeOrganizationUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
