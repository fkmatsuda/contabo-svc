package cmd

import (
	"fmt"

	"github.com/fkmatsuda/contabo-svc/lib/contabo"
	"github.com/fkmatsuda/contabo-svc/snapshot/services"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new snapshot for a VPS instance",
	Run: func(cmd *cobra.Command, args []string) {
		const traceID = "create-snapshot"

		snapshotName, err := cmd.Flags().GetString("name")
		if err != nil {
			panic(err)
		}

		snapshotDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			panic(err)
		}

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
		requestID, snapshotCreateReponse, err := service.CreateSnapshot(traceID, instanceID, snapshotName, snapshotDescription)
		if err != nil {
			fmt.Printf("Error creating snapshot: %s\nRequest ID: %s\n", err, requestID)
			panic(err)
		}
		fmt.Printf("Snapshot created successfully for vps %d\nRequest ID: %s\n", instanceID, requestID)
		for _, snapshot := range snapshotCreateReponse.Data {
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
	createCmd.Flags().Int64("instance-id", 0, "VPS Instance ID")
	createCmd.Flags().String("instance-name", "", "VPS Instance Name")

	createCmd.Flags().String("name", "", "Snapshot name")
	createCmd.Flags().String("description", "", "Snapshot description")

	createCmd.MarkFlagsOneRequired("instance-id", "instance-name")
	createCmd.MarkFlagRequired("name")
}
