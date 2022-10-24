sc
==

# help
<?>

# enter a numeric value
<e>, then type: let D5 = 666

# enter a string
<E><a>

# enter a date (enter this formula after <e>)
@dts(15,12,31)

# format seconds since epoch as a date (^ is ctrl, see format in man strftime)
<F><^d>%Y-%m-%y

# format numeric values, columns and rows
<f>, then
  <j>,<k>   change number of decimal places shown
  <h>,<l>   change column width

# fill a range with increment of one
<r><f>, (select range), tab, 1, 1

# save current file in spreadsheet format
<P> <filename.sc>
<ZZ>

# save current file in text format
<P> <filename.sc>

