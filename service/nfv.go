package service

import (
	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/request"
	"github.com/yunify/qingcloud-sdk-go/request/data"
	"time"
)

type NfvService struct {
	Config     *config.Config
	Properties *NfvServiceProperties
}

type NfvServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Nfv(zone string) (*NfvService, error) {
	properties := &NfvServiceProperties{
		Zone: &zone,
	}
	return &NfvService{Config: s.Config, Properties: properties}, nil
}

type DescribeNFVsOutput struct {
	Action     *string  `json:"action"`
	TotalCount *int     `json:"total_count"`
	NfvSet     []NfvSet `json:"nfv_set"`
	RetCode    *int     `json:"ret_code"`
	Message    *string  `json:"message"`
}

type InstanceInfo struct {
	InstanceID          string `json:"instance_id"`
	InstanceTime        string `json:"instance_time"`
	NodeIdx             int    `json:"node_idx"`
	EipNodeHealthStatus int    `json:"eip_node_health_status"`
}

type NfvSet struct {
	Status           string        `json:"status"`
	IsApplied        int           `json:"is_applied"`
	Description      string        `json:"description"`
	NfvName          string        `json:"nfv_name"`
	NfvType          int           `json:"nfv_type"`
	SubCode          int           `json:"sub_code"`
	TransitionStatus string        `json:"transition_status"`
	Cluster          []ClusterInfo `json:"cluster"`
	RootUserID       string        `json:"root_user_id"`
	CreateTime       time.Time     `json:"create_time"`
	NfvID            string        `json:"nfv_id"`
	NfvSpec          int           `json:"nfv_spec"`
	Owner            string        `json:"owner"`
	StatusTime       time.Time     `json:"status_time"`
	NodeCount        int           `json:"node_count"`
	ClusterMode      int           `json:"cluster_mode"`
	ZoneID           string        `json:"zone_id"`
}

type ClusterInfo struct {
	Instances []InstanceInfo `json:"instances"`
	EipId     string         `json:"eip_id"`
}

type DescribeNFVsInput struct {
	Limit   int    `json:"limit" name:"limit"`
	Offset  int    `json:"offset" name:"offset"`
	Verbose int    `json:"verbose" name:"verbose"`
	Zone    string `json:"zone" name:"zone"`
}

func (d *DescribeNFVsInput) Validate() error {
	return nil
}

func (e *NfvService) DescribeNFVs(i *DescribeNFVsInput) (*DescribeNFVsOutput, error) {
	if i == nil {
		i = &DescribeNFVsInput{}
	}
	o := &data.Operation{
		Config:        e.Config,
		Properties:    e.Properties,
		APIName:       "DescribeNFVs",
		RequestMethod: "GET",
	}

	x := &DescribeNFVsOutput{}
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

type (
	DescribeSnatRulesInput struct {
		NatgwIds []string `json:"natgw_ids" name:"natgw_ids"`
	}
	DescribeSnatRulesOutput struct {
		Action     *string   `json:"action"`
		SnatSet    []SnatSet `json:"snat_set"`
		RetCode    *int      `json:"ret_code"`
		TotalCount *int      `json:"total_count"`
		Message    *string   `json:"message"`
	}
	Eips struct {
		EipID   string `json:"eip_id"`
		EipAddr string `json:"eip_addr"`
	}
	SnatSet struct {
		NatgwID    string    `json:"natgw_id"`
		Enable     int       `json:"enable"`
		Name       string    `json:"name"`
		SnatID     string    `json:"snat_id"`
		TargetID   string    `json:"target_id"`
		TargetType int       `json:"target_type"`
		Eips       []Eips    `json:"eips"`
		CreateTime time.Time `json:"create_time"`
		TargetIP   string    `json:"target_ip"`
	}
)

func (d *DescribeSnatRulesInput) Validate() error {
	return nil
}

func (e *NfvService) DescribeSnatRules(i *DescribeSnatRulesInput) (*DescribeSnatRulesOutput, error) {
	if i == nil {
		i = &DescribeSnatRulesInput{}
	}
	o := &data.Operation{
		Config:        e.Config,
		Properties:    e.Properties,
		APIName:       "DescribeSnatRules",
		RequestMethod: "GET",
	}

	x := &DescribeSnatRulesOutput{}
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

type (
	DescribeDnatRulesInput struct {
		NatgwIds []string `json:"natgw_ids" name:"natgw_ids"`
	}
	DescribeDnatRulesOutput struct {
		Action     *string   `json:"action"`
		TotalCount *int      `json:"total_count"`
		DnatSet    []DnatSet `json:"dnat_set"`
		RetCode    *int      `json:"ret_code"`
		Message    *string   `json:"message"`
	}
	DnatSet struct {
		NatgwID     string    `json:"natgw_id"`
		Enable      int       `json:"enable"`
		EipID       string    `json:"eip_id"`
		Name        string    `json:"name"`
		PublicPort  string    `json:"public_port"`
		PrivatePort string    `json:"private_port"`
		CreateTime  time.Time `json:"create_time"`
		PrivateIP   string    `json:"private_ip"`
		Protocol    string    `json:"protocol"`
		DnatID      string    `json:"dnat_id"`
	}
)

func (d DescribeDnatRulesInput) Validate() error {
	return nil
}

func (e *NfvService) DescribeDnatRules(i *DescribeDnatRulesInput) (*DescribeDnatRulesOutput, error) {
	if i == nil {
		i = &DescribeDnatRulesInput{}
	}
	o := &data.Operation{
		Config:        e.Config,
		Properties:    e.Properties,
		APIName:       "DescribeDnatRules",
		RequestMethod: "GET",
	}

	x := &DescribeDnatRulesOutput{}
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
