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
)

type DisassociateNetworkAclRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* networkAclId ID  */
    NetworkAclId string `json:"networkAclId"`

    /* networkAcl要解绑的子网ID  */
    SubnetId string `json:"subnetId"`
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param subnetId: networkAcl要解绑的子网ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisassociateNetworkAclRequest(
    regionId string,
    networkAclId string,
    subnetId string,
) *DisassociateNetworkAclRequest {

	return &DisassociateNetworkAclRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/networkAcls/{networkAclId}:disassociateNetworkAcl",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NetworkAclId: networkAclId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param networkAclId: networkAclId ID (Required)
 * param subnetId: networkAcl要解绑的子网ID (Required)
 */
func NewDisassociateNetworkAclRequestWithAllParams(
    regionId string,
    networkAclId string,
    subnetId string,
) *DisassociateNetworkAclRequest {

    return &DisassociateNetworkAclRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:disassociateNetworkAcl",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NetworkAclId: networkAclId,
        SubnetId: subnetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisassociateNetworkAclRequestWithoutParam() *DisassociateNetworkAclRequest {

    return &DisassociateNetworkAclRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/networkAcls/{networkAclId}:disassociateNetworkAcl",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DisassociateNetworkAclRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param networkAclId: networkAclId ID(Required) */
func (r *DisassociateNetworkAclRequest) SetNetworkAclId(networkAclId string) {
    r.NetworkAclId = networkAclId
}

/* param subnetId: networkAcl要解绑的子网ID(Required) */
func (r *DisassociateNetworkAclRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisassociateNetworkAclRequest) GetRegionId() string {
    return r.RegionId
}

type DisassociateNetworkAclResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisassociateNetworkAclResult `json:"result"`
}

type DisassociateNetworkAclResult struct {
}