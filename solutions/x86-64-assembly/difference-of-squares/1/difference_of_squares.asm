section .text
global square_of_sum
square_of_sum:
    push    rbp
    mov     rbp, rsp

    mov     rcx, rdi
    xor     rax, rax

.loop:
    test    rcx, rcx
    jle     .done
    add     rax, rcx
    dec     rcx
    jmp     .loop

.done:
    imul    rax, rax
    pop     rbp
    ret

global sum_of_squares
sum_of_squares:
    push    rbp
    mov     rbp, rsp

    mov     rcx, rdi
    xor     rax, rax

.loop:
    test    rcx, rcx
    jle     .done
    mov     rdx, rcx
    imul    rdx, rdx
    add     rax, rdx
    dec     rcx
    jmp     .loop

.done:
    pop     rbp
    ret

global difference_of_squares
difference_of_squares:
    push    rbp
    mov     rbp, rsp

    call    square_of_sum

    mov     r10, rax
    call    sum_of_squares

    mov     r11, rax
    sub     r10, r11 
    mov     rax, r10

    pop     rbp
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif