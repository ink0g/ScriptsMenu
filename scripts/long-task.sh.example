#!/usr/bin/env bash

# Safely resolve the repository root directory
SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE}")" &> /dev/null && pwd)
PROJECT_DIR=$(dirname "$SCRIPT_DIR")
LOG_FILE="$PROJECT_DIR/demo_output.log"

echo "[$(date '+%H:%M:%S')] Background task started. Rofi window should close now." >> "$LOG_FILE"

# Simulate a 9-second background process (3 steps with a sleep interval)
for i in {1..3}
do
    sleep 3
    echo "[$(date '+%H:%M:%S')] Background task: Step $i completed..." >> "$LOG_FILE"
done

echo "[$(date '+%H:%M:%S')] Background task finished successfully!" >> "$LOG_FILE"

