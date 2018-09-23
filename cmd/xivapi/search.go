package main

import (
	"errors"

	"github.com/k0kubun/pp"
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
		case xivapi.IndexAchievement:
			entity, err := entry.GetAchievement()
			if err != nil {
				return err
			}
			pp.Println(entity)
		case xivapi.IndexAction:
			entity, err := entry.GetAction()
			if err != nil {
				return err
			}
			pp.Println(entity)

		default:
			pp.Println(entry)
		}
	} else {
		pp.Println(res)
	}

	return nil
}
