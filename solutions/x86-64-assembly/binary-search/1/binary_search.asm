section .text
global find
find:
    test rdi, rdi
    jz .nullptr
    xor r8d, r8d
    mov r10d, esi
    xor r9, r9
.loop:
    cmp r8d, r10d
    jge .nullptr
    lea r9d, [r8d + r10d - 1]
    shr r9d, 1
    mov eax, dword [rdi + 4 * r9]
    cmp eax, edx
    je .found
    jl .toright
    mov r10d, r9d
    jmp .loop
.toright:
    mov r8d, r9d
    inc r8d
    jmp .loop
.nullptr:
    mov rax, -1
    ret
.found:
    mov rax, r9
    ret