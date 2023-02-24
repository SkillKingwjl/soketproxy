package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"soketproxy/proxy"
	_ "soketproxy/proxy"
)

func main() {
	app := &cli.App{
		Name: "proxy",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "userName",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "server",
				Aliases:  []string{"s"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"po"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			return proxy.Run(ctx.String("userName"), ctx.String("password"), ctx.String("server"), ctx.Int("port"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	//proxy.Run1()
}
