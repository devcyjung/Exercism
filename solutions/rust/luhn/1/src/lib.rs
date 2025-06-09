use std::ops::ControlFlow;

pub fn is_valid(code: &str) -> bool {
    let ctrl_flow = code
        .as_bytes()
        .iter()
        .rev()
        .try_fold((0, 0), |(i, acc), byt| {
            match byt {
                b'0'..=b'9' => {
                    let num = byt - b'0';
                    let delta = if i % 2 == 0 {
                        num
                    } else {
                        let doubled = num * 2;
                        if doubled > 9 {
                            doubled - 9
                        } else {
                            doubled
                        }
                    };
                    ControlFlow::Continue(
                        (i + 1, (acc + delta) % 10)
                    )
                },
                b' ' => ControlFlow::Continue((i, acc)),
                _ => ControlFlow::Break(false),
            }
    });
    match ctrl_flow {
        ControlFlow::Break(fail) => fail,
        ControlFlow::Continue((i, result)) => i > 1 && result == 0,
    }
}
