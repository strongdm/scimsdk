module github.com/strongdm/scimsdk/internal/service

go 1.17

replace github.com/strongdm/scimsdk/internal/api v1.0.2 => ../api

require (
	bou.ke/monkey v1.0.2
	github.com/stretchr/testify v1.7.0
	github.com/strongdm/scimsdk/internal/api v1.0.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
