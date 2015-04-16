# gophers
Gophers is a tool you can use to scrape user information from a Github
organization or search results collection.

## Installation

No installation necessary, just download the executable and run it. Or clone
this repo and build it yourself. It's all open source.

## Usage

You can download the executable `gophers` and run it in your terminal.

```
$ ./gophers
```

Gophers will look for a file called `urls.txt` and pull each URL to parse in
sequential order.

Here's an example:

```
https://github.com/orgs/rails/people
https://github.com/orgs/php/people
https://github.com/thoughtbot/paperclip/stargazers
https://github.com/sergiotapia
https://github.com/yonasb
https://github.com/search?utf8=%E2%9C%93&q=location%3A%22San+Fransisco%22+location%3ACA+followers%3A%3E100&type=Users&ref=advsearch&l='
https://github.com/danieltapia
```

Gopher will know how to process each URL and output whatever users it finds into
`output.csv`.

## License

The MIT License (MIT)

Copyright (c) 2015 Sergio Tapia Gutierrez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
