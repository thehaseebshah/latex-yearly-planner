package cal

import (
	"strconv"
	"time"
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
	return m.Breadcrumb() + " $\\vert$ " + leaf
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

	// Build column specification: first column for row labels, then one for each day
	// We need days columns plus 1 for the label
	colSpec := "|l|"
	for i := 0; i < days; i++ {
		colSpec += "c|"
	}

	// Start the table
	result := `\begingroup\tiny
\setlength{\tabcolsep}{1pt}
\renewcommand{\arraystretch}{1.2}
\begin{tabular}{` + colSpec + `}
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

	result += `\end{tabular}
\endgroup`

	return result
}
