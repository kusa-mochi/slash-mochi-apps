#!/bin/bash

# run this script as super user.

LOG_DIR=/var/slash-mochi
LOG_FILENAME=log.txt

if [ -d $LOG_DIR ]; then
    rm -rf $LOG_DIR
fi
mkdir -p $LOG_DIR

chmod +x ./slash_mochi_server
./slash_mochi_server -LogPath ${LOG_DIR}${LOG_FILENAME}
