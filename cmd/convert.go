package cmd

import (
	"context"

	yahon "github.com/ha36d/yahon/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// convertCmd represents the copy command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert the infrastructure",
	Long:  `convert the infrastructure from a yaml file`,
	Run: func(cmd *cobra.Command, args []string) {

		source := viper.GetString("source")

		ctx := context.Background()
		yahon.Yahon(ctx, source)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
