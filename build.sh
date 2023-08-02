#!/bin/bash

command="python3 -m PyInstaller --onefile src/tools/subnet/subnet.py"
requirements_file="requirements.txt"


clean() {
    echo "removing..."
    rm -rf shell scanner ping dns subnet dist/ build/ subnet.spec whois sniffer
}

check_and_install_dependencies() {
    if ! python3 -m PyInstaller --version >/dev/null 2>&1; then
        echo "Installing dependencies..."
        pip3 install -r $requirements_file
    fi
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
    gcc src/main.c src/shell/shell.c -o shell
    go build -o scanner src/tools/port-scanner/portscanner.go
    go build -o ping src/tools/ping/icmp.go
    go build -o dns src/tools/dns/dns.go
    go build -o whois src/tools/whois/whois.go
    go build -o sniffer src/tools/packet-sniffer/packet-sniffer.go
    check_and_install_dependencies

    echo "Compiling Python code..."
    # Start the Python build command in the background
    $command >/dev/null 2>&1 &
    command_pid=$!

    # Display the progress bar until the command finishes
    while ps -p $command_pid >/dev/null; do
        print_progress_bar $(ps -o etimes= -p $command_pid) 60 "Compiling build:" "Pegasus" 0 30 "#"
        sleep 1
    done

    # Print the final progress bar at 100%
    print_progress_bar 60 60 "Compiling Pegasus:" "Complete" 0 30 "#"
    echo ""  # Move to the next line after the progress bar

    echo "Compilation completed."
    sudo "./shell"
fi
