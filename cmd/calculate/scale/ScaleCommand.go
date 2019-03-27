package scale

import (
	"github.com/bmaximilian/gutils/cmd/calculate/scale/length"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ScaleCommand = &cobra.Command{
	Use: "scale",
	Short: "Perform scale conversion stuff",
}

// Set the default viper values
func SetDefaults() {
	viper.SetDefault("calculate.scale", 0.00)
}

// Initializes the command line tool
func InitScaleCommand() {
	ScaleCommand.AddCommand(length.LengthCommand)
}
