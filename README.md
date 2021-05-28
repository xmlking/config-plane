# config-plane

Push config changes to all subscribing clients

## Requirements

- [ ] Provider source plugins (GitHub, Filesystem, database)
- [ ] Implement as Kubernetes Operator
- [ ] Management Dashboard (view subscribers, providers, current config set)
- [ ] Log config changes  to changelog file
- [ ] expose gRPC API to pull/push/stream config changes

## Client

```go
ctx := context.Background()
client, err := config.NewClient(ctx, "config-plane-url")
if err != nil {
        return fmt.Errorf("config.NewClient: %v", err)
}
defer client.Close()

var mu sync.Mutex
var changelog []string

sub := client.Subscription("/config/**/*")
cctx, cancel := context.WithCancel(ctx)

err = sub.Receive(cctx, func(ctx context.Context, cSet []string) {
        mu.Lock()
        defer mu.Unlock()
        fmt.Fprintf(w, "Got changeset: %v\n", cSet)

        // do something with changes 
        changelog = append(changelog, cSet) 
        // cancel when don't needed
        // cancel()
    })

if err != nil {
    fmt.Errorf("Receive: %v", err)
}

fmt.Fprintf(w, "changelog: %v\n", changelog)
```
