package agent_services

import (
	"context"
	"fmt"
	"github.com/headend/iptv-agent-service/model"
	agentpb "github.com/headend/iptv-agent-service/proto"
	static_config "github.com/headend/share-module/configuration/static-config"
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
	filterSupport := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
		Location:       in.Location,
	}
	agents, err := fillterAgent(filterSupport, c)

	if err != nil {
		return nil, err
	}
	var res []*agentpb.Agent

	for _, tmp := range agents {
		agent := ConvertModelToProtoType(tmp)
		res = append(res, &agent)
	}
	result := agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_SUCCESS,
		Agents: res,
	}
	return &result, nil
}

func fillterAgent(filterSupport IdentifyAgent, c *agentServer) ([]*model.Agent, error) {
	var agents []*model.Agent
	mcase := matchFilterCase(filterSupport)
	var err error
	switch mcase {
	case 1:
		err = c.DB.Db.Where("id = ?", filterSupport.AgentID).Find(&agents).Error
	case 2:
		err = c.DB.Db.Where("ip_control = ?", filterSupport.AgentControlIP).Find(&agents).Error
	case 3:
		err = c.DB.Db.Where("location = ?", filterSupport.Location).Find(&agents).Error
	default:
		err = fmt.Errorf("Filter not support")
	}
	return agents, err
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
		Status: 			  in.Status,
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
	if in.Status != oldAgentModel.Status {
		changeList["status"] = in.Status
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
	filterSupport := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
	}
	var agentModel model.Agent
	err, agentModel := CheckAgentExists(filterSupport, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}

	err2 := c.DB.Db.Delete(&agentModel).GetErrors()
	var res []*agentpb.Agent
	tmp := ConvertModelToProtoType(&agentModel)
	res = append(res, &tmp)
	if len(err2) == 0 {
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_SUCCESS, Agents: res}, nil
	} else {
		log.Print(err2)
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
}

type IdentifyAgent struct {
	AgentID int64
	AgentControlIP string
	Location string
}

func CheckAgentExists(in IdentifyAgent, c *agentServer) (err error, agent model.Agent) {
	filterSupport := IdentifyAgent{
		AgentID:        in.AgentID,
		AgentControlIP: in.AgentControlIP,
		Location:       in.Location,
	}
	var result model.Agent
	agents, err2 := fillterAgent(filterSupport, c)
	if err2 != nil {
		return err2, result
		// Return not found if DB response notfound err
		//if gorm.IsRecordNotFoundError(err2) {
			//return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found"), true

		//}
		// Return err
		//return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error"), true
	}
	// Return not found if can not found agent
	if len(agents) == 0 {
		//return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found"), true
		return fmt.Errorf("Not found"), result
	}
	result = *agents[0]
	return nil, result
}


//*
// Update Agent status
func (c *agentServer) UpdateStatus(ctx context.Context, in *agentpb.AgentUpdateStatus) (*agentpb.AgentResponse, error) {
	var agentModel model.Agent
	indentifyAgent := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
		Location:       "",
	}
	err, agentModel := CheckAgentExists(indentifyAgent, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	if in.Status != agentModel.Status {
		log.Println("Update status")
		//log.Printf("%#v", agentModel)
		err := c.DB.Db.Model(agentModel).Updates(map[string]interface{}{"status": in.Status}).Error
		//err := c.DB.Db.Save(&agentModel).Error
		//err := c.DB.Db.Model(&agentModel).Updates(model.Agent{
		//	Status: in.Status,
		//}).Error
		if err != nil {
			log.Println(err)
			return &agentpb.AgentResponse{
				Status: agentpb.AgentResponseStatus_FAIL,
				Agents: nil,
			}, err
		}
		//log.Println("update status Done")
	}
	return &agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_FAIL,
		Agents: nil,
	}, nil
}
//*
// Update RunThread
func (c *agentServer) UpdateRunthread(ctx context.Context, in *agentpb.AgentUpdateMonitorRunThread) (*agentpb.AgentResponse, error) {
	var agentModel model.Agent
	indentifyAgent := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
		Location:       "",
	}
	err, agentModel := CheckAgentExists(indentifyAgent, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	if in.RunThread != agentModel.Run_Thread {
		err := c.DB.Db.Model(&agentModel).Updates(model.Agent{
			Run_Thread: in.RunThread,
		}).Error
		if err != nil {
			return &agentpb.AgentResponse{
				Status: agentpb.AgentResponseStatus_FAIL,
				Agents: nil,
			}, err
		}
	}
	return &agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_FAIL,
		Agents: nil,
	}, nil
}
//*
// Active monitor
func (c *agentServer) UpdateActiveMonitor(ctx context.Context, in *agentpb.AgentActiveMonitor) (*agentpb.AgentResponse, error) {
	log.Println("Update is_monitor")
	var agentModel model.Agent
	indentifyAgent := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
		Location:       "",
	}
	err, agentModel := CheckAgentExists(indentifyAgent, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	if in.IsMonitor != agentModel.IsMonitor {
		log.Println("Do update")
		err := c.DB.Db.Model(&agentModel).Updates(map[string]interface{}{"is_monitor": in.IsMonitor}).Error
		if err != nil {
			return &agentpb.AgentResponse{
				Status: agentpb.AgentResponseStatus_FAIL,
				Agents: nil,
			}, nil
		}
	}
	return &agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_FAIL,
		Agents: nil,
	}, nil
}
//*
// Active monitor
func (c *agentServer) UpdateMonitorSignal(ctx context.Context, in *agentpb.AgentActiveMonitorSignal) (*agentpb.AgentResponse, error) {
	log.Println("Update monitor signal")
	var agentModel model.Agent
	indentifyAgent := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
		Location:       "",
	}
	err, agentModel := CheckAgentExists(indentifyAgent, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	if in.SignalMonitor != agentModel.Signal_Monitor {
		log.Println("Do update ")
		err := c.DB.Db.Model(&agentModel).Updates(map[string]interface{}{"signal_monitor": in.SignalMonitor}).Error
		if err != nil {
			return &agentpb.AgentResponse{
				Status: agentpb.AgentResponseStatus_FAIL,
				Agents: nil,
			}, err
		}
	}
	return &agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_FAIL,
		Agents: nil,
	}, nil
}
//*
// Active monitor video
func (c *agentServer) UpdateMonitorVideo(ctx context.Context, in *agentpb.AgentActiveMonitorVideo) (*agentpb.AgentResponse, error) {
	log.Println("Update monitor video")
	var agentModel model.Agent
	indentifyAgent := IdentifyAgent{
		AgentID:        in.Id,
		AgentControlIP: in.IpControl,
		Location:       "",
	}
	err, agentModel := CheckAgentExists(indentifyAgent, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(404, "Not found")
		}
		return &agentpb.AgentResponse{Status: agentpb.AgentResponseStatus_FAIL}, status.Error(500, "Internal server error")
	}
	if in.VideoMonitor != agentModel.Video_Monitor {
		log.Println("Do update ")
		err := c.DB.Db.Model(&agentModel).Updates(map[string]interface{}{"video_monitor": in.VideoMonitor}).Error
		if err != nil {
			return &agentpb.AgentResponse{
				Status: agentpb.AgentResponseStatus_FAIL,
				Agents: nil,
			}, err
		}
	}
	return &agentpb.AgentResponse{
		Status: agentpb.AgentResponseStatus_FAIL,
		Agents: nil,
	}, nil
}

func (c *agentServer) GetProfileMonitor(ctx context.Context, in *agentpb.ProfileMonitorRequest) (*agentpb.ProfileMonitorResponse, error) {
	var agentModel model.Agent
	indentifyAgent := IdentifyAgent{
		AgentID:        in.AgentId,
		AgentControlIP: in.IpControl,
		Location:       "",
	}
	err, agentModel := CheckAgentExists(indentifyAgent, c)
	if err != nil {
		log.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(404, "Not found")
		}
		if err.Error() == "Not found" {
			return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(404, "Not found")
		}
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(500, "Internal server error")
	}
	// go get progile monitor by agent info

	// Make where clause
	var whereClause string
	switch in.MonitorType {
	case static_config.Signal:
		whereClause = fmt.Sprintf("monitor.is_enable= %d AND monitor.agent_id=%d AND monitor.signal_monitor=%d",1, agentModel.Agent_id, 1)
	case static_config.Video:
		whereClause = fmt.Sprintf("monitor.is_enable= %d AND monitor.agent_id=%d AND monitor.video_monitor=%d",1, agentModel.Agent_id, 1)
	default:
		err = fmt.Errorf("Wait for support monitor type = %d", in.MonitorType)
		log.Print(err)
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(415, err.Error())
	}

	// make query
	rows, err := c.DB.Db.Table("monitor").Select(
		"monitor.id, monitor.agent_id, monitor.status_id, profile.id as profile_id, multicast_ip.ip as ip, monitor.signal_monitor, monitor.video_monitor, monitor.status_video, monitor.is_enable").Joins("" +
		"join profile on profile.id = monitor.profile_id").Joins(
		"join multicast_ip on multicast_ip.id = profile.multicast_ip_id").Where(whereClause).Rows()
	if err != nil {
		log.Println(err)
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(500, "Internal server error")
	}

	var fetchErr error
	var resulfResponse agentpb.ProfileMonitorResponse
	var resulfData []*agentpb.ProfileMonitorElement
	for rows.Next() {
		var (
			Id int64
			AgentID		int64
			StatusId	int64
			ProdileId	int64
			Ip	string
			SignalMonitor bool
			VideoMonitor	bool
			StatusVideo		bool
			IsEnable	bool
		)
		err := rows.Scan(&Id, &AgentID, &StatusId, &ProdileId, &Ip, &SignalMonitor, &VideoMonitor, &StatusVideo, &IsEnable)
		if err != nil {
			log.Println(err)
			break
		}
		dataElement := agentpb.ProfileMonitorElement{
			MonitorId:     Id,
			ProfileId:		ProdileId,
			AgentId:       AgentID,
			StatusId:      StatusId,
			IpControl:     "",
			MulticastIp:   Ip,
			SignalMonitor: SignalMonitor,
			VideoMonitor:  VideoMonitor,
			StatusVideo:   StatusVideo,
			IsEnable:      IsEnable,
		}
		resulfData = append(resulfData, &dataElement)
		log.Printf("monitor: %#v", dataElement)
		//log.Printf("data rows: %d, %d, %d, %d, %s, %b, %b, %b, %b", Id, AgentID, StatusId, ProdileId, Ip, SignalMonitor, VideoMonitor, StatusVideo, IsEnable)
	}
	if fetchErr != nil{
		log.Print(fetchErr)
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(500, "Internal server error")
	}
	resulfResponse.Status = agentpb.AgentResponseStatus_SUCCESS
	resulfResponse.Profiles = resulfData
	resulfResponse.MonitorType = in.MonitorType
	return &resulfResponse, nil
}

func (c *agentServer) UpdateMonitorStatus(ctx context.Context, in *agentpb.MonitorUpdateStatusRequest) (*agentpb.ProfileMonitorResponse, error)  {
	// check monitor id
	var monitors []*model.Monitor
	if in.MonitorId == 0 {
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(400, "Invalid params")
	}
	// get monitor from database
	err := c.DB.Db.Where("id = ?", in.MonitorId).Find(&monitors).Error
	if err != nil {
		log.Println(err)
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(404, "Not found")
	}
	if len(monitors) != 1 {
		errNotfound := fmt.Errorf("Not found")
		log.Println(errNotfound)
		return &agentpb.ProfileMonitorResponse{Status: agentpb.AgentResponseStatus_FAIL, MonitorType: in.MonitorType}, status.Error(404, errNotfound.Error())
	}
	thisMonitor := monitors[0]
	switch in.MonitorType {
	case static_config.Signal:
		// update signal status
		if in.NewStatus != 1 {
			_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status_signal": false}).Error
		} else {
			_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status_signal": true}).Error
		}
	case static_config.Video:
		// update video status
		if in.NewStatus != 1 {
			_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status_video": false}).Error
		} else {
			_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status_video": true}).Error
		}
	case static_config.Audio:
		// update video status
		if in.NewStatus != 1 {
			_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status_audio": false}).Error
		} else {
			_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status_audio": true}).Error
		}
	default:
		log.Printf("Not support monitor type: %d", in.MonitorType)
	}
	// update status
	_ = c.DB.Db.Model(&thisMonitor).Updates(map[string]interface{}{"status": in.NewStatus}).Error
	return &agentpb.ProfileMonitorResponse{
		Status:      agentpb.AgentResponseStatus_SUCCESS,
		Profiles:    nil,
		MonitorType: in.MonitorType,
	}, nil
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
		Status:				tmp.Status,
	}
	return agent
}

func matchFilterCase(in IdentifyAgent) (uint8) {

	if in.AgentID != 0 {
		return 1
	}
	if in.AgentControlIP != "" {
		return 2
	}
	if in.Location != "" {
		return 3
	}
	return 0
}
