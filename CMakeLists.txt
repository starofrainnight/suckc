project(suckc)

cmake_minimum_required(VERSION 3.22.1)

find_path(ANTLR4_RUNTIME_BASE_DIR antlr4-runtime)
if(ANTLR4_RUNTIME_BASE_DIR)
    include_directories("${ANTLR4_RUNTIME_BASE_DIR}/antlr4-runtime")
endif(ANTLR4_RUNTIME_BASE_DIR)

find_library(ANTLR4_RUNTIME_LIB antlr4-runtime)

aux_source_directory("${CMAKE_CURRENT_SOURCE_DIR}/src" PROJECT_SOURCE_FILES)

add_executable(${PROJECT_NAME} ${PROJECT_SOURCE_FILES})
target_link_libraries(${PROJECT_NAME} ${ANTLR4_RUNTIME_LIB})
