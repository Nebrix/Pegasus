#!/bin/bash

command="python3 -m PyInstaller --onefile src/tools/subnet/subnet.py"
requirements_file="requirements.txt"

clean() {
    echo "removing..."
    rm -rf pegasus scanner ping dns subnet dist/ build/ subnet.spec whois sniffer
}

install_pip() {
    if ! command -v pip3 >/dev/null; then
        echo "Installing pip..."
        # Detect the operating system and install pip based on the OS
        # You can modify this section with the appropriate package manager commands for your Unix distribution.
        if [[ "$(uname)" == "Linux" ]]; then
            if command -v apt-get >/dev/null; then
                sudo apt-get update
                sudo apt-get install -y python3-pip
            elif command -v yum >/dev/null; then
                sudo yum install -y python3-pip
            elif command -v dnf >/dev/null; then
                sudo dnf install -y python3-pip
            elif command -v pacman >/dev/null; then
                sudo pacman -S python-pip
            elif command -v zypper >/dev/null; then
                sudo zypper install -y python3-pip
            else
                echo "Package manager not found. Please install pip manually."
                exit 1
            fi
        elif [[ "$(uname)" == "Darwin" ]]; then
            if command -v brew >/dev/null; then
                brew install python
            else
                echo "Homebrew not found. Please install pip manually."
                exit 1
            fi
        elif [[ "$(uname)" == "FreeBSD" ]]; then
            sudo pkg install -y py38-pip
        elif [[ "$(uname)" == "OpenBSD" ]]; then
            doas pkg_add py3-pip
        elif [[ "$(uname)" == "NetBSD" ]]; then
            pkgin install py39-pip
        else
            echo "Unsupported operating system. Please install pip manually."
            exit 1
        fi
    fi
}

install_go() {
    if ! command -v go >/dev/null; then
        # Detect the operating system and install Go based on the OS
        # You can modify this section with the appropriate package manager commands for your Unix distribution.
        if [[ "$(uname)" == "Linux" ]]; then
            if command -v apt-get >/dev/null; then
                sudo apt-get update
                sudo apt-get install -y golang
            elif command -v yum >/dev/null; then
                sudo yum install -y golang
            elif command -v dnf >/dev/null; then
                sudo dnf install -y golang
            elif command -v pacman >/dev/null; then
                sudo pacman -S go
            elif command -v zypper >/dev/null; then
                sudo zypper install -y go
            else
                echo "Package manager not found. Please install Go manually."
                exit 1
            fi
        elif [[ "$(uname)" == "Darwin" ]]; then
            if command -v brew >/dev/null; then
                brew install go
            else
                echo "Homebrew not found. Please install Go manually."
                exit 1
            fi
        elif [[ "$(uname)" == "FreeBSD" ]]; then
            sudo pkg install -y go
        elif [[ "$(uname)" == "OpenBSD" ]]; then
            doas pkg_add go
        elif [[ "$(uname)" == "NetBSD" ]]; then
            pkgin install go
        else
            echo "Unsupported operating system. Please install Go manually."
            exit 1
        fi
    fi
}

check_and_install_dependencies() {
    if ! command -v python3 -m PyInstaller --version >/dev/null 2>&1; then
        echo "Installing dependencies..."
        pip3 install -r $requirements_file
    fi
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
    install_go
    install_pip

    gcc src/main.c src/shell/shell.c src/help/help.c src/ascii/ascii.c -o pegasus
    compile_go "src/tools/port-scanner/portscanner.go" "scanner"
    compile_go "src/tools/ping/icmp.go" "ping"
    compile_go "src/tools/dns/dns.go" "dns"
    compile_go "src/tools/whois/whois.go" "whois"
    compile_go "src/tools/packet-sniffer/packet-sniffer.go" "sniffer"
    check_and_install_dependencies

    echo "Compiling Python code..."
    # Start the Python build command in the background
    $command >/dev/null 2>&1 &
    command_pid=$!

    # Display the progress bar until the command finishes
    while ps -p $command_pid >/dev/null; do
        print_progress_bar $(ps -o etimes= -p $command_pid) 60 "Compiling build:" "Pegasus" 0 30 "█"
        sleep 1
    done

    # Print the final progress bar at 100%
    print_progress_bar 60 60 "Compiling Pegasus:" "Complete" 0 30 "█"
    echo ""  # Move to the next line after the progress bar

    echo "Compilation completed."
    sudo "./pegasus"
fi
