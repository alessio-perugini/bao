package main

import (
	"bao"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main()  {
	app := &cli.App{
		Name: "bao",
		Usage: "block that motherf***er bots",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	bao.GetIpFromLog()
}