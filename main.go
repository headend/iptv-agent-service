package main

import (
	"log"
	"github.com/headend/iptv-agent-service/agent-services"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Agent control service")
	agent_services.StartServer()
}
