section .text

global sum_to_n

; Which register contains the argument (n)?
; rdi
; In which register will the return value be placed?
; rax?
; What instruction is used to do the addition?
; ADD
; What instruction(s) can be used to construct the loop?
; conditional jump


sum_to_n:
    xor rax, rax
    cmp rdi, 2
    jl simple

    sum: ; loop
    add rax, rdi ; add argument to output
    dec rdi ; decrease argument by 1
    cmp rdi, 0 ; compare argument to 0
    jg sum ; start again if armgument > 0

	ret

simple:
    mov rax, rdi
    ret
