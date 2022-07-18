package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	type Config struct {
		Version string `config:"version"`
		Server  struct {
			Ports   []int64 `config:"ports"`
			Cleanup string  `config:"cleanup"`
		} `config:"server"`
		Logger struct {
			Level string `config:"level"`
			Trace bool   `config:"trace"`
		} `config:"logger"`
		Rules []struct {
			Host    string `config:"host"`
			Backend struct {
				ServiceName string `config:"service_name"`
				ServicePort int64  `config:"service_port"`
			} `config:"backend"`
		} `config:"rules"`
		Struc struct {
			Field1 string `config:"field1"`
			Field2 string `config:"field2"`
		} `config:"struct"`
	}
	var cfg Config
	if err := Load(&cfg); err != nil {
		t.Fatal(err)
	}

	// j, _ := json.MarshalIndent(cfg, "", "  ")
	// fmt.Println(string(j))
	if cfg.Version != "1.0.0" {
		t.Fatal("version is not 1.0.0")
	}

	if len(cfg.Server.Ports) != 2 {
		t.Fatal("ports is not 2")
	}

	if cfg.Server.Ports[0] != 8080 {
		t.Fatal("ports[0] is not 8080")
	}

	if cfg.Server.Ports[1] != 9090 {
		t.Fatal("ports[1] is not 9090")
	}

	if cfg.Server.Cleanup != "1h" {
		t.Fatal("cleanup is not 1h")
	}

	if cfg.Logger.Level != "warn" {
		t.Fatal("level is not warn")
	}

	if cfg.Logger.Trace != false {
		t.Fatal("trace is not false")
	}

	if len(cfg.Rules) != 2 {
		t.Fatal("rules is not 2")
	}

	if cfg.Rules[0].Host != "backend-a.example.com" {
		t.Fatal("rules[0].host is not backend-a.example.com")
	}

	if cfg.Rules[0].Backend.ServiceName != "backend-a" {
		t.Fatal("rules[0].backend.service_name is not backend-a")
	}

	if cfg.Rules[0].Backend.ServicePort != 80 {
		t.Fatal("rules[0].backend.service_port is not 80")
	}

	if cfg.Rules[1].Host != "backend-b.example.com" {
		t.Fatal("rules[1].host is not backend-b.example.com")
	}

	if cfg.Rules[1].Backend.ServiceName != "backend-b" {
		t.Fatal("rules[1].backend.service_name is not backend-b")
	}

	if cfg.Rules[1].Backend.ServicePort != 8080 {
		t.Fatal("rules[1].backend.service_port is not 8080")
	}

}
