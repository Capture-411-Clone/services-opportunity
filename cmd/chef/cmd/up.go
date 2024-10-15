package cmd

import (
	"fmt"

	"github.com/Capture-411/services-opportunity/database"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "migrate up",
	Long:  "migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.Connect()
		database.Migrate(db)
		fmt.Println("Finished!")
	},
}

func init() {
	RootCmd.AddCommand(migrateUpCmd)
}
