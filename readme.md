# Capture-go
1. Takes screenshots of given urls
2. Resizes .png .jpeg .jpg images

#### Development
```shell
docker run -v $(pwd):/app -v $GOPATH/pkg/mod:/go/pkg/mod -p 8887:8887 capture:go refresh start
```
