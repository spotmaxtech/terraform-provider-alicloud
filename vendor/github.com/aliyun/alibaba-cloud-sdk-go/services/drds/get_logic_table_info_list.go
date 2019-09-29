package drds

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

// GetLogicTableInfoList invokes the drds.GetLogicTableInfoList API synchronously
// api document: https://help.aliyun.com/api/drds/getlogictableinfolist.html
func (client *Client) GetLogicTableInfoList(request *GetLogicTableInfoListRequest) (response *GetLogicTableInfoListResponse, err error) {
	response = CreateGetLogicTableInfoListResponse()
	err = client.DoAction(request, response)
	return
}

// GetLogicTableInfoListWithChan invokes the drds.GetLogicTableInfoList API asynchronously
// api document: https://help.aliyun.com/api/drds/getlogictableinfolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLogicTableInfoListWithChan(request *GetLogicTableInfoListRequest) (<-chan *GetLogicTableInfoListResponse, <-chan error) {
	responseChan := make(chan *GetLogicTableInfoListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLogicTableInfoList(request)
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

// GetLogicTableInfoListWithCallback invokes the drds.GetLogicTableInfoList API asynchronously
// api document: https://help.aliyun.com/api/drds/getlogictableinfolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLogicTableInfoListWithCallback(request *GetLogicTableInfoListRequest, callback func(response *GetLogicTableInfoListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLogicTableInfoListResponse
		var err error
		defer close(result)
		response, err = client.GetLogicTableInfoList(request)
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

// GetLogicTableInfoListRequest is the request struct for api GetLogicTableInfoList
type GetLogicTableInfoListRequest struct {
	*requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// GetLogicTableInfoListResponse is the response struct for api GetLogicTableInfoList
type GetLogicTableInfoListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGetLogicTableInfoListRequest creates a request to invoke GetLogicTableInfoList API
func CreateGetLogicTableInfoListRequest() (request *GetLogicTableInfoListRequest) {
	request = &GetLogicTableInfoListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "GetLogicTableInfoList", "drds", "openAPI")
	return
}

// CreateGetLogicTableInfoListResponse creates a response to parse from GetLogicTableInfoList response
func CreateGetLogicTableInfoListResponse() (response *GetLogicTableInfoListResponse) {
	response = &GetLogicTableInfoListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}