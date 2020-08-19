#!/bin/bash
kubectl patch deployment test -p '{"spec":{"template":{"spec":{"containers":[{"name":"kping","env":[{"name":"KPING_DELAY","value":"60"}]}]}}}}'
