package model

import "time"

func (Agent) TableName() string {
	return "agent"
}

type Agent struct {
	Agent_id			int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Ip_Control			string `gorm:"column:ip_control;type:char(15);unique_index:ipcontrol;not null"`
	Ip_Receive_Multicast	string `gorm:"column:ip_receive_multicast;type:char(15)"`
	Cpu					int64 `gorm:"column:cpu;type:number;default:null"`
	Ram					int64 `gorm:"column:cpu;type:number;default:null"`
	Disk				int64 `gorm:"column:disk;type:number;default:null"`
	Location			string `gorm:"column:location;type:varchar;size:60;default:null"`
	IsMonitor			bool `gorm:"column:is_monitor;default:false"`
	IsAlarm				bool `gorm:"column:is_alarm;default:false"`
	Signal_Monitor		bool `gorm:"column:signal_monitor;default:true"`
	Video_Monitor		bool `gorm:"column:video_monitor;default:false"`
	Audio_Monitor		bool `gorm:"column:audio_monitor;default:false"`
	Run_Thread			int64 `gorm:"column:run_thread;type:number;default:1"`
	Date_create			time.Time `gorm:"column:date_create;type:datetime;default:null"`
	Date_update			time.Time `gorm:"column:date_update;type:datetime;default:null"`
	Status				bool `gorm:"column:status;default:false"`
}


func (Monitor) TableName() string {
	return "monitor"
}

type Monitor struct {
	Id  			int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Status_signal 	bool `gorm:"column:status_signal;default:false"`
	Status_video  	bool `gorm:"column:status_video ;default:false"`
	Status_audio  	bool `gorm:"column:status_audio ;default:false"`
	Signal_monitor  bool  `gorm:"column:signal_monitor ;default:false"`
	Video_monitor 	bool  `gorm:"column:video_monitor;default:false"`
	Audio_monitor 	bool  `gorm:"column:audio_monitor;default:false"`
	Is_enable 		bool  `gorm:"column:is_enable;default:false"`
	Agent_id  		int64 `gorm:"column:agent_id ;default:not null"`
	Profile_id  	int64 `gorm:"column:profile_id ;default:not null"`
	Status_id 		int64 `gorm:"column:status_id;default:null"`
	Date_update 	time.Time `gorm:"column:date_update;default:null"`
	Profile 		Profile `gorm:"foreignkey:id;references:profile_id"`
	Agent 			Agent `gorm:"foreignkey:id;references:agent_id"`
}


func (Profile) TableName() string {
	return "profile"
}

type Profile struct {
	Id  				int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Desc  				string `gorm:"column:desc;type:varchar(60);default:null"`
	Date_create			time.Time `gorm:"column:date_create;type:datetime;default:null"`
	Date_update			time.Time `gorm:"column:date_update;type:datetime;default:null"`
	Is_original 		bool  `gorm:"column:is_original;default:false"`
	Channel_id 			int64 `gorm:"column:channel_id;default:null"`
	Encoder_id  		int64 `gorm:"column:encoder_id ;default:null"`
	Multicast_ip_id  	int64 `gorm:"column:multicast_ip_id ;default:null"`
	Profile_quality_id  int64 `gorm:"column:profile_quality_id;default:not null"`
	Vlan_id  			int64 `gorm:"column:vlan_id ;default:not null"`
	Status_id 			int64 `gorm:"column:status_id;default:null"`
	Multicast_id 		MulticastIp `gorm:"foreignkey:id;references:multicast_ip_id"`
}


func (MulticastIp) TableName() string {
	return "multicast_ip"
}

type MulticastIp struct {
	Id  	int64 `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Ip 		string `gorm:"column:ip;type:char(15);default:null;"`
	Desc  	string `gorm:"column:desc;type:varchar(60);default:null"`
}


