project(suckc)

cmake_minimum_required(VERSION 3.22.1)

list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

include("${CMAKE_CURRENT_SOURCE_DIR}/cmake/options.cmake")

find_package(wxWidgets REQUIRED core base)
find_package(Antlr4Ex)

include(${wxWidgets_USE_FILE})

# generate lexer
antlr4_generate(
    suckc_lexer
    "${CMAKE_CURRENT_SOURCE_DIR}/src/SuckCLexer.g4"
    LEXER
    FALSE
    FALSE
)

# generate parser
antlr4_generate(
    suckc_parser
    "${CMAKE_CURRENT_SOURCE_DIR}/src/SuckCParser.g4"
    PARSER
    TRUE
    TRUE
    ""
    "${ANTLR4_TOKEN_FILES_suckc_lexer}"
    "${ANTLR4_TOKEN_DIRECTORY_suckc_lexer}"
)

    # add directories for generated include files
include_directories(
    "${CMAKE_CURRENT_SOURCE_DIR}/src"
    ${PROJECT_BINARY_DIR}
    ${ANTLR4_INCLUDE_DIR}
    ${ANTLR4_INCLUDE_DIR_suckc_lexer}
    ${ANTLR4_INCLUDE_DIR_suckc_parser})
aux_source_directory("${CMAKE_CURRENT_SOURCE_DIR}/src" PROJECT_SOURCE_FILES)
aux_source_directory("${CMAKE_CURRENT_SOURCE_DIR}/src/ast" PROJECT_SOURCE_FILES)
add_executable(
    ${PROJECT_NAME}
    ${PROJECT_SOURCE_FILES}
    ${ANTLR4_SRC_FILES_suckc_lexer}
    ${ANTLR4_SRC_FILES_suckc_parser}
)
add_dependencies(${PROJECT_NAME} antlr4_shared)
target_link_libraries(${PROJECT_NAME} PRIVATE antlr4_shared ${wxWidgets_LIBRARIES})
