# Form-Formula-Go
![Version: version](https://img.shields.io/badge/version-v0.0.1--alpha-success.svg)
![Tests: tests 100% coverage](https://img.shields.io/badge/tests-67_of_67=100%-success.svg)
[![License: GPL3.0](https://img.shields.io/badge/License-GPL3.0-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.html)

Library to iterate over all possible math functions to find acceptable one (GoLang implementation).<br>

Next arithmetics are supported:
* <b>Modular</b> with uint64 type
    (all values calculated by some module)
* <b>Iterational</b> with float64 type
    (allows you to iterate over several iterations which increase precision to solution)

### Search principle
There are several steps to make this search:
<details>
    <summary>1. Iterate over all formula forms (this step is clusterable, you can calculate all further steps on different computers)</summary>
    TBD<br>
</details>
<details>
    <summary>2. Iterate over all X combinations in constants in leafs (maximum number of X occurrences limit are configurable) and Previous Values (PV's)</summary>
    TBD<br>
</details>
<details>
    <summary>3. Iterate over all possible constant values in free places (where X is not appearing)</summary>
    TBD<br>
</details>
<details>
    <summary>4. Iterate over all possible functions (where number of child arguments = 1)</summary>
    TBD<br>
</details>
<details>
    <summary>5. Iterate over all possible operators (where number of child arguments = 2)</summary>
    TBD<br>
</details>
<details>
    <summary>6. Apply formula to itself several times</summary>
    TBD<br>
</details>
<br>
Finally, after all it included loops opens, you get some function form and could calculate your sample points.<br>
Based on this calculation you can estimate how close this function to your original pattern.<br>
