# ASCII Art Generator

This project is an ASCII Art Generator that converts input text into ASCII art using different-fonts/bannerfiles and optional color coding.  

## Prerequisites

This program was created entirely with Golang programming language so ensure your workstation supports golang version 1.18 or later before cloning the repository.

You can install golang from there website depending on the type of operating system your using at https://go.dev/doc/install and follow instructions on how to set-up

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone  https://github.com/jesee-kuya/ascii-art-fs.git
    ```

2. Navigate to the project directory:
    ```sh
    cd ascii-art-fs
    ```

3. Build the project:
    ```sh
    go build -o ascii-art-fs
    ```

## Usage

Run the binary with the required arguments:

```sh
./ascii-art-fs "Your text here" <Banner_file>
```
this other options also work because the program also supports optional color codes and letters to color. The banner file portion in optional.
```sh
./ascii-art-fs <color-flag> <letters to color> "your text here" <Banner_file>
```
```sh
./ascii-art-fs <color-flag> "your text here" <Banner_file>
```
and the normal base functionality
```sh
./ascii-art-fs "your text here"
```


### Parameters

- **-Banner**: The font file to use (default: `standard`).
- **-color**: The color to use for the text.
- **-letters**: The specific letters to color.
- **"Your text here"**: The text you want to convert to ASCII art.


# **Examples**
## **Example 1: Basic Usage**
```go
./ascii-art-fs "Hello, World!" standard.txt
 _    _          _   _                    __          __                 _       _   _  
| |  | |        | | | |                   \ \        / /                | |     | | | | 
| |__| |   ___  | | | |   ___              \ \  /\  / /    ___    _ __  | |   __| | | | 
|  __  |  / _ \ | | | |  / _ \              \ \/  \/ /    / _ \  | '__| | |  / _` | | | 
| |  | | |  __/ | | | | | (_) |  _           \  /\  /    | (_) | | |    | | | (_| | |_| 
|_|  |_|  \___| |_| |_|  \___/  ( )           \/  \/      \___/  |_|    |_|  \__,_| (_) 
                                |/                                                      
                                                                                        
```
## **Example 2: Using Color**
```go
./ascii-art-fs -color="rgb(255, 0, 0)" "Hello, World!" shadow                                                                                                   
```

## **Example 3: Coloring Specific Letters**
```go
./ascii-art-fs -color="#2E8B57" "HW" "Hello, World!" thinkertoy
```
One can enter rgb codes or hex codes of any color they choose at the color flag

## **Example 4: Coloring Specific Letters in a filename**
### **make sure you escape the name of the bannerfile**
```go
./ascii-art-fs -color="#FF0000" tk \\thinkertoy
```
 
# **Contributing**
We welcome contributions to improve the ASCII Art Generator. If you find a bug or have a feature request, please open an issue. If you would like to contribute code, please fork the repository and submit a pull request.

### **Steps to Contribute**

1. **Fork the repository.**
```sh
   https://github.com/jesee-kuya/ascii-art-fs.git
```
2. **Create a new branch**:
```sh
   git checkout -b feature-branch
```
3. **Make your changes.**
4. **Commit your changes**
```sh
git commit -am 'Add new feature
```
5. **Push to the branch**
```sh
git push origin feature-branch
```
5. **Open a pull request**

## License

This project is licensed under the MIT License. See the [LICENSE]( https://github.com/jesee-kuya/ascii-art-color.git/main/LICENSE)  file for details.

Brought to you by [Jesee Kuya](https://github.com/jesee-kuya) , [James Muchiri](https://github.com/j1mmy7z7) and [Fena Onditi][def]


[def]: https://github.com/konditi1

