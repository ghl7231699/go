#!/bin/bash
#startup.sh 会和编译出来的可执行文件在同一级
##############################
# /home/www/app or /home/www
# | startup.sh
# | app (可执行文件)
# | config
# |		| app.yaml
# | logs
##############################
exec ./app -c config/app.yaml