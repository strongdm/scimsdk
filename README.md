# SCIM SDK

SDM SDK for SDM SCIM API.

## Table of Contents

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [Support](#support)

## Installation

```bash
$ go get github.com/strongdm/scimsdk
```

or

```bash
$ git clone https://github.com/strongdm/scimsdk $GOPATH/src/github.com/strongdm/scimsdk
```

Ps.: in this second use, make sure that your `$GOPATH` is set to your `go` folder inside `$HOME` (`$HOME/go`) and have the path `$GOPATH/src/github.com/strongdm/` available.

## Authentication

To use the SDM SCIM API you'll need the SCIM Token. You can get one in the SCIM settings section if you have an organization with overhaul permissions. If you're not, please, contact strongDM support.

Once you have the SCIM API Token, you can use exporting as an environment var:

```bash
$ export SDM_SCIM_TOKEN=<YOUR ADMIN TOKEN>
```

## Getting Started

To get started with the SCIM SDK, you can try the following example scripts:

- Users:
  - [Create](./example/users/create/main.go)
  - [List](./example/users/list/main.go)
  - [Find](./example/users/find/main.go)
  - [Replace](./example/users/replace/main.go)
  - [Update](./example/users/update/main.go)
  - [Delete](./example/users/delete/main.go)
- Groups:
  - [Create](./example/groups/create/main.go)
  - [List](./example/groups/list/main.go)
  - [Find](./example/groups/find/main.go)
  - [Replace](./example/groups/replace/main.go)
  - [Update](./example/groups/update/main.go)
  - [Delete](./example/groups/delete/main.go)

## Contributing

Refer to the [contributing](CONTRIBUTING.md) guidelines or dump part of the information here.

## Support

Refer to the [support](SUPPORT.md) guidelines or dump part of the information here.
