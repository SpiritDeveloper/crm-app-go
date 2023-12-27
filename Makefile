check-swagger:
	which swagger || (go install -mod=mod github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor && swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml

check-compile-daemon:
	GO111MODULE=on go mod vendor && (go install -mod=mod github.com/githubnemo/CompileDaemon)

run-dev: check-compile-daemon
	CompileDaemon -command="./crm-app-go || crm-app-go.exe"


