# Capture-go

### Description
1. Takes screenshots of given urls
2. Resizes .png .jpeg .jpg images

### How to use
#### Page capture (screenshots)
1. Get request /capture/url?url=<yourpageurl>
    Available query params:
    ```
    width
    height
    device
    ```
    Device should be valid device from https://github.com/GoogleChrome/puppeteer/blob/master/DeviceDescriptors.js
    If device parameter is set width and height will be ignored

### Build dev image
```shell
docker build -f Dockerfile-dev -t capture:go .  
```
### Start development enviroment
```shell
docker run -v $(pwd):/app -v $GOPATH/pkg/mod:/go/pkg/mod -p 8887:8887 capture:go refresh start
```
