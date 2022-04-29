from datetime import datetime, timedelta, timezone

from git import RemoteReference

from git_helper.objects.author import Author
from git_helper.utils.string_utils import round_string


class Branch:
    """Holds main information about the branch."""

    def __init__(self, name, last_commit_date, last_commit_message, last_commit_author: Author):
        self.name = name
        self.last_commit_date = last_commit_date
        self.last_commit_message = last_commit_message
        self.last_commit_header = last_commit_message.splitlines()[0]
        self.last_commit_author = last_commit_author

    def get_row(self):
        return [self.name, self.last_commit_date, round_string(f"{self.last_commit_author}"),
                round_string(self.last_commit_header)]

    def is_older_than(self, days: int) -> bool:
        """
        Check if this branch is older than given days

        :param days: other branch
        :return: True if this branch is older than the other branch
        """
        return self.last_commit_date < (datetime.now(timezone.utc) - timedelta(days=days))

    @staticmethod
    def get_table_header() -> list:
        """
        Returns the header for a table with rows returned by ```get_row()```

        :return: list of columns
        """
        return ["Name", "Date", "Author", "Message"]

    @staticmethod
    def from_reference(reference: RemoteReference) -> "Branch":
        """
        Create a Branch object from a reference

        :param reference: reference
        :return: Branch object
        """
        return Branch(
            name=reference.remote_head,
            last_commit_date=reference.commit.committed_datetime,
            last_commit_message=reference.commit.message,
            last_commit_author=Author.from_reference(reference.commit)
        )
