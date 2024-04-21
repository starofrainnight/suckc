#!/bin/sh

# Print usage
usage() {
    echo "Usage: $0 -jar <ANTLR4_BIN> <ARGUMENTS>"
}

# If argument less than 1 then print usage and exit
if [ $# -lt 1 ]; then
    usage
    exit 1
fi

# If the second argument is equal to "-version" then pass "-version" to "java" program
if [ "$1" = "-version" ]; then
    java -version

    # Don't touch the exit value by "java -version"
    exit
fi

# If the second argument is not equal to "-jar" then print usage and exit
if [ "$1" != "-jar" ]; then
    usage
    exit 3
fi

ANTLR4_BIN=$2

# Pass all the rest arguments except the 2 arguments to"antlr4" program
shift 2
$ANTLR4_BIN "$@"
