// The Licensed Work is (c) 2024 Commune
// SPDX-License-Identifier: LGPL-3.0-only

package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/ChainSafe/chainbridge-core/flags"

	"github.com/ChainSafe/sygma-relayer/example/app"
)

var (
	rootCMD = &cobra.Command{
		Use: "",
	}
	runCMD = &cobra.Command{
		Use:   "run",
		Short: "Run example app",
		Long:  "Run example app",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.Run(); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	flags.BindFlags(runCMD)
}

func Execute() {
	rootCMD.AddCommand(runCMD)
	if err := rootCMD.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute root cmd")
	}
}
