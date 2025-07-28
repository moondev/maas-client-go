# maas-client-go
MAAS client for GO

## Usage

### Basic Usage

```go
c := NewAuthenticatedClientSet(os.Getenv("MAAS_ENDPOINT"), os.Getenv("MAAS_API_KEY"))

ctx := context.Background()

// List DNS Resources
res, err := c.DNSResources().List(ctx, nil)

// List DNS Resources filtered by fqdn
filters := ParamsBuilder().Add(FQDNKey, "bad-doesntexist.maas")
res, err := c.DNSResources().List(ctx, filters)

// Create DNS Resource
res, err := c.DNSResources().
	Builder().
	WithFQDN("test-unit1.maas.sc").
	WithAddressTTL("10").Create(ctx)

// Update DNS Resource
err = res.Modifier().
	SetIPAddresses([]string{"1.2.3.4", "5.6.7.8"}).
	Modify(ctx)

// Get DNS Resource by ID
res2 := c.DNSResources().DNSResource(res.ID())

// Delete DNS Resource
err = res.Delete(ctx)
```

### Machine Deployment

#### Standard Deployment (Persistent)

```go
// Allocate a machine
machine, err := c.Machines().Allocator().Allocate(ctx)

// Deploy with persistent storage (default)
machine, err = machine.Deployer().
	SetOSSystem("ubuntu").
	SetDistroSeries("focal").
	Deploy(ctx)
```

#### Ephemeral Deployment (In-Memory)

```go
// Allocate a machine
machine, err := c.Machines().Allocator().Allocate(ctx)

// Deploy in ephemeral mode (in-memory, no disk persistence)
machine, err = machine.Deployer().
	SetOSSystem("ubuntu").
	SetDistroSeries("focal").
	SetEphemeralDeploy(true).  // Enable ephemeral deployment
	Deploy(ctx)
```

Ephemeral deployment runs the operating system entirely in RAM without persisting changes to disk. This is useful for:
- Testing environments where you need a clean state
- Security testing without leaving traces
- Temporary workloads
- Development and prototyping

For more details, see [EPHEMERAL_DEPLOYMENT.md](EPHEMERAL_DEPLOYMENT.md).
