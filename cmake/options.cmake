option(SUCKC_TRACE_NODE_TREE "Enable tracing of node tree" OFF)
if(SUCKC_TRACE_NODE_TREE)
  add_definitions(-DSUCKC_TRACE_NODE_TREE)
endif(SUCKC_TRACE_NODE_TREE)