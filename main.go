package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/hiromoon/getLatest/checkers"
	"github.com/hiromoon/getLatest/github"
	"github.com/hiromoon/getLatest/parsers"
)

func main() {
	app := cli.NewApp()
	app.Name = "glv"
	app.Usage = "get the latest version from github.com"
	app.Action = func(c *cli.Context) error {
		res := github.Fetch("erlang", "otp")
		checker := &checkers.SemanticVersionChecker{
			Parser: &parsers.ErlangVersionParser{},
		}
		tag := res.GetLatestTag(checker)
		b, err := json.Marshal(tag)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
