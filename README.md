# ✨ SHINE

`shine` is a ridiculous command line utility for displaying text with a gradient background.

## Install

If you're **already using go**, then just run `go install github.com/rewdy/shine/cmd/shine@latest`

## Usage

[![asciicast](https://asciinema.org/a/JsmY7s4qNTRM2jCGAeefREj7e.png)](https://asciinema.org/a/JsmY7s4qNTRM2jCGAeefREj7e)

```bash
Usage: shine [options] <text>

  -e, --end-color string     The end color of the gradient (default "blue")
  -p, --pad                  Adds extra padding around the test to let it breathe (default true)
  -r, --random               Selects random start and end colors
  -s, --start-color string   The start color of the gradient (default "red")

Color swatches for you to choose from:

 red         yellow      gray        salmon      black

 white       persimmon   green       pink        teal

 brown       orange      blue        indigo      violet
```

### Want to override default colors? Env vars FTW

If you don't like pink-to-blue gradients, try this:

```bash
export SHINE_START_COLOR=teal
export SHINE_END_COLOR=violet

# live your life....

shine "shine bright like a ♦"

# voila!
```

### That's... so... 🤭 rAnDOm!

If you want `shine` to always use random colors, try this:

```bash
export SHINE_RANDOM=yasss  # set to anything. enabled by being set.

# and again

shine "wow, hi, yes"

# voila!
```

If goes without saying, you could add the above env vars to your `.zshrc` (or comparable) and enjoy persistent defaults!

---

😂 Made with joy by [rewdy](https://rewdy.lol) xoxo
