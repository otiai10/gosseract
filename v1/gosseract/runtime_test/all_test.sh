#!/bin/sh

runtimes=`ls ./runtime_test/*.Dockerfile`

for runtime in $runtimes; do
  testcase=`basename ${runtime} | sed -e s/\.Dockerfile$//`
  echo "=== ${testcase} ==="
  docker build . -f ${runtime} -t gosseract/test:${testcase} 1>/dev/null
  if docker run -i -t --rm gosseract/test:${testcase}; then
    echo "--- ${testcase}: Pass ---"
  else
    echo "--- ${testcase}: Failed ---"
    exit 1
  fi
done
