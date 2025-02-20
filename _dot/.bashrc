# unlimited history (value of 0 disables history)
HISTSIZE=""
HISTFILE="${HOME}/.bash_history"
HISTCONTROL=ignoredups:ignorespace
HISTIGNORE='l:ls:ll:la:lx:lm:d'
HISTTIMEFORMAT='%Y-%m-%d %H:%M:%S : '

green='\[\033[32m\]'
red='\[\033[31m\]'
white='\[\033[00m\]'

[ "${UID}" -eq 0 ] && color=${red} || color=${green}
PS1="${color}\u@\h: ${white}\W \$ "

# affects tmux copy mode
export EDITOR=vi

# limit the scope of gpg-agent
TTY="$(tty | cut -d'/' -f3- | sed 's|/||g')"
export GPG_TTY=$(tty)

shopt -s autocd
shopt -s checkhash
shopt -s histappend

alias d='pwd'
alias x='exit'
alias watch='watch'
alias cp='cp -a'
alias cl='clear'
alias l='ls -1FNL --group-directories-first'
alias ls='ls -FN --group-directories-first'
alias ll='ls -lhFN --group-directories-first'
alias lx='ls -lhFNB --group-directories-first'
alias la='ls -lhFNa --group-directories-first'
alias lm='ls -lFN --block-size=M --group-directories-first'
alias grep='grep -s --devices=skip --color=auto'
alias egrep='egrep -s --devices=skip --color=auto'
alias fgrep='fgrep -s --devices=skip --color=auto'
alias cal='ncal -M -b'
alias du='du -hs'
alias df='df -h'
alias free='free -h'
alias psgrep='ps aux | grep'

dusort() { du -h "${@}" | sort -h ;}

