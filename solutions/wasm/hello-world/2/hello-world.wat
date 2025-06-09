(module
  (memory (export "mem") 1)

  (data (i32.const 0) "Hello, World!")

  (func (export "hello") (result i32 i32)
    (i32.const 0) (i32.const 13)
  )
)
