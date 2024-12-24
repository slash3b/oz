section .text
global index
index:
	; rdi: matrix
	; esi: rows
	; edx: cols
	; ecx: rindex
	; r8d: cindex

    imul ecx, edx ; rowindex * columns
    add ecx, r8d
    mov eax, [rdi + rcx * 4]
	ret
