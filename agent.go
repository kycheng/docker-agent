package agent

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(Agent{})
}

// Gizmo is an example; put your own type here.
type Agent struct {
}

// CaddyModule returns the Caddy module information.
func (Agent) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "Agent",
		New: func() caddy.Module { return new(Agent) },
	}
}
