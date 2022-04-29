#!/usr/bin/python3

import argparse

from git import Remote, RemoteReference

from git_helper.objects.author import Author
from git_helper.objects.branch import Branch
from git_helper.utils.cli_utils import print_error, print_table
from git_helper.utils.collection_utils import sort_dict_by_value
from git_helper.utils.package_utils import get_version_from_package
from git_helper.utils.repo_utils import get_remote

SHORT = False


def print_all_branches(remote: Remote) -> None:
    """
    Print all branches info

    :param remote: Optional remote to get branches from
    :return: None
    """
    branches = find_all_branches(remote)
    print_branches(branches)


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


def count_commits_by_authors(remote: Remote) -> dict:
    """
    Count commits by authors. Divide authors by their name and email combination

    :param remote: Remote to count commits from
    :return: Dict of authors and their commits count
    """
    commits_count_by_author = {}

    for remote_ref in remote.refs:
        author = str(Author(remote_ref.commit.author.name, remote_ref.commit.author.email))
        if author not in commits_count_by_author:
            commits_count_by_author[author] = 0
        commits_count_by_author[author] += 1

    return commits_count_by_author


def print_all_authors_sum(remote: Remote) -> None:
    """
    Print all authors info

    :param remote: Remote to count commits from
    :return: None
    """
    commits_by_authors = sort_dict_by_value(count_commits_by_authors(remote))

    print_table(commits_by_authors, ["Author", "Commits"])


def find_branches_by_author(remote: Remote, search_filter: str):
    """
    Find branches by author. Compare authors by their name and email part

    :param remote: Remote to find branches from
    :param search_filter: Author name or email part
    :return: List of branches
    """
    all_branches = find_all_branches(remote)

    return [branch for branch in all_branches if search_filter in str(branch.last_commit_author)]


def print_branch_names_of_author(remote: Remote, author: str):
    """
    Print branches by author. Compare authors by their name and email part

    :param remote: Remote to find branches from
    :param author: Author name or email part
    :return: None
    """
    print("Branches of author: " + author)

    branches = find_branches_by_author(remote, author)
    print_branches(branches)


def find_branches_older_than(remote: Remote, days: int) -> list:
    """
    Find branches older than given days

    :param remote: Remote to find branches from
    :param days: Days to compare with
    :return: List of branches
    """
    all_branches = find_all_branches(remote)

    return [branch for branch in all_branches if branch.is_older_than(days)]


def print_old_branches(remote: Remote, days: int):
    """
    Print branches older than given days

    :param remote: Remote to find branches from
    :param days: Days to compare with
    :return: None
    """
    print(f"Branches older than (days): {days}")

    branches = find_branches_older_than(remote, days)
    print_branches(branches)


def print_branches(branches: list):
    branches.sort(key=lambda x: x.last_commit_date, reverse=True)
    if SHORT:
        print_branches_list(branches)
    else:
        print_branches_table(branches)

    print("\n\nUse `git push origin --delete ...` to delete branches")


def print_branches_table(branches: list):
    table = [branch.get_row() for branch in branches]
    print_table(table, Branch.get_table_header())


def print_branches_list(branches: list):
    for branch in branches:
        print(branch.name, end=" ")
    print()


def print_branch_info(branch: RemoteReference):
    print(branch.name)


def main():
    global SHORT

    parser = argparse.ArgumentParser(prog="git-helper")
    parser.add_argument('-v', '--version', action='version', help="show program version",
                        version=f"Version: {get_version_from_package()}")
    parser.add_argument('-b', '--branches', action='store_true', help="show repository branches")
    parser.add_argument('-a', '--authors', '-u', '--users', action='store_true', help="show repository authors")
    parser.add_argument('-f', '--filter', nargs=1, help="show commits of author (partly or full name or email)",
                        metavar="NAME")
    parser.add_argument('-s', '--short', action='store_true', help="print branch names as string instead of table")
    parser.add_argument('-o', '--old', nargs='?', help="print old branches", metavar="DAYS", type=int)

    args = parser.parse_args()
    SHORT = args.short

    try:
        remote = get_remote()

        if args.branches:
            print_all_branches(remote)

        elif args.authors:
            print_all_authors_sum(remote)

        elif args.old:
            print_old_branches(remote, args.old)

        elif args.filter:
            print_branch_names_of_author(remote, args.filter[0])

        else:
            print_error("Unknown options. Use -h or --help for help.")

    except Exception as e:
        # raise e  # for testing
        print_error(f"[{type(e).__name__}]: {e}")
        exit(1)


if __name__ == '__main__':
    main()
