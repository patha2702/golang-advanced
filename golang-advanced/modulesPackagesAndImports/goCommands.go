// Go commands:


// go mod init <modName> => Initializes and writes a new go.mod file in the current directory.

// go mod tidy => Ensures that the go.mod file matches the source code in the module.
//                Adds any missing module requirements necessary to build the current module packages, dependencies.
//                Removes the packages that are not used/ not relevant

// go run <filename> => Compiles and runs the program.
//                  Internally,
//                  - compiles and builds an executable file in temporary location.
//                  - launches the executable file
//                  - cleans the space, when the app exits
//                  Used for testing small programs during initial development stage of application.

// go build => Compiles the packages named by the import paths, along with their dependencies into an executable.
//             Executable file gets created in the current source directory.

// go install => Compiles and move the executable to $GOPATH/bin
//               run this executable from any path on the terminal

// go get => updates go.mod to require those versions, and
//           downloads the source code into the module cache.
//           - Add dependency for a package
//           - Upgrade/ Downgrade dependency to a specific version.

