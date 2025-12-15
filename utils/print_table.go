package utils

import (
	"fmt"
	"os"
	"project-app-inventaris-cli-fathoni/model"
	"text/tabwriter"
)

// warna
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
)

func PrintTabel(categories []model.Categories) {
	if len(categories) == 0 {
		fmt.Printf("\n%s!!! No Data !!!%s\n\n", Red, Reset)
		return
	}

	// Inisialisasi tabwriter
	table := tabwriter.NewWriter(os.Stdout, 25, 4, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(table, "ID\tNAME\tDESCRIPTION")
	fmt.Fprintln(table, "--------------------------------------------------------------------")

	for _, c := range categories {
		fmt.Fprintf(table, "%d \t%s\t%s\n",
			c.Id,
			c.Name,
			c.Description,
		)
	}

	table.Flush()
}

func PrintTabelItems(items []model.Items) {
	if len(items) == 0 {
		fmt.Printf("\n%s!!! No Data !!!%s\n\n", Red, Reset)
		return
	}

	// Inisialisasi tabwriter
	table := tabwriter.NewWriter(os.Stdout, 10, 2, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(table, "ID\tNAME\tCategory\tPrice\tPurchase Date\tReplaced\tTotal Usage")
	fmt.Fprintln(table, "--------------------------------------------------------------------")

	for _, i := range items {
		fmt.Fprintf(table, "%d \t%s\t%s\t%.f\t%s\t%v\t%d\n",
			i.Id,
			i.Name,
			i.Category,
			i.Price,
			i.PurchaseDate.Format("2006-01-02"),
			i.Replaced,
			i.TotalUsage,
		)
	}

	table.Flush()
}

func PrintTabelInvest(invest []model.TotalInvest) {
	if len(invest) == 0 {
		fmt.Printf("\n%s!!! No Data !!!%s\n\n", Red, Reset)
		return
	}

	// Inisialisasi tabwriter
	table := tabwriter.NewWriter(os.Stdout, 25, 4, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(table, "Total Item\tTotal Purchase Price\tTotal Current Investment\tTotal Depreciation")
	fmt.Fprintln(table, "--------------------------------------------------------------------")

	for _, i := range invest {
		fmt.Fprintf(table, "%d\t%d\t%.f\t%.f\n",
			i.TotalItem,
			i.TotalPurchasePrice,
			i.TotalInvest,
			i.TotalDepreciation,
		)
	}

	table.Flush()
}

func PrintTabelInvestItems(invest []model.Items) {
	if len(invest) == 0 {
		fmt.Printf("\n%s!!! No Data !!!%s\n\n", Red, Reset)
		return
	}

	// Inisialisasi tabwriter
	table := tabwriter.NewWriter(os.Stdout, 15, 4, 2, ' ', tabwriter.Debug)

	fmt.Fprintln(table, "ID\tName\tPurchase Price\tPurchase Date\tAge Item\tCurrent Investment\tDepreciation Value")
	fmt.Fprintln(table, "--------------------------------------------------------------------")

	for _, i := range invest {
		fmt.Fprintf(table, "%d \t%s\t%.f\t%s\t%d\t%.f\t%.f\n",
			i.Id,
			i.Name,
			i.Price,
			i.PurchaseDate.Format("2006-01-02"),
			i.AgeItem,
			i.CurrentInvestment,
			i.DepreciationValue,
		)
	}

	table.Flush()
}
