package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "s3-size",
		Usage: "cli tool to get the size of the s3 bucket and its container",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "endpoint",
				Usage:    "s3 endpoint",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "access-key",
				Usage:    "s3 access key",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "secret-key",
				Usage:    "s3 secret key",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "bucket",
				Usage:    "root bucket",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "use-ssl",
				Usage: "use ssl for connecting to s3",
				Value: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			err := showSize(s3Config{
				Endpoint:  ctx.String("endpoint"),
				AccessKey: ctx.String("access-key"),
				SecretKey: ctx.String("secret-key"),
				UseSSL:    ctx.Bool("use-ssl"),
				Bucket:    ctx.String("bucket"),
			})
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
