# dotdo (Linux)
> Minimalist, dot-matrix, focused.

`dotdo` is a developer-centric CLI/TUI tool. It manages your tasks through a monochromatic interface, featuring the signature pixel Bunny mascot.

---

## The Mascot
```text
    ⠏⢣ ⠏⢣
  ⢠⡶⠧⠧⠶⠧⠧⠶⢶⡄
  ⡜         ⢣
 ⢸   ⠛   ⠛  ⢣
  ⢣      Y  ⢸
  ⢸      "  ⡜
  ⡜        ⢸
⠺⡜         ⡜
  ⠙⠒⠤⣀⣀⣇⣸⣇⣸
```

# Install

1. Ensure you have [Go](https://go.dev/doc/install) installed
2. ```git clone``` this repository
3. ```go install```
4. Be sure your `go/bin` folder is added to your `PATH`:
    
    a. Get your go folder using
    ```bash
    go env GOPATH
    ```
    b. It should return something like `/home/yourname/go`. If you exedcuted 3. then in this path in `/bin` you should find an executable called `dotdo`
    c. Edit your bash config:
    ```bash
    nano ~/.bashrc
    ```
    d. Go to the last line and paste this line:
    ```bash
    export PATH=$PATH:$GOPATH/bin
    ```
    where `$GOPATH` is the path you got from a. and b. so for example:
    ```bash
    export PATH=$PATH:/home/USERNAME/go/bin
    ```
    e. Save and exit: Ctrl+O -> Enter -> Ctrl+X
    f. Reload your terminal!


# Usage

Initialize the local storage:
```bash
mkdir ~/.dotdo
dotdo init
```

If you want to sync your data, **dotdo** uses a hidden Git repository in your home directory to sync tasks across devices automatically.

```bash
cd ~/.dotdo
git init
# Link to a private GitHub repo
git remote add origin [https://github.com/yourusername/.dotdo.git](https://github.com/yourusername/.dotdo.git)
```

## Commands

- `dotdo` — Show all tasks not on done.

- `dotdo help` — Show all commands.

- `dotdo init` — Initilize the `tasks.json` file to start adding tasks.

- `dotdo add [task name] -d [date]` — Add a new task. The name does not need to be in "". Example:
```bash
dotdo add Feed bunny
```
You can optionally add a due date using `-d` and adding a date in the **YYYY-MM-DD** format. Example:
```bash
dotdo add Feed bunny -d 2026-01-01
```


- `dotdo doing [task name]` — Mark task as "currently doing". This command has auto-complete for the task name.

- `dotdo done [task name]` — Mark task as "done". This command has auto-complete for the task name.

- `dotdo remove` — Deletes all "done" tasks to clean up `tasks.json`.

- `dotdo sync` — Pulls changes for `tasks.json` & pushes new changes.

- `dotdo list` — Show all tasks. Even "done" ones.