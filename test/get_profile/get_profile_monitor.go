package main

import (
	"context"
	agentpb "github.com/headend/iptv-agent-service/proto"
	"github.com/headend/share-module/configuration"
	static_config "github.com/headend/share-module/configuration/static-config"
	"log"
	"time"
	"github.com/headend/share-module/mygrpc/connection"
)

func main()  {
	var config configuration.Conf
	config.LoadConf()
	var agentRPCConn connection.RpcClient
	agentRPCConn.InitializeClient(config.RPC.Agent.Host, config.RPC.Agent.Port)
	defer agentRPCConn.Client.Close()
	c, cancel := context.WithTimeout(context.Background(), time.Second)
	agentClient := agentpb.NewAgentServiceClient(agentRPCConn.Client)
	defer cancel()
	res, err := (agentClient).GetProfileMonitor(c, &agentpb.ProfileMonitorRequest{
		IpControl:   "10.0.1.11",
		MonitorType: static_config.Signal,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("data %v", res)
}