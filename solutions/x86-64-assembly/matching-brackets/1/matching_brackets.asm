section .data
stack: db 0
section .text
global is_paired
is_paired:
    lea r8, [rel stack]
    mov byte [r8], 0
    mov rsi, rdi
    xor rcx, rcx
.loop:
    lodsb
    test al, al
    jz .done
    cmp al, '['
    jz .append
    cmp al, '{'
    jz .append
    cmp al, '('
    jz .append
    cmp al, ']'
    jz .eval1
    cmp al, '}'
    jz .eval2
    cmp al, ')'
    jz .eval3
    jmp .loop
.append:
    mov byte [r8 + rcx], al
    inc rcx
    jmp .loop
.eval1:
    dec rcx
    mov al, byte [r8 + rcx]
    cmp al, '['
    jnz .false
    mov byte [r8 + rcx], 0
    jmp .loop
.eval2:
    dec rcx
    mov al, byte [r8 + rcx]
    cmp al, '{'
    jnz .false
    mov byte [r8 + rcx], 0
    jmp .loop
.eval3:
    dec rcx
    mov al, byte [r8 + rcx]
    cmp al, '('
    jnz .false
    mov byte [r8 + rcx], 0
    jmp .loop
.done:
    mov al, byte [r8]
    test al, al
    jnz .false
    mov rax, 1
    ret
.false:
    mov rax, 0
    ret