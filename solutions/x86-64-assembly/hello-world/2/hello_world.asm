section .text
    global hello

hello:
    ; --- mmap(NULL, 4096, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0)
    xor rdi, rdi                ; addr = NULL
    mov rsi, 4096               ; length
    mov rdx, 0x3                ; PROT_READ | PROT_WRITE
    mov r10, 0x22               ; MAP_PRIVATE | MAP_ANONYMOUS
    mov r8, -1                  ; fd = -1
    xor r9, r9                  ; offset = 0
    mov rax, 9                  ; syscall number for mmap
    syscall                     ; returns pointer in RAX

    ; --- Copy "Hello, World!\0" into allocated memory
    mov rdi, rax                ; save base ptr
    lea rcx, [rel hello_msg]    ; source pointer
    mov rsi, rdi                ; destination pointer

.copy_loop:
    mov al, byte [rcx]
    mov [rsi], al
    inc rcx
    inc rsi
    test al, al                 ; if null terminator reached, done
    jnz .copy_loop

    ; RAX still contains the pointer => return
    ret

section .rodata
hello_msg: db "Hello, World!", 0
