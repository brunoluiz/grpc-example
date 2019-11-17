package main

import (
	"context"
	"net/http"
	"os"

	"github.com/brunoluiz/grpc-example/simple/generated/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "grpc server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "gateway-address",
				Usage: ":8080",
				Value: ":8080",
			},
			&cli.StringFlag{
				Name:  "grpc-address",
				Usage: ":5000",
				Value: ":5000",
			},
		},
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			// Register gRPC server endpoint
			// Note: Make sure the gRPC server is running properly and accessible
			mux := runtime.NewServeMux()
			opts := []grpc.DialOption{grpc.WithInsecure()}
			err := api.RegisterIdentityHandlerFromEndpoint(ctx, mux, c.String("grpc-address"), opts)
			if err != nil {
				return err
			}

			// Start HTTP reverse proxy: sends calls to GRPC server
			return http.ListenAndServe(c.String("gateway-address"), mux)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
