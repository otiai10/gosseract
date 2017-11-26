#!/bin/bash

function help_message() {
  cat << EOF
This script is a driver for "runtime tests".
The "runtime test" is to test gosseract package in specific environments,
such as OS, Go version and Tesseract version.

Options:
  --driver|-d {name}  Specify VM driver software, either of [docker, vagrant].
  --verbose|-v        Show verbose logs of setting VMs up.
  --run|-R {case}     Run only specified cases which includes given pattern in the case name.
  --rm                Remove VMs which are created by this runtime test.
  --help|-h           Show this message ;)

Examples:

  ./test/script/runtime.sh -d docker -v -R CentOS --rm

EOF
}

DRIVER=
REMOVE=
QUIET="--quiet"
MATCH=
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
    --verbose|-v)
    QUIET=
    shift
    ;;
    --run|-R)
    MATCH="${2}"
    shift && shift
    ;;
    --help|-h)
    help_message
    exit 0
    ;;
    *)
    shift
    ;;
esac
done

function test_docker_runtimes() {
  for runtime in `ls ./test/runtime/*.Dockerfile`; do
    testcase=`basename ${runtime} | sed -e s/\.Dockerfile$//`
    if [ -n "${MATCH}" ]; then
      if [[ "${testcase}" != *${MATCH}* ]]; then
        continue
      fi
    fi
    echo "┌──────────── ${testcase}"
    echo "│ [Docker] Building image..."
    docker build . -f ${runtime} -t gosseract/test:${testcase} ${QUIET} | sed "s/^/│ /"
    echo "│ [Docker] Running tests..."
    SUCCEEDED=
    if docker run -i -t --rm gosseract/test:${testcase} | sed "s/^/│ /" ; then
      SUCCEEDED=YES
    fi
    if [ -n "${REMOVE}" ]; then
      echo "│ [Docker] Removing image..."
      docker rmi gosseract/test:${testcase} 1>/dev/null
    fi
    if [ -n "${SUCCEEDED}" ]; then
      echo "└───────────── ${testcase} [OK]"
    else
      echo ">>>>> ${testcase} [NG!] <<<<<"
      exit 1
    fi
  done
}

function test_vagrant_runtimes() {
  for runtime in `ls ./test/runtime/*.Vagrantfile`; do
    testcase=`basename ${runtime} | sed -e s/\.Vagrantfile$//`
    if [ -n "${MATCH}" ]; then
      if [[ "${testcase}" != *${MATCH}* ]]; then
        continue
      fi
    fi
    echo "┌───────────── ${testcase}"
    echo "│ [Vagrant] Making VM up..."
    vboxname=gosseract-test-${testcase}
    SUCCEEDED=
    if VAGRANT_VAGRANTFILE=${runtime} VIRTUALBOX_NAME=${vboxname} vagrant up --provision 2>&1 | sed "s/^/│ /" ; then
      SUCCEEDED=YES
    fi
    if [ -n "${REMOVE}" ]; then
      echo "│ [Vagrant] Removing VM..."
      VAGRANT_VAGRANTFILE=${runtime} vagrant destroy -f | sed "s/^/│ /"
    fi
    if [ -n "${SUCCEEDED}" ]; then
      echo "└───────────── ${testcase} [OK]"
    else
      echo ">>>>> ${testcase} [NG!] <<<<<"
      exit 1
    fi
  done
}

function __main__() {
  case ${DRIVER} in
    docker)
    test_docker_runtimes
    ;;
    vagrant)
    test_vagrant_runtimes
    ;;
    *)
    test_docker_runtimes
    test_vagrant_runtimes
    ;;
  esac
}

__main__
