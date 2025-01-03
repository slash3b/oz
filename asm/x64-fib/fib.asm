section .text
global fib
global fac

fac:
    mov rax, rdi

    cmp rdi, 0
    je .end

    call .factorial

    ret

.factorial:
    cmp rax, 1
    je .end

    push rax
    dec rax

    call .factorial

    ; ---------------------

    pop rcx ; poping 

    imul rax, rcx

.end:
    ret

.end_rec:
    pop rcx
    imul rax, rcx
    ret

; fib:
;     cmp rdi, 0
;     je .end

;     sub rdi, 1

;     push rdi
;     pop rdi

;     add rdi, 10

;     push rdi
;     pop rax

;     ret

; .end:
;     ret

; .fib2:
;     cmp rdi, 0
;     je .end

;     mov rax, 42

;     dec rdi

;     imul rax, rdi

;     jmp .fib2

; section .text
; global fib

; fib:
;     cmp rdi, 2
;     jl .short

;     ; keep variables somewhere

;     mov r8, 0
;     mov r9, 1

;     jmp .recfib

;     ret

; .fib2:
;     cmp rdi, 0
;     je .end

;     dec rdi

;     ; r10 is a temporary storage
;     xor r10, r10
;     mov r10, r8

;     mov r8, r9
;     mov r9, [r10 + r9]

;     jmp .recfib

; .short:
;     mov rax, rdi

;     ret

; .end:
;     xor rax, rax

;     ; final answer

;     add rax, r8
;     add rax, r9

;     ret
