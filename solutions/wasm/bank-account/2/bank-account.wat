(module
  (global $open (mut i32) (i32.const 0))
  (global $balance (mut i32) (i32.const 0))

  (func (export "open") (result i32)
    (i32.eqz (global.get $open))
    (if (then
      (global.set $open (i32.const 1))
      (return (i32.const 0))
    ))
    (return (i32.const -1))
  )

  (func (export "close") (result i32)
    (i32.eqz (global.get $open))
    (if (then
      (return (i32.const -1))
    ))
    (global.set $open (i32.const 0))
    (global.set $balance (i32.const 0))
    (return (i32.const 0))
  )

  (func (export "deposit") (param $amount i32) (result i32)
    (i32.eqz (global.get $open))
    (if (then
      (return (i32.const -1))
    ))
    (i32.lt_s (local.get $amount) (i32.const 0))
    (if (then
      (return (i32.const -2))
    ))
    (global.set $balance (i32.add (global.get $balance) (local.get $amount)))
    (return (i32.const 0))
  )

  (func (export "withdraw") (param $amount i32) (result i32)
    (i32.eqz (global.get $open))
    (if (then
      (return (i32.const -1))
    ))
    (i32.or
      (i32.lt_s (local.get $amount) (i32.const 0))
      (i32.lt_s (global.get $balance) (local.get $amount))
    )
    (if (then
      (return (i32.const -2))
    ))
    (global.set $balance (i32.sub (global.get $balance) (local.get $amount)))
    (return (i32.const 0))
  )

  (func (export "balance") (result i32)
    (i32.eqz (global.get $open))
    (if (then
      (return (i32.const -1))
    ))
    (return (global.get $balance))
  )
)