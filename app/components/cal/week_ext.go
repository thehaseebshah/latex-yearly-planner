package cal

import (
	"strconv"
	
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
)

func (w *Week) BreadcrumbWithLeaf(leaf string) string {
	return header.Items{
		header.NewIntItem(w.Year.Number),
		w.QuartersBreadcrumb(),
		w.MonthsBreadcrumb(),
		header.NewTextItem("Week " + strconv.Itoa(w.weekNumber())).RefText(w.ref()).Ref(true),
		header.NewTextItem(leaf),
	}.Table(true)
}
