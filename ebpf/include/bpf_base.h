#ifndef __BPF_BASE_H__
#define __BPF_BASE_H__

#if defined(__TARGET_ARCH_x86)

#define GO_PARAM1(x) ((x)->ax)
#define GO_PARAM2(x) ((x)->bx)
#define GO_PARAM3(x) ((x)->cx)
#define GO_PARAM4(x) ((x)->di)
#define GO_PARAM5(x) ((x)->si)
#define GO_PARAM6(x) ((x)->r8)
#define GO_PARAM7(x) ((x)->r9)
#define GO_PARAM8(x) ((x)->r10)
#define GO_PARAM9(x) ((x)->r11)

#define GOROUTINE(x) ((x)->r14)

#elif defined(__TARGET_ARCH_arm64)

#define PT_GO_PARAM1(x) (((PT_REGS_ARM64 *)(x))->regs[0])
#define PT_GO_PARAM2(x) (((PT_REGS_ARM64 *)(x))->regs[1])
#define PT_GO_PARAM3(x) (((PT_REGS_ARM64 *)(x))->regs[2])
#define PT_GO_PARAM4(x) (((PT_REGS_ARM64 *)(x))->regs[3])
#define PT_GO_PARAM5(x) (((PT_REGS_ARM64 *)(x))->regs[4])
#define PT_GO_PARAM6(x) (((PT_REGS_ARM64 *)(x))->regs[5])
#define PT_GO_PARAM7(x) (((PT_REGS_ARM64 *)(x))->regs[6])
#define PT_GO_PARAM8(x) (((PT_REGS_ARM64 *)(x))->regs[7])
#define PT_GO_PARAM9(x) (((PT_REGS_ARM64 *)(x))->regs[8])

#define GOROUTINE_PTR(x) (((PT_REGS_ARM64 *)(x))->regs[28])

#endif /*defined(__TARGET_ARCH_arm64)*/

#endif /*__BPF_BASE_H__*/


