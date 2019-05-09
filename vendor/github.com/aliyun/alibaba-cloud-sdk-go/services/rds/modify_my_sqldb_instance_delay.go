package rds

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

// ModifyMySQLDBInstanceDelay invokes the rds.ModifyMySQLDBInstanceDelay API synchronously
// api document: https://help.aliyun.com/api/rds/modifymysqldbinstancedelay.html
func (client *Client) ModifyMySQLDBInstanceDelay(request *ModifyMySQLDBInstanceDelayRequest) (response *ModifyMySQLDBInstanceDelayResponse, err error) {
	response = CreateModifyMySQLDBInstanceDelayResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMySQLDBInstanceDelayWithChan invokes the rds.ModifyMySQLDBInstanceDelay API asynchronously
// api document: https://help.aliyun.com/api/rds/modifymysqldbinstancedelay.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMySQLDBInstanceDelayWithChan(request *ModifyMySQLDBInstanceDelayRequest) (<-chan *ModifyMySQLDBInstanceDelayResponse, <-chan error) {
	responseChan := make(chan *ModifyMySQLDBInstanceDelayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMySQLDBInstanceDelay(request)
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

// ModifyMySQLDBInstanceDelayWithCallback invokes the rds.ModifyMySQLDBInstanceDelay API asynchronously
// api document: https://help.aliyun.com/api/rds/modifymysqldbinstancedelay.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMySQLDBInstanceDelayWithCallback(request *ModifyMySQLDBInstanceDelayRequest, callback func(response *ModifyMySQLDBInstanceDelayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMySQLDBInstanceDelayResponse
		var err error
		defer close(result)
		response, err = client.ModifyMySQLDBInstanceDelay(request)
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

// ModifyMySQLDBInstanceDelayRequest is the request struct for api ModifyMySQLDBInstanceDelay
type ModifyMySQLDBInstanceDelayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	SqlDelay             string           `position:"Query" name:"SqlDelay"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyMySQLDBInstanceDelayResponse is the response struct for api ModifyMySQLDBInstanceDelay
type ModifyMySQLDBInstanceDelayResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	SqlDelay       string `json:"SqlDelay" xml:"SqlDelay"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
}

// CreateModifyMySQLDBInstanceDelayRequest creates a request to invoke ModifyMySQLDBInstanceDelay API
func CreateModifyMySQLDBInstanceDelayRequest() (request *ModifyMySQLDBInstanceDelayRequest) {
	request = &ModifyMySQLDBInstanceDelayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyMySQLDBInstanceDelay", "rds", "openAPI")
	return
}

// CreateModifyMySQLDBInstanceDelayResponse creates a response to parse from ModifyMySQLDBInstanceDelay response
func CreateModifyMySQLDBInstanceDelayResponse() (response *ModifyMySQLDBInstanceDelayResponse) {
	response = &ModifyMySQLDBInstanceDelayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}