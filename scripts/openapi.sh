#!/bin/bash

# Debounce period in seconds
DEBOUNCE_PERIOD=2

# Get current directory
CURR_DIR=$(dirname "$(realpath "${0}")")

# Load environment variables
. "$CURR_DIR"/base/init.sh

# File to store the timestamp of the last execution
mkdir -p "$CURR_DIR/logs"
TIMESTAMP_FILE="$CURR_DIR/logs/last_execution_timestamp"

# Check if the timestamp file exists and read the last execution time
if [ -f "$TIMESTAMP_FILE" ]; then
    LAST_EXECUTION_TIME=$(cat "$TIMESTAMP_FILE")
else
    LAST_EXECUTION_TIME=0
fi

# Get current time in seconds since epoch
CURRENT_TIME=$(date +%s)

# Calculate time elapsed since last execution
TIME_DIFF=$((CURRENT_TIME - LAST_EXECUTION_TIME))

# Check if the time elapsed is less than the debounce period
if [ "$TIME_DIFF" -lt "$DEBOUNCE_PERIOD" ]; then
    # Just output a message if the debounce period has not elapsed
    echo "Debounce period not elapsed. Main action will not be performed."
else
    # Update the timestamp file with the current time
    echo "$CURRENT_TIME" > "$TIMESTAMP_FILE"

    # Execute the main script functionality
    "$PROJECT_ROOT_DIR"/scripts/swagger/generate.sh
fi


# End of script - parent process remains unaffected
