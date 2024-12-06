# Fibonacci App
## Commands
**Build**
```bash
$ go build -o bin/serverapp main.go
```

**Compile the container image**
```bash
$ podman build -f Dockerfile -t myserver:test
```

**Run tests**
```bash
$ go test -v ./...
```

**Run the container locally**
```bash
podman run -it --rm -p 8080:8080 localhost/myserver:test
```
Once running, the server will log to the command line and an interactive app will be available at `localhost:8080`