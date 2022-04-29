import sys
from typing import Union

from tabulate import tabulate


def print_error(message: any):
    """
    Print an error message to stderr

    :param message: message to print
    :return: None
    """
    print("Error: " + str(message), file=sys.stderr)


def print_table(table: list, header: Union[list, None] = None):
    """
    Print a table to stdout

    :param table: table to print
    :param header: optional header list
    :return: None
    """
    if header is None:
        header = []

    print(tabulate(table, headers=header, tablefmt='orgtbl'))
    print(f"\nTotal elements: {len(table)}")
