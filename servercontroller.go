package mscl

import "errors"

// ServerController manages and orchestrates all servers for a single MSCT instance
type ServerController struct {
	Servers []Server
	Config  *Config
}

// New initializes a nil ServerController to default values
func (sc *ServerController) New(cfg *Config) error {
	if sc != nil {
		return errors.New("non-nil struct object")
	}

	sc = &ServerController{
		Servers: []Server{},
		Config:  cfg,
	}
	return nil
}
