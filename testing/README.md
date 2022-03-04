# gotest

- Black-box Testing: Use package _packagename_test_, which will ensure you're only using the exported identifiers.

- White-box Testing: Use package _packagename_ so that you have access to the non-exported identifiers. Good for unit tests that require access to non-exported variables, functions, and methods.

- Example Testing: Unlike normal test functions, example functions take no arguments and begin with the work *Example* instead of *Test*.
