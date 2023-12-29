# How to deploy

## run.sh

put this file to `/mochi/server/` dir.

## slash-mochi-app.service

1. move this file to `/lib/systemd/system/` dir in product env.
1. make a symbolic link to `/etc/systemd/system/multi-user.target.wants/` dir that link to slash-mochi-app.service file.