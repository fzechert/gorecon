# GoRecon
OS independent implementation of a driver for the [BitFenix Recon Fan Controller](https://www.bitfenix.com/global/en/products/accessories/recon/) in GO

# Dependencies
## For Development / Compiling
To compile this application you will need the `gousb` package. See https://github.com/google/gousb on how to install `gousb`. You will also need the `libusb` library. See https://sourceforge.net/projects/libusb/.

## For Using / Running
Make sure your operating system can find a libusb library. E.g. `libusb-1.0.dll` or `libusb-1.0.so.0`. See https://sourceforge.net/projects/libusb/ on how to obtain one if you don't have one.
