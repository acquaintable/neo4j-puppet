#!/bin/sh
mkfs.ext3 /dev/xvdj -F
mkdir /backup
chmod 0755 /backup
mkdir /mnt/data
chown -R neo4j:nogroup /mnt/data/
