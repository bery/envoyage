#!/bin/bash

# Number of times the command will be executed
repeat_times=7

# Loop and execute the command
for i in $(seq 1 $repeat_times); do
    helm template ev-${i} ./helm --output-dir rendered/${i} -f helm/values-base.yaml --set ingress.host="ev-${i}.sandbox.xbery.online"
done
