package template

// Makefile is the Makefile template used for new projects.
var Makefile = `GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	{{- if .Tern}}
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	{{- end}}
	{{- if .Sqlc}}
	@go install github.com/jackc/tern@latest
	{{- end}}

.PHONY: proto
proto:
	{{- if .OldProtoc}}
	@protoc --proto_path=. --experimental_allow_proto3_optional --micro_out=. --go_out=:. proto/{{.Service}}.proto
	{{- else}}
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/{{.Service}}.proto
	{{- end}}
	{{- if .Health}}
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/health.proto
	{{end}}

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o {{.Service}}{{if .Client}}-client{{end}} *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@{{if .Buildkit}}DOCKER_BUILDKIT=1 {{end}}docker build -t {{.Service}}{{if .Client}}-client{{end}}:latest {{if .PrivateRepo}}--ssh=default {{end}}.

{{- if .Sqlc}}

.PHONY: sqlc
sqlc:
	@sqlc generate -f ./postgres/sqlc.yaml
{{- end -}}
`
