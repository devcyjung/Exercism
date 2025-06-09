section .text
global egg_count
egg_count:
    popcnt rax, rdi
    ret