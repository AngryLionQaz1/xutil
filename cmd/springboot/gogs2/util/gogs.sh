#!/bin/bash
unset $(git rev-parse --local-env-vars)
Path="/root/test/gogs"
Dir="file"
Jar="file-0.0.1-SNAPSHOT.jar"
Git="root@47.92.213.93:xiaoyiqaz1/file.git"
Arguments="-Xms256m,-Xmx512m"
gogs -p=$Path -d=$Dir -g=$Git -a=$Arguments update