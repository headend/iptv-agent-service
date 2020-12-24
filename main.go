package main

import (
	agent_services "github.com/headend/iptv-agent-service/agent-services"
	"log"
)


func main()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Agent service")
	agent_services.StartServer()
	//agent_services.GetProfile()
}
