// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package env

import (
	"context"

	"github.com/daytonaio/daytona/internal/util/apiclient"
	"github.com/daytonaio/daytona/pkg/cmd/format"
	"github.com/daytonaio/daytona/pkg/views"
	"github.com/daytonaio/daytona/pkg/views/env"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List profile environment variables",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		apiClient, err := apiclient.GetApiClient(nil)
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()

		profileData, res, err := apiClient.ProfileAPI.GetProfileData(ctx).Execute()
		if err != nil {
			log.Fatal(apiclient.HandleErrorResponse(res, err))
		}

		if format.FormatFlag != "" {
			if profileData.EnvVars == nil {
				profileData.EnvVars = map[string]string{}
			}
			formattedData := format.NewFormatter(profileData.EnvVars)
			formattedData.Print()
			return
		}

		if len(profileData.EnvVars) == 0 {
			views.RenderInfoMessageBold("No environment variables set")
			return
		}

		env.List(profileData.EnvVars)
	},
}

func init() {
	format.RegisterFormatFlag(listCmd)
}
