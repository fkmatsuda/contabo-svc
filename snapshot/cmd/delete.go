package cmd

import (
	"fmt"

	"github.com/fkmatsuda/contabo-svc/lib/contabo"
	"github.com/fkmatsuda/contabo-svc/snapshot/services"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot from a VPS instance",
	Run: func(cmd *cobra.Command, args []string) {
		const traceID = "delete-snapshots"

		snapshotID, err := cmd.Flags().GetString("snapshot-id")

		_, instanceID, err := getInstanceID(traceID, cmd)
		if err != nil {
			panic(err)
		}

		token, err := contabo.GetAccessToken()
		if err != nil {
			panic(err)
		}

		service := services.NewSnapshotService(token)
		requestID, err := service.DeleteSnapshot(traceID, instanceID, snapshotID)
		if err != nil {
			fmt.Printf("Error deleting snapshot ID: %s\nRequest ID: %s\n", snapshotID, requestID)
			panic(err)
		}
	},
}

func init() {
	deleteCmd.PersistentFlags().Int64("instance-id", 0, "VPS Instance ID")
	deleteCmd.PersistentFlags().String("instance-name", "", "VPS Instance Name")
	deleteCmd.Flags().String("snapshot-id", "", "Snapshot ID to delete")

	deleteCmd.MarkFlagRequired("snapshot-id")

	deleteCmd.MarkFlagsOneRequired("instance-id", "instance-name")

	deleteCmd.AddCommand(deleteOldCmd)
}
