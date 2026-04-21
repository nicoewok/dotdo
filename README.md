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
2. `git clone` this repository
3. `go install`

# Usage

dotdo uses a hidden Git repository in your home directory to sync tasks across devices automatically.

Initialize the local storage:

```bash
mkdir ~/.dotdo
cd ~/.dotdo
git init
# Link to your private GitHub repo
git remote add origin [https://github.com/yourusername/.dotdo.git](https://github.com/yourusername/.dotdo.git)
```

## Commands

- `dotdo` — Launch the dashboard.

- `dotdo add "task name"` — Add a new task.

- `dotdo list` — Show all tasks.