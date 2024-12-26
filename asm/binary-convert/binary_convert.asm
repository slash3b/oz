section .text
global binary_convert
binary_convert:
    xor eax, eax ; reset to 0

    ; 0x30 == 0
    ; 0x31 == 1

.loop:
    movzx ecx, byte [rdi] ; read byte from rdi and zero out the rest;
    ; how do we iterate to the next byte? answer — we increment rdi

    cmp ecx, 0; compare that byte to 0 byte
    ; unclear what this 0 is, decimal, binary, hexadecimal? 
    ; I think it is just null sort of.
    je .end ; jump to the end. This is our loop break here.

    shl eax, 1; left byte shift before we add anything, e.g. value << 1
    ; we shift on _every_ iteration.

    ; so eax contains a byte. Byte! You hear me? so it will be 
    ; one of those 0x30, 0x31, 0x0. See ascii table.
    ; assuming that incoming string is not malformed.

    cmp ecx, 0x31
    je .apply ; add to accumulator only when we have 1 on our hands.

    inc rdi
    jmp .loop

.apply:
    add eax, 1
    inc rdi
    jmp .loop
    ret

.end:
	ret
