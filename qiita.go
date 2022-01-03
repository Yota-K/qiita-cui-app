package main

import (
	// "fmt"
	// "github.com/Yota-K/qiita-cui-app/api"
	"github.com/Yota-K/qiita-cui-app/setting"
	"github.com/Yota-K/qiita-cui-app/ui"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	// hoge := api.GetQiitaItems(10, 10, "react")
	// fmt.Println(hoge)
	// fmt.Println(len(hoge))
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
			&cli.IntFlag{
				// LGTM数で記事をソート
				// -p 20
				Name:    "popular",
				Aliases: []string{"p"},
				Value:   0,
				Usage:   "Get Popular Posts. Please set LGTM count.",
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
			ui.QiitaUi(c.Int("number"), c.Int("popular"), c.String("word"))
			return nil
		},
	}

	app.Run(os.Args)
}
