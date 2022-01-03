package ui

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"github.com/Yota-K/qiita-cui-app/api"
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
		node := widgets.TreeNode{
			Value: nodeValue(item.Title),
			Nodes: []*widgets.TreeNode{
				{
					Value: nodeValue(fmt.Sprintf("cmd+click → %s", item.Url)),
					Nodes: nil,
				},
				{
					Value: nodeValue(fmt.Sprintf("CreatedAt: %s", item.CreatedAt)),
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
