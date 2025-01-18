package cmd

import (
	"fmt"

	"github.com/fkmatsuda/contabo-svc/instance/services"
	"github.com/fkmatsuda/contabo-svc/lib/contabo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all VPS instances",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := contabo.GetAccessToken()
		if err != nil {
			panic(err)
		}
		service := services.NewInstanceService(token)
		requestID, listInstanceResponse, err := service.ListInstances("list_instances")
		if err != nil {
			panic(err)
		}
		fmt.Printf("List all VPS instances\nRequest ID: %s\n", requestID)
		for _, instance := range listInstanceResponse.Data {
			fmt.Printf("Instance ID: %d\n", instance.InstanceId)
			fmt.Printf("Instance Name: %s\n", instance.Name)
			fmt.Printf("Instance Display Name: %s\n", instance.DisplayName)
			fmt.Printf("Instance Data Center: %s\n", instance.DataCenter)
			fmt.Printf("System Host ID: %d\n", instance.VHostId)
			fmt.Printf("System Host Number: %d\n", instance.VHostNumber)
			fmt.Printf("Instance Created At: %s\n", instance.CreatedDate)
			fmt.Printf("Instance Status: %s\n", instance.Status)
			fmt.Printf("Instance IPs:\n\tIPv4:\n\t\tIP: %s\n\t\tCIDR: %d\n\t\tGateway: %s\n\tIPv6:\n\t\tIP: %s\n\t\tCIDR: %d\n\t\tGateway: %s\n", instance.IpConfig.V4.Ip, instance.IpConfig.V4.NetmaskCidr, instance.IpConfig.V4.Gateway, instance.IpConfig.V6.Ip, instance.IpConfig.V6.NetmaskCidr, instance.IpConfig.V6.Gateway)
			fmt.Println("-------------------------------------------------")
		}
	},
}
