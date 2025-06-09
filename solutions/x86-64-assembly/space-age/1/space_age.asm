default rel
section .rodata
earth_sec: dd 31557600.0
planet_year: dd 0.2408467, 0.61519726, 1.0, 1.8808158,
             dd 11.862615, 29.447498, 84.016846, 164.79132
section .text
global age
age:
    lea r8, [earth_sec]
    cvtsi2ss xmm0, rsi
    divss xmm0, [r8]
    lea r8, [planet_year]
    divss xmm0, [r8 + 4 * rdi]
    ret