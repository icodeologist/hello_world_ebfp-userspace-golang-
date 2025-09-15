#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

SEC("xdp")
int xdp_hello(struct xdp_md *ctx) {
  bpf_printk("Please someone watch HAIKYUUU......");
  return XDP_PASS;
}

char __license[] SEC("license") = "GPL";
