# Civitai-Cli

A Go Client to batch download models from CivitAI with CivitAI Api V1

> NOTE:please see [API Key Page](/doc/api_key.md) for instructions

## Getting Started

### Installing

> NOTE: You need to have Go installed on your machine to install this package
> You can download Go from [here](https://go.dev/dl/)
> Preferably Go 1.18 and above

You can also install the package using the following command

```bash
go install github.com/zhaolion/civitai-cli@latest
```

<br>

### Quick Start

To get started quickly, copy the command below.

Set the `CIVITAI_API_KEY` environment variable to your CivitAI API Key
```bash
export CIVITAI_API_KEY=<Your API Token Here>
```

or you copy the command below.

```bash
civitai-cli api set_api_token <Your API Token Here>
```

<br/>

### Usage

Explore the various functionalities provided by CivitAI-CLI:

```bash
civitai-cli --help
```

> NOTE: You can also explore the various subcommands provided by CivitAI-CLI by running the command below:

```bash
civitai-cli is a simple Go client to batch download models from CivitAI with CivitAI Api V1.

Usage:
  civitai-cli [command]

Available Commands:
  api         help you to interact with CivitAI.
  completion  Generate the autocompletion script for the specified shell
  download    download files from CivitAI
  help        Help about any command

Flags:
  -h, --help   help for civitai-cli

Use "civitai-cli [command] --help" for more information about a command.
```

#### Download

> model_id: The model id of the model you want to download
> eg. https://civitai.com/models/352581, the model id is 352581
> ---
> version_id: The version id of the model you want to download
> eg. https://civitai.com/models/352581?modelVersionId=647401, the version id is 647401

To download models from CivitAI, you can use the `download` subcommand.

##### Cmd - Download model

**will download all files from the model with the given model id to the current directory**
```bash
civitai-cli download model --mid <model id> --dir </path/to/download>
```

or just

```bash
civitai-cli download model --mid <model id>
```

##### Cmd - Download model version

if you want to download one version of the model, you can use:

```bash
civitai-cli download model_ver --mid <model id> --vid <version id>
```

#### API

To interact with CivitAI API, you can use the `api` subcommand.

##### Cmd - Set API Token

set api token for authentication.

```bash
civitai-cli api set_api_token <Your API Token Here>
```

##### Cmd - View API Token

view which api token is currently set.

```bash
civitai-cli api view_api_token
```

##### Cmd - View Model Info

view model info by model id

```bash
civitai-cli api model --mid <model id>
```

will return the model info in terminal:

```
| ID | Name                | Type       | Creator | Stat          | Versions                                          |
|----|---------------------|------------|---------|---------------|---------------------------------------------------|
| 1  | Superhero Diffusion | Checkpoint |         | download:2684 | [1][V1] https://civitai.com/api/download/models/1 |
|    |                     |            |         | thumbsUp:444  |                                                   |
```

##### Cmd - View Model Version Info

view model's version info by version id.

```bash
civitai-cli api model_ver --vid <version id>
```

will return the version info in terminal:
```
| ID | Name | ModelID | ModelType  | ModelName           | BaseModel | Stat                       | Size   | FileURL                                   |
|----|------|---------|------------|---------------------|-----------|----------------------------|--------|-------------------------------------------|
| 1  | V1   | 1       | Checkpoint | Superhero Diffusion | SD 1.5    | download:2684 thumbsUp:444 | 2.1 GB | https://civitai.com/api/download/models/1 |
```


## Contributing

Thanks for the interest in the project!

Please create an issue if you encounter any problem, bugs or if you have a feature request.

To debug things, it is recommended to run with `--debug` option.
* Running in debug allows users to print tracebacks and other messages useful for debugging.
* Example: `civitai-cli api model --mid 1 --debug`

To work on an issue:
* Please create a fork.
* Then clone your fork locally.
* Then create a local branch that describes the issue.
* Once you have commited your changes, push the branch to your forked repository.
* Then open a pull request to this repository.

<br/>

## License

This project is licensed under the MIT License - see the [LICENSE.md](./License) file for details
