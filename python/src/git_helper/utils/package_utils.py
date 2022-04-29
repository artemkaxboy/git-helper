import pkg_resources


def get_version_from_package(package_name='git_helper'):
    """
    Get the version of a package from its name.
    """
    return pkg_resources.get_distribution(package_name).version
