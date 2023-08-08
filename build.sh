#!/bin/bash

command_subnet="python3 -m PyInstaller --onefile src/tools/subnet/subnet.py"
command_ip="python3 -m PyInstaller --onefile src/tools/ip-lookup/ip.py"
command_sniffer="python3 -m PyInstaller --onefile src/tools/packet-sniffer/sniffer.py"
command_hashid="python3 -m PyInstaller --onefile src/tools/hashident/hash.py"
requirements_file="requirements.txt"

clean() {
    echo "removing..."
    rm -rf pegasus scanner ping dns subnet dist/ build/ subnet.spec whois dirb ip.spec sniffer.spec hash.spec server nohup.out hash
}

compile_go() {
    local file=$1
    local output=$2

    go build -o $output $file >/dev/null 2>&1 &
    local command_pid=$!

    while ps -p $command_pid >/dev/null; do
        print_progress_bar $(ps -o etimes= -p $command_pid) 10 "Compiling $file:" "Pegasus" 0 30 "█"
        sleep 1
    done

    # Print the final progress bar at 100%
    print_progress_bar 10 10 "Compiling $file:" "Complete" 0 30 "█"
    echo ""  # Move to the next line after the progress bar
}

compile_python() {
    local command=$1
    local task=$2

    # Start the Python build command in the background
    $command >/dev/null 2>&1 &
    local command_pid=$!

    # Display the progress bar until the command finishes
    while ps -p $command_pid >/dev/null; do
        print_progress_bar $(ps -o etimes= -p $command_pid) 60 "Compiling $task:" "Pegasus" 0 30 "█"
        sleep 1
    done

    # Print the final progress bar at 100%
    print_progress_bar 60 60 "Compiling $task:" "Complete" 0 30 "█"
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
    chmod a+x install
    ./install
    sudo pip3 install -r requirements.txt
    gcc src/main.c src/shell/shell.c src/help/help.c src/ascii/ascii.c src/shell/helpers/helpers.c src/shell/command/command.c src/shell/history/history.c src/core-util/core.c -o pegasus
    compile_go "src/tools/port-scanner/portscanner.go" "scanner"
    compile_go "src/tools/ping/icmp.go" "ping"
    compile_go "src/tools/dns/dns.go" "dns"
    compile_go "src/tools/whois/whois.go" "whois"
    compile_go "src/tools/dirb/dirb.go" "dirb"
    compile_go "src/tools/chat-room/server.go" "server"
    compile_go "src/tools/hash/genhash.go" "hash"

    echo "Compiling Python scripts..."
    compile_python "$command_subnet" "subnet"
    compile_python "$command_ip" "ip"
    compile_python "$command_sniffer" "sniffer"
    compile_python "$command_hashid" "hashident"

    echo "Compilation completed."
    sudo "./pegasus"
fi