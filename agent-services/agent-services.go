package agent_services

import (
	"fmt"
	agentpb "github.com/headend/iptv-agent-service/proto"
	"github.com/headend/share-module/configuration"
	database "github.com/headend/share-module/databases"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type agentServer struct {
	config *configuration.Conf
	DB *database.Database
}


func StartServer()  {
	var config configuration.Conf

	config.LoadConf()
	// Make Database connect
	var db database.Database
	db.Connect(&config)
	if db.Err != nil {
		log.Fatal("Cannot connect DB: " + db.Err.Error())
	}

	listenerAdd := fmt.Sprintf("%s:%d", config.RPC.Agent.Host, config.RPC.Agent.Port)
	ln, err := net.Listen("tcp", listenerAdd)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer(grpc.MaxSendMsgSize(2 * 1024 *1024), grpc.ConnectionTimeout(10 * time.Second), grpc.MaxConcurrentStreams(2048))
	agentpb.RegisterAgentCTLServiceServer(rpcServer, &agentServer{config: &config, DB: &db})
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}