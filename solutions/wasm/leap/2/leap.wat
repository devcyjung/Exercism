(module
  (func (export "isLeap") (param $year i32) (result i32)
    (i32.and 
      (i32.eqz (i32.rem_u (local.get $year) (i32.const 4)))
      (i32.eqz
        (i32.xor
          (i32.eqz (i32.rem_u (local.get $year) (i32.const 100)))
          (i32.eqz (i32.rem_u (local.get $year) (i32.const 400)))
        )
      )
    )
  )
)