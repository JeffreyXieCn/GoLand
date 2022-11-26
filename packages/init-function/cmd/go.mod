module github.com/gopherguides/cmd

go 1.19

//If a version is present on the left side of the arrow (=>), only that specific version of the module is replaced;
//other versions will be accessed normally. If the left version is omitted, all versions of the module are replaced.
//Regardless of whether a replacement is specified with a local path or module path, if the replacement module has
//a go.mod file, its module directive must match the module path it replaces.
replace github.com/gopherguides/creature => ../creature

//https://go.dev/ref/mod#:~:text=A%20replace%20directive%20replaces%20the,a%20platform%2Dspecific%20file%20path.
//Note that a replace directive alone does not add a module to the module graph.
//A require directive that refers to a replaced module version is also needed, either in the main
//module’s go.mod file or a dependency’s go.mod file. A replace directive has no effect if the module
//version on the left side is not required.
require github.com/gopherguides/creature v0.0.0-00010101000000-000000000000
