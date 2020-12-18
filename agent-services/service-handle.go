package agent_services

import (
	"context"
	"fmt"
	agentpb "github.com/headend/iptv-agent-service/proto"
	"github.com/headend/iptv-agent-service/model"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/status"
	"log"
	"time"
)


func (c *agentServer) Gets(ctx context.Context, in *agentpb.AgentGetAll) (*agentpb.AgentResponse, error) {
	var agentList []*model.Agent
	err := c.DB.Db.Find(&agentList).Error
	if err != nil {
		return nil, err
	}
	var res []*agentpb.Agent
	for _, tmp := range agentList	{
		agent := ConvertModelToProtoType(tmp)
		res = append(res, &agent)
	}
	result := agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_SUCCESS,
		Agents: res,
	}
	return &result, nil
}

func (c *agentServer) Get(ctx context.Context, in *agentpb.AgentFilter) (*agentpb.AgentResponse, error) {
	var agents []*model.Agent
	mcase := matchFilterCase(in)
	var err error
	switch mcase {
	case 1:
		err = c.DB.Db.Where("id = ?", in.Id).Find(&agents).Error
	case 2:
		err = c.DB.Db.Where("ip_control = ?", in.IpControl).Find(&agents).Error
	case 3:
		err = c.DB.Db.Where("location = ?", in.Location).Find(&agents).Error
	default:
		log.Printf("Exeption on %#v", in)
		return nil,nil
	}

	if err != nil {
		return nil, err
	}
	var res [] *agentpb.Agent

	for _,tmp := range agents {
		agent := ConvertModelToProtoType(tmp)
		res = append(res, &agent)
	}
	result := agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_SUCCESS,
		Agents: res,
	}
	return &result, nil
}

func (c *agentServer) Add(ctx context.Context, in *agentpb.AgentRequest) (*agentpb.AgentResponse, error) {
	log.Println("params in: %v", in)
	var agentModel model.Agent
	notFound := c.DB.Db.Where("ip_control = ?", in.IpControl).First(&agentModel).RecordNotFound()
	if !notFound {
		err := fmt.Errorf("Agent is available! %#v", in)
		log.Println(err)
		return &agentpb.AgentResponse{
			Status: agentpb.AgentResponseStatus_FAIL,
			Agents: nil,
		}, status.Errorf(409, "Confilic %s", err.Error())
	}
	agentModel = model.Agent{
		Ip_Control:           in.IpControl,
		Ip_Receive_Multicast: in.IpReceiveMulticast,
		Cpu:                  in.Cpu,
		Ram:                  in.Ram,
		Disk:                 in.Disk,
		Location:             in.Location,
		IsMonitor:            in.IsMonitor,
		IsAlarm:              in.IsAlarm,
		Signal_Monitor:       in.SignalMonitor,
		Video_Monitor:        in.VideoMonitor,
		Audio_Monitor:        in.AudioMonitor,
		Run_Thread:           in.RunThread,
	}
	var agentProto agentpb.Agent
	err := c.DB.Db.Create(&agentModel).Updates(map[string]interface{}{"date_create": time.Now(), "date_update": time.Now()}).Scan(&agentProto).Error
	if err == nil {
		var res []*agentpb.Agent
		res = append(res, &agentProto)
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_SUCCESS, Agents:res}, nil
	} else {
		log.Print(err)
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Errorf(409, "Confilic %s", err.Error())
	}

}

func (c *agentServer) Update(ctx context.Context, in *agentpb.AgentRequest) (*agentpb.AgentResponse, error) {
	var oldAgentModel model.Agent
	var err error

	if in.Id != 0 {
		err = c.DB.Db.Where("id = ?", in.Id).First(&oldAgentModel).Error
	} else {
		err = c.DB.Db.Where("ip_control = ?", in.IpControl).First(&oldAgentModel).Error
	}
	if err != nil {
		// Return not found if DB response notfound err
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		// Return err
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	// Return not found if can not found agent
	if oldAgentModel == (model.Agent{}) {
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
	}
	var changeList map[string]interface{}
	if in.IpReceiveMulticast != oldAgentModel.Ip_Receive_Multicast {
		changeList["ip_receive_multicast"] = in.IpReceiveMulticast
	}
	if in.Cpu != oldAgentModel.Cpu {
		changeList["cpu"] = in.Cpu
	}
	if in.Ram != oldAgentModel.Ram {
		changeList["ram"] = in.Ram
	}
	if in.Disk != oldAgentModel.Disk {
		changeList["disk"] = in.Disk
	}
	if in.Location != oldAgentModel.Location {
		changeList["location"] = in.Location
	}
	if in.IsMonitor != oldAgentModel.IsMonitor {
		changeList["is_monitor"] = in.IsMonitor
	}
	if in.IsAlarm != oldAgentModel.IsAlarm {
		changeList["is_alarm"] = in.IsAlarm
	}
	if in.SignalMonitor != oldAgentModel.Signal_Monitor {
		changeList["signal_monitor"] = in.SignalMonitor
	}
	if in.VideoMonitor != oldAgentModel.Video_Monitor {
		changeList["video_monitor"] = in.VideoMonitor
	}
	if in.AudioMonitor != oldAgentModel.Audio_Monitor {
		changeList["audio_monitor"] = in.AudioMonitor
	}
	if in.RunThread != oldAgentModel.Run_Thread {
		changeList["run_thread"] = in.RunThread
	}
	if len(changeList) == 0 {
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_SUCCESS}, status.Errorf(200, "Not change")
	}
	changeList["date_update"] = time.Now()
	var newAgentPb agentpb.Agent
	err2 := c.DB.Db.Model(oldAgentModel).Updates(changeList).Scan(&newAgentPb).Error

	if err2 == nil {
		var res []*agentpb.Agent
		res = append(res, &newAgentPb)
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_SUCCESS, Agents:res}, nil
	} else {
		log.Print(err)
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Errorf(409, "Confilict: %v", err)
	}

}

func (c *agentServer) Delete(ctx context.Context, in *agentpb.AgentDelete) (*agentpb.AgentResponse, error) {
	var agentModel model.Agent
	var err error

	if in.Id != 0 {
		err = c.DB.Db.Where("id = ?", in.Id).First(&agentModel).Error
	} else {
		err = c.DB.Db.Where("ip_control = ?", in.IpControl).First(&agentModel).Error
	}
	if err != nil {
		// Return not found if DB response notfound err
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		// Return err
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	// Return not found if can not found agent
	if agentModel == (model.Agent{}) {
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
	}
	err2 := c.DB.Db.Delete(&agentModel).GetErrors()
	var res []* agentpb.Agent
	tmp := ConvertModelToProtoType(&agentModel)
	res = append(res, &tmp)
	if len(err2) == 0 {
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_SUCCESS, Agents:res}, nil
	} else {
		log.Print(err2)
		return &agentpb.AgentResponse{Status:agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
}

func ConvertModelToProtoType(tmp *model.Agent) agentpb.Agent {
	agent := agentpb.Agent{
		Id:                 tmp.Agent_id,
		IpControl:          tmp.Ip_Control,
		IpReceiveMulticast: tmp.Ip_Receive_Multicast,
		Cpu:                tmp.Cpu,
		Ram:                tmp.Ram,
		Disk:               tmp.Disk,
		Location:           tmp.Location,
		IsMonitor:          tmp.IsMonitor,
		IsAlarm:            tmp.IsAlarm,
		SignalMonitor:      tmp.Signal_Monitor,
		VideoMonitor:       tmp.Video_Monitor,
		AudioMonitor:       tmp.Audio_Monitor,
		RunThread:          tmp.Run_Thread,
		DateCreate:         tmp.Date_create.String(),
	}
	return agent
}

func matchFilterCase(in *agentpb.AgentFilter) (uint8) {

	if in.Id != 0 {
		return 1
	}
	if in.IpControl != "" {
		return 2
	}
	if in.Location != "" {
		return 3
	}
	return 0
}
