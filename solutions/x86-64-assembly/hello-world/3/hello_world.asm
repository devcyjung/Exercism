section .text
    global hello

hello:
    ; --- mmap(NULL, hello_len, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0)
    xor rdi, rdi                ; addr = NULL
    mov rsi, hello_len          ; length
    mov rdx, 0x3                ; PROT_READ | PROT_WRITE
    mov r10, 0x22               ; MAP_PRIVATE | MAP_ANONYMOUS
    mov r8, -1                  ; fd = -1
    xor r9, r9                  ; offset = 0
    mov rax, 9                  ; syscall number for mmap
    syscall                     ; returns pointer in RAX

    ; --- Copy "Hello, World!\0" into allocated memory
    lea rsi, [rel hello_msg]    ; source pointer
    mov rdi, rax                ; destination pointer
    mov rdx, rax                ; save return address

.copy_loop:
    lodsb
    test al, al                 ; if null terminator reached, done
    stosb
    jnz .copy_loop

    mov rax, rdx
    ret

section .rodata
hello_msg: db "Hello, World!", 0
hello_len: equ $ - hello_msg