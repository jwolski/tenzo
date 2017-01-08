# Instructions for the Vim Cook

## Create a command
```
command <CmdName> <ThingToDo>
```

For example, this custom command trims all trailing whitespaces:

```
command TrimAll %s/\s\+$//e
```

## :buffers or :ls

## Line break while in normal mode
Pressing `r` then `Enter` replaces character under the cursor
with line break.

# Folds
zR: open all folds
zO: open single fold
zM: close all fold
zC: close single fold

# :h
Help!

## :highlight
See currently set colors.

## :map
See keyboard mappings/shortcuts.

## :reg
See registers and values.

## Run an External Command
Example:

```
:!ls
```
