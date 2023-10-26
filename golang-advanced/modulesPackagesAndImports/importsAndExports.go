// Export:
// - The functions you want to export from the package must be capitalized.
// - In Go, there is no need to explicitly mention the exporting package,
//   instead the functions with Capitalized name is automatically exported.


// Import:
// - Use 'import' keyword for import packages
// - eg, import "github.com/patha2702/cryptit/encrypt"


// Importing packages from other modules locally:
// - Replace flag:
//   - Command: go mod edit -replace <global path of module>=<local path of module>
//   - for eg, go mod edit -repace github.com/patha2702/cryptit=../cryptit
