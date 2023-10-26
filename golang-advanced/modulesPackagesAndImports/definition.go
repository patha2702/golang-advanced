// Repository:
// - It is a place where the whole source code for the application is stored.
// - Eg, repository on platforms like github, local machine etc.


// Packages:
// - It is a collection of Go source files in the same directory that are compiled together.
// - Go code is grouped into packages.


// Modules:
// - Packages are grouped into modules.
// - Module specifies the dependencies needed to run our code, including the Go version
//   and the set of other modules it requires in the go.mod file.


// Creating a module:
// - go mod init <module_path> where, module_path => globally unique identifier.
// - eg. go mod init example.com/myProject


// go.mod file:
// - go mod init command creates a go.mod file to track the dependencies of the project.
// - A collection of Go source code becomes a module when there's a valid go.mod file in the 
//   root directory.
// - It contains module declaration, minimum compatible version for Go, and the dependencies 
//   for the imported third-party packages.


// downloading the dependencies for building the module, packages:
// go mod tidy
