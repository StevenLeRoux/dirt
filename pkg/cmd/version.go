package cmd

import (
	"fmt"
	"runtime"

	"github.com/StevenLeRoux/dirt/pkg/libs"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, arguments []string) {
		fmt.Printf(projectName+" version %s %s\n", libs.Version, libs.GitHash)
		fmt.Printf(projectName+" build date %s\n", libs.BuildDate)
		fmt.Printf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}
