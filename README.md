# Stencil

![Test](https://github.com/raystack/stencil/actions/workflows/test-server.yaml/badge.svg)
![Release](https://github.com/raystack/stencil/actions/workflows/release-server.yml/badge.svg)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?logo=apache)](LICENSE)
[![Version](https://img.shields.io/github/v/release/raystack/stencil?logo=semantic-release)](Version)

Stencil is a schema registry that provides schema mangement and validation to ensure data compatibility across applications. It enables developers to create, manage and consume schemas dynamically, efficiently, and reliably, and provides a simple way to validate data against those schemas. Stencil support multiple formats including Protobuf, Avro and JSON.

<p align="center"><img src="./docs/static/assets/overview.svg" /></p>

## Key Features

Discover why users choose Stencil as their main schema registry

- **Version history** Stencil stores versioned history of proto descriptor file on specified namespace and name
- **Backward compatibility** enforce backward compatibility check on upload by default
- **Flexbility** ability to skip some of the backward compatibility checks while upload
- **Descriptor fetch** ability to download proto descriptor files
- **Metadata** provides metadata API to retrieve latest version number given a name and namespace
- **Clients in multiple languages** Stencil provides clients in GO, Java, JS languages to interact with Stencil server and deserialize messages using dynamic schema

## Documentation

Explore the following resources to get started with Stencil:

- [Documentation](http://raystack.github.io/stencil) provides guidance on using stencil.
- [Server](https://raystack.github.io/stencil/docs/server/overview) provides details on getting started with stencil server.
- [Clients](https://raystack.github.io/stencil/docs/clients/overview) provides reference to supported stencil clients.

## Installation

Install Stencil on macOS, Windows, Linux, OpenBSD, FreeBSD, and on any machine.

#### Binary (Cross-platform)

Download the appropriate version for your platform from [releases](https://github.com/raystack/stencil/releases) page. Once downloaded, the binary can be run from anywhere.
You don’t need to install it into a global location. This works well for shared hosts and other systems where you don’t have a privileged account.
Ideally, you should install it somewhere in your PATH for easy use. `/usr/local/bin` is the most probable location.

#### MacOS

`stencil` is available via a Homebrew Tap, and as downloadable binary from the [releases](https://github.com/raystack/stencil/releases/latest) page:

```sh
brew install raystack/tap/stencil
```

To upgrade to the latest version:

```sh
brew upgrade stencil
```

#### Linux

`stencil` is available as downloadable binaries from the [releases](https://github.com/raystack/stencil/releases/latest) page. Download the `.deb` or `.rpm` from the releases page and install with `sudo dpkg -i` and `sudo rpm -i` respectively.

#### Windows

`stencil` is available via [scoop](https://scoop.sh/), and as a downloadable binary from the [releases](https://github.com/raystack/stencil/releases/latest) page:

```sh
scoop bucket add stencil https://github.com/raystack/scoop-bucket.git
```

To upgrade to the latest version:

```sh
scoop update stencil
```

#### Docker

We provide ready to use Docker container images. To pull the latest image:

```sh
docker pull raystack/stencil:latest
```

To pull a specific version:

```sh
docker pull raystack/stencil:v0.4.1
```

## Usage

Stencil has three major components. Server, CLI and clients. Stencil server and CLI are bundled in a single binary.

**Server**

Stencil server provides a way to store and fetch schemas and enforce compatibility rules. Run `stencil server --help` to see instructions to manage Stencil server.

Stencil server also provides a fully-featured GRPC and HTTP API to interact with Stencil server. Both APIs adheres to a set of standards that are rigidly followed. Please refer to [proton](https://github.com/raystack/proton/tree/main/raystack/stencil/v1beta1) for GRPC API definitions.

**CLI**

Stencil CLI allows users to iteract with server to create, view, and search schemas. CLI is fully featured but simple to use, even for those who have very limited experience working from the command line. Run `stencil --help` to see list of all available commands and instructions to use.

**Clients**

Stencil clients allows application to interact with stencil server to eserialize and deserialize messages using schema. Stencil supports clients in multiple languages.

- [Java](clients/java)
- [Go](clients/go)
- [Javascript](clients/js)
- [Clojure](clients/clojure)
- Ruby - Coming soon
- Python - Coming soon

## Running locally

<details>
  <summary>Dependencies:</summary>

    - Git
    - Go 1.16 or above
    - Yarn (Needed for UI)
    - PostgreSQL 13 or above

</details>

```sh
# Clone the repo
$ git clone git@github.com:raystack/stencil.git

# Check all build comamnds available
$ make help

# Build meteor binary file
$ make build

# Init server config
$ cp config/config.yaml config.yaml

# Run database migrations
$ ./stencil server migrate

# Start stencil server
$ ./stencil server start
```

## Running tests

```sh
# Running all unit tests
$ make test

# Print code coverage
$ make coverage
```

## Contribute

Development of Stencil happens in the open on GitHub, and we are grateful to the community for contributing bugfixes and improvements. Read below to learn how you can take part in improving stencil.

Read our [contributing guide](docs/docs/contribute/contribution.md) to learn about our development process, how to propose bugfixes and improvements, and how to build and test your changes to stencil.

To help you get your feet wet and get you familiar with our contribution process, we have a list of [good first issues](https://github.com/raystack/stencil/labels/good%20first%20issue) that contain bugs which have a relatively limited scope. This is a great place to get started.

This project exists thanks to all the [contributors](https://github.com/raystack/stencil/graphs/contributors).
