#include "Block.h"

namespace suckc {

class BlockPrivate {
  SUCKC_OBJECT_IMPL(Block)

public:
};
Block::Block() : Scope(BlockScope), dPtr_(new BlockPrivate(this)) {}

Block::~Block() {}

} // namespace suckc
