#ifndef _9M_CB6FB6C9_BA74_40FD_AF67_872BC5D211B0
#define _9M_CB6FB6C9_BA74_40FD_AF67_872BC5D211B0

class SuckCParser;

namespace suckc {

class World {
protected:
  World();

public:
  static World *getInstance();

public:
  void registerParser(SuckCParser *parser);
  SuckCParser *getParser();

private:
  SuckCParser *parser_;
};

} // namespace suckc

#endif /* _9M_CB6FB6C9_BA74_40FD_AF67_872BC5D211B0 */
