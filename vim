strings
=======

# delete a line containing a string (w/o /d it will print patterns to delete)
:g/profile/d

# delete only lines starting with a whitespace
:g/^\s

# delete lines starting with a string
:g/^string/d

# delete only empty lines
:g/^\s*$

# delete all lines except those that contain "error" or "warn" or "fail"
:v/error\|warn\|fail/d
---
:g!/error\|warn\|fail/d

# remove 5 first chars from each line
:%s/^.\{0,5\}//
---
:%s/^.\{5\}//g
---
:%s/^.....//

# replace double spaces in all lines (%) with a single space
:%s/  \+/ /g

# # replace double spaces in next 10 lines
:.,+10 s/  \+/ /g



files
=====

# save selected block in another file
#in visual mode type :w filename.txt

# open the file with recovery of data lost when system crashed
vi -r file

# commands
:r filename	read filename and insert after current line
:w filname	write current file to filename
:12,35w filen	write lines 12 to 35 to filen
:w! file	write current content to a preexisting filename



commands
========

# commands' list
:Explore		open the file explorer window
:set shell=/bin/bash	change shell
:shell		 	open shell window/start shell
set list		show tabs, EOL and other non printing chars

# spelling
:set spell spelllang=en

# spelling keys
]s or [s	move to previous or next misspelled word
z=		show suggestions
zg		add misspelled word to a dictionary
zw		mark word as incorrect

# autocomplete
<Ctrl>-<p> / <Ctrl>-<n>



exec
====

# commands' list
:!./%                          execute as a shell script
:w|!g++ % && ./a.out           build and execute a c++ program
:w|!g++ % -lpthread && ./a.out build and execute a program with multithreading
:w|!gcc % && ./a.out           build and execute a c++ program
:!date		execute date
:!g++ % && ./a.out  compile current file with g++ and run compiled program
:r !date	reads the output from executing the date command
:r filen	reads the contents of the filen to the current file
:w | !./%	save and execute current file
q:		show last vim commands
@:		execute last vim command
q/		show last searches



keys
====

# moving the cursor
0, ^    	move cursor to the beginning of the line
$		move cursor to the end of the line
w		move cursor to the beginning of the next word
e	    	move cursor to the end of the next word
b	    	move cursor to the beginning of the preceding word
:5 / 5G  	move cursor to line nr 5
%		move curosr to the matching parantheses/bracket
* / #		move cursor to next/previous instance of the current word
t<char>	move forward until the next occurrence of the character
f<char>		move forward over the next occurrence of the character
T<char>		move backward until the previous occurrence of the character
F<char>		move backward over the previous occurrence of the character
^f	    	move forward one screen
^b	    	move backward one screen
^d	    	move down half screen
^u	    	move up half screen
^y	    	move one line down
^e	    	move one line up
^l	    	redraws the screen
^r	    	redraws the screen removing deleted lines
{, }		move to previous/next paragragh
$ `		move to the last visited place
gi		move to the last place text was edited
g; / g,		move to the previous/next place text was edited
zz		move the current line to the middle of the screen
zt	    	move the current line to the top of the screen
zb	    	move the current line to the bottom of the screen

# add text
i		insert text before cursor
I		insert text at the beginning of the current line
a		append text after cursor
A	   	append text at the end of the current line
gi		insert text where edited last
o		open and put text in a new line below current line
O		open and put text in a new line above current line
u		undo last change
U		undo all changes to the line
R		redo undone changes
.		repeat last change
yiw		yank inner word (copy word under cursor, say "first").
viwp		select "second", then replace it with "first".
Vp		select line, paste what was yanked

# change text
r		replace a single character under the cursor
R		replace characters, starting with current cursor position
cw		change the current word, starting with the char under the cursor
cf<char>	change up to the first <char> (including);
ct<char>	change up to the first <char> (excluding);
c5w		change 5 words, starting with the char under the cursor
C		change (replace) the characters in the current line
cc		change (replace) entire line
5cc, c5c	change the next 5 lines starting with the current line
3s, c3l, 3cl	substitute 3 characters
J		append the line below to the current line
gUw		change the word to uppercase
guw		change the word to lowercase
gul		change the letter to lowercase
g~w		toggle lowercase/uppercase for a word
:set linebreak	wrap (adjust) text

# delete text
x		delete single character under cursor
5x	    	delete 5 characters, starting with the one under the cursor
dw	    	delete a single word starting wth the char under the cursor
d5w		delete 5 words beginning with the character under the cursor
d$		delete from cursor to the end of line
d/search<enter>	delete all characters until 'search' term
d^, d0		delete from cursor to the beginning of line
dd	    	delete entire line
5dd or d5d	delete 5 lines starting with the current one
:%normal 2x remove the first 2 characters of every line
:%s/^  /    remove first 2 characters of every line only if they're spaces, last slash is optional
:%normal << move indentation to left for every line

# delete many lines of text
go to the starting line and type ma (mark "a"), then go to the last line and enter d'a (delete to mark "a").

# copy, cut and paste
yw	 	copy the current word into the buffer
yy	    	copy the current line into the buffer
yi{	    	copy section between braces
ya{	    	copy section between braces with braces
y/search<enter>	copy all characters until 'search' term
5yy or y5y	copy 5 lines into the buffer, starting with the current one
y$         	copy cursor to the end of line
y^, y0     	copy  from cursor to the beginning of line
p	    	paste the line into the text after the current line
P	    	paste the line into the text above the current line
$ "*p`	    	paste the data from system register

# search and replace
/string		search forward for the string
?string		search backward for the string
n / N		move to next/previous occurence of search string
:s/the/THE/g	replaces the with THE in the current line
:s/ /\r/g	replace spaces with new lines
:s/THE/\rTHE/g	inserts new line before THE in the current line
:1,$ s/the/THE/g	replaces the with THE in lines from 1 till the end
:%s/the/THE/g	replaces the with THE in the document
&		repeat last substitution on the current line
g&		repeat last substitution on all lines
^g		returns the current and total line number



macros
======

# record macro at register 'a'
qa<commands>q

# run macro at register 'a'
@a

# run macro at register 'a' on five lines
5@a

# repeat last macro
@@



misc
====

# misc commands' list
:set all	displays the vim environment settings
:set opt	sets the option
:set tabstop=4	sets tab to 4 spaces
C-V, C-@	write an ascii control character
:set paste	turn on the paste mode to allow clean pasting from clipboard
:set nopaste	turn off the paste mode
:set textwidth=80
:set ruler	show cursor position
:C-g		show cursor position
ga              show hex and octal value of the character under the cursor

# set syntax highlighting for html
syntax on
set syntax=html



visual
======

# visual-operators*
aw	a word (with white space)
iw	inner word
aW	a WORD (with white space)
iW	inner WORD
as	a sentence (with white space)
is	inner sentence
ap	a paragraph (with white space)
ip	inner paragraph
ab	a () block (with parenthesis)
ib	inner () block
aB	a {} block (with braces)
iB	inner {} block
a<	a <> block (with <>)
i<	inner <> block
a[	a [] block (with [])
i[	inner [] block

# show spaces
:%s/ /█/g

# lowercase all text in a file
ggVGu
# Explanation:
	gg - goes to first line of text
	V - turns on Visual selection, in line mode
	G - goes to end of file (at the moment you have whole text selected)
	u - lowercase selected area
---
# or ggguG
# Explanation:
	gg - Goto the first line
	g  - start to converting from current line
	u  - Convert into lower case for all characters
	G  - To end of the file.

# switch text in subsitution to lowercase
:%s/[A-Z]/\L&/g

# switch text in subsitution to uppercase
:%s/[A-Z]/\U&/g

# enter into and exit from hex editor in binary mode
:set binary
:set textwidth=0
:set nomodeline
:set noexpandtab
:%!xxd
:%!xxd -r
or
vim -b file
:%!xxd -r



vimdiff
=======

# keys
ctrl+w,w	change window
zc		collapse (close) folding level
zo		expand (open) folding level
zR		expand (open) all folding levels everywhere



neovim
======

# run command in the background
:terminal ++hidden ./%

# escape from terminal mode
<Ctrl-\><Ctrl-n>

