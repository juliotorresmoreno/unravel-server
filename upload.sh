#/bin/sh

rsync -rtvzP . ubuntu@unravel.ga:/home/ubuntu/unravel-server --exclude .git