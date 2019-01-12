#!/bin/bash

cols=${1:-}
if [[ ! "$cols" ]]; then
    echo "number of columns are not specified"
    echo "you can obtain it by running tput cols"
    exit 1
fi

i3-msg 'workspace e' &>/dev/null

:terminal() {
    local name="$1"
    shift
    local cmd="$@"

    rm suite.out
    echo "terminal: $name"
    "${@}"
    cat suite.out
    echo "---"
}

export SUITE_COLS=$cols

:terminal rxvt-unicode urxvt -e ./suite.sh
:terminal xterm xterm -e ./suite.sh
:terminal alacritty alacritty -e ./suite.sh
:terminal kitty kitty -e ./suite.sh
