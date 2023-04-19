package main

import (
	"github.com/WALL-EEEEEEE/gagdets/core"
	_ "github.com/WALL-EEEEEEE/gagdets/pipes"
	_ "github.com/WALL-EEEEEEE/gagdets/sources"
)

func main() {
	core.Exec.Start()
}
