package call

import (
	"fmt"

	"github.com/hckia/shopify-seller-central-product-converter/sscli/pkg/sscli"
	"github.com/spf13/cobra"
)

var fetchprodCmd = &cobra.Command{
	Use:     "fetchProducts",
	Aliases: []string{"fd"},
	Short:   "Fetches Products",
	Args:    cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		res := sscli.FetchProducts(args[0], args[1])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(fetchprodCmd)
}
