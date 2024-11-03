section .data
msg db "Hello beautiful new world!", 0ah

section .text

global _start

_start:
    mov ebx, msg
    mov eax, msg


; so these "labels" do not mean anything, except a point we can jump to under a specific conditions.
nextchar:
    cmp byte [eax], 0 ; compare
    jz finished
    inc eax
    jmp nextchar

finished:
    sub eax, ebx

    ; print message
    mov edx, eax
    mov ecx, msg ; msg
    mov ebx, 1 ; 1 is stdout
    mov eax, 4 ; sys_write
    int 80h

    ; exit with "zero errors"
    mov ebx, 0
    mov eax, 1
    int 80h
