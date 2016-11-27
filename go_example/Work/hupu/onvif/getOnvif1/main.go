package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "getOnvif"
	app.Version = "1.1"
	app.Usage = "Using the onvif protocol for the camera information"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1",
			Usage: "host address",
		},
		cli.StringFlag{
			Name:  "username, u",
			Value: "admin",
			Usage: "username for the host",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "admin",
			Usage: "password for the host",
		},
	}

	onvif := NewOnvif()
	app.Action = func(c *cli.Context) error {
		argNum := len(c.Args())
		flagsNum := c.NumFlags()
		if argNum < 3 && flagsNum < 3 {
			return cli.NewExitError("please man help.", -1)
		}
		if argNum == 3 {
			onvif.IP = c.Args()[0]
			onvif.Username = c.Args()[1]
			onvif.Password = c.Args()[2]
		}

		if flagsNum == 5 {
			onvif.IP = c.String("host")
			onvif.Username = c.String("username")
			onvif.Password = c.String("password")
		}

		result, _ := onvif.OnvifDevice()
		// if err != nil {
		// 	return err
		// }
		fmt.Println(result)
		return nil
	}

	app.Run(os.Args)
}
