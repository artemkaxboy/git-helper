# Git Helper

Handy tool to find unused branches in your git repository

## What is it

It is a tool to simplify work with outdated git branches. It helps to find old branches that are not used anymore or branches made by certain people.

## Usage

**Use it in directories under version control**.

```bash
usage: git-helper [-h] [-v] [-b] [-a] [-f NAME] [-s] [-o [DAYS]]

optional arguments:
  -h, --help            show this help message and exit
  -v, --version         show program version
  -b, --branches        show repository branches
  -a, --authors, -u, --users
                        show repository authors
  -f NAME, --filter NAME
                        show commits of author (partly or full name or email)
  -s, --short           print branch names as string instead of table
  -o [DAYS], --old [DAYS]
                        print old branches
```

## Installation

`git-helper` is just a regular docker image it may be used **without installation**: 

```shell
docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper [COMMAND] [PARAMS]
```

Linux aliases can be used to **make the command shorter**. For only current terminal session:

```shell
alias opener='docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper'
```

Or permanently (re-login required):

```shell
echo "alias git-helper='docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper'" >> ~/.bash_aliases
```

### Alias usage

After adding alias you can use short form:

```shell
# long
docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper [COMMAND] [PARAMS]
# short
git-helper [COMMAND] [PARAMS]
```

## License

Apache License 2.0, see [LICENSE](https://github.com/artemkaxboy/docker-opener/blob/main/LICENSE).
