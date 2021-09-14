#!/bin/bash

if [ "$#" -eq "0" ]
  then
    echo 'Enter process name or PID'
    read ent
    echo 'How maximum lines do you want to output?'
    read lines
 elif [ "$#" -eq "1" ]
  then
    ent=$1
    lines=5
  else
    ent=$1
    lines=$2
  fi

netstat -tunapl 2>/dev/null | awk '"/"$ent {print $5}' | cut -d: -f1 | sort | uniq -c | sort | tail -n$lines | grep -oP '(\d+\.){3}\d+'| while read IP ; do whois $IP | awk -F':' '/^Organization/ {print $2}' ; done

