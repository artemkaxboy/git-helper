def round_string(string, max_length=50):
    """
    Cut the string if it is too long and add ... at the end.

    :param string: String to be rounded
    :param max_length: Maximum length of the string
    :return: Rounded string
    """
    return string[:max_length] + "..." if len(string) > max_length else string
