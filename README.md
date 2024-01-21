# Form-Formula-Go
![Version: version](https://img.shields.io/badge/version-v0.0.1--alpha-success.svg)
![Tests: tests 100% coverage](https://img.shields.io/badge/tests-67_of_67=100%-success.svg)
[![License: GPL3.0](https://img.shields.io/badge/License-GPL3.0-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.html)

Library to iterate over all possible math functions to find acceptable one (GoLang implementation).<br>

### idea
Every mathematical function f(x) could be represented as [Abstract Syntax Tree (AST)](https://en.wikipedia.org/wiki/Abstract_syntax_tree) where nodes of this tree could be an<br>
 - <b>operators</b> (has two input arguments): <b>*, +, -, /, mod, pow(x,p)</b>
 <br>or<br>
 - <b>functions</b> (has one input argument): <b>sqrt(x), sin(x), x!, log(x), exp(x),round(x)</b>
Complex formulas which contains [&#8721;](https://en.wikipedia.org/wiki/Summation),[&#8719;](https://en.wikipedia.org/wiki/Multiplication),[&#8970;...&#8971;](https://en.wikipedia.org/wiki/Continued_fraction), or other sequence operator could be represented as recursive function which accepts previous calculated value (<b>pv</b>), index(i) and argument(x) as input parameters.

Leafs of this tree are values of:
* constants or
* X argument or
* iteration number or
* previous calculated value

##### Example 1 - Computing exp(z) using [Euler formula](https://en.wikipedia.org/wiki/Euler%27s_formula):<br>
![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/6a91595ef0946463456b2d0184bdcdb2ae9da7a2)

recursive function would be:<br><br> ![alt tag](https://chart.googleapis.com/chart?cht=tx&chl=z^n/n!%2bpv0) (n=...,6,5,4,3,2,1,0)
(<b>pv0</b> means what first value would be equal to 0)
<br>
##### Example 2 - Computing Pi using [Wallis product](https://en.wikipedia.org/wiki/Wallis_product):<br>
![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/df59bf8aa67b6dff8be6cffb4f59777cea828454)<br>
last product could not be equal to 0, otherwise all final product will be equal to 0 too
<br><br><br>
* <b>pv = x<br></b>
##### Example 3 - Computing square root using [Geron iteration formula](https://ru.wikipedia.org/wiki/%D0%98%D1%82%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D0%BE%D0%BD%D0%BD%D0%B0%D1%8F_%D1%84%D0%BE%D1%80%D0%BC%D1%83%D0%BB%D0%B0_%D0%93%D0%B5%D1%80%D0%BE%D0%BD%D0%B0):<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/9935d6f7061161b29325d712518fb58496f58bfb)<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/cd0d9bc3389f73d8501bfef1303b06246d81f771)<br>
 ![alt tag](https://wikimedia.org/api/rest_v1/media/math/render/svg/a8788bf85d532fa88d1fb25eff6ae382a601c308) could not be 0 or 1 and should be equal to initial function argument x
<br><br><br>

### search principle
There are several steps to make this search:
#### 1. Iterate over all formula forms (this step is clusterable, you can calculate all further steps on different computers)</summary>
 

#### 2. Iterate over all X combinations in constants in leafs (maximum number of X occurrences limit are configurable) and Previous Values (PV's)
#### 3. Iterate over all possible constant values in free places (where X is not appearing)
#### 4. Iterate over all possible functions (where number of child arguments = 1)
#### 5. Iterate over all possible operators (where number of child arguments = 2)
#### 6. Apply formula to itself several times

Finally, after all it included loops opens, you get some function form and could calculate your sample points.<br>
Based on this calculation you can estimate how close this function to your original pattern.<br>

### API
Next arithmetics are supported:
* <b>Modular</b> with uint64 type
    (all values calculated by some module)
* <b>Iterational</b> with float64 type
    (allows you to iterate over several iterations which increase precision to solution)
