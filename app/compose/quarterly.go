package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Quarterly(cfg config.Config, tpls []string) (page.Modules, error) {
	modules := make(page.Modules, 0, 4)
	year := cal.NewYear(cfg.WeekStart, cfg.Year)

	hRight := header.Items{
		header.NewTextItem("Notes Index").RefText("Tasks Index"),
	}

	for _, quarter := range year.Quarters {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Year":         year,
				"Quarter":      quarter,
				"Breadcrumb":   quarter.Breadcrumb(),
				"HeadingMOS":   quarter.HeadingMOS(),
				"SideQuarters": year.SideQuarters(quarter.Number),
				"SideMonths":   year.SideMonths(0),
				"Extra":        hRight.WithTopRightCorner(cfg.ClearTopRightCorner),
				"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, nil, 0),
			},
		})

		if len(tpls) > 1 {
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[1],
				Body: map[string]interface{}{
					"Year":         year,
					"Quarter":      quarter,
					"Breadcrumb":   quarter.BreadcrumbWithLeaf("Tasks"),
					"HeadingMOS":   quarter.HeadingMOS(),
					"SideQuarters": year.SideQuarters(quarter.Number),
					"SideMonths":   year.SideMonths(0),
					"Extra":        hRight.WithTopRightCorner(cfg.ClearTopRightCorner),
					"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, nil, 0),
				},
			})
		}
	}

	return modules, nil
}
