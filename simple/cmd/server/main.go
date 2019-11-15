package main

import (
	"net"
	"os"

	"github.com/brunoluiz/grpc-example/simple/generated/api"
	"github.com/brunoluiz/grpc-example/simple/service"
	"github.com/sirupsen/logrus"
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
			lis, err := net.Listen("tcp", c.String("grpc-address"))
			if err != nil {
				return err
			}

			logrus.Infof("Listening at %s", c.String("grpc-address"))

			s := grpc.NewServer()
			api.RegisterIdentityServer(s, service.NewServer())
			return s.Serve(lis)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
