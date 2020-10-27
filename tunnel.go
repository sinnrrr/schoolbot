package main

import (
	"github.com/afdalwahyu/gonnel"
	"github.com/sinnrrr/schoolbot/config"
	"os"
)

func InitTunnel() error {
	if os.Getenv("PUBLIC_URL") == "" {
		client, err := gonnel.NewClient(config.TunnelOptions)
		if err != nil {
			return err
		}

		done := make(chan bool)
		go client.StartServer(done)

		client.AddTunnel(&config.Tunnel)
		err = client.ConnectAll()
		if err != nil {
			return err
		}

		err = os.Setenv("PUBLIC_URL", config.Tunnel.RemoteAddress[8:])
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}