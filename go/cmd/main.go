package main

import (
	"context"
	"log"
	"os"

	"github.com/for-the-kidz/diccionario"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "server",
		Usage: "Start the web server",
		Action: func(context.Context, *cli.Command) error {
			s := diccionario.NewServer()
			return s.Run(":8080")
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
