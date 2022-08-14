package sscli

import (
	"fmt"
	"github.com/hckia/shopify-seller-central-product-converter"
	"github.com/spf13/cobra"
)

var fetchswatchCmd = &cobra.Command{
	Use:		"fetchSwatches",
	Aliases:	[]string{"fs"},
	Short:		"Fetches Swatches",
	Args:		cobra.RangeArgs(1, 2),
	Run:		func(cmd *cobra.Command, args []string) {
		res := fetchSwatches(args[0], args[1])
		fmt.Println(res)
	},

}

func init() {
	rootCmd.AddCommand(fetchswatchCmd)
}