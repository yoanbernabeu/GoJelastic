# GoJelastic (Experimental)

A Alternative CLI for Jelastic in Go

## Why rewrite the CLI in Go?

The official Jelastic command-line interface (CLI) is written in Java, which can make it heavy and inconvenient to use in continuous integration/continuous deployment (CI/CD) environments.

Rewriting the CLI in Go will make it lighter and more portable, allowing it to be easily incorporated into CI/CD pipelines and used on a variety of systems without the need for additional dependencies.

This will make it easier for developers to manage and deploy applications on Jelastic, improving efficiency and streamlining the development process.

## How to use

### Installation

#### From source

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

### Get Environment Info

Get informations about one environment

```bash
./GoJelastic getEnv --appid <APPID> --token <TOKEN> --env <ENV> --url <URL>
```

### Get Environments Info

Get informations about all environments

```bash
./GoJelastic getEnvs --token <TOKEN> --env <ENV> --url <URL>
```

### Start Environment

Start an environment

```bash
./GoJelastic startEnv --appid <APPID> --token <TOKEN> --env <ENV> --url <URL>
```

### Stop Environment

Stop an environment

```bash
./GoJelastic stopEnv --appid <APPID> --token <TOKEN> --env <ENV> --url <URL>
```

## License

[MIT](LICENSE)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

![GoJelastic Logo](GoJelastic.png)
