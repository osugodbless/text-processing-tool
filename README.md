# Go Text Processor

This is a command-line Go application designed to parse and format text files. It reads an input file, perform a series of string manipulations, grammatical corrections, and format transformations, and writes the resulting formatted text to a specified output file.

---

## 1. Project Goals

- Demonstrate idiomatic string, rune, and number manipulation in Go.
- Utilize the Go File System (`os`) API and Regular Expressions (`regexp`) for robust text tokenization.
- Implement a modular architecture that cleanly separates file I/O from transformation logic.
- Provide highly testable business logic using in-place pointer mutations.

---

## 2. Features and Transformations

The processor automatically applies the following rules to the input text:

- **Hexadecimal & Binary to Decimal:** Converts the preceding number from hexadecimal to decimal when it encounters the marker `(hex)` and binary to decimal when the marker is `(bin)`.
- **Casing Transformations:** - `(up)`: Converts the preceding word to UPPERCASE.
  - `(low)`: Converts the preceding word to lowercase.
  - `(cap)`: Capitalizes the preceding word.
  - **Numbered Casing:** Supports numeric arguments (e.g., `(low, 3)`) to apply the transformation to a specified number of preceding words.
- **Punctuation Formatting:** - Fixes spacing around standard punctuation (`.`, `,`, `!`, `?`, `:`, `;`) so they attach to the previous word and are followed by a space.
- **Quote Formatting:** Automatically formats single quotes (`'`) to cleanly wrap the enclosed text without internal spaces.
- **Modify Article:** Intelligently changes the article "a" or "A" to "an" or "An" when preceding a word starting with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`.

---

## 3. Architecture & Flow

The application is structured to separate file tokenization from the core text manipulation logic.

```text
.
├── go.mod
├── main.go                 # Entry point, validates CLI arguments
├── processor/              # Core business logic and text manipulation
│   ├── file_io.go          # RegEx tokenization, File reading/writing
│   ├── lettercase.go       # String case conversion logic
│   ├── match.go            # Punctuation formatting logic
│   ├── number.go           # Hex/Bin conversion logic
│   ├── processor.go        # Orchestrator applying rules via switch statements
│   └── vowels.go           # Article adjustment logic
└── testdata/               # Target directory for input and output files
    ├── sample.txt          
    └── result.txt
```

### Implementation Highlights:
* **RegEx Tokenization:** `file_io.go` uses Regular Expressions to intelligently split the raw text into a slice of strings, isolating punctuation, words, and bracketed commands for precise iteration.
* **In-Place Mutation:** Helper functions (like `Lowercase(&s)`) accept string pointers to mutate values directly in memory, optimizing performance.
* **Clean Tag Removal:** `processor.go` tracks applied command tags using a `stringToDel` map, efficiently filtering them out before assembling the final text.

## 4. Usage

### Prerequisites
* Go 1.22.2 or higher installed on your system.
* Standard Go packages only (no third-party dependencies required).

### Execution Details
**Important:** The application is hardcoded to read from and write to the `testdata/` directory. You must place your input text file inside the `testdata/` folder before running the program.

Do not pass the directory path in the arguments, only the filenames:

```bash
# Valid execution:
go run . sample.txt result.txt
```

## 5. Testing Strategy
The project features a robust, isolated testing suite using idiomatic Go testing patterns.

* **Table-Driven Tests:** Used extensively (e.g., in `lettercase_test.go`) to cleanly validate multiple edge cases (mixed case, numbers, empty strings) in a single, readable block.
* **Subtests:** Uses `t.Run` to isolate specific test cases within the test tables, making debugging easier when a specific edge case fails.
* **Mutation Testing:** Ensures that functions correctly mutate the data via pointers rather than just returning new values.

To run the entire test suite:

```bash
go test ./processor -v
```

## 6. Trade-offs & Limitations

* Designed specifically for plain text files (.txt).
* Assumes the input file is formatted with correctly spaced command tags (e.g., `(cap, 2)`).