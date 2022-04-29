"""Minimal setup file for tasks project."""

from setuptools import setup, find_packages

setup(
    name="git-helper",
    version="0.1.0",
    license="Apache-2.0",
    description="Git helper",

    author="Artem Kolin",
    author_email="Please use github account to contact.",
    url="https://github.com/artemkaxboy/git-helper",

    packages=find_packages(where="src"),
    package_dir={"": "src"},
    python_requires=">=3.8",

    # install_requires=["click==7.1.2", "tinydb==3.15.1", "six", "pytest", "pytest-mock"],
    install_requires=["GitPython==3.1.27", "tabulate==0.8.9", "argparse==1.4.0"],
    # tests_require=["pytest", "pytest-mock"],

    entry_points={
        "console_scripts": [
            "git-helper=git_helper.cli:main",
        ]
    },
)
