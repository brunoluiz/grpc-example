package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func server() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "grpc server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "grpc-address",
				Usage:    "127.0.0.1:5000",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("server address: %s", c.String("grpc-address"))
			return nil
		},
	}
}
