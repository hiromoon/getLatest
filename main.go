package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/hiromoon/getLatest/checkers"
	"github.com/hiromoon/getLatest/github"
	"github.com/hiromoon/getLatest/parsers"
)

func main() {
	app := cli.NewApp()
	app.Name = "glv"
	app.Usage = "get the latest version from github.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "repo",
			Value: "erlang/otp",
			Usage: "Specify Repository.",
		},
	}
	app.Action = func(c *cli.Context) error {
		arg := strings.Split(c.String("repo"), "/")
		owner := arg[0]
		repo := arg[1]

		res := github.Fetch(owner, repo)
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
