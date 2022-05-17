# Git Helper

Handy tool to find unused branches in your git repository

## What is it

It is a tool to simplify work with outdated git branches. It helps to find old branches that are not used anymore or branches made by certain people.

## Usage

**Use it in directories under version control**.

```bash
Usage:
  git-helper [OPTIONS] <authors | list>

Application Options:
      --git-dir= The directory where the git repositories are stored [$GIT_DIR]
      --dbg      debug mode [$DEBUG]

Help Options:
  -h, --help     Show this help message

Available commands:
  authors  Show authors of branches in repository
  list     List branches in repository
```

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

### List branches

```shell
Usage:
  git-helper [OPTIONS] list [list-OPTIONS]

Application Options:
      --git-dir=     The directory where the git repositories are stored
                     [$GIT_DIR]
      --dbg          debug mode [$DEBUG]

Help Options:
  -h, --help         Show this help message

[list command options]
      -d, --git-dir= Path to git repository [$GIT_DIR]
      -f, --filter=  Filter by branch name, last commit author or message
                     [$FILTER]
      -s, --short    Short form of the output [$SHORT]
```

Shows all remote branches in all remotes in repository.

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

`--filter` accepts a string to filter by branch name, last commit author or message.

Example:

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

`--short` option shows only branch names. Can be used with `--filter` option. Output of command shows list of branches and an example of how to delete them from remote.

```shell
$ git-helper list --filter author1 --short
author1-branch1 author1-branch2 

Use `git push origin --delete ...` to delete branches
```

### List authors

```shell
Usage:
  git-helper [OPTIONS] authors [authors-OPTIONS]

Application Options:
      --git-dir=     The directory where the git repositories are stored
                     [$GIT_DIR]
      --dbg          debug mode [$DEBUG]

Help Options:
  -h, --help         Show this help message

[authors command options]
      -d, --git-dir= Path to git repository [$GIT_DIR]
      -f, --filter=  Filter by branch name, last commit author or message
                     [$FILTER]
```

Shows all authors of branches (last commit in each branch) with sum of their branches.

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

`--filter` accepts a string to filter by branch name, last commit author or message.

```shell
$ git-helper authors --filter maintainer
Remote = origin:
+---+----------------------------------------+----------+
| # |                 Author                 | Branches |
+---+----------------------------------------+----------+
| 1 | Maintainer <maintainer@git-helper.pro> |        1 |
+---+----------------------------------------+----------+
|                             Sum of 1 items |        1 |
+---+----------------------------------------+----------+
```

## Alias usage

After adding alias you can use short form:

```shell
# long
docker run --rm -it -v $(pwd):/data artemkaxboy/git-helper:go-main [OPTIONS] [COMMAND]
# short
git-helper [OPTIONS] [COMMAND]
```

## License

Apache License 2.0, see [LICENSE](https://github.com/artemkaxboy/git-helper/blob/main/LICENSE).
