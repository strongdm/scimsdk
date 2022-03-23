module github.com/strongdm/scimsdk

go 1.17

replace github.com/strongdm/scimsdk/internal/api v1.0.2 => ./internal/api

replace github.com/strongdm/scimsdk/internal/service v1.0.2 => ./internal/service

require (
	bou.ke/monkey v1.0.2
	github.com/getsentry/sentry-go v0.13.0
	github.com/stretchr/testify v1.7.0
	github.com/strongdm/scimsdk/internal/api v1.0.2
	github.com/strongdm/scimsdk/internal/service v1.0.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
