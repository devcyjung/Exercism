section .text
global can_create
can_create:
    cmp edi, 0
    jl .false
    cmp edi, 7
    jg .false
    cmp esi, 0
    jl .false
    cmp esi, 7
    jg .false
    mov rax, 1
    ret
.false:
    mov rax, 0
    ret

global can_attack
can_attack:
    call can_create
    test rax, rax
    jz .false
    mov r8, rdi
    mov r9, rsi
    mov rdi, rdx
    mov rsi, rcx
    call can_create
    test rax, rax
    jz .false
    cmp rdi, r8
    jz .true
    cmp rsi, r9
    jz .true
    sub rdi, r8
    sub rsi, r9
    cmp rdi, rsi
    jz .true
    add rdi, rsi
    test rdi, rdi
    jz .true
.false:
    mov rax, 0
    ret
.true:
    mov rax, 1
    ret