# #!/bin/bash
# CURRENT_DIR=$1
# for x in $(find ${CURRENT_DIR}/delever_protos/* -type d); do
#   protoc -I=${x} -I=${CURRENT_DIR}/delever_protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
# done

#!/bin/bash
CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genproto
for x in $(find ${CURRENT_DIR}/delever_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/delever_protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=require_unimplemented_servers=false:${CURRENT_DIR} ${x}/*.proto
done