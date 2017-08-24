#!/usr/bin/env bash

images=( "accounts" "events" "speakers" "talks")

for i in "${images[@]}"
do
  :
  docker build -t blondie/$i ../$i/
done
