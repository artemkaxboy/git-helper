# Git Helper

Handy tool to find unused branches in your git repository

## What is it

It is a tool to simplify work with outdated git branches. It helps to find old branches that are not used anymore or branches made by certain people.

## Installation

`git-helper` is just a regular docker image it may be used **without installation**: 

```shell
docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper:go-main [OPTIONS] [COMMAND]
```

Linux aliases can be used to **make the command shorter**. For only current terminal session:

```shell
alias git-helper='docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper:go-main'
```

Or permanently (re-login required):

```shell
echo "alias git-helper='docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper:go-main'" >> ~/.bash_aliases
```

## Usage

**Use it in directories under version control** (which contains inner `.git` directory).

```bash
Usage:
  git-helper [OPTIONS] <authors | list>

Application Options:
  -d, --git-dir= Path to git repository [$GIT_DIR]
  -f, --filter=  Filter by branch name, last commit author or message [$FILTER]
  -a, --age=     Minimal age of last commit in branch to show, e.g. 1d, 1w, 1m,
                 1y3m (default: 0) [$AGE]
      --dbg      debug mode [$DEBUG]

Help Options:
  -h, --help     Show this help message

Available commands:
  authors  Show authors of branches in repository
  list     List branches in repository
```

### List branches

`list` command shows all remote branches in all remotes in repository.

```shell
Usage:
  git-helper [OPTIONS] list [list-OPTIONS]

Application Options:
  -d, --git-dir=   Path to git repository [$GIT_DIR]
  -f, --filter=    Filter by branch name, last commit author or message [$FILTER]
  -a, --age=       Minimal age of last commit in branch to show, e.g. 1d, 1w, 1m, 1y3m [$AGE]
      --dbg        debug mode [$DEBUG]

Help Options:
  -h, --help       Show this help message

[list command options]
      -s, --short  Short form of the output [$SHORT]
```

Basic usage:

```shell
$ git-helper list
origin:
+---+-----------------+---------------------+----------------------------------------+-------------------+
| # |     Branch      |        Date         |                 Author                 |       Title       |
+---+-----------------+---------------------+----------------------------------------+-------------------+
| 1 | author2-branch1 | 2022-05-07 22:45+07 | Author2 <author2@git-helper.pro>       | Cleans README.md  |
| 2 | author1-branch1 | 2022-05-07 22:40+07 | Author1 <author1@git-helper.pro>       | Updates README.md |
| 3 | author1-branch2 | 2022-05-07 22:40+07 | Author1 <author1@git-helper.pro>       | Updates README.md |
| 4 | main            | 2022-05-07 21:52+07 | Maintainer <maintainer@git-helper.pro> | Init commit       |
+---+-----------------+---------------------+----------------------------------------+-------------------+
```

`--short` option shows only branch names. Can be used with `--filter` option. Output of command shows list of branches and an example of how to delete them from remote.

```shell
$ git-helper list --filter author1 --short
author1-branch1 author1-branch2 

Use `git push origin --delete ...` to delete branches
```

### List authors

`authors` command shows all authors of branches (last commit in each branch) with sum of their branches.

```shell
Usage:
  git-helper [OPTIONS] authors

Application Options:
  -d, --git-dir=  Path to git repository [$GIT_DIR]
  -f, --filter=   Filter by branch name, last commit author or message [$FILTER]
  -a, --age=      Minimal age of last commit in branch to show, e.g. 1d, 1w, 1m, 1y3m [$AGE]
      --dbg       debug mode [$DEBUG]

Help Options:
  -h, --help      Show this help message
```

Basic usage:

```shell
$ git-helper authors
Remote = origin:
+---+----------------------------------------+----------+
| # |                 Author                 | Branches |
+---+----------------------------------------+----------+
| 1 | Author1 <author1@git-helper.pro>       |        2 |
| 2 | Author2 <author2@git-helper.pro>       |        1 |
| 3 | Maintainer <maintainer@git-helper.pro> |        1 |
+---+----------------------------------------+----------+
|                             Sum of 3 items |        4 |
+---+----------------------------------------+----------+
```

### Common options

`--filter` accepts a string to filter by branch name, last commit author or message.

Usage:

```shell
$ git-helper list --filter author1
origin:
+---+-----------------+---------------------+----------------------------------+-------------------+
| # |     Branch      |        Date         |              Author              |       Title       |
+---+-----------------+---------------------+----------------------------------+-------------------+
| 1 | author1-branch1 | 2022-05-07 22:40+07 | Author1 <author1@git-helper.pro> | Updates README.md |
| 2 | author1-branch2 | 2022-05-07 22:40+07 | Author1 <author1@git-helper.pro> | Updates README.md |
+---+-----------------+---------------------+----------------------------------+-------------------+```
```

`--age` accepts a duration string to filter by last commit time. Duration string is a sequence of numbers and a unit of time, e.g. `1d`, `1w`, `1m`, `1y3m`.

Usage:

```shell
$ git-helper list --age 1w1d
origin:
+---+-----------------+---------------------+----------------------------------------+-------------------+
| # |     Branch      |        Date         |                 Author                 |       Title       |
+---+-----------------+---------------------+----------------------------------------+-------------------+
| 1 | main            | 2022-05-07 21:52+07 | Maintainer <maintainer@git-helper.pro> | Init commit       |
+---+-----------------+---------------------+----------------------------------------+-------------------+
```

## License

Apache License 2.0, see [LICENSE](https://github.com/artemkaxboy/git-helper/blob/main/LICENSE).
