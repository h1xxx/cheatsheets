evil-org mode
=============

# movement/select keys
gj     next item
gk     previous item
vae    select current items
vaR    select current subtree
(      previous table cell
)      next table cell
{      beginning of table
}      end of table
vae    select table cell
vaE    select table row
var    select whole table

# move item (header, column, row, etc) keys
M-j    move item down 
M-k    move item up 
M-k    move item right
M-k    move item l.eft

# header keys
M-ret  insert header
tab    unfold a header 
<<     promote a header
>>     demote a header
<aR    promote a header and subtree
>aR    demote a header and subtree

# table keys
M-ret   add a row

# number all the headings
M-x org-num-mode

# cycle todo state
C-t

# open agenda
C-c a



org-mode todo list
==================

# set a counter in a todo list
** a todo list [/]
---
** a todo list [%]

# add a checkbox in a todo list
** a todo list [/]
  - [ ] item 1

# toggle checkbox / recalculate total items and items done
C-c C-c



keys
====

# find which action is mapped under a key
C-h c

# show all key mappings
C-h b
---
C-h m

# movement
C-v		scroll forward one screen
M-v		scroll backward one screen
C-l		center the screen around the cursor
C-f		move cursor one character forward
C-b		move cursor one character backward
M-f		move cursor one word forward
M-b		move cursor one word backward
C-n		move cursor one line forward (next line)
C-p		move cursor one line backward (previous line)
C-a		move to beginning of line
C-e		move to end of line
M-a		move back to beginning of sentence
M-e		move forward to end of sentence
M-<		move to the beginning of whole text
M->		move to the end of the whole text

# commands
C-x C-c		quit emacs
C-g		stop executing command, clear command buffer
C-u 3 C-f	repeat command 3 times
M-3 C-f		repeat command 3 times

# editing
C-u 3 f		enter an 'f' character 3 times
<DEL>		delete the character just before the cursor
C-d		delete e character under the cursor
M-<DEL>		kill the word immediately before the cursor
M-d		kill the next word after the cursor
C-k		kill from the cursor position to end of line
M-k		kill to the end of the current sentence
C-<SPC>		set a mark
C-w		kill the text between the cursor and the mark
C-y		yank (paste) killed text
M-y		after yank - iterate through previously killed text
C-/		undo
C-_		undo
C-g C-_		redo

# files
C-x C-f		find a file
C-x C-s		save the file in current buffer
C-x s		save files in other buffers
C-x C-b		list buffers
C-x b		select other buffer

# windows, frames
C-x 0		kill current window
C-x 2		split the screen into two windows
C-M-v		scroll down the other window.
C-x o		move the cursor to the bottom window
C-x 5 2		create a new frame
C-x 5 0		remove a selected frame

# help
C-h ?		help
C-h r		read emacs manual
C-h c C-f	show short help on C-f command
C-h k C-f	open a new window with help on C-f command
C-h f		describe a function
C-h a		command apropos - find help by keyword
C-h i		read included manuals



tables
======

# autoformat table
C-c C-c

# recalculate all table formulas (need to be on #+TBLFM line)
C-c C-c

# syntax for a cell formula (tab or C-c C-c to calculate)
:= $1 * $2^2

sum of first and second column
$1+$2

sum of  column 1 and 2 in column 3 row 1 
@1$3=$1+$2

format result to two decimals
$1+$2;%.2f

Math functions can be used
exp($2)+exp($1)

Reformat current cell to 1 decimal
$0;%.1f

compute column range mean, using vector function
vmean($2..$7)

compute column range mean, using vector function, treat empty columns as 0
vmean($2..$7);EN

# field and name reference examples
@2$3            2nd row, 3rd column (same as ‘C2’)
$5              column 5 in the current row (same as ‘E&’)
@2              current column, row 2
@-1$-3          field one row up, three columns to the left
@-I$2           field just under hline above current row, column 2
@>$5            field in the last row, in column 5
$1..$3          first three fields in the current row
$P..$Q          range, using column names
$<<<..$>>       start in third column, continue to the last but one
@2$1..@4$3      nine fields between these two fields (same as ‘A2..C4’)
@-1$-2..@-1     3 fields in the row above, starting from 2 columns on the left
@I..II          between first and second hline, short for ‘@I..@II’

# example table with calculations
| 3.14 | 10 | := $1 * $2^2 |
---
| *Expected Income*          |                |
| Gross Salary Wanted        | 100000         |
| Health Insurance           | :=0.07 * @2$2  |
| Total Revenue              | :=vsum(@2..@3) |
| *Hourly rate*              |                |
| Unpaid vacation days       | 20             |
| Public holidays + weekends | 120            |
| Work days per year         | :=365-@6-@7    |
| Work hours per year        | :=@8 * 8       |



config
======

# list all colors used
M-x list-faces-display

# customize color of the face under the cursor
M-x customize-face

# list available colors
M-x list-colors-display

# show configuration option
C-h v make-backup-files

