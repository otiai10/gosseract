#!/bin/sh

# This script is a driver for "runtime tests".
# The "runtime test" is to test gosseract package in specific environments,
# such as OS, Go version and Tesseract version.

DRIVER=docker
REMOVE=
while [[ $# -gt 0 ]]; do
case "${1}" in
    --driver|-d)
    DRIVER="${2}"
    shift && shift
    ;;
    --rm)
    REMOVE=YES
    shift
    ;;
esac
done

function test_docker_runtimes() {
  for runtime in `ls ./test/runtime/*.Dockerfile`; do
    testcase=`basename ${runtime} | sed -e s/\.Dockerfile$//`
    echo "┌─────── ${testcase}"
    echo "│ [Docker] Building image..."
    docker build . -f ${runtime} -t gosseract/test:${testcase} 1>/dev/null
    echo "│ [Docker] Running tests..."
    SUCCEEDED=
    if docker run -i -t --rm gosseract/test:${testcase} 1>/dev/null ; then
      SUCCEEDED=YES
    fi
    if [ -n "${REMOVE}" ]; then
      echo "│ [Docker] Removing image..."
      docker rmi gosseract/test:${testcase} 1>/dev/null
    fi
    if [ -n "${SUCCEEDED}" ]; then
      echo "└─────── ${testcase} [OK]"
    else
      echo ">>>>> ${testcase} [NG!] <<<<<"
      exit 1
    fi
  done
}

function __main__() {
  case "${DRIVER}" in
    docker)
    test_docker_runtimes
    ;;
    vagrant)
    echo "# TODO: Set up vagrant test, especially for FreeBSD"
    ;;
    *)
    test_docker_runtimes
    ;;
  esac
}

__main__
