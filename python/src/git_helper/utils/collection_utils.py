def sort_dict_by_value(dictionary) -> list:
    """
    Sorts a dictionary by its values

    :param dictionary: dictionary to sort
    :return: list of tuples (key, value) sorted by value descending
    """
    return sorted(dictionary.items(), key=lambda x: x[1], reverse=True)
