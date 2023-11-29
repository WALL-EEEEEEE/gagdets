package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	. "github.com/WALL-EEEEEEE/Axiom/core"
	"github.com/WALL-EEEEEEE/gagdets/cmd/gadget/subcmd"
	_ "github.com/WALL-EEEEEEE/gagdets/sources/spider"

	//. "github.com/WALL-EEEEEEE/gadgets/cmd/gadget/subcmd"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/bobg/go-generics/maps"
	"github.com/pterm/pterm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func init_log() {
	log.SetReportCaller(true)
	log.SetFormatter(&zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ShowFullLevel:   true,
		},
	})
}

func start() {
	init_log()
	servs_map := Reg.GetByType()
	if len(servs_map) < 1 {
		log.Warn("No Spider avaliable here")
		return
	}
	servs := maps.Keys(servs_map)
	serv_name_selected, _ := pterm.DefaultInteractiveSelect.WithOptions(servs).Show("Select your services: ")
	serv_selected := servs_map[serv_name_selected].(ITask)
	Exec.Add(serv_selected)
	Exec.Start()
}

var rootCmd = &cobra.Command{
	Use:   "gagdet",
	Short: "Gadget is just a tool used in my daily life",
	Long:  `A Simple and useful tools for boosting my life efficience`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		start()
	},
}

func main() {
	rootCmd.AddCommand(subcmd.ListCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
