default rel
section .rodata
encoding:
    times 65 db 0
    db "U_G___C____________A______"
section .text
global to_rna
to_rna:
    push rbp
    mov rbp, rsp
    lea r8, [encoding]
    xor r9, r9
    xor r10, r10
.loop:
    movzx rax, byte [rdi + r9]
    test al, al
    jz .done
    inc r9
    movzx rbx, byte [r8 + rax]
    cmp rbx, '_'
    je .invalid
    mov byte[rsi + r10], bl
    inc r10
    jmp .loop
.done:
    mov byte [rsi + r10], 0
    pop rbp
    ret
.invalid:
    mov byte [rsi], 0
    pop rbp
    ret