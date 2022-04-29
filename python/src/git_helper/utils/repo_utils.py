import os

from git import Repo, Remote

from git_helper.objects.branch import Branch


def find_repo(directory: str = os.getcwd(), search_parent_directories: bool = True) -> Repo:
    """
    Find repo in the given directory and all parent directories

    :param directory: Directory to search in (default: current directory)
    :param search_parent_directories: Search for repo in parent directories (default: True)
    :return: Repo object
    """
    repo = Repo(directory, search_parent_directories=search_parent_directories)
    return repo


def get_remote(repo: Repo = find_repo()) -> Remote:
    """
    Get remote from the given repo or find it in the current directory

    :param repo: Optional repo to find remote from
    :return: Remote object
    """
    remote = repo.remote()
    return remote


def find_all_branches(remote: Remote) -> list:
    """
    Find all branches from given remote

    :param remote: Remote to find branches from
    :return: list of Branch objects
    """
    branches = []

    for remote_ref in remote.refs:
        branches.append(Branch.from_reference(remote_ref))

    return branches
