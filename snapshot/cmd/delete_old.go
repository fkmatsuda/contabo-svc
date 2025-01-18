package cmd

import (
	"fmt"

	"github.com/fkmatsuda/contabo-svc/lib/contabo"
	"github.com/fkmatsuda/contabo-svc/snapshot/services"
	"github.com/spf13/cobra"
)

var deleteOldCmd = &cobra.Command{
	Use:   "old",
	Short: "Delete old snapshots from a VPS instance",
	Run: func(cmd *cobra.Command, args []string) {
		const traceID = "delete-old-snapshot"

		_, instanceID, err := getInstanceID(traceID, cmd)
		if err != nil {
			panic(err)
		}

		token, err := contabo.GetAccessToken()
		if err != nil {
			panic(err)
		}

		snapshotsToKeep, err := cmd.Flags().GetInt("snapshots-to-keep")

		service := services.NewSnapshotService(token)

		requestID, err := service.DeleteOldSnapshots(traceID, instanceID, snapshotsToKeep)
		if err != nil {
			fmt.Printf("Error deleting old snapshots: \n\tRequest ID: %s\n", requestID)
			panic(err)
		}

	},
}

func init() {
	deleteOldCmd.Flags().Int("snapshots-to-keep", 0, "Number of snapshots to keep")
	deleteOldCmd.MarkFlagRequired("snapshots-to-keep")
}
