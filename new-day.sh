#!/bin/bash

# Parse command line arguments
while getopts "d:" opt; do
  case $opt in
  d)
    DAY=$OPTARG
    ;;
  \?)
    echo "Invalid option: -$OPTARG" >&2
    echo "Usage: $0 -d <day_number>"
    exit 1
    ;;
  esac
done

# Check if day was provided
if [ -z "$DAY" ]; then
  echo "Error: Day number is required"
  echo "Usage: $0 -d <day_number>"
  exit 1
fi

# Ensure day is zero-padded to 2 digits
DAY=$(printf "%02d" $DAY)

# Calculate previous day
PREV_DAY=00

NEW_DIR="day$DAY"
PREV_DIR="day$PREV_DAY"

# Check if previous day exists
if [ ! -d "$PREV_DIR" ]; then
  echo "Error: Previous day directory '$PREV_DIR' does not exist"
  exit 1
fi

# Check if new day already exists
if [ -d "$NEW_DIR" ]; then
  echo "Error: Directory '$NEW_DIR' already exists"
  exit 1
fi

# Copy previous day to new day
echo "Copying $PREV_DIR to $NEW_DIR..."
cp -r "$PREV_DIR" "$NEW_DIR"

# Update package names in all Go files
echo "Updating package names from $PREV_DIR to $NEW_DIR..."
find "$NEW_DIR" -type f -name "*.go" -print0 | xargs -0 perl -pi -e "s/$PREV_DIR/$NEW_DIR/g"

echo "Successfully created $NEW_DIR from $PREV_DIR"
