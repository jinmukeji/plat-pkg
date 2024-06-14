package service

import "go-micro.dev/v4/server"

type RegisterServerFunc func(srv server.Server) error

func RegisterServer(srv server.Server, registers ...RegisterServerFunc) error {
	err := srv.Init(
		server.Wait(nil),
	)
	if err != nil {
		return err
	}

	for _, r := range registers {
		err = r(srv)
		if err != nil {
			return err
		}
	}

	return nil
}
