#!/bin/bash

sudo rm -rf /tmp/rootfs
mkdir -p /tmp/rootfs/etc/zypp/repos.d/
cp /etc/zypp/repos.d/openSUSE-Leap-15.2-1.repo /tmp/rootfs/etc/zypp/repos.d
tree /tmp/rootfs
