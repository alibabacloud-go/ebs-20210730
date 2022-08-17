// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddDiskReplicaPairRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	ReplicaPairId  *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s AddDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairRequest) SetClientToken(v string) *AddDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetRegionId(v string) *AddDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetReplicaGroupId(v string) *AddDiskReplicaPairRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetReplicaPairId(v string) *AddDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type AddDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairResponseBody) SetRequestId(v string) *AddDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type AddDiskReplicaPairResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairResponse) SetHeaders(v map[string]*string) *AddDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *AddDiskReplicaPairResponse) SetStatusCode(v int32) *AddDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDiskReplicaPairResponse) SetBody(v *AddDiskReplicaPairResponseBody) *AddDiskReplicaPairResponse {
	s.Body = v
	return s
}

type CreateDedicatedBlockStorageClusterRequest struct {
	Azone    *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	Capacity *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	DbscId   *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetAzone(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.Azone = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetCapacity(v int64) *CreateDedicatedBlockStorageClusterRequest {
	s.Capacity = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetDbscId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.DbscId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetDbscName(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.DbscName = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetRegionId(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterRequest) SetType(v string) *CreateDedicatedBlockStorageClusterRequest {
	s.Type = &v
	return s
}

type CreateDedicatedBlockStorageClusterResponseBody struct {
	DbscId    *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetDbscId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.DbscId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetOrderId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponseBody) SetRequestId(v string) *CreateDedicatedBlockStorageClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateDedicatedBlockStorageClusterResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDedicatedBlockStorageClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDedicatedBlockStorageClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetHeaders(v map[string]*string) *CreateDedicatedBlockStorageClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetStatusCode(v int32) *CreateDedicatedBlockStorageClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetBody(v *CreateDedicatedBlockStorageClusterResponseBody) *CreateDedicatedBlockStorageClusterResponse {
	s.Body = v
	return s
}

type CreateDiskReplicaGroupRequest struct {
	Bandwidth           *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationZoneId   *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RPO                 *int64  `json:"RPO,omitempty" xml:"RPO,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceZoneId        *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
}

func (s CreateDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupRequest) SetBandwidth(v int64) *CreateDiskReplicaGroupRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetClientToken(v string) *CreateDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDescription(v string) *CreateDiskReplicaGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDestinationRegionId(v string) *CreateDiskReplicaGroupRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDestinationZoneId(v string) *CreateDiskReplicaGroupRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetGroupName(v string) *CreateDiskReplicaGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetRPO(v int64) *CreateDiskReplicaGroupRequest {
	s.RPO = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetRegionId(v string) *CreateDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetSourceZoneId(v string) *CreateDiskReplicaGroupRequest {
	s.SourceZoneId = &v
	return s
}

type CreateDiskReplicaGroupResponseBody struct {
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupResponseBody) SetReplicaGroupId(v string) *CreateDiskReplicaGroupResponseBody {
	s.ReplicaGroupId = &v
	return s
}

func (s *CreateDiskReplicaGroupResponseBody) SetRequestId(v string) *CreateDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaGroupResponse) SetStatusCode(v int32) *CreateDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskReplicaGroupResponse) SetBody(v *CreateDiskReplicaGroupResponseBody) *CreateDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type CreateDiskReplicaPairRequest struct {
	Bandwidth           *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ChargeType          *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationDiskId   *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationZoneId   *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	DiskId              *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	PairName            *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	Period              *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit          *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RPO                 *int64  `json:"RPO,omitempty" xml:"RPO,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceZoneId        *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
}

func (s CreateDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairRequest) SetBandwidth(v int64) *CreateDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetChargeType(v string) *CreateDiskReplicaPairRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetClientToken(v string) *CreateDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDescription(v string) *CreateDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationDiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationRegionId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationZoneId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPairName(v string) *CreateDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriod(v int64) *CreateDiskReplicaPairRequest {
	s.Period = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriodUnit(v string) *CreateDiskReplicaPairRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetRPO(v int64) *CreateDiskReplicaPairRequest {
	s.RPO = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetRegionId(v string) *CreateDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetSourceZoneId(v string) *CreateDiskReplicaPairRequest {
	s.SourceZoneId = &v
	return s
}

type CreateDiskReplicaPairResponseBody struct {
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponseBody) SetOrderId(v string) *CreateDiskReplicaPairResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetReplicaPairId(v string) *CreateDiskReplicaPairResponseBody {
	s.ReplicaPairId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetRequestId(v string) *CreateDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetStatusCode(v int32) *CreateDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetBody(v *CreateDiskReplicaPairResponseBody) *CreateDiskReplicaPairResponse {
	s.Body = v
	return s
}

type DeleteDiskReplicaGroupRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s DeleteDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupRequest) SetClientToken(v string) *DeleteDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) SetRegionId(v string) *DeleteDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) SetReplicaGroupId(v string) *DeleteDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type DeleteDiskReplicaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupResponseBody) SetRequestId(v string) *DeleteDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) SetStatusCode(v int32) *DeleteDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskReplicaGroupResponse) SetBody(v *DeleteDiskReplicaGroupResponseBody) *DeleteDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type DeleteDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s DeleteDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairRequest) SetClientToken(v string) *DeleteDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetRegionId(v string) *DeleteDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetReplicaPairId(v string) *DeleteDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type DeleteDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponseBody) SetRequestId(v string) *DeleteDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetStatusCode(v int32) *DeleteDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetBody(v *DeleteDiskReplicaPairResponseBody) *DeleteDiskReplicaPairResponse {
	s.Body = v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksRequest struct {
	DbscId     *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetDbscId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.DbscId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetMaxResults(v int64) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksRequest) SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksRequest {
	s.RegionId = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBody struct {
	Disks     *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	NextToken *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetDisks(v *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetRequestId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks struct {
	Disk []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) SetDisk(v []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks {
	s.Disk = v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk struct {
	AttachedTime              *string                                                               `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	BdfId                     *string                                                               `json:"BdfId,omitempty" xml:"BdfId,omitempty"`
	Category                  *string                                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteAutoSnapshot        *bool                                                                 `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	DeleteWithInstance        *bool                                                                 `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description               *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	DetachedTime              *string                                                               `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	Device                    *string                                                               `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskChargeType            *string                                                               `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	DiskId                    *string                                                               `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskName                  *string                                                               `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	EnableAutoSnapshot        *bool                                                                 `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	Encrypted                 *bool                                                                 `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	IOPS                      *int64                                                                `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	ImageId                   *string                                                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceId                *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KMSKeyId                  *string                                                               `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	MountInstanceNum          *int32                                                                `json:"MountInstanceNum,omitempty" xml:"MountInstanceNum,omitempty"`
	MultiAttach               *string                                                               `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	PerformanceLevel          *string                                                               `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Portable                  *bool                                                                 `json:"Portable,omitempty" xml:"Portable,omitempty"`
	RegionId                  *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size                      *int32                                                                `json:"Size,omitempty" xml:"Size,omitempty"`
	SourceSnapshotId          *string                                                               `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	Status                    *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageClusterId          *string                                                               `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	StorageSetId              *string                                                               `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	StorageSetPartitionNumber *int32                                                                `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	Tags                      []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Type                      *string                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	ZoneId                    *string                                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetAttachedTime(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetBdfId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.BdfId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetCategory(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDeleteAutoSnapshot(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDeleteWithInstance(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDescription(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDetachedTime(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DetachedTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDevice(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Device = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskChargeType(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskName(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetEnableAutoSnapshot(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetEncrypted(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetIOPS(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.IOPS = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetImageId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ImageId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetInstanceId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetKMSKeyId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetMountInstanceNum(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.MountInstanceNum = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetMultiAttach(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.MultiAttach = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetPerformanceLevel(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetPortable(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Portable = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetSize(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Size = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetSourceSnapshotId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStatus(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageClusterId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageClusterId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageSetId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageSetId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageSetPartitionNumber(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetTags(v []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetType(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetZoneId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) SetTagKey(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) SetTagValue(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	s.TagValue = &v
	return s
}

type DescribeDedicatedBlockStorageClusterDisksResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDedicatedBlockStorageClusterDisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetStatusCode(v int32) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponse) SetBody(v *DescribeDedicatedBlockStorageClusterDisksResponseBody) *DescribeDedicatedBlockStorageClusterDisksResponse {
	s.Body = v
	return s
}

type DescribeDedicatedBlockStorageClustersRequest struct {
	AzoneId                        *string   `json:"AzoneId,omitempty" xml:"AzoneId,omitempty"`
	Category                       *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientToken                    *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DedicatedBlockStorageClusterId []*string `json:"DedicatedBlockStorageClusterId,omitempty" xml:"DedicatedBlockStorageClusterId,omitempty" type:"Repeated"`
	MaxResults                     *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                      *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId                       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                         []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedBlockStorageClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetAzoneId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.AzoneId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetCategory(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetClientToken(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetDedicatedBlockStorageClusterId(v []*string) *DescribeDedicatedBlockStorageClustersRequest {
	s.DedicatedBlockStorageClusterId = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetMaxResults(v int32) *DescribeDedicatedBlockStorageClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetNextToken(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetRegionId(v string) *DescribeDedicatedBlockStorageClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersRequest) SetStatus(v []*string) *DescribeDedicatedBlockStorageClustersRequest {
	s.Status = v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBody struct {
	DedicatedBlockStorageClusters []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters `json:"DedicatedBlockStorageClusters,omitempty" xml:"DedicatedBlockStorageClusters,omitempty" type:"Repeated"`
	NextToken                     *string                                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                     *string                                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetDedicatedBlockStorageClusters(v []*DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.DedicatedBlockStorageClusters = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetNextToken(v string) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBody) SetRequestId(v string) *DescribeDedicatedBlockStorageClustersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters struct {
	Category                             *string                                                                                                             `json:"Category,omitempty" xml:"Category,omitempty"`
	CreateTime                           *string                                                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DedicatedBlockStorageClusterCapacity *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity `json:"DedicatedBlockStorageClusterCapacity,omitempty" xml:"DedicatedBlockStorageClusterCapacity,omitempty" type:"Struct"`
	DedicatedBlockStorageClusterId       *string                                                                                                             `json:"DedicatedBlockStorageClusterId,omitempty" xml:"DedicatedBlockStorageClusterId,omitempty"`
	DedicatedBlockStorageClusterName     *string                                                                                                             `json:"DedicatedBlockStorageClusterName,omitempty" xml:"DedicatedBlockStorageClusterName,omitempty"`
	Description                          *string                                                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredTime                          *string                                                                                                             `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	PerformanceLevel                     *string                                                                                                             `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	RegionId                             *string                                                                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                               *string                                                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedCategory                    *string                                                                                                             `json:"SupportedCategory,omitempty" xml:"SupportedCategory,omitempty"`
	Type                                 *string                                                                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	ZoneId                               *string                                                                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetCategory(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetCreateTime(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterCapacity(v *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterCapacity = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDedicatedBlockStorageClusterName(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.DedicatedBlockStorageClusterName = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetDescription(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetExpiredTime(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetPerformanceLevel(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetRegionId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetStatus(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetSupportedCategory(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.SupportedCategory = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetType(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters) SetZoneId(v string) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClusters {
	s.ZoneId = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity struct {
	AvailableCapacity *int64 `json:"AvailableCapacity,omitempty" xml:"AvailableCapacity,omitempty"`
	DeliveryCapacity  *int64 `json:"DeliveryCapacity,omitempty" xml:"DeliveryCapacity,omitempty"`
	TotalCapacity     *int64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	UsedCapacity      *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetAvailableCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.AvailableCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetDeliveryCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.DeliveryCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetTotalCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity) SetUsedCapacity(v int64) *DescribeDedicatedBlockStorageClustersResponseBodyDedicatedBlockStorageClustersDedicatedBlockStorageClusterCapacity {
	s.UsedCapacity = &v
	return s
}

type DescribeDedicatedBlockStorageClustersResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDedicatedBlockStorageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDedicatedBlockStorageClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetHeaders(v map[string]*string) *DescribeDedicatedBlockStorageClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetStatusCode(v int32) *DescribeDedicatedBlockStorageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClustersResponse) SetBody(v *DescribeDedicatedBlockStorageClustersResponseBody) *DescribeDedicatedBlockStorageClustersResponse {
	s.Body = v
	return s
}

type DescribeDiskReplicaGroupsRequest struct {
	GroupIds   *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Site       *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s DescribeDiskReplicaGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsRequest) SetGroupIds(v string) *DescribeDiskReplicaGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetMaxResults(v int64) *DescribeDiskReplicaGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetNextToken(v string) *DescribeDiskReplicaGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetRegionId(v string) *DescribeDiskReplicaGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsRequest) SetSite(v string) *DescribeDiskReplicaGroupsRequest {
	s.Site = &v
	return s
}

type DescribeDiskReplicaGroupsResponseBody struct {
	NextToken     *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ReplicaGroups []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups `json:"ReplicaGroups,omitempty" xml:"ReplicaGroups,omitempty" type:"Repeated"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetNextToken(v string) *DescribeDiskReplicaGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetReplicaGroups(v []*DescribeDiskReplicaGroupsResponseBodyReplicaGroups) *DescribeDiskReplicaGroupsResponseBody {
	s.ReplicaGroups = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBody) SetRequestId(v string) *DescribeDiskReplicaGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiskReplicaGroupsResponseBodyReplicaGroups struct {
	Bandwidth           *int64   `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description         *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationRegionId *string  `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationZoneId   *string  `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	GroupName           *string  `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LastRecoverPoint    *int64   `json:"LastRecoverPoint,omitempty" xml:"LastRecoverPoint,omitempty"`
	PairIds             [][]byte `json:"PairIds,omitempty" xml:"PairIds,omitempty" type:"Repeated"`
	PairNumber          *int64   `json:"PairNumber,omitempty" xml:"PairNumber,omitempty"`
	PrimaryRegion       *string  `json:"PrimaryRegion,omitempty" xml:"PrimaryRegion,omitempty"`
	PrimaryZone         *string  `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	RPO                 *int64   `json:"RPO,omitempty" xml:"RPO,omitempty"`
	ReplicaGroupId      *string  `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	Site                *string  `json:"Site,omitempty" xml:"Site,omitempty"`
	SourceRegionId      *string  `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	SourceZoneId        *string  `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	StandbyRegion       *string  `json:"StandbyRegion,omitempty" xml:"StandbyRegion,omitempty"`
	StandbyZone         *string  `json:"StandbyZone,omitempty" xml:"StandbyZone,omitempty"`
	Status              *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponseBodyReplicaGroups) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetBandwidth(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDescription(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDestinationRegionId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.DestinationRegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetDestinationZoneId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.DestinationZoneId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetGroupName(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetLastRecoverPoint(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.LastRecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPairIds(v [][]byte) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PairIds = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPairNumber(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PairNumber = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPrimaryRegion(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PrimaryRegion = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetPrimaryZone(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetRPO(v int64) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.RPO = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetReplicaGroupId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSite(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSourceRegionId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.SourceRegionId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetSourceZoneId(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.SourceZoneId = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStandbyRegion(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.StandbyRegion = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStandbyZone(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.StandbyZone = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponseBodyReplicaGroups) SetStatus(v string) *DescribeDiskReplicaGroupsResponseBodyReplicaGroups {
	s.Status = &v
	return s
}

type DescribeDiskReplicaGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiskReplicaGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiskReplicaGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) SetStatusCode(v int32) *DescribeDiskReplicaGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) SetBody(v *DescribeDiskReplicaGroupsResponseBody) *DescribeDiskReplicaGroupsResponse {
	s.Body = v
	return s
}

type DescribeDiskReplicaPairsRequest struct {
	MaxResults     *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PairIds        *string `json:"PairIds,omitempty" xml:"PairIds,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	Site           *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s DescribeDiskReplicaPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsRequest) SetMaxResults(v int64) *DescribeDiskReplicaPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetNextToken(v string) *DescribeDiskReplicaPairsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPageNumber(v int32) *DescribeDiskReplicaPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPageSize(v int32) *DescribeDiskReplicaPairsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPairIds(v string) *DescribeDiskReplicaPairsRequest {
	s.PairIds = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetRegionId(v string) *DescribeDiskReplicaPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetReplicaGroupId(v string) *DescribeDiskReplicaPairsRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetSite(v string) *DescribeDiskReplicaPairsRequest {
	s.Site = &v
	return s
}

type DescribeDiskReplicaPairsResponseBody struct {
	NextToken    *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNumber   *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ReplicaPairs []*DescribeDiskReplicaPairsResponseBodyReplicaPairs `json:"ReplicaPairs,omitempty" xml:"ReplicaPairs,omitempty" type:"Repeated"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBody) SetNextToken(v string) *DescribeDiskReplicaPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetPageNumber(v int32) *DescribeDiskReplicaPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetPageSize(v int32) *DescribeDiskReplicaPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetReplicaPairs(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairs) *DescribeDiskReplicaPairsResponseBody {
	s.ReplicaPairs = v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetRequestId(v string) *DescribeDiskReplicaPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetTotalCount(v int64) *DescribeDiskReplicaPairsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDiskReplicaPairsResponseBodyReplicaPairs struct {
	Bandwidth         *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ChargeType        *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	ExpiredTime       *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	LastRecoverPoint  *int64  `json:"LastRecoverPoint,omitempty" xml:"LastRecoverPoint,omitempty"`
	PairName          *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	PrimaryRegion     *string `json:"PrimaryRegion,omitempty" xml:"PrimaryRegion,omitempty"`
	PrimaryZone       *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	RPO               *int64  `json:"RPO,omitempty" xml:"RPO,omitempty"`
	ReplicaGroupId    *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	ReplicaGroupName  *string `json:"ReplicaGroupName,omitempty" xml:"ReplicaGroupName,omitempty"`
	ReplicaPairId     *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	Site              *string `json:"Site,omitempty" xml:"Site,omitempty"`
	SourceDiskId      *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	SourceRegion      *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	SourceZoneId      *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	StandbyRegion     *string `json:"StandbyRegion,omitempty" xml:"StandbyRegion,omitempty"`
	StandbyZone       *string `json:"StandbyZone,omitempty" xml:"StandbyZone,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage     *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetBandwidth(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetChargeType(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ChargeType = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetCreateTime(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDescription(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationZoneId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationZoneId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetExpiredTime(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetLastRecoverPoint(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.LastRecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPairName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PairName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPrimaryRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PrimaryRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPrimaryZone(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetRPO(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.RPO = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaGroupId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaGroupId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaGroupName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaGroupName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaPairId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaPairId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSite(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Site = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceZoneId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceZoneId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStandbyRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StandbyRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStandbyZone(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StandbyZone = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatus(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Status = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatusMessage(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StatusMessage = &v
	return s
}

type DescribeDiskReplicaPairsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiskReplicaPairsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiskReplicaPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetStatusCode(v int32) *DescribeDiskReplicaPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetBody(v *DescribeDiskReplicaPairsResponseBody) *DescribeDiskReplicaPairsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceType(v string) *DescribeRegionsRequest {
	s.ResourceType = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string                                    `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string                                    `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones          []*DescribeRegionsResponseBodyRegionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*DescribeRegionsResponseBodyRegionsZones) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsZones struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type FailoverDiskReplicaGroupRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s FailoverDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupRequest) SetClientToken(v string) *FailoverDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) SetRegionId(v string) *FailoverDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) SetReplicaGroupId(v string) *FailoverDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type FailoverDiskReplicaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupResponseBody) SetRequestId(v string) *FailoverDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type FailoverDiskReplicaGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FailoverDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FailoverDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) SetStatusCode(v int32) *FailoverDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDiskReplicaGroupResponse) SetBody(v *FailoverDiskReplicaGroupResponseBody) *FailoverDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type FailoverDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s FailoverDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairRequest) SetClientToken(v string) *FailoverDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetRegionId(v string) *FailoverDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetReplicaPairId(v string) *FailoverDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type FailoverDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponseBody) SetRequestId(v string) *FailoverDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type FailoverDiskReplicaPairResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FailoverDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FailoverDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetStatusCode(v int32) *FailoverDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetBody(v *FailoverDiskReplicaPairResponseBody) *FailoverDiskReplicaPairResponse {
	s.Body = v
	return s
}

type ModifyDedicatedBlockStorageClusterAttributeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbscId      *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	DbscName    *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetClientToken(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDbscId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.DbscId = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDbscName(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.DbscName = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetDescription(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeRequest) SetRegionId(v string) *ModifyDedicatedBlockStorageClusterAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyDedicatedBlockStorageClusterAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedBlockStorageClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDedicatedBlockStorageClusterAttributeResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDedicatedBlockStorageClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetHeaders(v map[string]*string) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetStatusCode(v int32) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponse) SetBody(v *ModifyDedicatedBlockStorageClusterAttributeResponseBody) *ModifyDedicatedBlockStorageClusterAttributeResponse {
	s.Body = v
	return s
}

type ModifyDiskReplicaGroupRequest struct {
	Bandwidth      *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RPO            *int64  `json:"RPO,omitempty" xml:"RPO,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s ModifyDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupRequest) SetBandwidth(v int64) *ModifyDiskReplicaGroupRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetClientToken(v string) *ModifyDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetDescription(v string) *ModifyDiskReplicaGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetGroupName(v string) *ModifyDiskReplicaGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetRPO(v int64) *ModifyDiskReplicaGroupRequest {
	s.RPO = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetRegionId(v string) *ModifyDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetReplicaGroupId(v string) *ModifyDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type ModifyDiskReplicaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupResponseBody) SetRequestId(v string) *ModifyDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskReplicaGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) SetStatusCode(v int32) *ModifyDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskReplicaGroupResponse) SetBody(v *ModifyDiskReplicaGroupResponseBody) *ModifyDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type ModifyDiskReplicaPairRequest struct {
	Bandwidth     *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PairName      *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	RPO           *int64  `json:"RPO,omitempty" xml:"RPO,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s ModifyDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairRequest) SetBandwidth(v int64) *ModifyDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetClientToken(v string) *ModifyDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetDescription(v string) *ModifyDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetPairName(v string) *ModifyDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRPO(v int64) *ModifyDiskReplicaPairRequest {
	s.RPO = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRegionId(v string) *ModifyDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetReplicaPairId(v string) *ModifyDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type ModifyDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponseBody) SetRequestId(v string) *ModifyDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetStatusCode(v int32) *ModifyDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetBody(v *ModifyDiskReplicaPairResponseBody) *ModifyDiskReplicaPairResponse {
	s.Body = v
	return s
}

type RemoveDiskReplicaPairRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	ReplicaPairId  *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s RemoveDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairRequest) SetClientToken(v string) *RemoveDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetRegionId(v string) *RemoveDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetReplicaGroupId(v string) *RemoveDiskReplicaPairRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetReplicaPairId(v string) *RemoveDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type RemoveDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairResponseBody) SetRequestId(v string) *RemoveDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type RemoveDiskReplicaPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairResponse) SetHeaders(v map[string]*string) *RemoveDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *RemoveDiskReplicaPairResponse) SetStatusCode(v int32) *RemoveDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDiskReplicaPairResponse) SetBody(v *RemoveDiskReplicaPairResponseBody) *RemoveDiskReplicaPairResponse {
	s.Body = v
	return s
}

type ReprotectDiskReplicaGroupRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s ReprotectDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupRequest) SetClientToken(v string) *ReprotectDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetRegionId(v string) *ReprotectDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaGroupRequest) SetReplicaGroupId(v string) *ReprotectDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type ReprotectDiskReplicaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupResponseBody) SetRequestId(v string) *ReprotectDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type ReprotectDiskReplicaGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReprotectDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReprotectDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) SetStatusCode(v int32) *ReprotectDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReprotectDiskReplicaGroupResponse) SetBody(v *ReprotectDiskReplicaGroupResponseBody) *ReprotectDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type ReprotectDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s ReprotectDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairRequest) SetClientToken(v string) *ReprotectDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetRegionId(v string) *ReprotectDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetReplicaPairId(v string) *ReprotectDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type ReprotectDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponseBody) SetRequestId(v string) *ReprotectDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type ReprotectDiskReplicaPairResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReprotectDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReprotectDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetStatusCode(v int32) *ReprotectDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetBody(v *ReprotectDiskReplicaPairResponseBody) *ReprotectDiskReplicaPairResponse {
	s.Body = v
	return s
}

type StartDiskReplicaGroupRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OneShot        *bool   `json:"OneShot,omitempty" xml:"OneShot,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s StartDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupRequest) SetClientToken(v string) *StartDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetOneShot(v bool) *StartDiskReplicaGroupRequest {
	s.OneShot = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetRegionId(v string) *StartDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetReplicaGroupId(v string) *StartDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type StartDiskReplicaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupResponseBody) SetRequestId(v string) *StartDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type StartDiskReplicaGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *StartDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaGroupResponse) SetStatusCode(v int32) *StartDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiskReplicaGroupResponse) SetBody(v *StartDiskReplicaGroupResponseBody) *StartDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type StartDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OneShot       *bool   `json:"OneShot,omitempty" xml:"OneShot,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s StartDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairRequest) SetClientToken(v string) *StartDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetOneShot(v bool) *StartDiskReplicaPairRequest {
	s.OneShot = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetRegionId(v string) *StartDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetReplicaPairId(v string) *StartDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type StartDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponseBody) SetRequestId(v string) *StartDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type StartDiskReplicaPairResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StartDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaPairResponse) SetStatusCode(v int32) *StartDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDiskReplicaPairResponse) SetBody(v *StartDiskReplicaPairResponseBody) *StartDiskReplicaPairResponse {
	s.Body = v
	return s
}

type StopDiskReplicaGroupRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s StopDiskReplicaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupRequest) SetClientToken(v string) *StopDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) SetRegionId(v string) *StopDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) SetReplicaGroupId(v string) *StopDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

type StopDiskReplicaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupResponseBody) SetRequestId(v string) *StopDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

type StopDiskReplicaGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDiskReplicaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDiskReplicaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaGroupResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupResponse) SetHeaders(v map[string]*string) *StopDiskReplicaGroupResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaGroupResponse) SetStatusCode(v int32) *StopDiskReplicaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDiskReplicaGroupResponse) SetBody(v *StopDiskReplicaGroupResponseBody) *StopDiskReplicaGroupResponse {
	s.Body = v
	return s
}

type StopDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s StopDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairRequest) SetClientToken(v string) *StopDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDiskReplicaPairRequest) SetRegionId(v string) *StopDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StopDiskReplicaPairRequest) SetReplicaPairId(v string) *StopDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type StopDiskReplicaPairResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponseBody) SetRequestId(v string) *StopDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type StopDiskReplicaPairResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StopDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaPairResponse) SetStatusCode(v int32) *StopDiskReplicaPairResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDiskReplicaPairResponse) SetBody(v *StopDiskReplicaPairResponseBody) *StopDiskReplicaPairResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ebs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDiskReplicaPairWithOptions(request *AddDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *AddDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDiskReplicaPair(request *AddDiskReplicaPairRequest) (_result *AddDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDiskReplicaPairResponse{}
	_body, _err := client.AddDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDedicatedBlockStorageClusterWithOptions(request *CreateDedicatedBlockStorageClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Azone)) {
		query["Azone"] = request.Azone
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		query["Capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		query["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.DbscName)) {
		query["DbscName"] = request.DbscName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDedicatedBlockStorageCluster"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDedicatedBlockStorageClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDedicatedBlockStorageCluster(request *CreateDedicatedBlockStorageClusterRequest) (_result *CreateDedicatedBlockStorageClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDedicatedBlockStorageClusterResponse{}
	_body, _err := client.CreateDedicatedBlockStorageClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDiskReplicaGroupWithOptions(request *CreateDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationRegionId)) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationZoneId)) {
		query["DestinationZoneId"] = request.DestinationZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceZoneId)) {
		query["SourceZoneId"] = request.SourceZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiskReplicaGroup(request *CreateDiskReplicaGroupRequest) (_result *CreateDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiskReplicaGroupResponse{}
	_body, _err := client.CreateDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDiskReplicaPairWithOptions(request *CreateDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *CreateDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationDiskId)) {
		query["DestinationDiskId"] = request.DestinationDiskId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationRegionId)) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationZoneId)) {
		query["DestinationZoneId"] = request.DestinationZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.PairName)) {
		query["PairName"] = request.PairName
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceZoneId)) {
		query["SourceZoneId"] = request.SourceZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiskReplicaPair(request *CreateDiskReplicaPairRequest) (_result *CreateDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CreateDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDiskReplicaGroupWithOptions(request *DeleteDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDiskReplicaGroup(request *DeleteDiskReplicaGroupRequest) (_result *DeleteDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiskReplicaGroupResponse{}
	_body, _err := client.DeleteDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDiskReplicaPairWithOptions(request *DeleteDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *DeleteDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDiskReplicaPair(request *DeleteDiskReplicaPairRequest) (_result *DeleteDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.DeleteDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedBlockStorageClusterDisksWithOptions(request *DescribeDedicatedBlockStorageClusterDisksRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		query["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedBlockStorageClusterDisks"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedBlockStorageClusterDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedBlockStorageClusterDisks(request *DescribeDedicatedBlockStorageClusterDisksRequest) (_result *DescribeDedicatedBlockStorageClusterDisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedBlockStorageClusterDisksResponse{}
	_body, _err := client.DescribeDedicatedBlockStorageClusterDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDedicatedBlockStorageClustersWithOptions(request *DescribeDedicatedBlockStorageClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AzoneId)) {
		body["AzoneId"] = request.AzoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedicatedBlockStorageClusterId)) {
		body["DedicatedBlockStorageClusterId"] = request.DedicatedBlockStorageClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDedicatedBlockStorageClusters"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDedicatedBlockStorageClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDedicatedBlockStorageClusters(request *DescribeDedicatedBlockStorageClustersRequest) (_result *DescribeDedicatedBlockStorageClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDedicatedBlockStorageClustersResponse{}
	_body, _err := client.DescribeDedicatedBlockStorageClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiskReplicaGroupsWithOptions(request *DescribeDiskReplicaGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Site)) {
		query["Site"] = request.Site
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskReplicaGroups"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskReplicaGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiskReplicaGroups(request *DescribeDiskReplicaGroupsRequest) (_result *DescribeDiskReplicaGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskReplicaGroupsResponse{}
	_body, _err := client.DescribeDiskReplicaGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiskReplicaPairsWithOptions(request *DescribeDiskReplicaPairsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PairIds)) {
		query["PairIds"] = request.PairIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Site)) {
		query["Site"] = request.Site
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskReplicaPairs"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiskReplicaPairs(request *DescribeDiskReplicaPairsRequest) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.DescribeDiskReplicaPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FailoverDiskReplicaGroupWithOptions(request *FailoverDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *FailoverDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FailoverDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FailoverDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FailoverDiskReplicaGroup(request *FailoverDiskReplicaGroupRequest) (_result *FailoverDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDiskReplicaGroupResponse{}
	_body, _err := client.FailoverDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FailoverDiskReplicaPairWithOptions(request *FailoverDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *FailoverDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FailoverDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FailoverDiskReplicaPair(request *FailoverDiskReplicaPairRequest) (_result *FailoverDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.FailoverDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDedicatedBlockStorageClusterAttributeWithOptions(request *ModifyDedicatedBlockStorageClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DbscId)) {
		query["DbscId"] = request.DbscId
	}

	if !tea.BoolValue(util.IsUnset(request.DbscName)) {
		query["DbscName"] = request.DbscName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDedicatedBlockStorageClusterAttribute"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDedicatedBlockStorageClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDedicatedBlockStorageClusterAttribute(request *ModifyDedicatedBlockStorageClusterAttributeRequest) (_result *ModifyDedicatedBlockStorageClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDedicatedBlockStorageClusterAttributeResponse{}
	_body, _err := client.ModifyDedicatedBlockStorageClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDiskReplicaGroupWithOptions(request *ModifyDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDiskReplicaGroup(request *ModifyDiskReplicaGroupRequest) (_result *ModifyDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskReplicaGroupResponse{}
	_body, _err := client.ModifyDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDiskReplicaPairWithOptions(request *ModifyDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bandwidth)) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PairName)) {
		query["PairName"] = request.PairName
	}

	if !tea.BoolValue(util.IsUnset(request.RPO)) {
		query["RPO"] = request.RPO
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDiskReplicaPair(request *ModifyDiskReplicaPairRequest) (_result *ModifyDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.ModifyDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDiskReplicaPairWithOptions(request *RemoveDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *RemoveDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDiskReplicaPair(request *RemoveDiskReplicaPairRequest) (_result *RemoveDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDiskReplicaPairResponse{}
	_body, _err := client.RemoveDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReprotectDiskReplicaGroupWithOptions(request *ReprotectDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReprotectDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReprotectDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReprotectDiskReplicaGroup(request *ReprotectDiskReplicaGroupRequest) (_result *ReprotectDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReprotectDiskReplicaGroupResponse{}
	_body, _err := client.ReprotectDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReprotectDiskReplicaPairWithOptions(request *ReprotectDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReprotectDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReprotectDiskReplicaPair(request *ReprotectDiskReplicaPairRequest) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.ReprotectDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDiskReplicaGroupWithOptions(request *StartDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *StartDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OneShot)) {
		query["OneShot"] = request.OneShot
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDiskReplicaGroup(request *StartDiskReplicaGroupRequest) (_result *StartDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDiskReplicaGroupResponse{}
	_body, _err := client.StartDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDiskReplicaPairWithOptions(request *StartDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *StartDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OneShot)) {
		query["OneShot"] = request.OneShot
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDiskReplicaPair(request *StartDiskReplicaPairRequest) (_result *StartDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.StartDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDiskReplicaGroupWithOptions(request *StopDiskReplicaGroupRequest, runtime *util.RuntimeOptions) (_result *StopDiskReplicaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaGroupId)) {
		query["ReplicaGroupId"] = request.ReplicaGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDiskReplicaGroup"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDiskReplicaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDiskReplicaGroup(request *StopDiskReplicaGroupRequest) (_result *StopDiskReplicaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDiskReplicaGroupResponse{}
	_body, _err := client.StopDiskReplicaGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDiskReplicaPairWithOptions(request *StopDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *StopDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicaPairId)) {
		query["ReplicaPairId"] = request.ReplicaPairId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDiskReplicaPair(request *StopDiskReplicaPairRequest) (_result *StopDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.StopDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
