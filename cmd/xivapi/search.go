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
		var entity interface{}
		var err error
		switch entry.Type {
		case xivapi.IndexAchievement:
			entity, err = entry.GetAchievement()
		case xivapi.IndexAction:
			entity, err = entry.GetAction()
		case xivapi.IndexBNPCName:
			entity, err = entry.GetBNPCName()
		case xivapi.IndexCompanion:
			entity, err = entry.GetCompanion()
		case xivapi.IndexENPCResident:
			entity, err = entry.GetENPCResident()
		case xivapi.IndexEmote:
			entity, err = entry.GetEmote()
		case xivapi.IndexFate:
			entity, err = entry.GetFate()
		case xivapi.IndexInstanceContent:
			entity, err = entry.GetInstanceContent()
		case xivapi.IndexItem:
			entity, err = entry.GetItem()
		case xivapi.IndexLeve:
			entity, err = entry.GetLeve()
		case xivapi.IndexMount:
			entity, err = entry.GetMount()
		case xivapi.IndexPlaceName:
			entity, err = entry.GetPlaceName()
		case xivapi.IndexQuest:
			entity, err = entry.GetQuest()
		case xivapi.IndexRecipe:
			entity, err = entry.GetRecipe()
		case xivapi.IndexStatus:
			entity, err = entry.GetStatus()
		case xivapi.IndexTitle:
			entity, err = entry.GetTitle()
		case xivapi.IndexWeather:
			entity, err = entry.GetWeather()
		case xivapi.IndexBuddyEquip:
			entity, err = entry.GetBuddyEquip()
		case xivapi.IndexOrchestrion:
			entity, err = entry.GetOrchestrion()

		default:
			pp.Println(entry)
			return nil
		}
		if err != nil {
			return err
		}
		pp.Println(entity)
	} else {
		pp.Println(res)
	}

	return nil
}
