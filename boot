grub
====

# start kernel from grub console
set root=(usb1,hd1)
linux /vmlinuz-linux root=/dev/sdb2 rw
initrd /initramfs-linux.img
boot



misc
====

# inittab - start getty on the serial port
ttyFIQ0::respawn:/sbin/getty -L -n ttyFIQ0 115200 vt100

