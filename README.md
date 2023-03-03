# Birthday Challenge written in Go:

Given a JSON file with a list of people and their dates of birth, write a program to print out the people whose birthday is today.

If a person was born on Feb 29th, then their birthday during a non-leap year should be considered to be Feb 28th.

Input sample file:

```
[
    ["Doe", "John", "1982/10/08"],
    ["Wayne", "Bruce", "1965/01/30"],
    ["Gaga", "Lady", "1986/03/28"],
    ["Curry", "Mark", "1988/02/29"]
]
```

You can use whichever programming language you like. The assignment should take 60 to 90 min. If it’s taking longer, consider whether you’re complicating things.

If you make any assumptions, trade-offs or de-prioritise features for timeliness, please document these decisions.

Your submission must have:

* Instructions to run the code

* Tests to check if your code is correct, robust and complete

* Edge cases handled and tested

* Iterative development, backed by commit history

* Modular, cohesive code with sensible separation of concerns

Bonus points for following Test-Driven Development.

Please do not overcomplicate the code. You don’t need a web framework, database or message queues for this submission. Keep it simple!

<br>

# Running Instructions & Comments Below:

*NB: I used Cobra CLI to run this program as a CLI application. Please run on a Unix environment. I haven't tested this on a Windows environment.*

How to run project:

* Git clone the repo to your local machine.

* Cd into the root directory & type **```"go mod download"```** to install dependencies.

* Type **```"go run ."```** to see the CLI's main menu.

* Type **```"go run . read"```** to read the input json file from its default location "./static/files/input.json"

* Use the **```"go run . read --csv"```** csv flag to read input from the csv file under static > files > input.csv

* I've written unit tests for the validation file - please type **```"go test -v ./internal/validator"```** to run them. I've tried to cover a couple of edge cases here.

How to build the project:

* Run the following: **```"go build -o /Your/desired/location && cp -r ./static /Your/desired/location"```**

* Then go to **```/Your/desired/location```** via the terminal & run it like: **```"./gobirth" instead of "go run ."```**

* *NB: The files under static will change quite often.*

How is the code structured:

```
├── LICENSE
├── README.md
├── cmd
│   ├── read.go
│   └── root.go
├── go.mod
├── go.sum
├── internal
│   ├── input
│   │   └── input.go
│   ├── output
│   │   └── output.go
│   └── validator
│       ├── validator.go
│       └── validator_test.go
├── main.go
└── static
    └── files
        ├── input.csv
        └── input.json
```
*NB: Please star this repo if you liked this short demo of how to write a CLI app in Go*