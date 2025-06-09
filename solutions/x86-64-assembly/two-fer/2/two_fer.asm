section .data
prefix: db "One for ", 0
default_name: db "you", 0
suffix: db ", one for me.", 0
section .text
global two_fer
two_fer:
    push rdi
    mov rdi, rsi
    lea rsi, [rel prefix]
    call write
    pop rsi
    test rsi, rsi
    jz .null
    mov al, byte [rsi]
    test al, al
    jz .null
.continue:
    call write
    lea rsi, [rel suffix]
    call write
    xor al, al
    stosb
    ret
.null:
    lea rsi, [rel default_name]
    jmp .continue

write:
    lodsb
    test al, al
    jz .done
    stosb
    jmp write
.done:
    ret