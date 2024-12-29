default rel

section .data
    pi dd 3.14159
    v dd 3.00

section .text
global volume


volume:
    ; how to hardcode pi?
    ; %define PI 3.14159

    ; what registries hold radius, height?
    ; xmm registries xmm0 - xmm15 hold 16 bytes of data
    ; ymm registries ymm0 - ymm15 hold 32 bytes of data
    ; xmm0 holds the first param -- radius
    ; xmm1 holds the second param -- height

    ; what registry is used to hold return value?
    ; xmm0

    ; formula
    ; r * r * PI * (h/3)

    divss xmm1, [v] ; h/3

    xorps xmm3, xmm3
    movss xmm3, xmm0 ; r
    mulss xmm3, xmm3 ; .. * r
    mulss xmm3, [pi] ; .. * PI
    mulss xmm3, xmm1 ; .. * (h/3)

    ; could have used xmm0
    movss xmm0, xmm3 ; final response

 	ret

