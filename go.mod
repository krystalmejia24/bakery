module github.com/cbsinteractive/bakery

replace github.com/zencoder/go-dash/v3 => github.com/cbsinteractive/go-dash/v3 v3.0.1-0.20210815181819-c356e550b00d

replace github.com/grafov/m3u8 => github.com/cbsinteractive/m3u8 v0.11.2-0.20210702182805-556854f1e40f

go 1.16

require (
	github.com/aws/aws-sdk-go v1.38.64 // indirect
	github.com/cbsinteractive/pkg/tracing v0.0.0-20210104155054-0951c16e08d6
	github.com/cbsinteractive/pkg/xrayutil v0.0.0-20210104155054-0951c16e08d6
	github.com/cbsinteractive/propeller-go v0.0.0-20200828205819-d02e4f54c5bc
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.6
	github.com/grafov/m3u8 v0.11.1
	github.com/justinas/alice v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/zerolog v1.23.0
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/zencoder/go-dash/v3 v3.0.1
)
