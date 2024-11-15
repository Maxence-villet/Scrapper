# Scrap Package

This package defines functionalities for scraping content from a website.

### Imports

`fmt`: Used for formatted printing to the console.
`log`: Used for logging messages.
`os`: Used for file operations.
`strings`: Used for string manipulation.
`github.com/gocolly/colly`: A popular library for web scraping in Golang.

## ScrapHandler Interface

Defines a single method `Scrap()` that every scraper implementation must fulfill. This method represents the scraping logic.

## Scrap Struct

A simple struct used for scraping functionalities. It currently doesn't have any fields.

## Scrap.Scrap() Function

This function performs the web scraping logic:

### 1. Create a Colly Collector: 
- A new collector object is created using `colly.NewCollector()`. This object handles the scraping process.
### 2. Define Callback for Specific HTML Element:
- The collector uses `c.OnHTML(".group", func(e *colly.HTMLElement) { ... })` to define a callback function that executes whenever an element with the class ".group" is encountered.
- Inside the callback, `e.ForEach(".group.relative a[href]", func(_ int, link *colly.HTMLElement) { ... })` iterates over all anchor tags (`a`) with the attributes `class="group.relative"` and `href` within the matched ".group" element.
### 3. Process Links:
- The `href` attribute of each link is extracted using `link.Attr("href")`.
- A check is performed whether the link starts with "/udemy/".
- If the link starts with "/udemy/":
    - Two lists are defined:
    - `list_search`: An array containing desired keywords ("project", "build", "apps", "scrum").
        - `blacklist_search`: An array containing keywords to exclude ("project-management").
    - The code iterates through both lists:
        - For each element in `list_search`:
            - It checks if the link's uppercase version (using `strings.ToUpper(href)`) contains any keyword from the `blacklist_search` list (also uppercased).
                - If a blacklisted keyword is found, no further action is taken (skips the link).
            - If no blacklist match occurs and the link contains a term from `list_search`:
                - A data string is created by prepending "comidoc.net" to the `href` with a newline character.
                - The file "data.txt" is opened in append mode with write permissions using `os.OpenFile`.
                - If an error occurs during file opening, it's logged using fmt.Println.
            - The data string is written to the file using `f.Write`. Any errors encountered during writing are logged with `fmt.Println`.
            - The file is closed using `defer f.Close()`.
### 4. Logging:
- `c.OnRequest(func(r *colly.Request) { ... })` defines a callback that logs the visited URL for each request.
- `c.OnError(func(r *colly.Response, err error) { ... })` defines a callback that logs errors encountered during requests.
### 5. Start Scraping:
- Finally, the `c.Visit("https://comidoc.net/daily")` call initiates the scraping process starting from the provided URL ("https://comidoc.net/daily").

## NewScrap Function

This function creates and returns a new `Scrap` instance.

## Notes:

- This scraper targets links with specific keywords within elements having the class ".group" on the website "comidoc.net/daily".
- It writes scraped links (prepended with "comidoc.net") to a file named "data.txt".
- Error handling is included for file operations.
Consider customizing the scraper logic based on your specific scraping needs.

## Usage Example 

    package main

    import (
        "fmt"
        "log"

        "scrap.com/scrap"
    )

    func main() {
        s := scrap.NewScrap()

        if err := s.Scrap(); err != nil {
            log.Fatal(err)
        }

        fmt.Println("Scraping completed!")
    }