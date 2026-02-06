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
		"Wake Time",
		"Sleep Time",
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
	genTable := func(startDay, endDay int) string {
		numDays := endDay - startDay + 1
		
		// Build column specification
		dayCols := ""
		for i := 0; i < numDays; i++ {
			dayCols += `|>{\centering\arraybackslash}X`
		}
		dayCols += "|"

		// Start the table
		out := `
\begingroup\scriptsize
\renewcommand{\arraystretch}{2.2}
\begin{tabularx}{\linewidth}{|l` + dayCols + `}
\hline
`
		// Header row
		out += " "
		for i := startDay; i <= endDay; i++ {
			out += " & " + strconv.Itoa(i)
		}
		out += ` \\ \hline
`
		// Data rows
		for _, row := range rows {
			out += row
			for i := 0; i < numDays; i++ {
				out += " & "
			}
			out += ` \\ \hline
`
		}

		out += `\end{tabularx}
\endgroup
`
		return out
	}

	// Split into two tables: 1-15 and 16-end
	result := genTable(1, 15)
	result += `\vspace{1cm}` // Add some vertical space between tables
	result += genTable(16, days)

	return result
}
