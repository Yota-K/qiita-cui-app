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

const Title = "Qiita CUI"

func QiitaUi(n, p int, w string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	items := api.GetQiitaItems(n, p, w)

	if len(items) != 0 {
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
		t.Title = Title
		t.TextStyle = ui.NewStyle(ui.ColorGreen)
		x, y := ui.TerminalDimensions()
		t.SetRect(0, 0, x, y)
		t.SetNodes(nodes)

		Keybindings(t)
	} else {
		notFoundPosts()
	}
}

// 条件に合致する投稿が0件の時の処理
func notFoundPosts() {
	p := widgets.NewParagraph()
	p.Title = Title
	p.Text = "Not found posts."
	x, y := ui.TerminalDimensions()
	p.SetRect(0, 0, x, y)
	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
