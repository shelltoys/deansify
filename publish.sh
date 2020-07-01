#!/bin/bash

set -eo pipefail
set -x

VERSION=$(git tag | grep "^v" | tail -1)

curl -T "./${BIN_NAME}" \
  -u"dpritchett:$BINTRAY_API_KEY" \
  -H "X-Bintray-Package:${PACKAGE_NAME}" \
  -H "X-Bintray-Version:${VERSION}" \
  "https://api.bintray.com/content/shelltoys/binaries/${PACKAGE_NAME}/${VERSION}/${BIN_NAME}"
