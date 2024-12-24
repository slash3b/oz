section .text
global binary_convert
binary_convert:
    ; go through every char if 1
    xor eax, eax

    ; 0x30 == 0
    ; 0x31 == 1

.loop:
    movzx ecx, byte [rdi] ; read byte from rdi and zero out the rest;
    ; how do we iterate to the next byte??
    ; ecx is now byte...

    cmp ecx, 0; compare that byte to 0 byte
    ; unclear what this 0 is, decimal, binary, hexadecimal?
    je .end ; jump to the end. This is our loop break here.

    shl eax, 1

    ; so eax contains a byte. Byte! You hear me? so it will be either
    ; one of those 0x30, 0x31, 0x0. See ascii table.
    ; assuming that incoming string is not malformed.

    cmp ecx, 0x31
    je .apply

    inc rdi
    jmp .loop

.apply:
    add eax, 1
    inc rdi
    jmp .loop
    ret

.end:
	ret
