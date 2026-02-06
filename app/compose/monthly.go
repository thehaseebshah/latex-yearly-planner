package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Monthly(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 24) // 12 months * 2 pages each

	hRight := header.Items{
		header.NewTextItem("Mamoolaat"),
	}

	for _, quarter := range year.Quarters {
		for _, month := range quarter.Months {
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[0],
				Body: map[string]interface{}{
					"Year":         year,
					"Quarter":      quarter,
					"Month":        month,
					"Breadcrumb":   month.Breadcrumb(),
					"HeadingMOS":   month.HeadingMOS(),
					"SideQuarters": year.SideQuarters(quarter.Number),
					"SideMonths":   year.SideMonths(month.Month),
					"Extra":        hRight.WithTopRightCorner(cfg.ClearTopRightCorner),
					"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, nil, 0),
				},
			})

			// Add Mamoolaat page after each monthly page
			if len(tpls) > 1 {
				modules = append(modules, page.Module{
					Cfg: cfg,
					Tpl: tpls[1],
					Body: map[string]interface{}{
						"Year":         year,
						"Quarter":      quarter,
						"Month":        month,
						"Breadcrumb":   month.Breadcrumb() + ` $\vert$ Mamoolaat`,
						"HeadingMOS":   month.HeadingMOS(),
						"SideQuarters": year.SideQuarters(quarter.Number),
						"SideMonths":   year.SideMonths(month.Month),
						"Extra":        month.PrevNext().WithTopRightCorner(cfg.ClearTopRightCorner),
						"Extra2":       extra2(cfg.ClearTopRightCorner, false, false, nil, 0),
					},
				})
			}
		}
	}

	return modules, nil
}
