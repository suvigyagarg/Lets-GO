

# CLI To-Do App


Whenever I’m working on something, I love to track my progress with a simple to-do list. That’s why I built this **To-Do App in the Terminal** using **Go**! It was my first CLI project and a great way to get my hands dirty with Go while building something genuinely useful.

---

##  Getting Started

Clone the repo:

```bash
git clone https://github.com/suvigyagarg/todo-cli.git
cd ./todo-cli
```

Run the app (for testing):

```bash
go run ./main.go <command>
```

If you wish to **build** it:

```bash
go build todo
./todo <command>
```

### Optional: Use It From Anywhere in your system

Want to run `todo` from *anywhere* in your terminal? Follow these steps:

```bash
go install
```

Then add your Go `bin` path to your system’s PATH. Example (for Windows Bash shell):

```bash
export PATH=$PATH:/c/Users/<YourName>/go/bin
```

Now, you can run:

```bash
todo -help
```

from *anywhere*! 😎



##  Commands

Here's what you can do with the app:

```bash
todo -help
```
See all commands and usage.


```bash
todo -add "<item>"
```
Add a new to-do item.

```bash
todo -list
```
 List all to-dos.

```bash
todo -del <index>
```
 Delete a to-do by its index.

```bash
todo -toggle <index>
```
Mark a to-do as done or undone.

```bash
todo -edit "<index:new_title>"
```
Edit an existing to-do's title.


## Where it Stores?

Your todos are safely stored as a JSON file here:

- **Windows:**  
  `C:\Users\<YourName>\.todoapp\todo.json`
- **Linux/macOS:**  
  `~/.todoapp/todo.json`



## What more I want to add

Here’s what I’d love to add — and you’re welcome to contribute!

- **Workspace Support**: Different to-do lists for different projects.
- **Due Dates & Priorities**: For all you organized souls out there.
- **Sync with Cloud or GitHub**: Because we love backups.




# todo-cli
