package cmd

import (
	"fmt"

	"github.com/Capture-411/services-opportunity/database"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed data",
	Long:  "seed data to database",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.Connect()
		database.SeedAllTables(db)
		fmt.Println("Finished!")
	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
}
