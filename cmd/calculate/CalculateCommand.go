package calculate

import (
	"github.com/bmaximilian/gutils/cmd/calculate/scale"
	"github.com/spf13/cobra"
)

var CalculateCommand = &cobra.Command{
	Use: "calculate",
	Short: "Perform calculation stuff",
}

// Set the default viper values
func SetDefaults() {
	scale.SetDefaults()
}

// Initializes the command line tool
func InitCalculateCommand() {
	scale.InitScaleCommand()
	CalculateCommand.AddCommand(scale.ScaleCommand)
}
