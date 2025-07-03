section .text
global append
append:
    mov r9, rdi            ; src1 address
    mov r10, rsi           ; remaining src1 length
    mov r11, rdx           ; src2 address
    mov r12, rcx           ; remaining src2 length
    mov r13, r8            ; dst address
    lea r14, [r10 + r12]
    mov rcx, r10
    mov rsi, r9
    mov rdi, r13
    rep movsd
    lea r13, [r13 + 4 * r10]
    mov rcx, r12
    mov rsi, r11
    mov rdi, r13
    rep movsd
    mov rax, r14
    ret

global filter
filter:
    mov r8, rdi        ; src address
    mov r9, rdx        ; function pointer
    mov r10, rsi       ; remaining src length
    mov r11, rcx       ; dst address
    xor r12, r12       ; result length
.loop:
    test r10, r10
    jz .done
    dec r10
    mov edi, [r8]
    push rdi
    add r8, 4
    call r9
    test al, al
    pop rdi
    jnz .match
    jmp .loop
.match:
    mov [r11], edi
    add r11, 4
    inc r12
    jmp .loop
.done:
    mov rax, r12 
    ret

global map
map:
    mov r8, rdi        ; src address
    mov r9, rdx        ; function pointer
    mov r10, rsi       ; remaining src length
    mov r11, rcx       ; dst address
    mov r12, r10       ; result length
.loop:
    test r10, r10
    jz .done
    dec r10
    mov edi, [r8]
    add r8, 4
    call r9
    mov [r11], eax
    add r11, 4
    jmp .loop
.done:
    mov rax, r12 
    ret

global foldl
foldl:
    mov r8, rdi        ; src address
    mov r9, rsi        ; remaining src length
    mov r10d, edx      ; accumulate
    mov r11, rcx       ; function pointer
.loop:
    test r9, r9
    jz .done
    dec r9
    mov edi, r10d
    mov esi, [r8]
    add r8, 4
    call r11
    mov r10d, eax
    jmp .loop
.done:
    mov eax, r10d
    ret

global foldr
foldr:
    lea r8, [rdi + 4 * (rsi - 1)]  ; src address
    mov r9, rsi                    ; remaining src length
    mov r10d, edx                  ; accumulate
    mov r11, rcx                   ; function pointer
.loop:
    test r9, r9
    jz .done
    dec r9
    mov edi, r10d
    mov esi, [r8]
    sub r8, 4
    call r11
    mov r10d, eax
    jmp .loop
.done:
    mov eax, r10d
    ret

global reverse
reverse:
    lea r8, [rdi + 4 * (rsi - 1)]  ; src address
    mov r9, rsi                    ; remaining src length
    mov r10, rdx                   ; dst address
    mov r11, r9                    ; result length
.loop:
    test r9, r9
    jz .done
    dec r9
    mov eax, [r8]
    mov [r10], eax
    sub r8, 4
    add r10, 4
    jmp .loop
.done:
    mov rax, r11
    ret