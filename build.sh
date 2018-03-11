#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o aare-exporter .
sudo docker build . -t ckevi/aare-exporter:latest
sudo docker tag ckevi/aare-exporter:latest ckevi/aare-exporter:1.2
sudo docker push ckevi/aare-exporter:latest
sudo docker push ckevi/aare-exporter:1.2

