#!/bin/bash

# Define default output directory
output_dir="scratch"
dev_flag=""

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -o|--output)
            output_dir="$2"
            shift 2
            ;;
        --dev)
            dev_flag="--dev"
            shift
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

#
pids=()

# Define a function to download and run command on each repository
function download_and_run_command {
    repo=$1
    if [ ! -d "$output_dir/$(basename "$repo")" ]; then
        git clone "https://$repo.git" "$output_dir/$(basename "$repo")" &
        pid=$!
        wait $pid
    else
        echo "already cloned $repo"
    fi
    cd "$output_dir/$(basename "$repo")"
    echo "running   : $repo"
    output=$(../../scip-go $dev_flag 2>&1 &)
    pid=$!
    pids+=($pid)

    if [ $? -eq 0 ]; then
        echo "completed : $repo"
    else
        echo "failed    : $repo"
        echo "==========================================="
        echo "$output"
        echo "==========================================="
    fi
    cd ../..
}

# Create output directory if it doesn't exist
if [ ! -d "$output_dir" ]; then
    mkdir -p "$output_dir"
fi

# Define an array of repositories to download
repos=(
    "github.com/sourcegraph/scip-go"
    "github.com/sourcegraph-testing/etcd"
    "github.com/sourcegraph-testing/tidb"
    "github.com/sourcegraph-testing/titan"
    "github.com/sourcegraph-testing/zap"
    "github.com/sourcegraph-testing/nacelle"
    "github.com/sourcegraph-testing/nacelle-config"
    "github.com/sourcegraph-testing/nacelle-service"
    "github.com/sourcegraph/code-intel-extensions"
)

# Download and run command on each repository in parallel
for repo in "${repos[@]}"; do
    download_and_run_command "$repo" &
done

# Wait for all commands to finish
while true; do
    echo "checking jobs..."

    for pid in "${pids[@]}"; do
        if jobs -l | grep -q $pid; then
            echo "stil waiting for $pid"
        fi
    done

    sleep 1;
done
