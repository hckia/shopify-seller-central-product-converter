package sscli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sscli",
	Short: "sscli - a CLI to transform TimeForTouchup Shopify product data into swatches, and products for Amazon seller central",
	Long: `sscli is an app to transform Vehicle Data into Swatches & Products)
   
One can use sscli to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
