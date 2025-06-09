section .text
global steps
steps:
    mov rcx, 0    ; int iter = 0
    test edi, edi
    jle .wrong       ; if n <= 0 return -1

.loop:
    cmp rdi, 1    
    je .found     ; while n > 1
    test rdi, 1   
    jz .even      ; if n % 2 == 0 (if n & 1 == 0)
    lea rdi, [rdi*2+rdi]
    inc rdi       ; else n = 3*n+1 (if n % 2 == 1)
    jmp .inc

.even:
    shr rdi, 1    ; n /= 2 (if n % 2 == 0)
    jmp .inc

.inc:
    inc rcx       ; iter += 1
    jmp .loop

.found:
    mov rax, rcx  ; return iter
    ret  

.wrong:
    mov rax, -1
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
