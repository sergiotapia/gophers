# gophers
Gophers is a tool you can use to scrape user information from a Github
organization or search results collection.

## Installation

Instructions go here.

## Usage

You can download the executable `gophers` and run it in your terminal.

```
// Scrape an entire Github organization
$ gophers -github_url=https://github.com/orgs/rails/people

// Scrape a single Github profile
$ gophers -github_url=https://github.com/sergiotapia

// Scrape the results of a Github search
$ gophers -github_url=https://github.com/search?utf8=%E2%9C%93&q=ruby&type=Users&ref=searchresults
```

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
