#!/bin/sh
eval $(ssh-agent -s)
ssh-add
ssh git@github.com -T
echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > ~/.ssh/config

export SSH_KNOWN_HOSTS=~/.ssh/known_hosts
cd /config && /go/bin/goorgeous-george
