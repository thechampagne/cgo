package main

/*
#include <stdlib.h>
#include <stdint.h>

typedef struct {
    unsigned int sum;
    unsigned int carry_out;
} bits_add_t;

typedef struct {
    uint32_t sum;
    uint32_t carry_out;
} bits_add32_t;

typedef struct {
    uint64_t sum;
    uint64_t carry_out;
} bits_add64_t;

typedef struct {
    unsigned int quo;
    unsigned int rem;
} bits_div_t;

typedef struct {
    uint32_t quo;
    uint32_t rem;
} bits_div32_t;

typedef struct {
    uint64_t quo;
    uint64_t rem;
} bits_div64_t;

typedef struct {
    unsigned int hi;
    unsigned int lo;
} bits_mul_t;

typedef struct {
    uint32_t hi;
    uint32_t lo;
} bits_mul32_t;

typedef struct {
    uint64_t hi;
    uint64_t lo;
} bits_mul64_t;

typedef struct {
    unsigned int diff;
    unsigned int borrow_out;
} bits_sub_t;

typedef struct {
    uint32_t diff;
    uint32_t borrow_out;
} bits_sub32_t;

typedef struct {
    uint64_t diff;
    uint64_t borrow_out;
} bits_sub64_t;
*/
import "C"
import (
  "unsafe"
	"math/bits"
)

/**
 * Add returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; 
 * otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1. 
 *
 * This function's execution time does not depend on the inputs.
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_add_t* add = bits_add(12, 32, 0);
 *  printf("Sum: %d\n", add->sum);
 *  printf("Carry out: %d", add->carry_out);
 *  free(add);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @param carry
 * @return bits_add_t
 */
//export bits_add
func bits_add(x C.uint, y C.uint, carry C.uint) *C.bits_add_t {
  self := (*C.bits_add_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_add_t{}))))
  re, se := bits.Add(uint(x), uint(y), uint(carry))
  self.sum = C.uint(re)
  self.carry_out = C.uint(se)
  return self
}

/**
 * Add32 returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; 
 * otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_add32_t* add = bits_add32(12, 32, 0);
 *  printf("Sum: %d\n", add->sum);
 *  printf("Carry out: %d", add->carry_out);
 *  free(add);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @param carry
 * @return bits_add32_t
 */
//export bits_add32
func bits_add32(x C.uint32_t, y C.uint32_t, carry C.uint32_t) *C.bits_add32_t {
  self := (*C.bits_add32_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_add32_t{}))))
  re, se := bits.Add32(uint32(x), uint32(y), uint32(carry))
  self.sum = C.uint32_t(re)
  self.carry_out = C.uint32_t(se)
  return self
}

/**
 * Add64 returns the sum with carry of x, y and carry: sum = x + y + carry. The carry input must be 0 or 1; 
 * otherwise the behavior is undefined. The carryOut output is guaranteed to be 0 or 1.
 *
 * This function's execution time does not depend on the inputs.  
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_add64_t* add = bits_add64(12, 32, 0);
 *  printf("Sum: %d\n", add->sum);
 *  printf("Carry out: %d", add->carry_out);
 *  free(add);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @param carry
 * @return bits_add64_t
 */
//export bits_add64
func bits_add64(x C.uint64_t, y C.uint64_t, carry C.uint64_t) *C.bits_add64_t {
  self := (*C.bits_add64_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_add64_t{}))))
  re, se := bits.Add64(uint64(x), uint64(y), uint64(carry))
  self.sum = C.uint64_t(re)
  self.carry_out = C.uint64_t(se)
  return self
}

/**
 * Div returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y 
 * with the dividend bits' upper half in parameter hi and the lower half in parameter lo. Div panics for y == 0 
 * (division by zero) or y <= hi (quotient overflow).
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_div_t* div = bits_div(0, 6, 3);
 *  printf("%d\n", div->quo);
 *  printf("%d", div->rem);
 *  free(div);
 *  return 0;
 * }
 * * *
 *
 * @param hi
 * @param lo
 * @param y
 * @return bits_div_t
 */
//export bits_div
func bits_div(hi C.uint, lo C.uint, y C.uint) *C.bits_div_t {
  self := (*C.bits_div_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_div_t{}))))
  re, se := bits.Div(uint(hi), uint(lo), uint(y))
  self.quo = C.uint(re)
  self.rem = C.uint(se)
  return self
}

/**
 * Div32 returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y 
 * with the dividend bits' upper half in parameter hi and the lower half in parameter lo. Div32 panics for y == 0 
 * (division by zero) or y <= hi (quotient overflow). 
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_div32_t* div = bits_div32(0, 6, 3);
 *  printf("%d\n", div->quo);
 *  printf("%d", div->rem);
 *  free(div);
 *  return 0;
 * }
 * * *
 *
 * @param hi
 * @param lo
 * @param y
 * @return bits_div32_t
 */
//export bits_div32
func bits_div32(hi C.uint32_t, lo C.uint32_t, y C.uint32_t) *C.bits_div32_t {
  self := (*C.bits_div32_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_div32_t{}))))
  re, se := bits.Div32(uint32(hi), uint32(lo), uint32(y))
  self.quo = C.uint32_t(re)
  self.rem = C.uint32_t(se)
  return self
}

/**
 * Div64 returns the quotient and remainder of (hi, lo) divided by y: quo = (hi, lo)/y, rem = (hi, lo)%y 
 * with the dividend bits' upper half in parameter hi and the lower half in parameter lo. Div64 panics for y == 0
 * (division by zero) or y <= hi (quotient overflow). 
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_div64_t* div = bits_div64(0, 6, 3);
 *  printf("%d\n", div->quo);
 *  printf("%d", div->rem);
 *  free(div);
 *  return 0;
 * }
 * * *
 *
 * @param hi
 * @param lo
 * @param y
 * @return bits_div64_t
 */
//export bits_div64
func bits_div64(hi C.uint64_t, lo C.uint64_t, y C.uint64_t) *C.bits_div64_t {
  self := (*C.bits_div64_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_div64_t{}))))
  re, se := bits.Div64(uint64(hi), uint64(lo), uint64(y))
  self.quo = C.uint64_t(re)
  self.rem = C.uint64_t(se)
  return self
}

/**
 * LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_leading_zeros(1));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_leading_zeros
func bits_leading_zeros(x C.uint) C.int {
  return C.int(bits.LeadingZeros(uint(x)))
}

/**
 * LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_leading_zeros16(1));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_leading_zeros16
func bits_leading_zeros16(x C.uint16_t) C.int {
  return C.int(bits.LeadingZeros16(uint16(x)))
}

/**
 * LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_leading_zeros32(1));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_leading_zeros32
func bits_leading_zeros32(x C.uint32_t) C.int {
  return C.int(bits.LeadingZeros32(uint32(x)))
}

/**
 * LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_leading_zeros64(1));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_leading_zeros64
func bits_leading_zeros64(x C.uint64_t) C.int {
  return C.int(bits.LeadingZeros64(uint64(x)))
}

/**
 * LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_leading_zeros8(1));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_leading_zeros8
func bits_leading_zeros8(x C.uint8_t) C.int {
  return C.int(bits.LeadingZeros8(uint8(x)))
}

/**
 * Len returns the minimum number of bits required to represent x; the result is 0 for x == 0. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_len(8));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_len
func bits_len(x C.uint) C.int {
  return C.int(bits.Len(uint(x)))
}

/**
 * Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_len16(8));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_len16
func bits_len16(x C.uint16_t) C.int {
  return C.int(bits.Len16(uint16(x)))
}

/**
 * Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_len32(8));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_len32
func bits_len32(x C.uint32_t) C.int {
  return C.int(bits.Len32(uint32(x)))
}

/**
 * Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_len64(8));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_len64
func bits_len64(x C.uint64_t) C.int {
  return C.int(bits.Len64(uint64(x)))
}

/**
 * Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_len8(8));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_len8
func bits_len8(x C.uint8_t) C.int {
  return C.int(bits.Len8(uint8(x)))
}

/**
 * Mul returns the full-width product of x and y: (hi, lo) = x * y with the product bits'
 * upper half returned in hi and the lower half returned in lo.
 * 
 * This function's execution time does not depend on the inputs.
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_mul_t* mul = bits_mul(12, 12);
 *  printf("%d\n", mul->hi);
 *  printf("%d", mul->lo);
 *  free(mul);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @return bits_mul_t
 */
//export bits_mul
func bits_mul(x C.uint, y C.uint) *C.bits_mul_t {
  self := (*C.bits_mul_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_mul_t{}))))
  re, se := bits.Mul(uint(x), uint(y))
  self.hi = C.uint(re)
  self.lo = C.uint(se)
  return self
}

/**
 * Mul32 returns the 64-bit product of x and y: (hi, lo) = x * y with the product bits'
 * upper half returned in hi and the lower half returned in lo. 
 * 
 * This function's execution time does not depend on the inputs.
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_mul32_t* mul = bits_mul32(12, 12);
 *  printf("%d\n", mul->hi);
 *  printf("%d", mul->lo);
 *  free(mul);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @return bits_mul32_t
 */
//export bits_mul32
func bits_mul32(x C.uint32_t, y C.uint32_t) *C.bits_mul32_t {
  self := (*C.bits_mul32_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_mul32_t{}))))
  re, se := bits.Mul32(uint32(x), uint32(y))
  self.hi = C.uint32_t(re)
  self.lo = C.uint32_t(se)
  return self
}

/**
 * Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y with the product bits'
 * upper half returned in hi and the lower half returned in lo.
 * 
 * This function's execution time does not depend on the inputs.
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_mul64_t* mul = bits_mul64(12, 12);
 *  printf("%d\n", mul->hi);
 *  printf("%d", mul->lo);
 *  free(mul);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @return bits_mul64_t
 */
//export bits_mul64
func bits_mul64(x C.uint64_t, y C.uint64_t) *C.bits_mul64_t {
  self := (*C.bits_mul64_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_mul64_t{}))))
  re, se := bits.Mul64(uint64(x), uint64(y))
  self.hi = C.uint64_t(re)
  self.lo = C.uint64_t(se)
  return self
}

/**
 * OnesCount returns the number of one bits ("population count") in x.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_ones_count(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_ones_count
func bits_ones_count(x C.uint) C.int {
  return C.int(bits.OnesCount(uint(x)))
}

/**
 * OnesCount16 returns the number of one bits ("population count") in x.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_ones_count16(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_ones_count16
func bits_ones_count16(x C.uint16_t) C.int {
  return C.int(bits.OnesCount16(uint16(x)))
}

/**
 * OnesCount32 returns the number of one bits ("population count") in x. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_ones_count32(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_ones_count32
func bits_ones_count32(x C.uint32_t) C.int {
  return C.int(bits.OnesCount32(uint32(x)))
}

/**
 * OnesCount64 returns the number of one bits ("population count") in x.  
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_ones_count64(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_ones_count64
func bits_ones_count64(x C.uint64_t) C.int {
  return C.int(bits.OnesCount64(uint64(x)))
}

/**
 * OnesCount8 returns the number of one bits ("population count") in x.   
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_ones_count8(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_ones_count8
func bits_ones_count8(x C.uint8_t) C.int {
  return C.int(bits.OnesCount8(uint8(x)))
}

/**
 * Rem returns the remainder of (hi, lo) divided by y. Rem panics for y == 0 (division by zero) but
 * unlike Div, it doesn't panic on a quotient overflow.
 *
 * @param hi
 * @param lo
 * @param y
 * @return unsigned int
 */
//export bits_rem
func bits_rem(hi C.uint, lo C.uint, y C.uint) C.uint {
  return C.uint(bits.Rem(uint(hi), uint(lo), uint(y)))
}

/**
 * Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics for y == 0 (division by zero) but,
 * unlike Div32, it doesn't panic on a quotient overflow.
 *
 * @param hi
 * @param lo
 * @param y
 * @return uint32_t
 */
//export bits_rem32
func bits_rem32(hi C.uint32_t, lo C.uint32_t, y C.uint32_t) C.uint32_t {
  return C.uint32_t(bits.Rem32(uint32(hi), uint32(lo), uint32(y)))
}

/**
 * Rem64 returns the remainder of (hi, lo) divided by y. Rem64 panics for y == 0 (division by zero) but,
 * unlike Div64, it doesn't panic on a quotient overflow. 
 *
 * @param hi
 * @param lo
 * @param y
 * @return uint64_t
 */
//export bits_rem64
func bits_rem64(hi C.uint64_t, lo C.uint64_t, y C.uint64_t) C.uint64_t {
  return C.uint64_t(bits.Rem64(uint64(hi), uint64(lo), uint64(y)))
}

/**
 * Reverse returns the value of x with its bits in reversed order.   
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse(19));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return unsigned int
 */
//export bits_reverse
func bits_reverse(x C.uint) C.uint {
  return C.uint(bits.Reverse(uint(x)))
}

/**
 * Reverse16 returns the value of x with its bits in reversed order.    
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse16(19));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint16_t
 */
//export bits_reverse16
func bits_reverse16(x C.uint16_t) C.uint16_t {
  return C.uint16_t(bits.Reverse16(uint16(x)))
}

/**
 * Reverse32 returns the value of x with its bits in reversed order.    
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse32(19));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint32_t
 */
//export bits_reverse32
func bits_reverse32(x C.uint32_t) C.uint32_t {
  return C.uint32_t(bits.Reverse32(uint32(x)))
}

/**
 * Reverse64 returns the value of x with its bits in reversed order.    
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse64(19));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint64_t
 */
//export bits_reverse64
func bits_reverse64(x C.uint64_t) C.uint64_t {
  return C.uint64_t(bits.Reverse64(uint64(x)))
}

/**
 * Reverse8 returns the value of x with its bits in reversed order.    
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse8(19));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint8_t
 */
//export bits_reverse8
func bits_reverse8(x C.uint8_t) C.uint8_t {
  return C.uint8_t(bits.Reverse8(uint8(x)))
}

/**
 * ReverseBytes returns the value of x with its bytes in reversed order.    
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse_bytes(15));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return unsigned int
 */
//export bits_reverse_bytes
func bits_reverse_bytes(x C.uint) C.uint {
  return C.uint(bits.ReverseBytes(uint(x)))
}

/**
 * ReverseBytes16 returns the value of x with its bytes in reversed order.    
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse_bytes16(15));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint16_t
 */
//export bits_reverse_bytes16
func bits_reverse_bytes16(x C.uint16_t) C.uint16_t {
  return C.uint16_t(bits.ReverseBytes16(uint16(x)))
}

/**
 * ReverseBytes32 returns the value of x with its bytes in reversed order.    
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse_bytes32(15));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint32_t
 */
//export bits_reverse_bytes32
func bits_reverse_bytes32(x C.uint32_t) C.uint32_t {
  return C.uint32_t(bits.ReverseBytes32(uint32(x)))
}

/**
 * ReverseBytes64 returns the value of x with its bytes in reversed order.    
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_reverse_bytes64(15));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return uint64_t
 */
//export bits_reverse_bytes64
func bits_reverse_bytes64(x C.uint64_t) C.uint64_t {
  return C.uint64_t(bits.ReverseBytes64(uint64(x)))
}

/**
 * RotateLeft returns the value of x rotated left by (k mod UintSize) bits.
 * To rotate x right by k bits, call RotateLeft(x, -k).
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_rotate_left(15, 2));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param k
 * @return unsigned int
 */
//export bits_rotate_left
func bits_rotate_left(x C.uint, k C.int) C.uint {
  return C.uint(bits.RotateLeft(uint(x), int(k)))
}

/**
 * RotateLeft16 returns the value of x rotated left by (k mod 16) bits.
 * To rotate x right by k bits, call RotateLeft16(x, -k).
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_rotate_left16(15, 2));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param k
 * @return uint16_t
 */
//export bits_rotate_left16
func bits_rotate_left16(x C.uint16_t, k C.int) C.uint16_t {
  return C.uint16_t(bits.RotateLeft16(uint16(x), int(k)))
}

/**
 * RotateLeft32 returns the value of x rotated left by (k mod 32) bits.
 * To rotate x right by k bits, call RotateLeft32(x, -k).
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_rotate_left32(15, 2));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param k
 * @return uint32_t
 */
//export bits_rotate_left32
func bits_rotate_left32(x C.uint32_t, k C.int) C.uint32_t {
  return C.uint32_t(bits.RotateLeft32(uint32(x), int(k)))
}

/**
 * RotateLeft64 returns the value of x rotated left by (k mod 64) bits.
 * To rotate x right by k bits, call RotateLeft64(x, -k).
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_rotate_left64(15, 2));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param k
 * @return uint64_t
 */
//export bits_rotate_left64
func bits_rotate_left64(x C.uint64_t, k C.int) C.uint64_t {
  return C.uint64_t(bits.RotateLeft64(uint64(x), int(k)))
}

/**
 * RotateLeft8 returns the value of x rotated left by (k mod 8) bits.
 * To rotate x right by k bits, call RotateLeft8(x, -k).
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d\n", bits_rotate_left8(15, 2));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param k
 * @return uint8_t
 */
//export bits_rotate_left8
func bits_rotate_left8(x C.uint8_t, k C.int) C.uint8_t {
  return C.uint8_t(bits.RotateLeft8(uint8(x), int(k)))
}

/**
 * Sub returns the difference of x, y and borrow: diff = x - y - borrow. The borrow input must be 0 or 1;
 * otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_sub_t* sub = bits_sub(23, 12, 0);
 *  printf("%d\n", sub->diff);
 *  printf("%d", sub->borrow_out);
 *  free(sub);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @param borrow
 * @return bits_sub_t
 */
//export bits_sub
func bits_sub(x C.uint, y C.uint, borrow C.uint) *C.bits_sub_t {
  self := (*C.bits_sub_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_sub_t{}))))
  re, se := bits.Sub(uint(x), uint(y), uint(borrow))
  self.diff = C.uint(re)
  self.borrow_out = C.uint(se)
  return self
}

/**
 * Sub32 returns the difference of x, y and borrow, diff = x - y - borrow. The borrow input must be 0 or 1;
 * otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_sub32_t* sub = bits_sub32(23, 12, 0);
 *  printf("%d\n", sub->diff);
 *  printf("%d", sub->borrow_out);
 *  free(sub);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @param borrow
 * @return bits_sub32_t
 */
//export bits_sub32
func bits_sub32(x C.uint32_t, y C.uint32_t, borrow C.uint32_t) *C.bits_sub32_t {
  self := (*C.bits_sub32_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_sub32_t{}))))
  re, se := bits.Sub32(uint32(x), uint32(y), uint32(borrow))
  self.diff = C.uint32_t(re)
  self.borrow_out = C.uint32_t(se)
  return self
}

/**
 * Sub64 returns the difference of x, y and borrow: diff = x - y - borrow. The borrow input must be 0 or 1;
 * otherwise the behavior is undefined. The borrowOut output is guaranteed to be 0 or 1.
 *
 * This function's execution time does not depend on the inputs. 
 *
 * Example:
 * * *
 * int main()
 * {
 *  bits_sub64_t* sub = bits_sub64(23, 12, 0);
 *  printf("%d\n", sub->diff);
 *  printf("%d", sub->borrow_out);
 *  free(sub);
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @param y
 * @param borrow
 * @return bits_sub64_t
 */
//export bits_sub64
func bits_sub64(x C.uint64_t, y C.uint64_t, borrow C.uint64_t) *C.bits_sub64_t {
  self := (*C.bits_sub64_t) (C.malloc(C.size_t(unsafe.Sizeof(C.bits_sub64_t{}))))
  re, se := bits.Sub64(uint64(x), uint64(y), uint64(borrow))
  self.diff = C.uint64_t(re)
  self.borrow_out = C.uint64_t(se)
  return self
}

/**
 * TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d", bits_trailing_zeros(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_trailing_zeros
func bits_trailing_zeros(x C.uint) C.int {
  return C.int(bits.TrailingZeros(uint(x)))
}

/**
 * TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d", bits_trailing_zeros16(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_trailing_zeros16
func bits_trailing_zeros16(x C.uint16_t) C.int {
  return C.int(bits.TrailingZeros16(uint16(x)))
}

/**
 * TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d", bits_trailing_zeros32(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_trailing_zeros32
func bits_trailing_zeros32(x C.uint32_t) C.int {
  return C.int(bits.TrailingZeros32(uint32(x)))
}

/**
 * TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d", bits_trailing_zeros64(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_trailing_zeros64
func bits_trailing_zeros64(x C.uint64_t) C.int {
  return C.int(bits.TrailingZeros64(uint64(x)))
}

/**
 * TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
 *
 * Example:
 * * *
 * int main()
 * {
 *  printf("%d", bits_trailing_zeros8(14));
 *  return 0;
 * }
 * * *
 *
 * @param x
 * @return int
 */
//export bits_trailing_zeros8
func bits_trailing_zeros8(x C.uint8_t) C.int {
  return C.int(bits.TrailingZeros8(uint8(x)))
}

func main() {}
