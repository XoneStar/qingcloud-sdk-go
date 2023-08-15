package service

import (
	"encoding/json"
	"fmt"
	"github.com/yunify/qingcloud-sdk-go/config"
	"github.com/yunify/qingcloud-sdk-go/request"
	"github.com/yunify/qingcloud-sdk-go/request/data"
	"time"
)

var _ fmt.State
var _ time.Time

type Boss2Service struct {
	Config     *config.Config
	Properties *Boss2ServiceProperties
}

type Boss2ServiceProperties struct {
	// QingCloud Zone ID
	Zone *string `json:"zone" name:"zone"` // Required
}

func (s *QingCloudService) Boss2Api(config *config.Config, zone string) (*Boss2Service, error) {
	properties := &Boss2ServiceProperties{
		Zone: &zone,
	}
	return &Boss2Service{Config: config, Properties: properties}, nil
}

func (b *Boss2Service) DescribeBoss2Bots(i *DescribeBoss2BotsArgs) (*DescribeBoss2BotsOutput, error) {
	if i == nil {
		i = &DescribeBoss2BotsArgs{
			Action: "Boss2DescribeBots",
		}
	}
	if i.Zone == "" && b.Properties.Zone != nil {
		i.Zone = *b.Properties.Zone
	}
	o := &data.Operation{
		Config:        b.Config,
		Properties:    b.Properties,
		APIName:       "Boss2DescribeBots",
		RequestMethod: "POST",
	}

	x := &DescribeBoss2BotsOutput{}
	r, err := request.New(o, i.ToBoss2RequestParams(), x)
	if err != nil {
		return nil, err
	}

	err = r.Send()
	if err != nil {
		return nil, err
	}

	return x, err
}

type DescribeBoss2BotsInput struct {
	Action *string `json:"action" name:"action" location:"action"`
	Params *string `json:"params" name:"params" location:"params"`
}

type DescribeBoss2BotsArgs struct {
	Action               string   `json:"action"`
	Limit                int      `json:"limit,omitempty"`
	Offset               int      `json:"offset,omitempty"`
	Verbose              int      `json:"verbose,omitempty"`
	Reverse              int      `json:"reverse,omitempty"`
	InitialSelectedIndex int      `json:"initialSelectedIndex,omitempty"`
	Status               []string `json:"status,omitempty"`
	RealtimeStatus       []int    `json:"realtime_status,omitempty"`
	Zone                 string   `json:"zone"`
}

func (v *DescribeBoss2BotsInput) Validate() error {
	return nil
}

func (v *DescribeBoss2BotsArgs) ToBoss2RequestParams() *DescribeBoss2BotsInput {
	raw, _ := json.Marshal(v)
	params := string(raw)
	return &DescribeBoss2BotsInput{Action: &v.Action, Params: &params}
}

type (
	DescribeBoss2BotsOutput struct {
		Action     *string  `json:"action"`
		TotalCount *int     `json:"total_count"`
		BotSet     []BotSet `json:"bot_set"`
		RetCode    *int     `json:"ret_code"`
		Message    *string  `json:"message" name:"message"`
	}
	Dvrs struct {
		RouterID    string `json:"router_id"`
		HyperNodeID string `json:"hyper_node_id"`
		Role        int    `json:"role"`
		VxnetID     string `json:"vxnet_id"`
	}
	PlaceGroups struct {
		PlaceGroupID   string `json:"place_group_id"`
		PlaceGroupName string `json:"place_group_name"`
	}
	BotSet struct {
		HostMachine     string             `json:"host_machine"`
		MemoryUsed      float64            `json:"memory_used"`
		Dvrs            []Dvrs             `json:"dvrs"`
		TotalMemory     float64            `json:"total_memory"`
		VcpuUsed        float64            `json:"vcpu_used"`
		ContainerMode   string             `json:"container_mode"`
		PlaceGroups     []PlaceGroups      `json:"place_groups"`
		HyperNodeType   int                `json:"hyper_node_type"`
		Conntrack       int                `json:"conntrack"`
		FreeDisk        float64            `json:"free_disk"`
		MemoryMax       float64            `json:"memory_max"`
		MasterNetwork   string             `json:"master_network"`
		HasBrks         int                `json:"has_brks"`
		FreeVcpu        float64            `json:"free_vcpu"`
		UsedMemory      float64            `json:"used_memory"`
		FreeGpu         map[string]float64 `json:"free_gpu"`
		HyperNodeID     string             `json:"hyper_node_id"`
		MirrorNode      string             `json:"mirror_node"`
		CPUCores        int                `json:"cpu_cores"`
		FreeMemory      float64            `json:"free_memory"`
		VirtualDisk     float64            `json:"virtual_disk"`
		CPUIdle         float64            `json:"cpu_idle"`
		VipNetwork      string             `json:"vip_network"`
		HyperNodeName   string             `json:"hyper_node_name"`
		Provider        string             `json:"provider"`
		ServiceType     string             `json:"service_type"`
		StatusTime      time.Time          `json:"status_time"`
		UsedVcpu        float64            `json:"used_vcpu"`
		CPUModels       int                `json:"cpu_models"`
		Status          string             `json:"status"`
		RealUsedMemory  float64            `json:"real_used_memory"`
		TotalVcpu       float64            `json:"total_vcpu"`
		FlashNode       int                `json:"flash_node"`
		VcpuMax         float64            `json:"vcpu_max"`
		UsedDisk        float64            `json:"used_disk"`
		Running         int                `json:"running"`
		Remarks         string             `json:"remarks"`
		Kernel          []int              `json:"kernel"`
		KernelVerbose   string             `json:"kernel_verbose"`
		Box             string             `json:"box"`
		StorageMax      int                `json:"storage_max"`
		TotalGpu        map[string]float64 `json:"total_gpu"`
		ZoneID          string             `json:"zone_id"`
		RegionID        string             `json:"region_id"`
		RealFreeMemory  int                `json:"real_free_memory"`
		UsedGpu         map[string]float64 `json:"used_gpu"`
		RealTotalMemory float64            `json:"real_total_memory"`
		CPUWa           float64            `json:"cpu_wa"`
		CtnUtil         float64            `json:"ctn_util"`
		PlaceGroupIds   []string           `json:"place_group_ids"`
		CPUSi           float64            `json:"cpu_si"`
		Rack            string             `json:"rack"`
		BotID           string             `json:"bot_id"`
		TotalDisk       float64            `json:"total_disk"`
		IP              string             `json:"ip"`
		RealtimeStatus  int                `json:"realtime_status"`
	}
)
