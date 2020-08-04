#!/usr/bin/python
###########################################################################
# Description:
#    Setup file to install the generated python protobuf files
#    as a python project
# Copyright (c) 2018 Nokia
###########################################################################
from setuptools import setup, find_packages

setup(
    name='sdkmgr-proto',
    version="0.1",
    description="Protobuf libraries for srlinux sdkmgr",
    author="",
    packages=['sdk_protos'],
    include_package_data=True,
)
