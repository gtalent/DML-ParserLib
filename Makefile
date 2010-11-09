build: src/6.out
	mv src/6.out htmler
src/6.out:
	cd src;
	gomake;
	6l _go_.6
	cd ../;
