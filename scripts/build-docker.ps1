param (
    [string]$choice
)

function Local {
    docker build -t pegasus .

    docker run -it --privileged pegasus
}

function DockerGlobal {
    docker pull nebrix/pegasus:latest

    docker run -it --privileged nebrix/pegasus:latest
}

# Main script logic
if ($choice -eq "local") {
    Local
}
elseif ($choice -eq "global") {
    DockerGlobal
}
else {
    Write-Host "Invalid choice. Please specify 'local' or 'global'."
}