#!/bin/sh
eval $(ssh-agent -s)
cd /config && /go/bin/goorgeous-george
