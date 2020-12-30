build:
		GOOS=linux GOARCH=amd64 go build -o deployments/agent-service main.go
		docker build -t agent-service deployments
proto:
		protoc -I proto/ proto/agent.proto --go_out=plugins=grpc:proto/
run:
		docker run -p 50006:50006 --name agent-service --hostname agent-service@iptv -v /etc/hosts:/etc/hosts -v /opt/iptv/:/opt/iptv/ --log-driver json-file --log-opt max-size=1m --log-opt max-file=10 -d agent-service

docs:
		protoc --doc_out=./doc  proto/*.proto
