section .text
global leap_year
leap_year:
    mov rax, rdi
    mov rdx, 0
    mov rcx, 4
    div rcx
    cmp rdx, 0
    jne is_not_leap
    mov rax, rdi
    mov rdx, 0
    mov rcx, 400
    div rcx
    cmp rdx, 0
    je is_leap
    mov rax, rdi
    mov rdx, 0
    mov rcx, 100
    div rcx
    cmp rdx, 0
    je is_not_leap
is_leap:
    mov rax, 1
    ret
is_not_leap:
    mov rax, 0
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
