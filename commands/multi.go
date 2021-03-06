package commands

import (
	"flag"
	"fmt"
	"github.com/PagerDuty/nut/container"
	"github.com/mitchellh/cli"
	log "github.com/sirupsen/logrus"
	"strings"
)

type MultiCommand struct {
}

func Multi() (cli.Command, error) {
	command := &MultiCommand{}
	return command, nil
}

func (command *MultiCommand) Help() string {
	helpText := `
	Usage: nut multi [options]

		Create a set of LXC container from docker-compose like yaml specifiction

  Options:

		-specfile=docker-compose.yml  Path of the specification file
	`
	return strings.TrimSpace(helpText) + AddCommonHelp()
}

func (command *MultiCommand) Synopsis() string {
	synopsis := "Build multi container environment from docker compose specification"
	return synopsis
}

func (command *MultiCommand) Run(args []string) int {
	flagSet := flag.NewFlagSet("multi", flag.ExitOnError)
	flagSet.Usage = func() { fmt.Println(command.Help()) }
	file := flagSet.String("specfile", "docker-compose.yml", "Multi-container specification file")
	AddCommonFlags(flagSet)
	if err := flagSet.Parse(args); err != nil {
		log.Errorln(err)
		return -1
	}
	ConfigureLogging()
	g, err := container.GroupFromYAML(*file)
	if err != nil {
		log.Errorln(err)
		return -1
	}
	log.Debugf("Successfully created group specification form yaml file")
	if err := g.Create(); err != nil {
		log.Errorln(err)
		return -1
	}
	return 0
}
