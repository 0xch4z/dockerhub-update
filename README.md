# Dockerhub update [![Build Status][travis-ci-badge]][travis-ci]

> Update your Dockerhub repository

## Usage

```sh
$ dockerhub-update

NAME:
   dockerhub-update - Update your Dockerhub repository

USAGE:
   dockerhub-update [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --username value, -u value     Dockerhub username [$DOCKERHUB_USERNAME]
   --password value, -p value     Dockerhub password [$DOCKERHUB_PASSWORD]
   --token value, -t value        Dockerhub API token [$DOCKERHUB_API_TOKEN]
   --readme value, -r value       Path to readme update [$README_PATH]
   --description value, -d value  Description update [$README_PATH]
   --help, -h                     show help (default: false)
```

### Example

```sh
dockerhub-update -u myusername -p ********** \
  -r ./myreadmefile.md -d "My description" \
  myusername/myrepo
```

Updates description and readme file (to locally specified file) for repository located at `hub.docker.com/r/myusername/myrepo`.

### Via docker

```sh
docker run -e DOCKERHUB_USERNAME -e DOCKERHUB_PASSWORD \
  -v $(PWD)/myreadmefile.md:/data/myreadmefile.md:ro \
  charliekenney23/dockerhub-update \
  -d "My description" -r /data/myreadmefile.md \
  myusername/myrepo
```

## License

Copyright &copy; 2019 [Charles Kenney][me]

This project is [MIT][license] licensed.

[me]: https://github.com/charliekenney23
[license]: https://github.com/charliekenney23/dockerhub-update/tree/master/LICENSE
[travis-ci-badge]: https://travis-ci.org/Charliekenney23/dockerhub-go.svg?branch=master
[travis-ci]: https://travis-ci.org/Charliekenney23/dockerhub-go
