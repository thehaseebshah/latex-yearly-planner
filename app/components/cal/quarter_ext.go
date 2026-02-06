package cal

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
)

func (q *Quarter) BreadcrumbWithLeaf(leaf string) string {
	return header.Items{header.NewIntItem(q.Year.Number), header.NewItemsGroup(
		header.NewTextItem("Q1").Bold(q.Number == 1).Ref(q.Number == 1),
		header.NewTextItem("Q2").Bold(q.Number == 2).Ref(q.Number == 2),
		header.NewTextItem("Q3").Bold(q.Number == 3).Ref(q.Number == 3),
		header.NewTextItem("Q4").Bold(q.Number == 4).Ref(q.Number == 4),
	), header.NewTextItem(leaf)}.Table(true)
}
