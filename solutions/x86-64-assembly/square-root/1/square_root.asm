section .text
global square_root
square_root:
    xor rax, rax
.loop:
    mov r8, rax
    mul r8
    cmp rax, rdi
    mov rax, r8
    jz .done
    inc rax
    jmp .loop
.done:
    ret