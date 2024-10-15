package openengine

import "github.com/Capture-411/services-opportunity/openengine/engine"

func (p *openEngine) AddServers(servers engine.ApiServers) OpenEngine {
	p.Servers = servers
	return p
}
