package serve

import (
	berlin "berlin/cmd/http"
	koanfx "berlin/utils/koanfx"
	"github.com/spf13/cobra"
	"sync"
)

func serveAll() func(cmd *cobra.Command) {
	return func(cmd *cobra.Command) {
		var wg sync.WaitGroup
		wg.Add(1)
		go berlin.Run(&wg)
		wg.Wait()
	}
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Berlin HTTP Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		koanfx.LoadConfig()
		serveAll()(cmd)
	},
}

func RegisterCommandRecursive(parent *cobra.Command) {
	parent.AddCommand(serveCmd)
}
