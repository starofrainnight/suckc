project(suckc)

cmake_minimum_required(VERSION 3.22.1)

aux_source_directory("${CMAKE_CURRENT_SOURCE_DIR}/src" PROJECT_SOURCE_FILES)

add_executable(${PROJECT_NAME} ${PROJECT_SOURCE_FILES})
