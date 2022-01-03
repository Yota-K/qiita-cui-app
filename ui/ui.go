package ui

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/Yota-K/qiita-cui-app/api"
	"strconv"
)

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func QiitaUi(n int, word string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	items := api.GetQiitaItems(n, word)
	var nodes []*widgets.TreeNode

	for _, item := range items {
		createdAt := item.CreatedAt.Format("2006/01/02")

		node := widgets.TreeNode{
			Value: nodeValue(fmt.Sprintf("%s: %s", createdAt, item.Title)),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue(fmt.Sprintf("LGTM: %s", strconv.Itoa(item.LikesCount))),
					Nodes: nil,
				},
				{
					Value: nodeValue(fmt.Sprintf("cmd+click → %s", item.Url)),
					Nodes: nil,
				},
			},
		}

		nodes = append(nodes, &node)
	}

	t := widgets.NewTree()
	t.Title = "Qiita CUI"
	t.TextStyle = ui.NewStyle(ui.ColorGreen)
	x, y := ui.TerminalDimensions()
	t.SetRect(0, 0, x, y)
	t.SetNodes(nodes)

	Keybindings(t)
}
