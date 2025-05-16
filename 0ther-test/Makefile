MNTPNT = ./env
MNTSIZE = 8G
TIMER = time -v

ifeq (, $(shell which time))
	TIMER =
endif

all: env test checksum done

avg:
	dd if=/dev/random of=$(MNTPNT)/files/first.txt bs=1M count=1024
	dd if=/dev/random of=$(MNTPNT)/files/second.txt bs=1M count=1024
lrg:
	dd if=/dev/random of=$(MNTPNT)/files/first.txt bs=1M count=2048
	dd if=/dev/random of=$(MNTPNT)/files/second.txt bs=1M count=2048
mount:
	sudo mkdir -p $(MNTPNT)
	sudo mount -o size=$(MNTSIZE) -t tmpfs ramdisk $(MNTPNT)
	mkdir -p $(MNTPNT)/files
umount:
	sudo umount $(MNTPNT)
clean:
	find $(MNTPNT)/files ! -name first.txt ! -name second.txt -type f -delete
diff:
	echo "1" >> $(MNTPNT)/files/first.txt
	echo "01" >> $(MNTPNT)/files/second.txt

done: umount
	rm -rf $(MNTPNT)

lnk:
	rm -f $(MNTPNT)/go-xor
	rm -f $(MNTPNT)/go-xor-one-thread
	rm -f $(MNTPNT)/c-xor
	rm -f $(MNTPNT)/php-xor
	rm -f $(MNTPNT)/python-xor

	ln -s ../../bin/xorfiles-go $(MNTPNT)/go-xor
	ln -s ../../0ther-impl/Go-one/bin/xorfiles-go-one $(MNTPNT)/go-xor-one-thread
	ln -s ../../0ther-impl/C/bin/xorfiles-c $(MNTPNT)/c-xor
	ln -s ../../0ther-impl/PHP/xorfiles.php $(MNTPNT)/php-xor
	ln -s ../../0ther-impl/Python/xorfiles.py $(MNTPNT)/python-xor

env: mount lrg diff lnk

test:
	@echo "\n********** Go implementation with 2 goroutines **********"
	$(TIMER) $(MNTPNT)/go-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt | pv > /dev/null

	@echo "\n********** Go implementation with 1 goroutine **********"
	$(TIMER) $(MNTPNT)/go-xor-one-thread $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt | pv > /dev/null

	@echo "\n********** C implementation, no threads  **********"
	$(TIMER) $(MNTPNT)/c-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt | pv > /dev/null

	@echo "\n********** PHP implementation, no threads  **********"
	$(TIMER) $(MNTPNT)/php-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt | pv > /dev/null

	@echo "\n********** Python implementation, no threads  **********"
	$(TIMER) $(MNTPNT)/python-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt | pv > /dev/null

checksum:
	@echo "\n********** Go implementation with 2 goroutines **********"
	$(MNTPNT)/go-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt > $(MNTPNT)/files/go-output.txt
	sha256sum $(MNTPNT)/files/go-output.txt
	@rm -f $(MNTPNT)/files/go-output.txt

	@echo "\n********** Go implementation with 1 goroutine **********"
	$(MNTPNT)/go-xor-one-thread $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt > $(MNTPNT)/files/go-one-output.txt
	sha256sum $(MNTPNT)/files/go-one-output.txt
	@rm -f $(MNTPNT)/files/go-one-output.txt

	@echo "\n********** C implementation, no threads  **********"
	$(MNTPNT)/c-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt > $(MNTPNT)/files/c-output.txt
	sha256sum $(MNTPNT)/files/c-output.txt
	@rm -f $(MNTPNT)/files/c-output.txt

	@echo "\n********** PHP implementation, no threads  **********"
	$(MNTPNT)/php-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt > $(MNTPNT)/files/php-output.txt
	sha256sum $(MNTPNT)/files/php-output.txt
	@rm -f $(MNTPNT)/files/php-output.txt

	@echo "\n********** Python implementation, no threads  **********"
	$(MNTPNT)/python-xor $(MNTPNT)/files/first.txt $(MNTPNT)/files/second.txt > $(MNTPNT)/files/python-output.txt
	sha256sum $(MNTPNT)/files/python-output.txt
	@rm -f $(MNTPNT)/files/python-output.txt
