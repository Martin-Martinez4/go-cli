#!/bin/bash
# Only works on Linux/Unix

FHASH=`md5sum $1`
while true; do
    NHASH=`md5sum $1`
    if [ "$FHASH" != "$FHASH" ]; then
        ./mdp -file $1
        FHASH=$NHASH
    fi
    sleep 5
done