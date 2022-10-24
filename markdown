pandoc
======

# create a pdf from markdown
pandoc -s --pdf-engine xelatex -o out.pdf file.md



commonmark
==========

# new line
<two trailing spaces>
---
<backslash at the end of line>

# double new line in vim
enter non breaking space <C-k><space><space> followed by an empty line

# draw a line
---------------

# italic text*
*text*

# bold text
**text**

# bold and italic
**bold _italic_**

# show a star
\*

# blockquote
> text

# numbered bullet points
* bullet 1
   Text that's in a paragraph
   > blockquote with a paragraph
* bullet 2

# unordered list
- item 1
- item 2

# ordered list
1. item 1
2. item 2

# escape bullet point / list item
1996\. line that is not a list

# nested bullet points and lists
* list 1
  * nested list 1
  * nested list 2
* list 2
  * nested list 1

# nested bullet points and lists
+ list 1
  1. nested list 1
  2. nested list 2
+ list 2
  1. nested list 1

# link
<http://www.duckduckgo.com>

# inline link
[google](http://www.duckduckgo.com)

# reference style link link
text [link] text
...
[link]: http://commonmark.org/help/images/favicon.png

# a table (colons mark the alignment)
| header | header | header |
|--------|:------:|-------:|
| a      |    b   |      c |
| 1      |    2   |      3 |
| foo    |   bar  |    baz |

