#! /bin/bash

# Deploy\
# copy web statics resources
cp -R ./templates ./bin/ 
# video temp dir
mkdir ./bin/videos

# start server
cd bin

# & means start in background
nohup ./api &
nohup ./scheduler &
nohup ./stream &
nohup ./web &

# log
echo "deploy finished...."