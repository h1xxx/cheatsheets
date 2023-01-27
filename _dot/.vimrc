" ---- general settings

" do not execute defaults.vim
let skip_defaults_vim = 1

" disable vi compatible mode (noop in most cases)
set nocompatible

" set filetype detection
filetype off

" dont' expand tabs
set noexpandtab

" set tabs to have 8 spaces (tabstop)
set ts=8

" when using the >> or << commands, shift lines by 8 spaces
set shiftwidth=8

" expand tab to spaces for text alignment, don't expand for indentation
function! TabOnlyIndent()
if strpart(getline('.'), 0, col('.') - 1) =~ '^\s*$'
    return "\<Tab>"
  else
    return "    "
  endif
endfunction
imap <Tab> <C-r>=TabOnlyIndent()<CR>

" indent when moving to the next line while writing code
set autoindent

" disable syntax highlighting
syntax off

" set line numbers: nu/nonu (numbers/nonumbers)
set nonu

" show the matching part of the pair for [] {} and ()
set showmatch

" limit number of characters in one line
set textwidth=79

" always use 16 colors because linux console is the best
set t_Co=16

" set colors
colorscheme industry-vintage

" stop search at the end of the file
set nowrapscan 

" unset the "last search pattern" register by hitting <Enter>
nnoremap <CR> :noh<CR><CR>

" command history buffer size
set history=512

" show cursor position
set ruler

" disable insecure modelines
set nomodeline

" disable backup files
set nobackup nowritebackup



" ---- mutt settings

" limit char size of line when editing in mutt
au BufRead /tmp/mutt-* set tw=1024



" ---- c settings

" disable converting tabs to spaces
au FileType c set noexpandtab

" when using the >> or << commands, shift lines by 8 spaces
au FileType c set shiftwidth=8
au FileType c set ts=8

" assign keybindings
au FileType c nnoremap <silent> <F5> :w<CR>:!tcc -run %<CR>



" ---- python settings

" disable all Python syntax highlighting features
let python_highlight_all = 0

" assign keybindings
au FileType python nnoremap <silent> <F5> :w<CR>:!python %<CR>

" set tabs to have 4 spaces (tabstop)
au FileType python set ts=4

" expand tabs into spaces
au FileType python set expandtab

" when using the >> or << commands, shift lines by 4 spaces
au FileType python set shiftwidth=4



" ---- yaml settings

" when using the >> or << commands, shift lines by 2 spaces
au FileType yaml set shiftwidth=2
au FileType yaml set ts=2



" ---- url handling with w3m - eperimental and not working

function! Browser ()
  let line = getline (".")
  let line = matchstr (line, "\%(http://\|www\.\)[^ ,;\t]*")
  exec "!echo" .line
  "exec 
  "!w3m ".line
endfunction
map <Leader>w :call Browser ()<CR>

function! HandleURL()
  let s:uri = matchstr(getline("."), '[a-z]*:\/\/[^ >,;]*')
  echo s:uri
  if s:uri != ""
    silent exec "!echo '".s:uri."'"
  else
    echo "No URI found in line."
  endif
endfunction
map <leader>u :call HandleURL()<cr>



" ---- not used settings

" column highlighting

"set colorcolumn=81
"highlight ColorColumn ctermbg=7 ctermfg=7
"highlight ColorColumn ctermbg=LightGrey guibg=LightGrey

" change highlighting of bold and italic text in markdown
"hi link htmlBold Storage
"hi link htmlItalic Identifier

" paste text as is in tmux buffer
"set paste 

" set a macro
"let @a="iHello World!\<CR>bye\<Esc>"

