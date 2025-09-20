generate_dates() {
    local start_date="$1"
    
    # Validate input date format
    if ! date -d "$start_date" >/dev/null 2>&1; then
        echo "Error: Invalid date format. Please use yyyy-mm-dd" >&2
        return 1
    fi
    
    # Get yesterday's date
    local end_date=$(date -d "yesterday" +%Y-%m-%d)
    
    # Validate that start date is before end date
    if [[ $(date -d "$start_date" +%s) -gt $(date -d "$end_date" +%s) ]]; then
        echo "Error: Start date must be before yesterday" >&2
        return 1
    fi
    
    # Initialize current date
    local current_date="$start_date"
    
    # Generate dates until yesterday
    while [[ $(date -d "$current_date" +%s) -le $(date -d "$end_date" +%s) ]]; do
        echo "$current_date"
        # Increment date by 1 day
        current_date=$(date -d "$current_date + 1 day" +%Y-%m-%d)
    done
}


generate_dates "$1"
