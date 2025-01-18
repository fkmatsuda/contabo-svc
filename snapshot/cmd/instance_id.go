package cmd

import (
	"fmt"

	instanceSevices "github.com/fkmatsuda/contabo-svc/instance/services"
	"github.com/fkmatsuda/contabo-svc/lib/contabo"
	"github.com/spf13/cobra"
)

func getInstanceID(traceID string, cmd *cobra.Command) (string, int64, error) {
	instanceID, err := cmd.Flags().GetInt64("instance-id")
	if err != nil {
		return "", 0, err
	}

	instanceName, err := cmd.Flags().GetString("instance-name")
	if err != nil {
		return "", 0, err
	}

	if instanceID == 0 && instanceName == "" {
		return "", 0, fmt.Errorf("Either instance-id or instance-name must be provided")
	}

	var requestID string
	token, err := contabo.GetAccessToken()
	if err != nil {
		return "", 0, err
	}

	if instanceID == 0 {
		instanceService := instanceSevices.NewInstanceService(token)
		requestID, instanceID, err = instanceService.GetInstanceIDByName(traceID, instanceName)
		if err != nil {
			fmt.Printf("Error getting instance ID: %s\nRequest ID: %s\n", err, requestID)
			return requestID, 0, err
		}
	}

	return requestID, instanceID, nil
}
