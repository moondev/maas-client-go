package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/moondev/maas-client-go/maasclient"
)

func main() {
	// Get MAAS configuration from environment variables
	maasEndpoint := os.Getenv("MAAS_ENDPOINT")
	maasAPIKey := os.Getenv("MAAS_API_KEY")

	if maasEndpoint == "" || maasAPIKey == "" {
		log.Fatal("MAAS_ENDPOINT and MAAS_API_KEY environment variables must be set")
	}

	// Create authenticated client
	// client := maasclient.NewAuthenticatedClientSet(maasEndpoint, maasAPIKey)
	// ctx := context.Background()

	c := maasclient.NewAuthenticatedClientSet(os.Getenv("MAAS_ENDPOINT"), os.Getenv("MAAS_API_KEY"))

	ctx := context.Background()

	// List DNS Resources
	res, err := c.DNSResources().List(ctx, nil)
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Println(res)
	// // List DNS Resources filtered by fqdn
	// filters := ParamsBuilder().Add(FQDNKey, "bad-doesntexist.maas")
	// res, err := c.DNSResources().List(ctx, filters)

	// // Create DNS Resource
	// res, err := c.DNSResources().
	// 	Builder().
	// 	WithFQDN("test-unit1.maas.sc").
	// 	WithAddressTTL("10").Create(ctx)

	// // Update DNS Resource
	// err = res.Modifier().
	// 	SetIPAddresses([]string{"1.2.3.4", "5.6.7.8"}).
	// 	Modify(ctx)

	// // Get DNS Resource by ID
	// res2 := c.DNSResources().DNSResource(res.ID())

	// // Delete DNS Resource
	// err = res.Delete(ctx)

	// // Create authenticated client
	// client := maasclient.NewAuthenticatedClientSet(maasEndpoint, maasAPIKey)
	// ctx := context.Background()

	// fmt.Println("=== Ephemeral Deployment Example ===")

	// // Example 1: Deploy in ephemeral mode (in-memory)
	// fmt.Println("\n1. Deploying machine in ephemeral mode...")
	// machine, err := deployMachine(client, ctx, true)
	// if err != nil {
	// 	log.Printf("Failed to deploy machine in ephemeral mode: %v", err)
	// } else {
	// 	fmt.Printf("Successfully deployed machine %s in ephemeral mode\n", machine.SystemID())
	// 	// Clean up
	// 	releaseMachine(machine, ctx)
	// }

	// // Example 2: Deploy in persistent mode (on disk)
	// fmt.Println("\n2. Deploying machine in persistent mode...")
	// machine, err = deployMachine(client, ctx, false)
	// if err != nil {
	// 	log.Printf("Failed to deploy machine in persistent mode: %v", err)
	// } else {
	// 	fmt.Printf("Successfully deployed machine %s in persistent mode\n", machine.SystemID())
	// 	// Clean up
	// 	releaseMachine(machine, ctx)
	// }

	// fmt.Println("\n=== Example completed ===")
}

// func deployMachine(client maasclient.ClientSetInterface, ctx context.Context, ephemeral bool) (maasclient.Machine, error) {
// 	// Allocate a machine
// 	fmt.Printf("Allocating machine...\n")
// 	machine, err := client.Machines().Allocator().Allocate(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to allocate machine: %w", err)
// 	}

// 	fmt.Printf("Allocated machine: %s\n", machine.SystemID())

// 	// Update machine settings
// 	fmt.Printf("Updating machine settings...\n")
// 	machine, err = machine.Modifier().SetSwapSize(0).Update(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update machine: %w", err)
// 	}

// 	// Deploy the machine
// 	deployMode := "persistent"
// 	if ephemeral {
// 		deployMode = "ephemeral"
// 	}
// 	fmt.Printf("Deploying machine in %s mode...\n", deployMode)

// 	machine, err = machine.Deployer().
// 		SetOSSystem("ubuntu").
// 		SetDistroSeries("noble").
// 		SetEphemeralDeploy(ephemeral).
// 		Deploy(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to deploy machine: %w", err)
// 	}

// 	fmt.Printf("Machine deployed successfully in %s mode\n", deployMode)
// 	return machine, nil
// }

// func releaseMachine(machine maasclient.Machine, ctx context.Context) {
// 	fmt.Printf("Releasing machine %s...\n", machine.SystemID())
// 	_, err := machine.Releaser().
// 		WithComment("Example cleanup").
// 		Release(ctx)
// 	if err != nil {
// 		fmt.Printf("Warning: Failed to release machine: %v\n", err)
// 	} else {
// 		fmt.Printf("Machine %s released successfully\n", machine.SystemID())
// 	}
// }
