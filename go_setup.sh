#!/bin/bash

# Check if argument is provided
if [ -z "$1" ]
then
    echo "Please provide a folder name"
    exit 1
fi

# Extract the name part after the number
TOPIC_NAME=$(echo "$1" | sed 's/^[0-9]*//')

# Create directory
mkdir -p "$1"

# Create main.go with basic template
cat > "$1/main.go" << EOL
package main

import "fmt"

func main() {
    fmt.Println("$TOPIC_NAME in golang")
}
EOL

# Initialize go module with extracted name
cd "$1"
go mod init "$TOPIC_NAME"

echo "Created Go project in ./$1 with module name $TOPIC_NAME"