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

World::World() : parser_(nullptr), isEnabledDebug_(false) {}

void World::registerParser(SuckCParser *parser) { parser_ = parser; }

SuckCParser *World::getParser() { return parser_; }
void World::setEnabledDebug(bool enabled) { isEnabledDebug_ = enabled; }
} // namespace suckc