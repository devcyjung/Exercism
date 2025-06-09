section .text
global is_pangram
is_pangram:
    sub rsp, 32
    lea r8, [rsp]
    xor r9, r9
    mov rsi, rdi
.zero_pangram:
    mov byte [r8 + r9], 0
    inc r9
    cmp r9, 26
    jl .zero_pangram
    xor rax, rax
.loop:
    lodsb
    test al, al
    jz .done
    cmp al, 'A'
    jl .loop
    cmp al, 'Z'
    jle .capital
    cmp al, 'a'
    jl .loop
    cmp al, 'z'
    jg .loop
    sub al, 'a'
.continue:
    mov byte [r8 + rax], 1
    jmp .loop
.capital:
    sub rax, 'A'
    jmp .continue
.done:
    xor r9, r9
.check:
    mov al, [r8 + r9]
    test al, al
    jz .false
    inc r9
    cmp r9, 26
    jl .check
.true:
    mov rax, 1
    add rsp, 32
    ret
.false:
    xor rax, rax
    add rsp, 32
    ret