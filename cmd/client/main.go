package main

import (
	"context"
	"os"
	"time"

	"github.com/brunoluiz/grpc-example/simple/generated/api"
	"github.com/brunoluiz/grpc-example/simple/sigint"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := &cli.App{
		Usage: "grpc client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "grpc-address",
				Usage: "127.0.0.1:5000",
				Value: "127.0.0.1:5000",
			},
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name: "users",
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:   "get",
						Usage:  "get user by id",
						Action: getUser,
					},
					&cli.Command{
						Name:   "get-all",
						Usage:  "get all users",
						Action: getUsers,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func getUsers(c *cli.Context) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(c.String("grpc-address"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create GRPC Client
	client := api.NewIdentityClient(conn)

	// Cancels context after 1 second (timeout) and handles SIGINT (CTRL + C)
	ctx := sigint.WithSignalHandler(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Call GRPC Service
	u, err := client.GetUsers(ctx, &api.GetUsersRequest{})
	if err != nil {
		return errors.Wrap(err, "issue on retrieving users")
	}

	// Print results
	for _, user := range u.Users {
		logrus.Infof("User: %s", user)
	}
	return nil
}

func getUser(c *cli.Context) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(c.String("grpc-address"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create GRPC Client
	client := api.NewIdentityClient(conn)

	// Cancels context after 1 second (timeout) and handles SIGINT (CTRL + C)
	ctx := sigint.WithSignalHandler(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Call GRPC Service
	u, err := client.GetUser(ctx, &api.GetUserRequest{
		UserId: c.Args().Get(0),
	})
	if err != nil {
		return errors.Wrap(err, "issue on retrieving users")
	}

	logrus.Infof("User: %s", u.User)
	return nil
}
