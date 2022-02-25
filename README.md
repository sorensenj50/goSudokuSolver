<p>
  <img src="https://img.shields.io/badge/Go-%20go1.18beta1-blue">
    <img src="https://img.shields.io/badge/Gin-1.7.7-red">
</p>

# Sudoku Solver
Learning Go by building a Sudoku Solver. 

### Features

* Solver
* Random Puzzle Generator
* API that delivers puzzle and solution

## Contents

This project has three substantive source files that each define a `struct`.

* `Puzzle` contains the cell-values in a 9x9 2d array. Has methods for displaying, traversing, and solving the given `Puzzle`.
* `Generator` delivers random `int`s for testing as valid cell values. Does so in an efficient and constrained way.
* `Tracker` is a wrapper over a `map` that ensures that the solver doesn't mutate cell values that are given in the initial puzzle. 


Two other files are used for delivering the content via the `REST API`.
* `main` contains the `Gin` router and function handlers for `/puzzle` and `/info`
* `apiUtils` contains several helper functions to unclutter the `main` file.



