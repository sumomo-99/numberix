package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "numberpad",
		Usage: "Prints numbers from start to end with zero padding",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "start",
				Aliases:  []string{"s"},
				Usage:    "Start number",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "end",
				Aliases:  []string{"e"},
				Usage:    "End number",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			start := c.Int("start")
			end := c.Int("end")

			if start > end {
				return fmt.Errorf("Invalid input. Start number must be less than or equal to end number.")
			}

			endStr := strconv.Itoa(end)
			paddingLength := len(endStr)

			for i := start; i <= end; i++ {
				fmt.Println(strings.Repeat("0", paddingLength-len(strconv.Itoa(i))) + strconv.Itoa(i))
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
