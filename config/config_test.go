package config

import (
	c "github.com/s-shin/gobbs/config"
	"testing"
)

func TestConfig(t *testing.T) {
	if c.Config == nil {
		t.Errorf("Config is nil.")
	}
	if c.Config.DB.Master == "" {
		t.Errorf("Config don't have DB information.")
	}
}
