package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "The current gutils version",
	Long:  "Print the installed gutils version",
	Run: func(cmd *cobra.Command, args []string) {
		version := viper.Get("VERSION").(string)
		build := viper.Get("BUILD").(string)

		fmt.Println("CLI Utils of bmaximilian")
		fmt.Println("\tVersion: " + version)
		fmt.Println("\tBuild: " + build)
	},
}
