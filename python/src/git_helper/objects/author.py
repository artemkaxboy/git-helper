class Author:
    """Holds main information about the commit author."""

    def __init__(self, name, email):
        self.name = name
        self.email = email

    def __str__(self) -> str:
        return self.get_name_email_string()

    def get_name_email_string(self) -> str:
        """
        Get the name and email of the author as a string

        :return: Pretty string
        """
        return f"{self.name} <{self.email}>"

    @classmethod
    def from_reference(cls, commit) -> "Author":
        """
        Create an Author object from the given commit

        :param commit: Commit to get author from
        :return: Author object
        """
        return cls(commit.author.name, commit.author.email)
