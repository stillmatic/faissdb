module faissdb

go 1.15

replace github.com/crumbjp/faissdb/server => ./

replace local.packages/go-faiss => ../go-faiss

require (
	github.com/crumbjp/faissdb/server v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.1.2
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/sevlyar/go-daemon v0.1.5
	github.com/stretchr/testify v1.5.1
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c
	golang.org/x/net v0.0.0-20210331212208-0fccb6fa2b5c
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
	local.packages/go-faiss v0.0.0-00010101000000-000000000000
)
