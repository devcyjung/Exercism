section .text
global leap_year
leap_year:
    mov rax, rdi
    mov rdx, 0
    mov rcx, 4
    div rcx
    cmp rdx, 0
    jne false_return
    mov rax, rdi
    mov rdx, 0
    mov rcx, 400
    div rcx
    cmp rdx, 0
    je true_return
    mov rax, rdi
    mov rdx, 0
    mov rcx, 100
    div rcx
    cmp rdx, 0
    je false_return
true_return:
    mov rax, 1
    ret
false_return:
    mov rax, 0
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
