#include "World.h"
#include "SuckCParser.h"

namespace suckc {

static World *sWorld;

World *World::getInstance() {
  if (!sWorld) {
    sWorld = new World();
  };

  return sWorld;
}

World::World() : parser_(nullptr) {}

void World::registerParser(SuckCParser *parser) { parser_ = parser; }

SuckCParser *World::getParser() { return parser_; }
}