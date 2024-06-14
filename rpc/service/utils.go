package service

import "time"

const (
	// defaultRegisterTTL specifies how long a registration should exist in
	// discovery after which it expires and is removed
	defaultRegisterTTL = 30 * time.Second

	// defaultRegisterInterval is the time at which a service should
	// re-register to preserve it's registration in service discovery.
	defaultRegisterInterval = 15 * time.Second
)

func die(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
