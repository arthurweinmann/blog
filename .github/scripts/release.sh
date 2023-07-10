#!/bin/bash
set -e

mkdir -p ~/.ssh

echo "$DEPLOY_SSH_KEY" > ~/.ssh/id_rsa

chmod 600 ~/.ssh/id_rsa

tmp=$(pwd)
make build
cd $tmp

ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i ~/.ssh/id_rsa root@5.161.204.97 "/bin/bash -c 'systemctl stop arthurweinmann.service'"

rsync -av --delete -e "ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i ~/.ssh/id_rsa" --exclude 'storage' ./build root@5.161.204.97:/arthurweinmann

ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i ~/.ssh/id_rsa root@5.161.204.97 "/bin/bash -c 'systemctl start arthurweinmann.service'"