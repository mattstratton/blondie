#!/usr/bin/env bash

images=( "accounts" "events" "speakers" "talks")

for i in "${images[@]}"
do
  :
  docker push blondie/$i 
done
