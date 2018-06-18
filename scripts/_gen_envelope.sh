#!/usr/bin/env bash

SRC_FILE=${1:2}
DST_ROOT=${2}

DST_FILE=${DST_ROOT}/${SRC_FILE}
DST_DIR=$(dirname ${DST_FILE})

BANNER=$(cat <<-END
// =================================================================================================
// WARNING! This file is auto-generated. Do not hand-edit.
// =================================================================================================

syntax = "proto3";

END
)

SRC=$(cat ${SRC_FILE})

PROTO_NAME_REGEX="message[[:space:]]+([[:alnum:]_]+)[[:space:]]*{"
[[ ${SRC} =~ ${PROTO_NAME_REGEX} ]]
PROTO_NAME=${BASH_REMATCH[1]}

PACKAGE_REGEX="package[[:space:]]+([[:alnum:]\._]+);"
[[ ${SRC} =~ ${PACKAGE_REGEX} ]]
SRC_PACKAGE_NAME=${BASH_REMATCH[1]}

GO_PACKAGE_REGEX="go_package[[:space:]]*=[[:space:]]*\"([[:alnum:]\./_]+)\";"
[[ ${SRC} =~ ${GO_PACKAGE_REGEX} ]]
SRC_GO_PACKAGE_NAME=${BASH_REMATCH[1]}

if ! [ -z "${PROTO_NAME}" ]
then

    PACKAGE="istio.config.protocol.envelope.${SRC_PACKAGE_NAME:6}"
    GO_PACKAGE="istio.io/api/config/protocol/envelope/${SRC_GO_PACKAGE_NAME:13}"

    BODY=$(cat <<END
import "${SRC_FILE}";
import "config/protocol/envelope/metadata.proto";

package ${PACKAGE};

option go_package = "${GO_PACKAGE}";

message ${PROTO_NAME} {
    istio.config.protocol.envelope.Metadata Metadata = 1;

    ${SRC_PACKAGE_NAME}.${PROTO_NAME} Contents = 2;
}
END
)
fi


mkdir -p ${DST_DIR}
printf "%s\n\n%s\n" "${BANNER}" "${BODY}" >${DST_FILE}
