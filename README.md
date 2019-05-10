# go-docker-containerid
Fetch a docker container ID from inside a Docker container.

Uses the `/proc/self/cgroup` to identify the container ID and returns it as a string 

## Usage
Import the library
```
import "github.com/discobean/go-docker-containerid/containerid"
```

Then use it like so
```
id, err := containerid.Find()

if err != nil {
  log.Fatalf("Failed to get container ID: %v\n", err)
}

log.Printf("Got Container ID: %s\n", id)
```
