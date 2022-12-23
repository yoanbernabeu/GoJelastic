# GoJelastic (Experimental)

An Alternative CLI for Jelastic in Go

## Why rewrite the CLI in Go?

The official Jelastic command-line interface (CLI) is written in Java, which can make it heavy and inconvenient to use in continuous integration/continuous deployment (CI/CD) environments.

Rewriting the CLI in Go will make it lighter and more portable, allowing it to be easily incorporated into CI/CD pipelines and used on a variety of systems without the need for additional dependencies.

This will make it easier for developers to manage and deploy applications on Jelastic, improving efficiency and streamlining the development process.

## How to use

### Installation

#### From binary

* Debian/Ubuntu

```bash
wget https://github.com/yoanbernabeu/GoJelastic/releases/download/v0.1.2/GoJelastic-0.1.2-linux-amd64.tar.gz
tar -xvzf GoJelastic-0.1.2-linux-amd64.tar.gz
sudo mv GoJelastic /usr/local/bin/
sudo chmod +x /usr/local/bin/GoJelastic
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

## Usage

### Variables definition

* APPID: The application ID
* TOKEN: Get your token [here](https://www.virtuozzo.com/application-platform-ops-docs/platform-access-token/)
* URL: Your Jelastic API URL
* NODEID: The unique ID of the node/container

### Configure Environment into a file

If you want to use the CLI without passing the token and the URL each time, you can configure it into a file (stored in your home directory in a file named `gojelastic.env`)


```bash
GoJelastic configure --token <TOKEN> --url <URL>
```

### Get Environment Info

Get informations about one environment

```bash
GoJelastic getEnv --appid <APPID>
```

### Get Environments Info

Get informations about all environments

```bash
GoJelastic getEnvs
```

### Start Environment

Start an environment

```bash
GoJelastic startEnv --appid <APPID>
```

### Stop Environment

Stop an environment

```bash
GoJelastic stopEnv --appid <APPID>
```

### Redeploy Container by ID

Redeploy a container by ID and target tag

```bash
GoJelastic redeployEnv --nodeid <NODEID> --tag <TAG> --appid <APPID>
```

## License

[MIT](LICENSE)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

![GoJelastic Logo](GoJelastic.png)
