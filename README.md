# Vector

This is a proposed way of expressing vector operations in Go programs.  Rather than introduce new types,
instead introduce functions over slices that can be treated specially, either by the compiler, or
perhaps by a preprocessor.  There are several paths to good performance, but my favored approach would
be for the compiler to recognize "expression trees" over these functions, and

- Declare that the order of evaluation across the slices is not guaranteed, and that aliasing inputs and outputs is a potential source of nonportable surprises.  (This implies perhaps a checking mode that would diagnose this issue).  The current proposal avoids this by "always returning a new slice", except for the possible case of a `copyInto` or `assign` function.
- Hoist all the checks outside the computation (this may or may not be possible without special treatment for these checks).
  It might make sense to rewrite the reference implementation into a checked set of functions and interior intrinsics, and
  a combination of inlining, common expression elimination and value propagation might be adequate.  Or, given that these are intrinsics, use knowledge about intrinsics in the analysis (so, special analysis, but not special semantics).
- Recognize all the slices that exist only as temporaries, as temporaries.
- Vectorize and chunk across the expression tree.

Do not depend on this code; it is for discussion only and subject to change as I/we learn better ways to do this.
The interface is more important than the implementation; the implementation, though it is good if it can be improved,
is intended as a consistency check on the documentation and a proof-of-concept, if there is no fancy compiler support.

Known issues:

 - the dynamic check for same slice length is unsatisfying.  It would be better if the compiler were allowed to complain statically whenever it could detect problems early.
 - there have been (very) informal discussions about both overloading and support for ML (in some vague form).  This change might be helpful to those possible features, or it might be a hindrance.
 - this is not directly helpful to "I want access to SIMD from Go", and whether it is or not probably depends heavily on the quality of the compiler's implementation.

 As an experiment, I tried running a bit of code through the current compiler with inlining turned off to see what it would look like right now in SSA, if there might be a chance of optimizing this, and I think that there could be:

For this input
```
  x := []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}
  ones := Fill(len(x), int8(1))
  y := Times(Plus(x, ones), Minus(x, ones))
  z := Minus(Times(x, x), ones)
```
the two lines initializing y and z turn into this
```
v66 (+45) = StaticLECall <[]go.shape.int8_0,mem> {AuxCall{github.com/dr2chase/vector.Fill[go.shape.int8_0]}} [24] v64 v63 v10 v55
v67 (45) = SelectN <mem> [1] v66
v68 (45) = SelectN <[]go.shape.int8_0> [0] v66 (ones[[]int8])
v70 (+46) = StaticLECall <[]go.shape.int8_0,mem> {AuxCall{github.com/dr2chase/vector.Plus[go.shape.int8_0]}} [56] v69 v62 v68 v67
v71 (46) = SelectN <mem> [1] v70
v72 (46) = SelectN <[]go.shape.int8_0> [0] v70
v74 (46) = StaticLECall <[]go.shape.int8_0,mem> {AuxCall{github.com/dr2chase/vector.Minus[go.shape.int8_0]}} [56] v73 v62 v68 v71
v75 (46) = SelectN <mem> [1] v74
v76 (46) = SelectN <[]go.shape.int8_0> [0] v74
v78 (46) = StaticLECall <[]go.shape.int8_0,mem> {AuxCall{github.com/dr2chase/vector.Times[go.shape.int8_0]}} [56] v77 v72 v76 v75
v79 (46) = SelectN <mem> [1] v78
v80 (46) = SelectN <[]go.shape.int8_0> [0] v78 (y[[]int8])
v82 (+47) = StaticLECall <[]go.shape.int8_0,mem> {AuxCall{github.com/dr2chase/vector.Times[go.shape.int8_0]}} [56] v81 v62 v62 v79
v83 (47) = SelectN <mem> [1] v82
v84 (47) = SelectN <[]go.shape.int8_0> [0] v82
v86 (47) = StaticLECall <[]go.shape.int8_0,mem> {AuxCall{github.com/dr2chase/vector.Minus[go.shape.int8_0]}} [56] v85 v84 v68 v83
v87 (47) = SelectN <mem> [1] v86
v88 (47) = SelectN <[]go.shape.int8_0> [0] v86 (z[[]int8])
```
Using knowledge about the intrinsic semantics, the bounds checks for `x` and `ones` can trivially be hoisted before the calls to `Plus` and `Minus`, those results have known size equal to the sizes of their inputs and thus need no further checking, and all the checks have been done before the calls that compute `z`, so without changes to Go semantics, the entire graph of calls and results can be done as unbroken vector operations (i.e., SIMD, or software pipelined, etc).
