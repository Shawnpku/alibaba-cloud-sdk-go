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

package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeTieringJobs invokes the nas.DescribeTieringJobs API synchronously
// api document: https://help.aliyun.com/api/nas/describetieringjobs.html
func (client *Client) DescribeTieringJobs(request *DescribeTieringJobsRequest) (response *DescribeTieringJobsResponse, err error) {
	response = CreateDescribeTieringJobsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTieringJobsWithChan invokes the nas.DescribeTieringJobs API asynchronously
// api document: https://help.aliyun.com/api/nas/describetieringjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTieringJobsWithChan(request *DescribeTieringJobsRequest) (<-chan *DescribeTieringJobsResponse, <-chan error) {
	responseChan := make(chan *DescribeTieringJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTieringJobs(request)
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

// DescribeTieringJobsWithCallback invokes the nas.DescribeTieringJobs API asynchronously
// api document: https://help.aliyun.com/api/nas/describetieringjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTieringJobsWithCallback(request *DescribeTieringJobsRequest, callback func(response *DescribeTieringJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTieringJobsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTieringJobs(request)
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

// DescribeTieringJobsRequest is the request struct for api DescribeTieringJobs
type DescribeTieringJobsRequest struct {
	*requests.RpcRequest
	Volume string `position:"Query" name:"Volume"`
}

// DescribeTieringJobsResponse is the response struct for api DescribeTieringJobs
type DescribeTieringJobsResponse struct {
	*responses.BaseResponse
	RequestId   string                          `json:"RequestId" xml:"RequestId"`
	TotalCount  int                             `json:"TotalCount" xml:"TotalCount"`
	PageSize    int                             `json:"PageSize" xml:"PageSize"`
	PageNumber  int                             `json:"PageNumber" xml:"PageNumber"`
	TieringJobs DescribeTieringJobs_TieringJobs `json:"TieringJobs" xml:"TieringJobs"`
}

type DescribeTieringJobs_TieringJobs struct {
	TieringJob []DescribeTieringJobs_TieringJob `json:"TieringJob" xml:"TieringJob"`
}

type DescribeTieringJobs_TieringJob struct {
	Name           string `json:"Name" xml:"Name"`
	Volume         string `json:"Volume" xml:"Volume"`
	Path           string `json:"Path" xml:"Path"`
	Recursive      bool   `json:"Recursive" xml:"Recursive"`
	Type           string `json:"Type" xml:"Type"`
	Policy         string `json:"Policy" xml:"Policy"`
	Weekday        int    `json:"Weekday" xml:"Weekday"`
	Hour           int    `json:"Hour" xml:"Hour"`
	Enabled        bool   `json:"Enabled" xml:"Enabled"`
	Status         string `json:"Status" xml:"Status"`
	LastUpdateTime int64  `json:"LastUpdateTime" xml:"LastUpdateTime"`
}

// CreateDescribeTieringJobsRequest creates a request to invoke DescribeTieringJobs API
func CreateDescribeTieringJobsRequest() (request *DescribeTieringJobsRequest) {
	request = &DescribeTieringJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeTieringJobs", "nas", "openAPI")
	return
}

// CreateDescribeTieringJobsResponse creates a response to parse from DescribeTieringJobs response
func CreateDescribeTieringJobsResponse() (response *DescribeTieringJobsResponse) {
	response = &DescribeTieringJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
