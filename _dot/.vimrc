" ---- function definitions

" expand tab to spaces for text alignment, don't expand for indentation
function! TabOnlyIndent()
  if strpart(getline('.'), 0, col('.') - 1) =~ '^\s*$'
    return "\<Tab>"
  else
    set expandtab
    execute "normal! i\<Tab>"
    execute "normal l"
    set noexpandtab
  endif
  return ""
endfunction

" when editing a file, always jump to the last known cursor position;
" don't do it when the position is invalid, when inside an event handler
" and for a commit message (it's " likely a different one than last time)
function! PrevPosition()
  if line("'\"") >= 1 && line("'\"") <= line("$") && &ft !~# 'commit'
    execute "normal! g`\""
  endif
endfunction

" toggle pasting with a nice message on the current state
function! PasteToggle()
  if &paste
    call feedkeys(":set nopaste\<CR>", "n")
  else
    call feedkeys(":set paste\<CR>", "n")
  endif
endfunction

" toggle spell check - first to "en", then to "pl", then turn if off
function! SpellToggle()
  if &spell && &spelllang == "pl"
    call feedkeys(":set nospell\<CR>", "n")
  elseif &spell && &spelllang == "en"
    call feedkeys(":set spelllang=pl\<CR>", "n")
  else
    set spell
    call feedkeys(":set spelllang=en\<CR>", "n")
  endif
endfunction

" toggle spell check in insert mode (doesn't print commands)
function! SpellToggleInsert()
  if &spell && &spelllang == "pl"
    set nospell
  elseif &spell && &spelllang == "en"
    set spell
    set spelllang=pl
  else
    set spell
    set spelllang=en
  endif
endfunction

" toggle syntax highlighting
function! SyntaxToggle()
  if exists("g:syntax_on")
    call feedkeys(":syntax off\<CR>", "n")
  else
    call feedkeys(":syntax on\<CR>", "n")
  endif
endfunction

" toggle syntax highlighting in insert mode
function! SyntaxToggleInsert()
  if exists("g:syntax_on")
    syntax off
  else
    syntax on
  endif
endfunction

" toggle highlighting the column just behind the text width
function! ColorColumnToggle()
  if &colorcolumn != ""
    set colorcolumn=
  else
    set colorcolumn=+1
    highlight colorcolumn ctermbg=lightgrey ctermfg=black
  endif
endfunction


" ---- general settings

" do not execute defaults.vim
let skip_defaults_vim = 1

" disable vi compatible mode (noop in most cases)
set nocompatible

" set filetype detection
filetype on

" dont' expand tabs
set noexpandtab

" set tabs to have 8 spaces (tabstop)
set ts=8

" when using the >> or << commands, shift lines by 8 spaces
set shiftwidth=8

" indent when moving to the next line while writing code
set autoindent

" enable syntax highlighting
syntax on

" set line numbers: nu/nonu (numbers/nonumbers)
set nonu

" show the matching part of the pair for [] {} and ()
set showmatch

" limit number of characters in one line
set textwidth=79

" always use 16 colors because linux console is the best
set t_Co=16

" set color scheme
colorscheme industry

" stop search at the end of the file
set nowrapscan 

" unset the "last search pattern" register by hitting <Enter>
nnoremap <CR> :noh<CR><CR>

" command history buffer size
set history=128

" show cursor position
set ruler

" disable insecure modelines
set nomodeline

" disable backup files
set nobackup nowritebackup

" allow backspacing over everything in insert mode
set backspace=indent,eol,start

" display incomplete commands
set showcmd

" display completion matches in a status line
set wildmenu

" time out for key codes - wait up to 100ms after Esc for special key
set ttimeout
set ttimeoutlen=100

" show @@@ in the last line if it is truncated
set display=truncate

" always show a few lines of context around the cursor
set scrolloff=5

" do incremental searching
set incsearch

" when editing a file, always jump to the last known cursor position
autocmd BufReadPost * call PrevPosition()


" ---- vim file settings

" use common convention for tabs
au FileType vim set expandtab ts=2 shiftwidth=2


" ---- c file settings

" expand tab to spaces for text alignment, don't expand for indentation
au FileType c inoremap <Tab> <C-r>=TabOnlyIndent()<CR>

" save, compile and run
au FileType c nnoremap <silent> <F5> :w<CR>:!tcc -run %<CR>


" ---- shell file settings

" use google convention for tabs
au FileType sh set expandtab ts=2 shiftwidth=2

" save, compile and run
au FileType sh nnoremap <silent> <F5> :w<CR>:!go run %<CR>


" ---- go file settings

" expand tab to spaces for text alignment, don't expand for indentation
au FileType go inoremap <Tab> <C-r>=TabOnlyIndent()<CR>

" save, compile and run
au FileType go nnoremap <silent> <F5> :w<CR>:!go run %<CR>


" ---- python file settings

" disable all Python syntax highlighting features
let python_highlight_all = 0

" save and run the script
au FileType python nnoremap <silent> <F5> :w<CR>:!python %<CR>

" use common convention for tabs
au FileType python set expandtab ts=4 shiftwidth=4


" ---- yaml file settings

" use common convention for tabs
au FileType yaml set expandtab ts=2 shiftwidth=2


" ---- html file settings

" set html syntax in gohtml files
au BufRead *.gohtml set syntax=html
au FileType html set expandtab ts=2 shiftwidth=2


" ---- mutt file settings

" limit char size of line when editing in mutt
au BufRead /tmp/mutt-* set tw=1024


" ---- cheatsheet file settings

" expand tab to spaces for text alignment, don't expand for indentation
au BufRead /*/cheatsheets/* inoremap <Tab> <C-r>=TabOnlyIndent()<CR>


" ---- misc files
au BufRead hosts set expandtab ts=2 shiftwidth=2


" ---- key mappings

" toggle pasting
set pastetoggle=<F1>
nnoremap <silent> <F1> :call PasteToggle()<CR>

" toggle spell check - first to "en", then to "pl", then turn if off
nnoremap <silent> <F2> :call SpellToggle()<CR>
inoremap <silent> <F2> <C-o>:call SpellToggleInsert()<CR>

" toggle syntax highlighting
nnoremap <silent> <F3> :call SyntaxToggle()<CR>
inoremap <silent> <F3> <C-o>:call SyntaxToggleInsert()<CR>

" toggle highlighting the column just behind the text width
nnoremap <silent> <F4> :call ColorColumnToggle()<CR>
inoremap <silent> <F4> <C-o>:call ColorColumnToggle()<CR>

