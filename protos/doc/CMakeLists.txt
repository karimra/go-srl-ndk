############################################################################
#
#   Description:        Cmake definition file
#
#   Copyright (c) 2018 Nokia
#
############################################################################

# Target name
set(proto_target sdkmgr_proto)
set(cust_target sdkmgr_all)
message(STATUS "Building executable ${proto_target}")

############################################################################
#                                                                          #
#                        PROTO FILE BUILDING                               #
#                                                                          #
############################################################################

include_directories(${CMAKE_CURRENT_SOURCE_DIR})
include_directories(${CMAKE_CURRENT_BINARY_DIR})


set(PROTO_SRCS
  sdk_common.proto
  sdk_service.proto
  config_service.proto
  interface_service.proto
  lldp_service.proto
  bfd_service.proto
  mpls_service.proto
  networkinstance_service.proto
  route_service.proto
  telemetry_service.proto
  nexthop_group_service.proto
  appid_service.proto
)

set(PROTO_DEPS
    protobuf
    gRPC::grpc
    gRPC::grpc++
    # gRPC::grpc++_reflection
    gRPC::gpr
    z
)

proto_library(${proto_target} SHARED SRCS ${PROTO_SRCS} DEPS ${PROTO_DEPS})
#target_link_libraries(${proto_target} gRPC::grpc++_reflection)

set(protos_output_dir ${PROJECT_BINARY_DIR}/sdk/service/sdk_protos)
set(protobufs_installed_stamp ${PROJECT_BINARY_DIR}/sdk/service/sdk_protos/protobufs_installed.stamp)
add_custom_target(${cust_target} ALL DEPENDS
    ${proto_target}_gen_proto_python
    ${protobufs_installed_stamp}
)
add_custom_command(
    OUTPUT  ${protobufs_installed_stamp}
    COMMAND cp ${CMAKE_CURRENT_SOURCE_DIR}/setup.py  ${protos_output_dir}/../setup.py
    COMMAND touch ${protobufs_installed_stamp}
    DEPENDS ${proto_target}
    COMMENT "${proto_target} Copying setup.py to destination"
)

#RPM
if(${BUILD_RPM})
    include("${SRLINUX_SOURCE_DIR}/cmake/Srlinuxpackage.cmake")
    configure_file("${CMAKE_CURRENT_SOURCE_DIR}/rpm.spec.in" "${CMAKE_CURRENT_BINARY_DIR}/rpm.spec" @ONLY IMMEDIATE)
    srlinux_add_package(${proto_target} CUSTOM_TARGET)
    add_dependencies(${proto_target}_rpm ${cust_target})
endif(${BUILD_RPM})
