package trace_test

import (
	trace "trace"
)

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}

func c() {
	defer trace.Trace()()
	d()
}

func d() {
	defer trace.Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	// g[00001]:    ->trace/example_test.a
	// g[00001]:        ->trace/example_test.b
	// g[00001]:            ->trace/example_test.c
	// g[00001]:                ->trace/example_test.d
	// g[00001]:                <-trace/example_test.d
	// g[00001]:            <-trace/example_test.c
	// g[00001]:        <-trace/example_test.b
	// g[00001]:    <-trace/example_test.a
}
