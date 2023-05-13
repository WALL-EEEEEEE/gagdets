package cmd

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	_ "github.com/WALL-EEEEEEE/gagdets/pipes"
	_ "github.com/WALL-EEEEEEE/gagdets/sources"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available plugins",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		servTypes := core.GetSupportedServTypes()
		for _, servType := range servTypes {
			servs := core.Reg.GetByType(servType)
			switch servType {
			case core.SPIDER:
				rootCmd.Println(" Spider\n")
			case core.PIPE:
				rootCmd.Println(" Pipe\n")
			case core.TASK:
				rootCmd.Println(" Task\n")
			}
			for serv_name, _ := range servs {
				rootCmd.Println("   " + serv_name)
			}
			rootCmd.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
