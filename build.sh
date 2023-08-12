#!/bin/bash

clean() {
    echo "removing..."
    rm -rf pegasus scanner ping dns subnet whois dirb sniffer server nohup.out hash id ip pegasusedit traceroute revshell webserver
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

install_perl() {
    if ! command -v perl >/dev/null; then
        # Detect the operating system and install Perl based on the OS
        # You can modify this section with the appropriate package manager commands for your Unix distribution.
        if [[ "$(uname)" == "Linux" ]]; then
            if command -v apt-get >/dev/null; then
                sudo apt-get update
                sudo apt-get install -y perl
            elif command -v yum >/dev/null; then
                sudo yum install -y perl
            elif command -v dnf >/dev/null; then
                sudo dnf install -y perl
            elif command -v pacman >/dev/null; then
                sudo pacman -S perl --noconfirm
            elif command -v zypper >/dev/null; then
                sudo zypper install -y perl
            else
                echo "Package manager not found. Please install Perl manually."
                exit 1
            fi
        elif [[ "$(uname)" == "Darwin" ]]; then
            if command -v brew >/dev/null; then
                brew install perl
            else
                echo "Homebrew not found. Please install Perl manually."
                exit 1
            fi
        elif [[ "$(uname)" == "FreeBSD" ]]; then
            sudo pkg install -y perl
        elif [[ "$(uname)" == "OpenBSD" ]]; then
            doas pkg_add perl
        elif [[ "$(uname)" == "NetBSD" ]]; then
            pkgin install perl
        else
            echo "Unsupported operating system. Please install Perl manually."
            exit 1
        fi
    fi
}

install_tools() {
    if [[ "$(uname)" == "Linux" ]]; then
        if command -v apt-get >/dev/null; then
            sudo apt-get update
            sudo apt-get install -y perl-Module-CoreList-tools
        elif command -v yum >/dev/null; then
            sudo yum install -y perl-Module-CoreList-tools
        elif command -v dnf >/dev/null; then
            sudo dnf install -y perl-Module-CoreList-tools
        elif command -v pacman >/dev/null; then
            sudo pacman -S perl-Module-CoreList-tools --noconfirm
        elif command -v zypper >/dev/null; then
            sudo zypper install -y perl-Module-CoreList-tools
        else
            echo "Package manager not found. Please install perl-Module-CoreList-tools manually."
            exit 1
        fi
    elif [[ "$(uname)" == "Darwin" ]]; then
        if command -v brew >/dev/null; then
            brew install perl-Module-CoreList-tools
        else
            echo "Homebrew not found. Please install perl-Module-CoreList-tools manually."
            exit 1
        fi
    elif [[ "$(uname)" == "FreeBSD" ]]; then
        sudo pkg install -y p5-Module-CoreList
    elif [[ "$(uname)" == "OpenBSD" ]]; then
        doas pkg_add perl-Module-CoreList-tools
    elif [[ "$(uname)" == "NetBSD" ]]; then
        pkgin install perl-Module-CoreList-tools
    else
        echo "Unsupported operating system. Please install perl-Module-CoreList-tools manually."
        exit 1
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
    install_tools
    install_go
    install_perl
    sudo cpan -i 
    sudo cpan JSON
    cc src/main.c src/shell/shell.c src/help/help.c src/ascii/ascii.c src/shell/helpers/helpers.c src/shell/command/command.c src/shell/history/history.c src/core-util/core.c -o pegasus
    cc -o pegasusedit src/pegasus-edit/editor.c -Wall -W -pedantic -std=c99
    cc -o traceroute src/tools/traceroute/route.c
    cc -o sniffer src/tools/packet-sniffer/sniffer.c -lpcap
    compile_go "src/tools/port-scanner/portscanner.go" "scanner"
    compile_go "src/tools/ping/icmp.go" "ping"
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
    sudo ./pegasus
fi