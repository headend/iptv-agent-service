package agent_services

import (
	"fmt"
	"github.com/headend/share-module/configuration"
	database "github.com/headend/share-module/databases"
	"log"
)

func GetProfile()  {
	var config configuration.Conf
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

	config.LoadConf()
	// Make Database connect
	var db database.Database
	db.Connect(&config)
	if db.Err != nil {
		log.Fatal("Cannot connect DB: " + db.Err.Error())
	}
	whereClause := fmt.Sprintf("monitor.agent_id=%d", 1)
	rows, err := db.Db.Table("monitor").Select(
		"monitor.id, monitor.agent_id, monitor.status_id, profile.id as profile_id, multicast_ip.ip as ip, monitor.signal_monitor, monitor.video_monitor, monitor.status_video, monitor.is_enable").Joins("" +
			"join profile on profile.id = monitor.profile_id").Joins(
				"join multicast_ip on multicast_ip.id = profile.multicast_ip_id").Where(whereClause).Rows()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&Id, &AgentID, &StatusId, &ProdileId, &Ip, &SignalMonitor, &VideoMonitor, &StatusVideo, &IsEnable)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("data rows: %d, %d, %d, %d, %s, %b, %b, %b, %b", Id, AgentID, StatusId, ProdileId, Ip, SignalMonitor, VideoMonitor, StatusVideo, IsEnable)
		}
	}
}

