section .text
global square_of_sum
square_of_sum:
    xor    rax, rax
    mov    rdx, rdi
    cmp    rdx, 0
    jl     .negative
    jz     .done
.positive:
    test   rdx, rdx
    jz     .done
    add    rax, rdx
    dec    rdx
    jmp    .positive
.negative:
    test   rdx, rdx
    jz     .done
    add    rax, rdx
    inc    rdx
    jmp    .negative
.done:
    mul    rax
    ret

global sum_of_squares
sum_of_squares:
    xor    rax, rax
    mov    rdx, rdi
    cmp    rdx, 0
    jz     .done
    jg     .positive
    neg    rdx
.positive:
    test   rdx, rdx
    jz     .done
    mov    r8, rdx
    imul   r8, r8
    add    rax, r8
    dec    rdx
    jmp    .positive
.done:
    ret

global difference_of_squares
difference_of_squares:
    push   rdi
    call   sum_of_squares
    pop    rdi
    push   rax
    call   square_of_sum
    pop    rdx
    sub    rax, rdx
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif