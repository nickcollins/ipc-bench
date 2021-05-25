## project(gocli)

###########################################################
## GOLANG CLIENT CONFIG
###########################################################

set(GOCLI_LIB gocli.lib)
set(GOCLI_SRCS Client.go Server.go)
set(GOCLI_LIB_NAME libgocli.a)
set(GOCLI_LIB_OUTS libgocli.a libgocli.h)

# Custom command for 'go build -buildmode=c-archive ...'
# to create a library from Go codes.
add_custom_command(
  OUTPUT ${GOCLI_LIB_OUTS}
  DEPENDS ipc-bench-common ${GOCLI_SRCS}
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
  COMMAND go build -buildmode=c-archive
  -o "${CMAKE_CURRENT_BINARY_DIR}/${GOCLI_LIB_NAME}"
  ${CMAKE_GO_FLAGS} ${GOCLI_SRCS}
  COMMENT "Building Golang client library")

# Add a custom target for the library.
add_custom_target(${GOCLI_LIB} DEPENDS ${GOCLI_LIB_OUTS})

# A library and a header are generated in the binary directory.
include_directories(${CMAKE_CURRENT_BINARY_DIR} .)
link_directories(${CMAKE_CURRENT_BINARY_DIR})

###########################################################
## TARGETS
###########################################################

add_executable(wtmp-client client.c)
add_dependencies(wtmp-client ${GOCLI_LIB})
add_executable(wtmp-server server.c)
add_dependencies(wtmp-server ${GOCLI_LIB})
add_executable(wtmp wtmp.c)

###########################################################
## COMMON
###########################################################

target_link_libraries(wtmp-client gocli ipc-bench-common pthread)
target_link_libraries(wtmp-server gocli ipc-bench-common pthread)
target_link_libraries(wtmp ipc-bench-common)
