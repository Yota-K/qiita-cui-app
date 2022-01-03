package main

import (
	"github.com/Yota-K/qiita-cui-app/setting"
	"github.com/Yota-K/qiita-cui-app/ui"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	appSetting()
}

func appSetting() {
	app := &cli.App{
		Name:  "qiita",
		Usage: "Qiita CUI Application.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "number",
				Aliases: []string{"n"},
				Value:   10,
				Usage:   "Set number of posts.",
			},
			&cli.StringFlag{
				Name:    "word",
				Aliases: []string{"w"},
				Value:   "",
				Usage:   "Set search word.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Initial setting.",
				Action: func(c *cli.Context) error {
					setting.InitSetting(c.Args().First())
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			// ui.QiitaUi(c.Int("number"), c.String("word"))
			ui.QiitaUi(c.Int("number"), "react")
			return nil
		},
	}

	app.Run(os.Args)
}
