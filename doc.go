/*
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

/*
Package pmc provides a Probabilistic Multiplicity Counting Sketch, a novel data structure that is capable of accounting traffic per flow probabilistically, that can be used as an alternative to Count-min sketch.
The stream processing algorithm — Probabilistic Multiplicity Counting (PMC) — uses probabilistic counting techniques to determine the approximate multiplicity of each element in large streams. It is particularly well suited for traffic measurements on high-speed communication links and likewise applicable for many other purposes.

Count-Min Sketches hold counters in a matrix-like organization. A big caveat for both Spectral Bloom Filters and Count-Min Sketches is that the maximum multiplicity has to be known a priori quite accurately, to provide large enough counters without wasting too much memory. PMC does not need to know the maximum frequency beforehand, and its counting operation is much simpler.

For details about the algorithm and citations please use this article:
"High-Speed Per-Flow Traffic Measurement with Probabilistic Multiplicity Counting" by Peter Lieven & Björn Scheuermann
(https://wwwcn.cs.uni-duesseldorf.de/publications/publications/library/Lieven2010a.pdf)
*/
package pmc
