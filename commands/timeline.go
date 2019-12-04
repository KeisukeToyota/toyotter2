package commands

import (
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/KeisukeToyota/toyotter/twitter"
	"github.com/urfave/cli"
)

// TimelineCommand timeline command function
func TimelineCommand(a *anaconda.TwitterApi, val url.Values) cli.Command {
	api = a
	v = val
	return cli.Command{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Usage:   "toyotter timeline [...option]",
		Flags:   timelineFlags(),
		Action:  timelineAction,
	}
}

func timelineFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "count, c",
			Value: "5",
			Usage: "toyotter timeline --count=[count]",
		},
		cli.StringFlag{
			Name: "list, li",
			Usage: "toyotter timeline --list=[listID]",
		},
	}
}

func timelineAction(c *cli.Context) error {
	v.Set("count", c.String("count"))

	if len(c.String("list")) > 0 {
		listID, _ := strconv.ParseInt(c.String("list"), 10, 64)
		twitter.ListTimeline(api, listID, v)
	} else {
		twitter.HomeTimeline(api, v)
	}
	return nil
}
