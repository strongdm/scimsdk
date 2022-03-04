module github.com/strongdm/scimsdk

go 1.17

replace github.com/strongdm/scimsdk/internal/api v0.0.0 => ./internal/api

replace github.com/strongdm/scimsdk/internal/service v0.0.0 => ./internal/service

require (
	bou.ke/monkey v1.0.2
	github.com/stretchr/testify v1.7.0
	github.com/strongdm/scimsdk/internal/api v0.0.0
	github.com/strongdm/scimsdk/internal/service v0.0.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
