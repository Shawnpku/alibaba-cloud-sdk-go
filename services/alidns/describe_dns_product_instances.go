package alidns

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

func (client *Client) DescribeDnsProductInstances(request *DescribeDnsProductInstancesRequest) (response *DescribeDnsProductInstancesResponse, err error) {
	response = CreateDescribeDnsProductInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDnsProductInstancesWithChan(request *DescribeDnsProductInstancesRequest) (<-chan *DescribeDnsProductInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsProductInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsProductInstances(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeDnsProductInstancesWithCallback(request *DescribeDnsProductInstancesRequest, callback func(response *DescribeDnsProductInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsProductInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsProductInstances(request)
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

type DescribeDnsProductInstancesRequest struct {
	*requests.RpcRequest
	PageSize     string `position:"Query" name:"PageSize"`
	PageNumber   string `position:"Query" name:"PageNumber"`
	VersionCode  string `position:"Query" name:"VersionCode"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type DescribeDnsProductInstancesResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize    int64  `json:"PageSize" xml:"PageSize"`
	DnsProducts struct {
		DnsProduct []struct {
			InstanceId            string `json:"InstanceId" xml:"InstanceId"`
			VersionCode           string `json:"VersionCode" xml:"VersionCode"`
			VersionName           string `json:"VersionName" xml:"VersionName"`
			StartTime             string `json:"StartTime" xml:"StartTime"`
			EndTime               string `json:"EndTime" xml:"EndTime"`
			StartTimestamp        int64  `json:"StartTimestamp" xml:"StartTimestamp"`
			EndTimestamp          int64  `json:"EndTimestamp" xml:"EndTimestamp"`
			Domain                string `json:"Domain" xml:"Domain"`
			BindCount             int64  `json:"BindCount" xml:"BindCount"`
			BindUsedCount         int64  `json:"BindUsedCount" xml:"BindUsedCount"`
			TTLMinValue           int64  `json:"TTLMinValue" xml:"TTLMinValue"`
			SubDomainLevel        int64  `json:"SubDomainLevel" xml:"SubDomainLevel"`
			DnsSLBCount           int64  `json:"DnsSLBCount" xml:"DnsSLBCount"`
			URLForwardCount       int64  `json:"URLForwardCount" xml:"URLForwardCount"`
			DDosDefendFlow        int64  `json:"DDosDefendFlow" xml:"DDosDefendFlow"`
			DDosDefendQuery       int64  `json:"DDosDefendQuery" xml:"DDosDefendQuery"`
			OverseaDDosDefendFlow int64  `json:"OverseaDDosDefendFlow" xml:"OverseaDDosDefendFlow"`
			SearchEngineLines     string `json:"SearchEngineLines" xml:"SearchEngineLines"`
			ISPLines              string `json:"ISPLines" xml:"ISPLines"`
			ISPRegionLines        string `json:"ISPRegionLines" xml:"ISPRegionLines"`
			OverseaLine           string `json:"OverseaLine" xml:"OverseaLine"`
			MonitorNodeCount      int64  `json:"MonitorNodeCount" xml:"MonitorNodeCount"`
			MonitorFrequency      int64  `json:"MonitorFrequency" xml:"MonitorFrequency"`
			MonitorTaskCount      int64  `json:"MonitorTaskCount" xml:"MonitorTaskCount"`
			RegionLines           bool   `json:"RegionLines" xml:"RegionLines"`
			Gslb                  bool   `json:"Gslb" xml:"Gslb"`
			InClean               bool   `json:"InClean" xml:"InClean"`
			InBlackHole           bool   `json:"InBlackHole" xml:"InBlackHole"`
		} `json:"DnsProduct" xml:"DnsProduct"`
	} `json:"DnsProducts" xml:"DnsProducts"`
}

func CreateDescribeDnsProductInstancesRequest() (request *DescribeDnsProductInstancesRequest) {
	request = &DescribeDnsProductInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductInstances", "", "")
	return
}

func CreateDescribeDnsProductInstancesResponse() (response *DescribeDnsProductInstancesResponse) {
	response = &DescribeDnsProductInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}