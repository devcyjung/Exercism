section .text
global allergic_to
allergic_to:
    mov cx, di
    sar rsi, cl
    and rsi, 1
    mov rax, rsi
    ret

global list
list:
    mov r8, rdi
    mov r9, rsi
    xor r10, r10
    xor r11, r11
.loop:
    mov rdi, r10
    mov rsi, r8
    call allergic_to
    test al, al
    jnz .add
.continue:
    cmp r10, 7
    je .done
    inc r10
    jmp .loop
.add:
    lea rcx, [r9 + 4 + r11 * 4]
    inc r11
    mov [rcx], r10d
    jmp .continue
.done:
    mov [r9], r11d
    ret