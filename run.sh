#!/bin/sh
eval $(ssh-agent -s)
ssh-add
cd /config && /go/bin/goorgeous-george
