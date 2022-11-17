#!/bin/bash

SERVICES=("inspection" "controller" "runtime-filter" "web")
AFFECTED_RAW=`npx nx affected:apps --plain`
AFFECTED=()

for SERVICE in "${SERVICES[@]}"
do
  if [[ "$AFFECTED_RAW" == *"$SERVICE"* ]]; then
    AFFECTED+=($SERVICE)
  fi
done
if [[ "$AFFECTED_RAW" == *"invoker"* ]]; then
  AFFECTED+=("invoker-linux")
fi

json_array() {
  echo -n '['
  while [ $# -gt 0 ]; do
    x=${1//\\/\\\\}
    echo -n \"${x//\"/\\\"}\"
    [ $# -gt 1 ] && echo -n ', '
    shift
  done
  echo -n ']'
}

json_array "${AFFECTED[@]}"