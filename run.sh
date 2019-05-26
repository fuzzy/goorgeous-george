#!/bin/sh
eval $(ssh-agent -s)
ssh-add
export SSH_KNOWN_HOSTS=/root/.ssh/known_hosts
cd /config && /go/bin/goorgeous-george
