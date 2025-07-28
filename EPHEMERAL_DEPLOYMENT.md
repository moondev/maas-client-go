# Ephemeral Deployment Support

This document describes how to use the ephemeral deployment feature in the MAAS client.

## Overview

Ephemeral deployment allows you to deploy machines in an in-memory mode where the operating system runs entirely in RAM without persisting changes to disk. This is useful for temporary testing, development environments, or scenarios where you need a clean state after each reboot.

## Usage

### Basic Ephemeral Deployment

```go
package main

import (
    "context"
    "github.com/moondev/maas-client-go/maasclient"
)

func main() {
    // Create authenticated client
    client := maasclient.NewAuthenticatedClientSet("http://your-maas-server", "your-api-key")
    
    ctx := context.Background()
    
    // Allocate a machine
    machine, err := client.Machines().Allocator().Allocate(ctx)
    if err != nil {
        panic(err)
    }
    
    // Deploy with ephemeral mode enabled
    machine, err = machine.Deployer().
        SetOSSystem("ubuntu").
        SetDistroSeries("focal").
        SetEphemeralDeploy(true).  // Enable ephemeral deployment
        Deploy(ctx)
    if err != nil {
        panic(err)
    }
    
    // Machine is now deployed in ephemeral mode
    // All changes will be lost on reboot
}
```

### Ephemeral vs Persistent Deployment

```go
// Ephemeral deployment (changes lost on reboot)
machine, err = machine.Deployer().
    SetOSSystem("ubuntu").
    SetDistroSeries("focal").
    SetEphemeralDeploy(true).
    Deploy(ctx)

// Persistent deployment (changes saved to disk)
machine, err = machine.Deployer().
    SetOSSystem("ubuntu").
    SetDistroSeries("focal").
    SetEphemeralDeploy(false).  // or omit this line (default is false)
    Deploy(ctx)
```

## API Parameters

The ephemeral deployment feature uses the `ephemeral_deploy` parameter:

- `true`: Deploy in ephemeral mode (in-memory, no disk persistence)
- `false`: Deploy in persistent mode (normal disk-based deployment)

## Benefits of Ephemeral Deployment

1. **Clean State**: Every reboot provides a fresh, clean environment
2. **Security**: No data persistence means sensitive information is not stored
3. **Testing**: Ideal for testing scenarios where you need consistent starting conditions
4. **Performance**: Can be faster for read-heavy workloads
5. **Resource Efficiency**: Reduces disk I/O and wear

## Limitations

1. **Data Loss**: All changes are lost on reboot
2. **Storage**: Limited to available RAM for temporary storage
3. **Persistence**: Cannot save files or configurations permanently
4. **Compatibility**: Not all applications work well in ephemeral mode

## Use Cases

- **CI/CD Testing**: Run tests in isolated, clean environments
- **Security Testing**: Penetration testing without leaving traces
- **Development**: Quick prototyping and experimentation
- **Demo Environments**: Presentations with guaranteed clean state
- **Temporary Workloads**: Short-term processing tasks 