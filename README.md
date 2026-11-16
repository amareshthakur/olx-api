not a production approach ->
PORT=9999 make run
<-

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o bin/api .cmd/api