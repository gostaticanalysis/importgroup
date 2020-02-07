package a

import _ "fmt" // OK
import _ "io"  // want "the import statement is not grouped"
