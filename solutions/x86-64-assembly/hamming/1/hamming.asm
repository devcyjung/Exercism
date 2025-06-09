section .text
global distance
distance:
    xor rax, rax
.loop:
    movzx rbx, byte[rdi]
    mov bh, byte[rsi]
    test bx, bx
    jz .done
    test bh, bh
    jz .invalid
    test bl, bl
    jz .invalid
    inc rdi
    inc rsi
    cmp bh, bl
    jz .loop
    inc rax
    jmp .loop
.invalid:
    mov rax, -1
.done:
    ret