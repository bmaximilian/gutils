package length

import (
	"github.com/bmaximilian/gutils/pkg/calculate/scale/length"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/fatih/color"
	googleLogger "github.com/google/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
)

var LengthCommand = &cobra.Command{
	Use:   "length [length] [sourceUnit] [destinationUnit]",
	Short: "Calculate a length",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)
		scale := viper.Get("calculate.scale").(float64)
		sourceUnit := args[1]
		destinationUnit := args[2]
		passedLength, parseError := strconv.ParseFloat(strings.Replace(args[0], ",", ".", -1), 64)

		if parseError != nil {
			l.Fatalln(parseError)
		}

		converted, conversionError := length.ConvertForScale(passedLength, sourceUnit, scale, destinationUnit)
		if conversionError != nil {
			l.Fatalln(conversionError)
		}

		l.Infof(
			"%v%v sind im Maßstab %v:\t %v%v\n\n",
			color.CyanString("%.2f", passedLength),
			color.CyanString("%v", sourceUnit),
			color.MagentaString("\"1 zu %.1f\"", scale),
			color.GreenString("%.4f", converted),
			color.GreenString("%v", destinationUnit),
		)
	},
}
