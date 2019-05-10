#!/bin/bash
unset $(git rev-parse --local-env-vars)
Path="/root/test/gogs"
Dir="file"
Jar="file-0.0.1-SNAPSHOT.jar"
Actuator="http://47.92.213.93:9082/actuator/shutdown"
Git="root@47.92.213.93:xiaoyiqaz1/file.git"
Arguments="-Xms256m,-Xmx512m"
gogs -p=$Path -d=$Dir -g=$Git -j=$Jar -ac=$Actuator -a=$Arguments update