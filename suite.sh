#!/bin/bash

exec benchmark-terminal -w ${SUITE_COLS} -t 5 -c 'X' -o /dev/stderr > suite.out
