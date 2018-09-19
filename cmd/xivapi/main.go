package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/paars/xiv/xivapi-go"

	"github.com/spf13/cobra"
)

var (
	details int
)

var rootCmd = &cobra.Command{
	Use:              "xivapi",
	Short:            "xivapi queries xivapi",
	Long:             `Query xivapi for different data from Final Fantasy 14 and Lodestone`,
	Run:              Run,
	TraverseChildren: true,
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Issue A Search Request",
	Args:  cobra.MinimumNArgs(1),
	Run:   RunSearch,
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

func main() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("key", "k", "", "The XIVAPI key")
	rootCmd.AddCommand(searchCmd, patchCmd, fcCmd)

	searchCmd.PersistentFlags().IntVarP(&details, "details", "d", -1, "Show details for the specific entry")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run xivapi, error was %v", err)
		os.Exit(2)
	}
}

func initConfig() {
}

func Run(cmd *cobra.Command, args []string) {

}

func RunFreeCompanySearch(cmd *cobra.Command, args []string) {
	c := xivapi.New(cmd.Flag("key").Value.String())
	res, err := c.FreeCompanySearch(args[0], "", "")
	if err != nil {
		log.Fatal(err)
	}

	if details > 0 {
		if details > res.Pagination.Results {
			log.Fatal("Failed to fetch details for requested index")
		}
	} else {
		fmt.Println(res)
	}

}

func RunSearch(cmd *cobra.Command, args []string) {
	c := xivapi.New(cmd.Flag("key").Value.String())
	res, err := c.Search(args[0], "", "", "", 0, 0, nil)
	if err != nil {
		log.Fatal(err)
	}

	if details > 0 {
		if details > res.Pagination.Results {
			log.Fatal("Failed to fetch details for requested index")
		}

		entry := res.Results[details-1]
		switch entry.Type {
		case xivapi.IndexCompanion:
			if entity, err := entry.GetCompanion(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(entity)
			}
		}
	} else {
		fmt.Println(res)
	}
}

func RunPatch(cmd *cobra.Command, args []string) {
	c := xivapi.New(cmd.Flag("key").Value.String())
	res, err := c.PatchList()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
