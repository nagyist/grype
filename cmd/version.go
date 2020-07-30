package cmd

import (
	"fmt"

	v1 "github.com/anchore/grype-db/pkg/db/v1"
	"github.com/anchore/grype/internal"
	"github.com/anchore/grype/internal/version"
	"github.com/spf13/cobra"
)

var showVerboseVersionInfo bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version",
	Run:   printVersion,
}

func init() {
	versionCmd.Flags().BoolVarP(&showVerboseVersionInfo, "verbose", "v", false, "show additional version information")

	rootCmd.AddCommand(versionCmd)
}

func printVersion(_ *cobra.Command, _ []string) {
	versionInfo := version.FromBuild()
	if showVerboseVersionInfo {
		fmt.Println("Application:         ", internal.ApplicationName)
		fmt.Println("Version:             ", versionInfo.Version)
		fmt.Println("BuildDate:           ", versionInfo.BuildDate)
		fmt.Println("GitCommit:           ", versionInfo.GitCommit)
		fmt.Println("GitTreeState:        ", versionInfo.GitTreeState)
		fmt.Println("Platform:            ", versionInfo.Platform)
		fmt.Println("GoVersion:           ", versionInfo.GoVersion)
		fmt.Println("Compiler:            ", versionInfo.Compiler)
		fmt.Println("Supported DB Schema: ", v1.SchemaVersion)
	} else {
		fmt.Printf("%s %s\n", internal.ApplicationName, versionInfo.Version)
	}
}