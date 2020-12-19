package model

import "time"

func (Agent) TableName() string {
	return "agent"
}

type Agent struct {
	Agent_id			int64 `gorm:"column:agent_id;AUTO_INCREMENT;primary_key"`
	Ip_Control			string `gorm:"column:ip_control;type:char(15);unique_index:ipcontrol;not null"`
	Ip_Receive_Multicast	string `gorm:"column:ip_receive_multicast;type:char(15)"`
	Cpu					int64 `gorm:"column:cpu;type:number;default:null"`
	Ram					int64 `gorm:"column:cpu;type:number;default:null"`
	Disk				int64 `gorm:"column:disk;type:number;default:null"`
	Location			string `gorm:"column:location;type:varchar;size:60;default:null"`
	IsMonitor			bool `gorm:"column:is_monitor;default:false"`
	IsAlarm				bool `gorm:"column:is_monitor;default:false"`
	Signal_Monitor		bool `gorm:"column:is_monitor;default:true"`
	Video_Monitor		bool `gorm:"column:is_monitor;default:false"`
	Audio_Monitor		bool `gorm:"column:is_monitor;default:false"`
	Run_Thread			int64 `gorm:"column:run_thread;type:number;default:1"`
	Date_create			time.Time `gorm:"column:date_create;type:datetime;default:null"`
	Date_update			time.Time `gorm:"column:date_update;type:datetime;default:null"`
	Status				bool `gorm:"column:status;default:false"`
}

