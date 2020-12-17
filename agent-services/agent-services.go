package agent_services

import (
	"fmt"
	"github.com/headend/share-module/configuration"
	agentctlpb "github.com/headend/iptv-agent-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type agentServer struct {
	config *configuration.Conf
}


func StartServer()  {
	var config configuration.Conf

	config.LoadConf()
	ln, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	agentctlpb.RegisterAgentServiceServer(rpcServer, &agentServer(config:&config))
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}