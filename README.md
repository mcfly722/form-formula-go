# Form-Formula-Go
![](https://img.shields.io/badge/version-v0.0.1--alpha-success.svg)
![](https://img.shields.io/badge/tests-69|69-success.svg)
[![](https://img.shields.io/badge/License-GPL3.0-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.html)

Library to iterate over all possible math functions to find acceptable one (GoLang implementation).<br>

### idea
Every mathematical function f(x) could be represented as [Abstract Syntax Tree (AST)](https://en.wikipedia.org/wiki/Abstract_syntax_tree) where nodes of this tree could be an<br>
 - operators (has two input arguments): <b>*, +, -, /, mod, pow(x,p)</b>
 - functions (has one input argument): <b>sqrt(x), sin(x), x!, log(x), exp(x),round(x)</b>

Leafs of this tree are:<br>
- constants or
- argument x
- iteration number
- previous calculated value
   
Complex formulas which contains [&#8721;](https://en.wikipedia.org/wiki/Summation),[&#8719;](https://en.wikipedia.org/wiki/Multiplication),
[&#8970;...&#8971;](https://en.wikipedia.org/wiki/Continued_fraction),
or other sequence operator could be represented as a recursive function which accepts previous calculated value (<b>pv</b>),
index(i) and argument(x) as input parameters.

---
##### <u>Example 1 - Computing exp(z) using [Euler formula](https://en.wikipedia.org/wiki/Euler%27s_formula):</u>
![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/6a91595ef0946463456b2d0184bdcdb2ae9da7a2)<br>

recursive function would be:<br><br> f(z) = z^n / n! + pv0<br>
n=...,6,5,4,3,2,1,0<br><b>pv0</b> means what first value would be equal to 0

---
##### <u>Example 2 - Computing Pi using [Wallis product](https://en.wikipedia.org/wiki/Wallis_product):</u>
![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/df59bf8aa67b6dff8be6cffb4f59777cea828454)<br><br>

recursive function would be:<br><br> f() = pv1 * (2*n)^2 / ((2*n-1)*(2*n+1)) <br>
n=...,6,5,4,3,2,1,0<br>
last product could not be equal to 0, otherwise all final product will be equal 0 too, so here we using pv1 instead pv0, <b>pv1</b> means what first value would be equal to 1

---
##### <u>Example 3 - Computing square root using [Geron iteration formula](https://ru.wikipedia.org/wiki/%D0%98%D1%82%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D0%B0%D1%8F_%D1%84%D0%BE%D1%80%D0%BC%D1%83%D0%BB%D0%B0_%D0%93%D0%B5%D1%80%D0%BE%D0%BD%D0%B0):</u>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/9935d6f7061161b29325d712518fb58496f58bfb)<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/cd0d9bc3389f73d8501bfef1303b06246d81f771)<br>
 here x1 could not be 0 or 1 and should be equal to initial function argument x, so, we using PVx for this purposes

recursive function would be:<br><br> f(x) = (PVx + A/PVx) / 2<br>
n=...,6,5,4,3,2,1,0<br>
A - could be your special constant

---
### search principle
There are several steps to make this search:
#### 1. Iterate over all formula forms (this step is clusterable, you can calculate all further steps on different computers)</summary>
```
nextChild,err := formFormula.GetNextBracketsSequence(currentChild, maxChilds uint)
```
Function generates new sequence based on [Catalan method](https://en.wikipedia.org/wiki/Catalan_number) and skips all sequences where nextMaxChilds > maxChilds.

| currentChild |  nextChild  | nextMaxChilds    |
|:------------:|:-----------:|:----------------:|
| ()           | ()()        | 2                |
| ()()         | (())        | 1                |
| (())         | ()()()      | 3                |
| ()()()       | ()(())      | 2                |
| ()(())       | (()())      | 2                |
| (()())       | (())()      | 2                |
| (())()       | ((()))      | 1                |
| ((()))       | ()()()()    | 4                |
| ()()()()     | ()()(())    | 3                |

#### 2. Iterate over all combinations for X in leafs constants (maximum number of X occurrences limit are configurable) and Previous Values (PV's)
You cannot use formula without X argument, so it is required at least several leafs filled with it.
To recombine all possible values, used lexicographic method<br>(see Donald Knuth: The Art of computer Programming Volume 4 Fascicle 3A - Generation All Combinations (7.2.1.3 Generating all combinations on page 17))
```
formFormula.RecombineRequiredX(input *[]*uint, maxOccurrences uint, setXValue uint, ready func(remained *[]*uint))
```
Example for one, two and three X arguments:
```
=== RUN   Test_RecombineRequiredX_3of5
 
for one X:
 1 [1 0 0 0 0]
 2 [0 1 0 0 0]
 3 [0 0 1 0 0]
 4 [0 0 0 1 0]
 5 [0 0 0 0 1]

for two X:
 6 [1 1 0 0 0]
 7 [1 0 1 0 0]
 8 [1 0 0 1 0]
 9 [1 0 0 0 1]
10 [0 1 1 0 0]
11 [0 1 0 1 0]
12 [0 1 0 0 1]
13 [0 0 1 1 0]
14 [0 0 1 0 1]
15 [0 0 0 1 1]

for three X:
16 [1 1 1 0 0]
17 [1 1 0 1 0]
18 [1 1 0 0 1]
19 [1 0 1 1 0]
20 [1 0 1 0 1]
21 [1 0 0 1 1]
22 [0 1 1 1 0]
23 [0 1 1 0 1]
24 [0 1 0 1 1]
25 [0 0 1 1 1]
```

#### 3. Iterate over all possible constant values in free places (where X is not appearing)
For this steps used recombination through simple recursion.
```
RecombineValues(input *[]*uint, possibleValues *[]uint, ready func())
```
```
=== RUN   Test_Recombine_3x3
  1 [1 1 1]
  2 [2 1 1]
  3 [3 1 1]
  4 [1 2 1]
  5 [2 2 1]
  6 [3 2 1]
  7 [1 3 1]
  8 [2 3 1]
  9 [3 3 1]
 10 [1 1 2]
 11 [2 1 2]
 12 [3 1 2]
 13 [1 2 2]
 14 [2 2 2]
 15 [3 2 2]
 16 [1 3 2]
 17 [2 3 2]
 18 [3 3 2]
 19 [1 1 3]
 20 [2 1 3]
 21 [3 1 3]
 22 [1 2 3]
 23 [2 2 3]
 24 [3 2 3]
 25 [1 3 3]
 26 [2 3 3]
 27 [3 3 3]
```
#### 4. Iterate over all possible functions (where number of child arguments = 1)
Used the same principle as in step 3, but applied for functions.
#### 5. Iterate over all possible operators (where number of child arguments = 2)
Used the same principle as in step 3, but applied for operations.<br>
#### 6. Apply formula to itself several times
Finally, after all it included loops opens, you get some function-form<br>
F.e:<br>
```
=== RUN   Test_RecombineModularProgram_ForSingleX
    1 ((x+1)+(1!)) mod 15
    2 ((x*1)+(1!)) mod 15
    3 ((x^1)+(1!)) mod 15
    4 ((gcd(x,1))+(1!)) mod 15
    5 ((x+1)*(1!)) mod 15
    6 ((x*1)*(1!)) mod 15
    7 ((x^1)*(1!)) mod 15
    8 ((gcd(x,1))*(1!)) mod 15
    9 ((x+1)^(1!)) mod 15
   10 ((x*1)^(1!)) mod 15
   11 ((x^1)^(1!)) mod 15
   12 ((gcd(x,1))^(1!)) mod 15
   13 (gcd((x+1),(1!))) mod 15
   14 (gcd((x*1),(1!))) mod 15
   15 (gcd((x^1),(1!))) mod 15
   16 (gcd((gcd(x,1)),(1!))) mod 15
   17 ((x+1)+(inverse(1))) mod 15
   18 ((x*1)+(inverse(1))) mod 15
   19 ((x^1)+(inverse(1))) mod 15
   20 ((gcd(x,1))+(inverse(1))) mod 15
   21 ((x+1)*(inverse(1))) mod 15
   22 ((x*1)*(inverse(1))) mod 15
   23 ((x^1)*(inverse(1))) mod 15
   24 ((gcd(x,1))*(inverse(1))) mod 15
   25 ((x+1)^(inverse(1))) mod 15
   26 ((x*1)^(inverse(1))) mod 15
   27 ((x^1)^(inverse(1))) mod 15
   28 ((gcd(x,1))^(inverse(1))) mod 15
   29 (gcd((x+1),(inverse(1)))) mod 15
   30 (gcd((x*1),(inverse(1)))) mod 15
   31 (gcd((x^1),(inverse(1)))) mod 15
   ... and so on...
```
You can apply this function-form several times or just once, and based on this calculation estimate how close this function to your original pattern.<br>

### API
Next arithmetics are supported:
* <b>Modular</b> with uint64 type
    (all values calculated by some module)
* <b>Iterational</b> with float64 type
    (allows you to iterate over several iterations which increase precision to solution)
