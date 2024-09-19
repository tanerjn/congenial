#!/bin/sh

docker build -t knapsack .

docker tag knapsack tanermetin/knapsack:latest

docker push tanermetin/knapsack:latest

oc new-app tanermetin/knapsack 

oc expose service/knapsack
