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
    ; rbx contains the bit set.
    xor rbx, rbx

.loop:
    movzx rbx, byte [rdi]

    ; if we have read 0 (null), break
    cmp rbx, 0
    je .end

    ; hacky way to lowercase
    cmp rbx, 'a' ; 0x60 is "`", right before "a"
    jl .tolowercase

    ; if less than 'a' -- continue
    cmp rbx, 'a'
    inc rdi
    jl .loop

    ; if greater than 'z' -- continue
    cmp rbx, 'z'
    inc rdi
    jg .loop

    sub rbx, 0x61
    xor rcx, rcx
    mov rcx, 1
    ;"bl" are lower bits from rbx 
    ; immediately accessible
    ;
    ; 8 bits : AH AL BH BL CH CL DH DL
    ; the "H" and "L" suffix on the 8 bit registers 
    ; stand for high byte and low byte
    ; mov rax, rbx
    mov bl, [rbx]
    shl rcx, bl ; e.g. 1 << 1 for 'b'. 


    ; actual writing to rax
    or rax, rcx

    inc rdi
    jmp .loop

.tolowercase:
    add rbx, 0x20
    ret

.end:
    cmp rax, 0x03ffffff
    mov rax, 0
    je .itispangram

	ret

.itispangram:
    mov rax, 1

    ret
