section .text
global pangram

; my original C code
; ch = *p > 0x40 && *p < 0x5b ? *p | 0x20 : *p;
; if (ch > 0x60 && ch < 0x7b) {
;   res |=  1 << ch - 0x61;
; }

pangram:
    ; rax contains the anaswer, 1 or 0.
    xor rax, rax
    ; rdx contains the bit set.
    xor rdx, rdx

.loop:
    movzx rcx, byte [rdi]

    ; if we have read 0 (null), break
    cmp rcx, 0x0
    je .end

    inc rdi ; safe to increment rdi

    ; way to lowercase
    cmp rcx, 'a' ; 0x60 is "`", right before "a"
    jl .tolowercase

    ; if less than 'a' -- continue
    ; cmp rcx, 0x61
    ; jl .loop

    ; if greater than 'z' -- continue
    ; cmp rcx, 0x7a
    ; jg .loop

    sub rcx, 0x61
    xor r8, r8
    mov r8, 1
    shl r8, cl
    or rdx, r8
    ; bts rdx, rcx

    jmp .loop

.tolowercase:
        add rcx, 0x20

.end:
    mov rax, rdx
    ;cmp rdx, 0x03ffffff
    ;je .ispangram


    ; mov rax, 42
    ;cmp rax, 0x03ffffff ; 26 bits
    ; 3ffffff or 67108863 | 26 bits

    ;je .itispangram
    ;jl .nullify

	ret

.ispangram:
   mov rax, 1
   ret

