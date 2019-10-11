// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeQuotasRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* resourceTypes - 资源类型，支持多个[instance，keypair，image，instanceTemplate，imageShare]
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* 私有镜像Id，查询镜像共享(imageShare)配额时，此参数必传 (Optional) */
    ImageId *string `json:"imageId"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeQuotasRequest(
    regionId string,
) *DescribeQuotasRequest {

	return &DescribeQuotasRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quotas",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param filters: resourceTypes - 资源类型，支持多个[instance，keypair，image，instanceTemplate，imageShare]
 (Optional)
 * param imageId: 私有镜像Id，查询镜像共享(imageShare)配额时，此参数必传 (Optional)
 */
func NewDescribeQuotasRequestWithAllParams(
    regionId string,
    filters []common.Filter,
    imageId *string,
) *DescribeQuotasRequest {

    return &DescribeQuotasRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
        ImageId: imageId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeQuotasRequestWithoutParam() *DescribeQuotasRequest {

    return &DescribeQuotasRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeQuotasRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param filters: resourceTypes - 资源类型，支持多个[instance，keypair，image，instanceTemplate，imageShare]
(Optional) */
func (r *DescribeQuotasRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param imageId: 私有镜像Id，查询镜像共享(imageShare)配额时，此参数必传(Optional) */
func (r *DescribeQuotasRequest) SetImageId(imageId string) {
    r.ImageId = &imageId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeQuotasRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeQuotasResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeQuotasResult `json:"result"`
}

type DescribeQuotasResult struct {
    Quotas []vm.Quota `json:"quotas"`
}