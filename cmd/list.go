package cmd

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	_ "github.com/WALL-EEEEEEE/gagdets/pipes"
	_ "github.com/WALL-EEEEEEE/gagdets/sources"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available plugins",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		servTypes := GetSupportedServTypes()
		for _, servType := range servTypes {
			servs := Reg.GetByType(servType)
			switch servType {
			case SPIDER:
				rootCmd.Println(" Spider\n")
			case PIPE:
				rootCmd.Println(" Pipe\n")
			case TASK:
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
