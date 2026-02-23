package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Config struct {
	Port int
	Host string
	Mode string
}

type ConfigManager struct {
	value atomic.Value
}

func NewConfigManager(initial Config) *ConfigManager {
	cm := &ConfigManager{}
	cm.value.Store(&initial)
	return cm
}

func (cm *ConfigManager) GetConfig() *Config {
	return cm.value.Load().(*Config)
}

func (cm *ConfigManager) Reload(newConfig Config) {
	cm.value.Store(&newConfig)
}

func main() {
	initial := Config{Port: 8080, Host: "localhost", Mode: "dev"}
	cm := NewConfigManager(initial)

	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				cfg := cm.GetConfig()
				fmt.Printf("Reader %d: %+v\n", id, cfg)
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			newCfg := Config{
				Port: 8080,
				Host: "localhost",
				Mode: "prod",
			}
			cm.Reload(newCfg)
			fmt.Println("Reloaded config to prod")
		}
	}()

	time.Sleep(10 * time.Second)
}
