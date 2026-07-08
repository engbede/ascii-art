# ASCII Art

A Go (Golang) application that converts plain text into visually appealing ASCII art using predefined banner templates. This project was completed as part of the **ASCII Art** curriculum and focuses on file processing, string manipulation, and command-line application development.

---

## About

**ASCII Art** is a command-line application that transforms text into stylized ASCII art by mapping each printable character to its graphical representation stored in banner files.

The project introduces concepts such as file parsing, character mapping, error handling, and efficient text rendering while encouraging clean, modular Go development.

---

## Objectives

* Learn file reading and parsing in Go.
* Strengthen string manipulation skills.
* Build command-line applications.
* Understand character encoding and mapping.
* Practice modular software design.
* Improve debugging and problem-solving abilities.

---

## Features

* Convert text into ASCII art.
* Support multiple banner styles.
* Read banner templates from text files.
* Preserve line breaks in the input.
* Validate user input and banner files.
* Handle printable ASCII characters (32–126).
* Produce consistent output formatting.
* Display meaningful error messages for invalid input.

---

## Project Structure

```text
ascii-art/
├── main.go
├── go.mod
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│   ├── GenerateArt.go
│   ├── RendererLine.go
│   ├── Banner_test.go
│   └── RenderLine.go
├── README.md
└── ...
```

> The directory structure may evolve as additional features and banner styles are added.

---

## How It Works

The application follows these steps:

1. Read the input text from the command line.
2. Load the selected banner template.
3. Parse the banner into character patterns.
4. Match each input character with its ASCII art representation.
5. Render each line of the banner.
6. Display the completed ASCII art in the terminal.

---

## Technologies

* Go (Golang)
* Git
* GitHub

---

## Concepts Practiced

* File handling
* String manipulation
* Character mapping
* ASCII encoding
* Error handling
* Algorithms
* Data structures
* Modular programming
* Command-line interface (CLI) development

---

## Running the Project

Clone the repository:

```bash
git clone https://github.com/nemmanuel/ascii-art.git
```

Navigate into the project directory:

```bash
cd ascii-art
```

Run the application:

```bash
go run . "Hello"
```

Use a specific banner style:

```bash
go run . "Hello" standard
```

Replace `standard` with another supported banner such as `shadow` or `thinkertoy`, depending on your implementation.

---

## Learning Outcomes

Working on this project helped me:

* Understand how ASCII art generation works.
* Build reusable Go packages.
* Parse structured text files efficiently.
* Improve debugging and testing skills.
* Design maintainable command-line applications.
* Strengthen problem-solving using Go.

---

## Contributions

This repository is intended for educational purposes and to document my learning journey. Suggestions, improvements, and constructive feedback are welcome.

---

## License

This project is shared for educational purposes. Please respect your institution's academic integrity policies when using or referencing the code.

---

## Author

**Emmanuel Ngbede**

GitHub: https://github.com/engbede

---

*"Every line of ASCII art begins with a single character and a well-designed algorithm."*
