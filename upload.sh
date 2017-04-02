#/bin/sh

rsync -rtvzP . ubuntu@unravel.ga:/home/ubuntu/go/src/github.com/unravel-server --exclude .git