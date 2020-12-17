package agent_services

import (
	"context"
	agentpb "github.com/headend/iptv-agent-service/proto"
)


func (c *agentServer) Gets(ctx context.Context, in *agentpb.AgentRequest) (*agentpb.AgentResponse, error) {
	return nil, nil
}

func (c *agentServer) Get(ctx context.Context, in *agentpb.AgentFilter) (*agentpb.AgentResponse, error) {
	return nil, nil
}

func (c *agentServer) Add(ctx context.Context, in *agentpb.AgentRequest) (*agentpb.AgentResponse, error) {
	return nil, nil
}

func (c *agentServer) Update(ctx context.Context, in *agentpb.AgentRequest) (*agentpb.AgentResponse, error) {
	return nil, nil
}

func (c *agentServer) Delete(ctx context.Context, in *agentpb.AgentFilter) (*agentpb.AgentResponse, error) {
	return nil, nil
}
