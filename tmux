tmux
====

# recover crashed session to recreate all sockets
killall -10 tmux

# change the number of window 4 to 3, do this in window 4
<ctrl-char> move-window -t 3

# run a command in pane 0
tmux send-keys -t 0 'echo message' 'C-m'

# display info according to FORMATS spec in 'man tmux'
tmux display-message -p '#I'
tmux display-message -p '#{window_number}'
tmux display-message -p '#{pane_tty}'

# display info in session 0, window 1, pane 1
tmux display-message -p -t 0:1.2 '#{pane_tty}'
tmux -S /tmp/tmux-1000/default  display-message -p -t 0:0.0 '#{pane_tty}'

# list all panes in all windows
tmux list-panes -s

# list all panes with session and window numbers
tmux list-panes -s -F '#{session_name}:#{window_index}.#{pane_index}'

# autostart tmux after ssh
ssh hostname -t "tmux -CC attach || tmux -CC"

# tmux detach other clients when attaching (too many dots issue)
tmux a -d

