
#!/bin/bash

echo 'Enter process name or PID'
read ent
echo 'How maximum lines do you want to output?'
read lines

netstat -tunapl 2>/dev/null | awk '$ent {print $5}' | cut -d: -f1 | sort | uniq -c | sort | tail -n$lines | grep -oP '(\d+\.){3}\d+'| while read IP ; do whois $IP | awk -F':' '/^Organization/ {print $2}' ; done

