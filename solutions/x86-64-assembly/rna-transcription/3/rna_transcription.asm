default rel
section .rodata
encoding:
    times 65 db 0
    db "U_G___C____________A______"
section .text
global to_rna
to_rna:
    lea r8, [encoding]
.loop:
    movzx rax, byte [rdi]
    mov al, [r8 + rax]
    mov [rsi], al
    inc rdi
    inc rsi
    test al, al
    jnz .loop
    ret