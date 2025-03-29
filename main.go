package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "s3-size",
		Usage: "cli tool to get the size of the s3 bucket and its objects (or directory)",
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
				Usage: "use HTTPS connection when communicating with S3",
				Value: false,
			},
			&cli.IntFlag{
				Name:  "max-depth",
				Usage: "print the size for a directory (or file, with --all) only if it is N or fewer levels below the command line argument",
			},
		},
		Action: func(ctx *cli.Context) error {
			err := showSize(s3Config{
				endPoint:  ctx.String("endpoint"),
				accessKey: ctx.String("access-key"),
				secretKey: ctx.String("secret-key"),
				useSSL:    ctx.Bool("use-ssl"),
				bucket:    ctx.String("bucket"),
				maxDepth:  ctx.Int("max-depth"),
			})
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
