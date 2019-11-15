package main

import (
	"net"
	"os"

	"github.com/brunoluiz/grpc-example/simple/generated/api"
	"github.com/brunoluiz/grpc-example/simple/service"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "grpc server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "grpc-address",
				Usage: ":5000",
				Value: ":5000",
			},
		},
		Action: func(c *cli.Context) error {
			// Listen to a specific port
			lis, err := net.Listen("tcp", c.String("grpc-address"))
			if err != nil {
				return err
			}

			// Create a new GRPC Server
			s := grpc.NewServer()

			// Register our service implementation against the GRPC service
			api.RegisterIdentityServer(s, service.NewServer())

			// Start serving
			return s.Serve(lis)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
