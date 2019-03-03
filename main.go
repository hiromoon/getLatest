package main

import (
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
		r := c.String("repo")
		arg := strings.Split(r, "/")
		if len(arg) < 2 {
			return fmt.Errorf("Invalid repository name: ", r)
		}
		owner := arg[0]
		repo := arg[1]

		res := github.Fetch(owner, repo)
		var checker *checkers.SemanticVersionChecker
		if owner == "erlang" {
			checker = &checkers.SemanticVersionChecker{
				Parser: &parsers.ErlangVersionParser{},
			}
		} else if owner == "elixir-lang" {
			checker = &checkers.SemanticVersionChecker{
				Parser: &parsers.ElixirVersionParser{},
			}
		}
		tag := res.GetLatestTag(checker)
		fmt.Println(tag.Name)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
