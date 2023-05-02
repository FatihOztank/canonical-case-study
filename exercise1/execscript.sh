#!/bin/bash
set -e

# Create a minimal filesystem, assuming host machine
# has busybox installed
mkdir -p rootfs/{bin,dev,etc,lib,lib64,proc,sbin,sys}
cp /bin/busybox rootfs/bin/
echo -e '#!/bin/busybox sh\n\n/bin/busybox --install -s\n' > rootfs/init

# insert hello world to init for printing "Hello World!" on startup
echo -e 'echo "hello world"\nexec /bin/sh' >> rootfs/init
chmod +x rootfs/init

# Create initramfs
cd rootfs
find . | cpio -o -H newc | gzip -9 > ../initramfs.gz
cd ..

# Run QEMU
qemu-system-x86_64 \
  -kernel bzImage \
  -initrd initramfs.gz \
  -nographic \
  -append "console=ttyS0 init=/init" \
