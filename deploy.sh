#!/bin/bash

kill -9 `ps -e -o pid,cmd | grep go-build| grep -v grep | awk '{ print $1 }'`
rm -rf go/src/workspace

cd go/src
git clone git@github.com:ITCreate/yorozuya-serverside.git workspace
cd workspace

git clone git@github.com:ITCreate/yorozuya-client.git client
cd client
npm install
npm run build

mv public/* ../public/
cd ..

/home/itc/.anyenv/envs/goenv/shims/go get
nohup /home/itc/.anyenv/envs/goenv/shims/go run app.go > my.log 2>&1 &