#!/bin/zsh
#ab -n 100 -c 10 'localhost:8080/sleep?r=1000'
ab -n 100 -c 1 'localhost:8080/mutex?r=1000'
