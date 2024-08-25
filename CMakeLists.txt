project(suckc)

cmake_minimum_required(VERSION 3.22.1)

list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

include("${CMAKE_CURRENT_SOURCE_DIR}/cmake/options.cmake")

find_package(Antlr4Ex)

# generate lexer & parser
antlr4_generate(
    suckc
    ${CMAKE_CURRENT_SOURCE_DIR}/src/SuckC.g4
    BOTH
    TRUE
    TRUE
)

# add directories for generated include files
include_directories(
    ${PROJECT_BINARY_DIR}
    ${ANTLR4_INCLUDE_DIR}
    ${ANTLR4_INCLUDE_DIR_suckc})
aux_source_directory("${CMAKE_CURRENT_SOURCE_DIR}/src" PROJECT_SOURCE_FILES)
aux_source_directory("${CMAKE_CURRENT_SOURCE_DIR}/src/ast" PROJECT_SOURCE_FILES)
add_executable(
    ${PROJECT_NAME}
    ${PROJECT_SOURCE_FILES}
    ${ANTLR4_SRC_FILES_suckc}
)
add_dependencies(${PROJECT_NAME} antlr4_shared)
target_link_libraries(${PROJECT_NAME} PRIVATE antlr4_shared)