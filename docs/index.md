# GoJelastic (Experimental)

An Alternative CLI for Jelastic in Go

## Why rewrite the CLI in Go?

The [official Jelastic command-line interface (CLI)](https://www.virtuozzo.com/application-platform-docs/cli/) is written in Java, which can make it heavy and inconvenient to use in continuous integration/continuous deployment (CI/CD) environments.

Rewriting the CLI in Go will make it lighter and more portable, allowing it to be easily incorporated into CI/CD pipelines and used on a variety of systems without the need for additional dependencies.

This will make it easier for developers to manage and deploy applications on Jelastic, improving efficiency and streamlining the development process.

## How to use

### Installation

#### From binary

* Linux/Darwin

```bash
# With wget
wget -qO- https://raw.githubusercontent.com/yoanbernabeu/GoJelastic/main/install.sh | bash
```

```bash
# With curl
curl -sL https://raw.githubusercontent.com/yoanbernabeu/GoJelastic/main/install.sh | bash
```

* Other Operating Systems

Please download the binary from the [release page](https://github.com/yoanbernabeu/GoJelastic/releases) and move it to your PATH.

#### From source

GoJelastic is written in Go, so you need to install it first.

```bash
git clone git@github.com:yoanbernabeu/GoJelastic.git
cd GoJelastic
go build -o GoJelastic
```

## Documentation

Please check the [documentation](documentation/GoJelastic) for more information.

## License

[MIT](LICENSE)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

![GoJelastic Logo](GoJelastic.png)
