#!/bin/bash

local() {
    docker build -t pegasus .

    docker run -it --privileged pegasus
}

docker_global() {
    docker pull nebrix/pegasus:latest

    docker run -it --privileged nebrix/pegasus:latest
}

case $choice in
    "local")
        local
        ;;
    "global")
        docker_global
        ;;
    *)
        echo "Invalid choice. Please specify 'local' or 'global'."
        exit 1
        ;;
esac
