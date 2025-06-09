(module
  ;;
  ;; Determine if a year is a leap year
  ;;
  ;; @param {i32} year - The year to check
  ;;
  ;; @returns {i32} 1 if leap year, 0 otherwise
  ;;
  (func (export "isLeap") (param $year i32) (result i32)
    local.get $year
    i32.const 400
    i32.rem_u
    i32.eqz
    (if
      (then
        i32.const 1
        return
      )
    )
    local.get $year
    i32.const 4
    i32.rem_u
    i32.eqz       ;; stack: [year % 4 == 0]
        
    local.get $year
    i32.const 100
    i32.rem_u
    i32.eqz       ;; stack: [year % 4 == 0, year % 100 == 0]
    
    i32.xor       ;; stack: [year % 4 ==0 && year % 100 != 0]
    
    return
  )
)
