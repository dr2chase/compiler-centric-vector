# Vector

Do not depend on this code; it is for discussion only and subject to change as I/we learn better ways to do this.

This is a proposed way of expressing vector operations in Go programs.  Rather than introduce new types,
instead introduce functions over slices that are can be treated specially, either by the compiler, or
perhaps by a preprocessor.  There are several paths to good performance, but my favorite approach would
be for the compiler to recognize "expression trees" over these functions, and

- Declare that the order of evaluation across the slices is not guaranteed, and aliasing inputs and outputs is a potential source of nonportable surprises.  (This implies perhaps a checking mode that would diagnose this issue).
- Hoist all the checks outside the computation (this may or may not be possible without special treatment for these checks).
   It might make sense to rewrite the reference implementation into a checked set of functions and interior intrinsics, and
   a combination of inlining, common expression elimination and value propagation might be adequate.
- Recognize all the slices that exist only as temporaries, as temporaries.
- Vectorize and chunk across the expression tree.

