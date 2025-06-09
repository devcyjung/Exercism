section .text
global reverse
reverse:
    push rdi
    mov rsi, rdi
    
.str_len:
    cmp byte [rsi], 0
    je .str_len_done
    inc rsi
    jmp .str_len

.str_len_done:
    dec rsi

.swap_loop:
    cmp rdi, rsi
    jge .swap_loop_done
    mov ah, [rdi]
    mov al, [rsi]
    mov [rdi], al
    mov [rsi], ah
    inc rdi
    dec rsi
    jmp .swap_loop

.swap_loop_done:
    pop rdi
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
