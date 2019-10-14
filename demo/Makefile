godev:
	GOPROXY=off go run main.go

check:
	GOPROXY=off go fmt ./...
	GOPROXY=off go test ./... -cover

coverprofile:
	GOPROXY=off go test ./... -coverprofile=c.out && GOPROXY=off go tool cover -html=c.out

linux:
	GOPROXY=off GOOS=linux GOARCH=amd64 go build -o bin/iptool main.go

mac:
	GOPROXY=off go build -o bin/iptool main.go

deploy:
	GOPROXY=off GOOS=linux GOARCH=amd64 go build -o bin/server main.go
	cd bin && tar czvf iptool.tar.gz server
	scp bin/iptool.tar.gz root@47.92.115.56:/web/www/iptool/bin/iptool.tar.gz
	ssh root@47.92.115.56 "cd /web/www/iptool/bin && mv server server2 && tar -xvf iptool.tar.gz && sudo supervisorctl restart iptool"
