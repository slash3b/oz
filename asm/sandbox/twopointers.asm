section .data
msg db "Hello beautiful new world!", 0ah

section .text

global _start

_start:
    mov ebx, msg ; ebx points to first byte in msg
    mov eax, msg ; eax points to first byte in msg


; so these "labels" do not mean anything, except a point we can jump to under a specific conditions.
nextchar:
    cmp byte [eax], 0 ; compares and throws away result, but fills in "flags", for example it fills in ZF "zero flag" 
    jz finished ; contidional jump "jump zero", does jump when cmp fills in ZF to 0.
    inc eax ; increment eax
    jmp nextchar ; jump back to nextchar label

finished:
    sub eax, ebx ; subtracts eax from ebx, keep result in eax

    ; print message
    mov edx, eax ; num of bytes to print
    mov ecx, msg ; msg
    mov ebx, 1 ; 1 is stdout
    mov eax, 4 ; sys_write
    int 80h

    ; exit with "zero errors"
    mov ebx, 0
    mov eax, 1
    int 80h
