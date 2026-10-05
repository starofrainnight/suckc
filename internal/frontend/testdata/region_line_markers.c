# 0 "first.suckc"
# 0 "<built-in>"
# 0 "<command-line>"
# 1 "/usr/include/stdc-predef.h" 1 3 4
# 0 "<command-line>" 2
# 1 "first.suckc"
#pragma c_scope_begin
# 1 "vendor.h" 1



typedef unsigned long vendor_size_t;

vendor_size_t vendor_len(const char *s) __attribute__((nonnull));

struct vendor_cfg {
    int flags;
} __attribute__((packed));

static double vendor_get_half(void) { return 1.5; }
static long double vendor_get_ld(void) { return 1.5L; }

#pragma pack(1)
struct vendor_packed {
    char a;
    int b;
};
#pragma pack()
# 3 "first.suckc" 2
#pragma c_scope_end

void f(void) {
    vendor_get_half();
    vendor_get_ld();
}
