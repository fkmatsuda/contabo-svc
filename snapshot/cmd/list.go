package cmd

import (
	"fmt"

	"github.com/fkmatsuda/contabo-svc/lib/contabo"
	"github.com/fkmatsuda/contabo-svc/snapshot/services"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snapshots for a VPS instance",
	Run: func(cmd *cobra.Command, args []string) {
		const traceID = "list-snapshots"

		requestID, instanceID, err := getInstanceID(traceID, cmd)
		if err != nil {
			fmt.Printf("Error getting instance ID: \n\tRequest ID: %s\n", requestID)
			panic(err)
		}

		token, err := contabo.GetAccessToken()
		if err != nil {
			panic(err)
		}

		service := services.NewSnapshotService(token)
		requestID, listSnapshotResponse, err := service.GetInstanceSnapshots(traceID, instanceID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("List all snapshots for vps: %d\nRequest ID: %s\n", instanceID, requestID)
		for _, snapshot := range listSnapshotResponse.Data {
			fmt.Printf("Snapshot ID: %s\n", snapshot.SnapshotID)
			fmt.Printf("Snapshot Name: %s\n", snapshot.Name)
			fmt.Printf("Snapshot Description: %s\n", snapshot.Description)
			fmt.Printf("Snapshot Created At: %s\n", snapshot.CreatedDate)
			fmt.Printf("Snapshot Size: %s\n", snapshot.AutoDeleteDate)
			fmt.Println("-------------------------------------------------")

		}
	},
}

func init() {
	listCmd.Flags().Int64("instance-id", 0, "VPS Instance ID")
	listCmd.Flags().String("instance-name", "", "VPS Instance Name")

	listCmd.MarkFlagsOneRequired("instance-id", "instance-name")
}
