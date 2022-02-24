# SCIM Integrations

SDM SDK for SDM SCIM API.

## Table of Contents

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [Support](#support)

## Installation

```bash
$ go get github.com/strongdm/scim-integrations
```

## Authentication

To use the SDM SCIM API you'll need a SCIM Token. You can get one in the SCIM settings section if you have an organization with overhaul permissions. If you're not, please, contact strongDM support.

Once you have the Admin Token, you can use exporting as an environment var:

```bash
$ export SDM_ADMIN_TOKEN=<YOUR ADMIN TOKEN>
```

## Getting Started

To get started with the SCIM SDK, you can try to use the User Scripts in [example users folder](./example/users/). We have examples for:

- [Create](./example/users/create/main.go)
- [List](./example/users/list/main.go)
- [Find](./example/users/find/main.go)
- [Replace](./example/users/replace/main.go)
- [Delete](./example/users/delete/main.go)

## Contributing

Refer to the [contributing](CONTRIBUTING.md) guidelines or dump part of the information here.

## Support

Refer to the [support](SUPPORT.md) guidelines or dump part of the information here.
