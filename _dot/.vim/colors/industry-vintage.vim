" Vim color file

" Reset to dark background, then reset everything to defaults:
set background=dark
highlight clear
if exists("syntax_on")
    syntax reset
endif

let colors_name = "industry"

" First set Normal to regular white on black text colors:
hi Normal ctermfg=LightGray ctermbg=Black

" Syntax highlighting (other color-groups using default, see :help group-name):
hi Comment    cterm=NONE ctermfg=DarkGrey
hi Constant   cterm=NONE ctermfg=DarkCyan
hi Identifier cterm=NONE ctermfg=DarkGreen
hi Function   cterm=NONE ctermfg=White
hi Statement  cterm=NONE ctermfg=Blue
hi PreProc    cterm=NONE ctermfg=DarkGreen
hi Type       cterm=NONE ctermfg=DarkGreen
hi Special    cterm=NONE ctermfg=DarkMagenta
hi Delimiter  cterm=NONE ctermfg=DarkMagenta

