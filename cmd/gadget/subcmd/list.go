package subcmd

import (
	. "github.com/WALL-EEEEEEE/Axiom/core"
	_ "github.com/WALL-EEEEEEE/gagdets/pipes"
	_ "github.com/WALL-EEEEEEE/gagdets/sources"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available plugins",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		servTypes := GetSupportedServTypes()
		for _, servType := range servTypes {
			servs := Reg.GetByType(servType)
			switch servType {
			case SPIDER:
				cmd.Println(" Spider\n")
			case PIPE:
				cmd.Println(" Pipe\n")
			case TASK:
				cmd.Println(" Task\n")
			}
			for serv_name, _ := range servs {
				cmd.Println("   " + serv_name)
			}
			cmd.Println()
		}

	},
}
