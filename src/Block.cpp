#include "Block.h"

namespace suckc {

class BlockPrivate {
  SUCKC_OBJECT_IMPL(Block)

public:
};
Block::Block() : dPtr_(new BlockPrivate(this)) {}

Block::~Block() { delete dPtr_; }

} // namespace suckc
