#!/bin/bash

clean() {
    echo "removing..."
    rm -rf dist/
}

compile_go() {
    local file=$1
    local output=$2

    go build -o dist/$output $file >/dev/null 2>&1 &
    local command_pid=$!

    while ps -p $command_pid >/dev/null; do
        print_progress_bar $(ps -o etimes= -p $command_pid) 10 "Compiling $file:" "Pegasus" 0 30 "█"
        sleep 1
    done

    # Print the final progress bar at 100%
    print_progress_bar 10 10 "Compiling $file:" "Complete" 0 30 "█"
    echo ""  # Move to the next line after the progress bar
}

print_progress_bar() {
    local iteration=$1
    local total=$2
    local prefix=$3
    local suffix=$4
    local decimals=$5
    local length=$6
    local fill=$7

    local percent=$(awk "BEGIN { pc=100*${iteration}/${total}; i=int(pc); print (pc-i<0.5)?i:i+1 }")
    local filled_length=$(awk "BEGIN { fl=${length}*${iteration}/${total}; i=int(fl); print (fl-i<0.5)?i:i+1 }")
    local bar=$(printf "%-${filled_length}s" "${fill}")
    bar=${bar// /"${fill}"}
    local spaces=$(printf "%-$((length-filled_length))s")

    echo -ne "\r${prefix} [${bar}${spaces}] ${percent}% ${suffix}"
}

if [ "$1" == "clean" ]; then
    clean
    echo "Successfully removed files"
else
    mkdir dist/
    cc src/main.c src/shell/shell.c src/help/help.c src/ascii/ascii.c src/shell/helpers/helpers.c src/shell/command/command.c src/shell/history/history.c src/core-util/core.c -o dist/pegasus
    cc src/pegasus-edit/editor.c -Wall -W -pedantic -std=c99 -o dist/pegasusedit
    cc src/tools/traceroute/route.c -o dist/traceroute
    cc src/tools/packet-sniffer/sniffer.c -lpcap -o dist/sniffer
    compile_go "src/tools/port-scanner/portscanner.go" "scanner" 
    compile_go "src/tools/ping/icmp.go" "ping" > dist/
    compile_go "src/tools/dns/dns.go" "dns" 
    compile_go "src/tools/whois/whois.go" "whois"
    compile_go "src/tools/dirb/dirb.go" "dirb"
    compile_go "src/tools/chat-room/server.go" "server"
    compile_go "src/tools/hash/genhash.go" "hash"
    compile_go "src/tools/hash/id.go" "id"
    compile_go "src/tools/subnet/subnet.go" "subnet"
    compile_go "src/tools/ip-lookup/ip.go" "ip"
    compile_go "src/tools/web-server/web.go" "webserver"
    compile_go "src/tools/rev-shell/revshell.go" "revshell"

    echo "Compilation completed."
    sudo ./dist/pegasus
fi