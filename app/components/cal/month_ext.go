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

	// Build column specification:
	// First column: explicit left aligned (l)
	// Day columns: specific centered X columns (>{\centering\arraybackslash}X)
	// We construct the string for the day columns
	dayCols := ""
	for i := 0; i < days; i++ {
		dayCols += `|>{\centering\arraybackslash}X`
	}
	dayCols += "|" // Closing vertical bar

	// Start the table
	// \linewidth ensures it uses the full text width
	// arraystretch 2.5 makes the rows much taller (approx 1cm+)
	// \scriptsize is larger than tiny but small enough for 31 columns
	result := `
\begingroup\scriptsize
\renewcommand{\arraystretch}{2.2}
\begin{tabularx}{\linewidth}{|l` + dayCols + `}
\hline
`

	// Header row with day numbers
	result += " "
	for i := 1; i <= days; i++ {
		result += " & " + strconv.Itoa(i)
	}
	result += ` \\ \hline
`

	// Data rows for each mamoolaat
	for _, row := range rows {
		result += row
		for i := 0; i < days; i++ {
			result += " & "
		}
		result += ` \\ \hline
`
	}

	result += `\end{tabularx}
\endgroup
`

	return result
}
