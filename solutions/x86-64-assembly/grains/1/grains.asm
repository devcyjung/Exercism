section .text
global square
square:
    cmp rdi, 64
    jg .invalid
    cmp rdi, 1
    jl .invalid
    mov rax, 1
    mov rcx, rdi
    sub rcx, 1
    shl rax, cl
    ret
.invalid:
    xor rax, rax
    ret

global total
total:
    xor rax, rax
    not rax
    ret