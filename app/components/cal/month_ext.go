package cal

import (
	"strconv"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
)

// DaysInMonth returns the number of days in the month
func (m *Month) DaysInMonth() int {
	// Go to the first day of next month and subtract one day
	firstOfNextMonth := time.Date(m.Year.Number, m.Month+1, 1, 0, 0, 0, 0, time.UTC)
	lastOfThisMonth := firstOfNextMonth.AddDate(0, 0, -1)
	return lastOfThisMonth.Day()
}

// DayNumbers returns a slice of day numbers for this month (1 to DaysInMonth)
func (m *Month) DayNumbers() []int {
	days := m.DaysInMonth()
	result := make([]int, days)
	for i := 0; i < days; i++ {
		result[i] = i + 1
	}
	return result
}

// BreadcrumbWithLeaf returns the breadcrumb with an additional leaf label
func (m *Month) BreadcrumbWithLeaf(leaf string) string {
	return header.Items{
		header.NewIntItem(m.Year.Number),
		header.NewTextItem("Q" + strconv.Itoa(m.Quarter.Number)),
		header.NewMonthItem(m.Month).Ref(),
		header.NewTextItem(leaf),
	}.Table(true)
}

// MamoolaatRows returns the mamoolaat items with their display names
func (m *Month) MamoolaatRows() []string {
	return []string{
		"Timely Wake",
		"Timely Sleep",
		"Tahajjud Protocol",
		"Fajr Protocol",
		"Ishraq Protocol",
		"Zuhr Protocol",
		"Asr Protocol",
		"Maghrib Protocol",
		"Isha Protocol",
		"Prohibitions",
		"Mandatory Pomos",
		"Optional Pomos",
	}
}


// MamoolaatTable generates the LaTeX code for the mamoolaat table
func (m *Month) MamoolaatTable() string {
	days := m.DaysInMonth()
	rows := m.MamoolaatRows()

	// Helper function to generate a table for a range of days
	// It basically forces 16 columns layout to ensure even widths
	genTable := func(startDay, endDay int) string {
		const forcedCols = 16
		
		// Build column specification
		dayCols := ""
		for i := 0; i < forcedCols; i++ {
			dayCols += `|>{\centering\arraybackslash}X`
		}
		dayCols += "|"

		// Start the table
		out := `\begingroup\scriptsize
\renewcommand{\arraystretch}{2.2}
\begin{tabularx}{\linewidth}{|l` + dayCols + `}
\hline
`
		// Header row
		out += " " // First empty cell for labels
		for i := 0; i < forcedCols; i++ {
			dayNum := startDay + i
			cellContent := ""
			if dayNum <= endDay {
				cellContent = strconv.Itoa(dayNum)
			}
			out += " & " + cellContent
		}
		out += ` \\ \hline
`
		// Data rows
		for _, row := range rows {
			out += row
			for i := 0; i < forcedCols; i++ {
				out += " & "
			}
			out += ` \\ \hline
`
		}

		out += `\end{tabularx}
\endgroup`
		return out
	}

	// Split into two tables: 1-16 and 17-end
	// Apply \noindent before each table for consistent alignment
	result := `\noindent` + genTable(1, 16)
	result += `

\vspace{1cm}

\noindent` + genTable(17, days)

	return result
}
