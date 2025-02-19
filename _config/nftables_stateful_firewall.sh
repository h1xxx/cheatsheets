#!/bin/bash

# delete all rules
nft flush ruleset

# add main_table
nft add table inet main_table

# add the input, forward, and output base chains 
nft add chain inet main_table input   '{ type filter hook input   priority 0 ; policy drop ;   }'
nft add chain inet main_table forward '{ type filter hook forward priority 0 ; policy drop ;   }'
nft add chain inet main_table output  '{ type filter hook output  priority 0 ; policy accept ; }'

# add two regular chains for udp and tcp
nft add chain inet main_table tcp_chain
nft add chain inet main_table udp_chain

# accept related and established traffic
nft add rule inet main_table input ct state related,established accept

# accept all loopback interface traffic
nft add rule inet main_table input iif lo accept

# drop any invalid traffic
nft add rule inet main_table input ct state invalid drop

# accept icmp and optionally igmp for multicasting
nft add rule inet main_table input meta l4proto  icmp      accept
nft add rule inet main_table input meta l4proto  ipv6-icmp accept
nft add rule inet main_table input ip   protocol igmp      accept

# jump new udp traffic to the udp chain
nft add rule inet main_table input meta l4proto udp ct state new jump udp_chain

# jump new tcp traffic to the tcp chain
nft add rule inet main_table input 'meta l4proto tcp tcp flags & (fin|syn|rst|ack) == syn ct state new jump tcp_chain'

# reject all traffic that was not processed by other rules
nft add rule inet main_table input meta l4proto udp reject
nft add rule inet main_table input meta l4proto tcp reject with tcp reset

# count rejected traffic
nft add rule inet main_table input counter reject with icmpx port-unreachable

# accept ssh traffic on port 22
nft add rule inet main_table tcp_chain tcp dport 22 accept

# accept incoming dns requests
nft add rule inet main_table tcp_chain tcp dport 53 accept
nft add rule inet main_table udp_chain udp dport 53 accept

# block outgoing dns request (if there's a local dns resolver)
nft add rule inet main_table output oif != lo tcp dport 53 drop
nft add rule inet main_table output oif != lo udp dport 53 drop


# add nat with port forwarding
# ============================

# create nat table and routing chains
nft add table ip nat
nft add chain ip nat prerouting  '{ type nat hook prerouting  priority -100; }'
nft add chain ip nat postrouting '{ type nat hook postrouting priority  100; }'

# forward port to a different host
nft add rule ip nat prerouting iif enp6s0 tcp dport 10443 dnat to 192.0.2.1

# forward multiple ports to a different host
nft add rule ip nat prerouting iif enp6s0 tcp dport '{1022, 1080}' dnat to 192.0.2.1

# forward local port to a different local port
nft add rule ip nat prerouting tcp dport 8022 redirect to :22

# masquarade forwarded traffic going through a wan port enp6s0
nft add rule ip nat postrouting oif enp6s0 masquerade


