#!/bin/bash

while getopts 'abc:h' opt; do
  case "$opt" in
    a)
      echo "Processing option 'a'"
      ;;
    b)
      echo "Processing option 'b'"
      ;;
    c)
      arg="$OPTARG"
      echo "Processing option 'c' with '${OPTARG}' argument"
      ;;
    ?|h)
      echo "Usage: $(basename $0) [-a] [-b] [-c arg]"
      exit 1
      ;;
  esac
done
shift "$(($OPTIND -1))"

echo $1
