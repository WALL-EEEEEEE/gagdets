package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	. "github.com/WALL-EEEEEEE/Axiom/core"
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
	var spiders []string
	spider_servs := Reg.GetByType(SPIDER)
	spiders = maps.Keys(spider_servs)
	spiderSelected, _ := pterm.DefaultInteractiveSelect.WithOptions(spiders).Show("Select your spider: ")
	spider_serv_selected := spider_servs[spiderSelected].(ITask)
	Exec.Add(spider_serv_selected)
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
