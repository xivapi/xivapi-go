package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/paars/xiv/xivapi"
)

var (
	searchDetails int
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Issue A Search Request",
	Args:  cobra.MinimumNArgs(1),
	RunE:  RunSearch,
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.PersistentFlags().IntVarP(&searchDetails, "details", "d", -1, "Show searchDetails for the specific entry")
}

// RunSearch runs the search command
func RunSearch(cmd *cobra.Command, args []string) error {
	c := xivapi.New(cmd.Flag("key").Value.String())
	res, err := c.Search(args[0], "", "", "", 0, 0, nil)
	if err != nil {
		return err
	}

	if searchDetails > 0 {
		if searchDetails > res.Pagination.Results {
			return errors.New("failed to fetch details for requested index")
		}

		entry := res.Results[searchDetails-1]
		switch entry.Type {
		case xivapi.IndexCompanion:
			entity, err := entry.GetCompanion()
			if err != nil {
				return err
			}
			fmt.Println(entity)
		}
	} else {
		fmt.Println(res)
	}

	return nil
}
