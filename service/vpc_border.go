package service

import (
	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/request"
	"github.com/yunify/qingcloud-sdk-go/request/data"
	"time"
)

type VpcBorderService struct {
	Config     *config.Config
	Properties *VpcBorderServiceProperties
}

type VpcBorderServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) VpcBorder(zone string) (*VpcBorderService, error) {
	properties := &VpcBorderServiceProperties{
		Zone: &zone,
	}
	return &VpcBorderService{Config: s.Config, Properties: properties}, nil
}

type VpcBorder struct {
	RouterId            string        `json:"router_id"`
	Status              string        `json:"status"`
	BorderName          string        `json:"border_name"`
	ZoneId              string        `json:"zone_id"`
	Tags                []interface{} `json:"tags"`
	VpcBorderId         string        `json:"vpc_border_id"`
	BorderType          int           `json:"border_type"`
	CreateTime          time.Time     `json:"create_time"`
	Owner               string        `json:"owner"`
	StatusTime          time.Time     `json:"status_time"`
	ResourceProjectInfo []interface{} `json:"resource_project_info"`
	Description         string        `json:"description"`
}

type DescribeVpcBorderOutput struct {
	Action       *string     `json:"action" name:"action"`
	TotalCount   *int        `json:"total_count" name:"total_count"`
	VpcBorderSet []VpcBorder `json:"vpc_border_set" name:"vpc_border_set"`
	RetCode      *int        `json:"ret_code" name:"ret_code"`
	Message      *string     `json:"message" name:"message"`
}

type DescribeVpcBorderInput struct {
	// 边界路由器的类型, 1-vpc
	BorderType *string `json:"border_type" name:"border_type" location:"params"`
	Status     *string `json:"status" name:"status" location:"params"`
	Offset     *int    `json:"offset" name:"offset" location:"params"`
	Limit      *int    `json:"limit" name:"limit" location:"params"`
	Verbose    *int    `json:"verbose" name:"verbose" location:"params"`
}

func (d *DescribeVpcBorderInput) Validate() error {
	return nil
}

func (v *VpcBorderService) DescribeVpcBorders(i *DescribeVpcBorderInput) (*DescribeVpcBorderOutput, error) {
	if i == nil {
		i = &DescribeVpcBorderInput{}
	}
	o := &data.Operation{
		Config:        v.Config,
		Properties:    v.Properties,
		APIName:       "DescribeVpcBorders",
		RequestMethod: "GET",
	}

	x := &DescribeVpcBorderOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeBorderVxnetInput struct {
	Border *string `json:"border" name:"border" location:"params"`
	Vxnet  *string `json:"vxnet" name:"vxnet" location:"params"`
	Offset *int    `json:"offset" name:"offset" location:"params"`
	Limit  *int    `json:"limit" name:"limit" location:"params"`
}

func (b *DescribeBorderVxnetInput) Validate() error {
	return nil
}

type BorderVxnet struct {
	VxnetId          string    `json:"vxnet_id"`
	DynIpStart       string    `json:"dyn_ip_start"`
	DynIpv6End       string    `json:"dyn_ipv6_end"`
	ConsoleId        string    `json:"console_id"`
	CreateTime       time.Time `json:"create_time"`
	Owner            string    `json:"owner"`
	DhcpServerIp     string    `json:"dhcp_server_ip"`
	Features         int       `json:"features"`
	ManagerIp        string    `json:"manager_ip"`
	BorderId         string    `json:"border_id"`
	Ipv6Network      string    `json:"ipv6_network"`
	VxnetName        string    `json:"vxnet_name"`
	BorderPrivateIp  string    `json:"border_private_ip"`
	DhcpServerIpv6   string    `json:"dhcp_server_ipv6"`
	RouterId         string    `json:"router_id"`
	IpNetwork        string    `json:"ip_network"`
	DynIpEnd         string    `json:"dyn_ip_end"`
	TransitionStatus string    `json:"transition_status"`
	Controller       string    `json:"controller"`
	DomainServers    string    `json:"domain_servers"`
	BorderZoneId     string    `json:"border_zone_id"`
	DynIpv6Start     string    `json:"dyn_ipv6_start"`
	ManagerIpv6      string    `json:"manager_ipv6"`
	VpcRouterId      string    `json:"vpc_router_id"`
	VxnetZoneId      string    `json:"vxnet_zone_id"`
	RootUserId       string    `json:"root_user_id"`
	Mode             int       `json:"mode"`
}

type DescribeBorderVxnetOutput struct {
	Action         *string       `json:"action"`
	BorderVxnetSet []BorderVxnet `json:"border_vxnet_set"`
	RetCode        *int          `json:"ret_code"`
	TotalCount     *int          `json:"total_count"`
	Message        *string       `json:"message"`
}

func (v *VpcBorderService) DescribeBorderVxnets(i *DescribeBorderVxnetInput) (*DescribeBorderVxnetOutput, error) {
	if i == nil {
		i = &DescribeBorderVxnetInput{}
	}
	o := &data.Operation{
		Config:        v.Config,
		Properties:    v.Properties,
		APIName:       "DescribeBorderVxnets",
		RequestMethod: "GET",
	}

	x := &DescribeBorderVxnetOutput{}
	r, err := request.New(o, i, x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}
