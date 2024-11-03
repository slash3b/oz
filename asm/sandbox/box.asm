section .text

_start:
    mov edx, 10 ; how many bytes do we want to print.
    mov ecx, msg ; loaded with address of first char from msg.
    mov ebx, 1 ; write to stdout
    mov eax, 4 ; syscall to sys_write
    int 80h ; request interrupt

    ; note above a,b,c,d sequence.

    xor ebx, ebx; 0 return
    mov eax, 1; syscall to sys_exit
    int 80h ;request an interrupt on libc using INT 80h.


section .data
msg db 'Hello, new brave world!'

global _start
