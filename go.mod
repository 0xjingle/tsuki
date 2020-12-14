module github.com/TsukiCore/tsuki

go 1.14

require (
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/cosmos/cosmos-sdk v0.40.0-rc4
	github.com/go-delve/delve v1.5.1 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
	github.com/iancoleman/strcase v0.1.2
	github.com/otiai10/copy v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.0-rc6
	github.com/tendermint/tm-db v0.6.2
	gitlab.com/c0b/go-ordered-json v0.0.0-20201030195603-febf46534d5a
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
	google.golang.org/genproto v0.0.0-20201014134559-03b6142f0dc9
	google.golang.org/grpc v1.33.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
