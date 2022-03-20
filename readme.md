# BMP

A simple BMP creation and edition go library.

## Ressources

https://en.wikipedia.org/wiki/BMP_file_format
https://medium.com/sysf/bits-to-bitmaps-a-simple-walkthrough-of-bmp-image-format-765dc6857393
http://www.ece.ualberta.ca/~elliott/ee552/studentAppNotes/2003_w/misc/bmp_file_format/bmp_file_format.htm
https://dotink.co/posts/bmp/

## Notes

#### Padding of scan lines

Padding bytes (not necessarily 0) must be appended to the end of the rows in order to bring up the length of the rows to a multiple of four bytes. When the pixel array is loaded into memory, each row must begin at a memory address that is a multiple of 4. This address/offset restriction is mandatory only for Pixel Arrays loaded in memory. For file storage purposes, only the size of each row must be a multiple of 4 bytes while the file offset can be arbitrary.[[5]](https://en.wikipedia.org/wiki/BMP_file_format#cite_note-DIBhelp-5) A 24-bit bitmap with Width=1, would have 3 bytes of data per row (blue, green, red) and 1 byte of padding, while Width=2 would have 6 bytes of data and 2 bytes of padding, Width=3 would have 9 bytes of data and 3 bytes of padding, and Width=4 would have 12 bytes of data and no padding.

#### How to use localy

`go mod edit -replace github.com/bmp=./local-path-to-bmp-folder`  
`go mod tidy`
