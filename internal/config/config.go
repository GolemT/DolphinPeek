package config

import "time"

const (
	DefaultCycleInterval = 10 * time.Second
	WindowWidth          = 400
	WindowHeight         = 300
)

func GetCycleInterval() time.Duration {
	// Later you can load this from a config file
	return DefaultCycleInterval
}
