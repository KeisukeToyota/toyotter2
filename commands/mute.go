package commands

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/ksk001100/toyotter/twitter"
	"github.com/urfave/cli/v2"
)

// MuteCommand mute command function
func MuteCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "mute",
		Aliases: []string{"mu"},
		Usage:   "toyotter mute [screenName]",
		Flags:   muteFlags(),
		Action:  muteAction,
	}
}

func muteFlags() []cli.Flag {
	deleteFlag := cli.BoolFlag{
		Name:  "delete, del, d",
		Usage: "toyotter mute [screenName] --delete",
	}
	return []cli.Flag{
		&deleteFlag,
	}
}

func muteAction(c *cli.Context) error {
	screenName := c.Args().First()
	if c.Bool("delete") {
		twitter.UnMute(api, screenName, v)
	} else {
		twitter.Mute(api, screenName, v)
	}
	return nil
}
