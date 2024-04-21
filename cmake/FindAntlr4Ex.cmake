# A trick to fix the FindAntlr4.cmake does not supports exists 
# antlr4 program in the PATH
if(NOT ANTLR4_JAR_LOCATION)
    find_program(Antlr4Ex_BIN antlr4)
    if(Antlr4Ex_BIN)
        set(ANTLR_EXECUTABLE "${Antlr4Ex_BIN}")
        set(Java_JAVA_EXECUTABLE "${CMAKE_CURRENT_LIST_DIR}/fake-java.sh")
        set(ANTLR4_JAR_LOCATION "${Antlr4Ex_BIN}")
    endif(Antlr4Ex_BIN)
endif()

find_package(antlr4-runtime REQUIRED)
find_package(antlr4-generator REQUIRED)

