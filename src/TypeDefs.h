#ifndef _9M_D6AADC0C_A023_45DF_9B3D_59998193748C
#define _9M_D6AADC0C_A023_45DF_9B3D_59998193748C

#include <memory>

#define SUCKC_EMPTY /* Empty */

#define SUCKC_OBJECT_DECL(name)                                                \
private:                                                                       \
  std::unique_ptr<class name##Private> dPtr_;                                  \
  friend class name##Private;

#define SUCKC_OBJECT_IMPL(name)                                                \
public:                                                                        \
  name *pPtr_;                                                                 \
  name##Private(name *p) : pPtr_(p) {}                                         \
  ~name##Private() {}

#define SUCKC_D() auto &d = dPtr_;
#define SUCKC_P() auto &p = pPtr_;

#endif /* _9M_D6AADC0C_A023_45DF_9B3D_59998193748C */
