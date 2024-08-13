package cmd

import (
	"context"
	"fmt"
	"reflect"

	yahon "github.com/ha36d/yahon/pkg/yahon"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// convertCmd represents the copy command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert the infrastructure",
	Long:  `convert the infrastructure from a yaml file`,
	Run: func(cmd *cobra.Command, args []string) {

		file := viper.GetString("file")
		from := viper.GetString("from")
		to := viper.GetString("to")

		ctx := context.Background()

		f := reflect.ValueOf(yahon.Holder{}).
			MethodByName(strcase.ToCamel(fmt.Sprintf("%s to %s", from, to))).Interface().(func(context.Context, string))
		f(ctx, file)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
