package cmd

import (
	. "reward/internal"

	"github.com/spf13/cobra"
)

var svcCmd = &cobra.Command{
	Use:                "svc",
	Short:              "Orchestrates global services such as traefik, portainer and dnsmasq via docker-compose",
	Long:               `Orchestrates global services such as traefik, portainer and dnsmasq via docker-compose`,
	ValidArgsFunction:  DockerComposeCompleter(),
	DisableFlagParsing: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := CheckIfInstalled(); err != nil {
			return err
		}

		if err := CheckDocker(); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return SvcCmd(args)
	},
}

func init() {
	rootCmd.AddCommand(svcCmd)
}
