package main

import (
	"fmt"
	"log"

	"github.com/k0kubun/pp"

	"github.com/spf13/cobra"
	"gitlab.com/paars/xiv/xivapi"
)

var rootCmd = &cobra.Command{
	Use:              "xivapi",
	Short:            "xivapi queries xivapi",
	Long:             `Query xivapi for different data from Final Fantasy 14 and Lodestone`,
	PersistentPreRun: preRun,
	TraverseChildren: true,
}

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Get a list of patches",
	Run:   RunPatch,
}

var fcCmd = &cobra.Command{
	Use:   "fcsearch",
	Short: "Search for a Free Company",
	Run:   RunFreeCompanySearch,
}

func preRun(*cobra.Command, []string) {
	cobra.OnInitialize(initConfig)
}

func init() {
	pp.BufferFoldThreshold = 4
	rootCmd.PersistentFlags().StringP("key", "k", "", "The XIVAPI key")
	rootCmd.AddCommand(patchCmd, fcCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run xivapi, error was %v", err)
	}
}

func initConfig() {
}

// RunFreeCompanySearch runs the free company search command
func RunFreeCompanySearch(cmd *cobra.Command, args []string) {
	c := xivapi.New(cmd.Flag("key").Value.String())
	res, err := c.FreeCompanySearch(args[0], "", 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func RunPatch(cmd *cobra.Command, args []string) {
	c := xivapi.New(cmd.Flag("key").Value.String())
	res, err := c.PatchList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
