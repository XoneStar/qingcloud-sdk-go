package service

import (
	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/request"
	"github.com/yunify/qingcloud-sdk-go/request/data"
	"time"
)

type RouteService struct {
	Config     *config.Config
	Properties *RouteServiceProperties
}

type RouteServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Route(zone string) (*RouteService, error) {
	properties := &RouteServiceProperties{
		Zone: &zone,
	}
	return &RouteService{Config: s.Config, Properties: properties}, nil
}

type DescribeRoutesInput struct {
	// 路由表规则的下一跳类型, 可选值：2,3,4
	//2 - 路由器，nexthop取值应该是路由器的ID
	//3 - IP地址，nexthop取值应该是IP地址
	//4 - NAT网关的ID
	NexthopType *string `json:"nexthop_type" name:"nexthop_type" location:"params"`
}

func (d *DescribeRoutesInput) Validate() error {
	return nil
}

type DescribeRoutesOutput struct {
	Action              *string               `json:"action"`
	TotalCount          *int                  `json:"total_count"`
	RoutingTableRuleSet []RoutingTableRuleSet `json:"routing_table_rule_set"`
	RetCode             *int                  `json:"ret_code"`
	Message             *string               `json:"message"`
}

type RoutingTableRuleSet struct {
	RtableRuleID   string     `json:"rtable_rule_id"`
	Network        string     `json:"network"`
	Nexthop        string     `json:"nexthop"`
	Disabled       int        `json:"disabled"`
	RootUserID     string     `json:"root_user_id"`
	CreateTime     *time.Time `json:"create_time"`
	RtableRuleName string     `json:"rtable_rule_name"`
	Owner          string     `json:"owner"`
	StatusTime     *time.Time `json:"status_time"`
	NexthopType    int        `json:"nexthop_type"`
	RtableID       string     `json:"rtable_id"`
}

func (e *RouteService) DescribeRoutes(i *DescribeRoutesInput) (*DescribeRoutesOutput, error) {
	if i == nil {
		i = &DescribeRoutesInput{}
	}
	o := &data.Operation{
		Config:        e.Config,
		Properties:    e.Properties,
		APIName:       "DescribeRoutes",
		RequestMethod: "GET",
	}

	x := &DescribeRoutesOutput{}
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
