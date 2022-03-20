package bmp

// BMP HEADER
// 66 77            field (42h, 4Dh)
// 128 0 0 0        Size of the BMP file
// 0 0              Application specific
// 0 0              Application specific
// 118 0 0 0        Offset where the pixel array (bitmap data) can be found

// 40 0 0 0         Number of bytes in the DIB header (from this point)
// 2 0 0 0          Width of the bitmap in pixels
// 2 0 0 0          Height of the bitmap in pixels. Positive for bottom to top pixel order.
// 1 0              Number of color planes being used
// 4 0              Number of bits per pixel
// 0 0 0 0          BI_RGB, no pixel array compression used
// 10 0 0 0         Size of the raw bitmap data (including padding)
// 18 11 0 0        horizontal	Print resolution of the image
// 18 11 0 0        vertical Print resolution of the image
// 0 0 0 0          Number of colors in the palette
// 0 0 0 0          means all colors are important

// bitmap data
// 0 255 255 255 0 0 0 255 0 0 255 0 0 255 0 0 0

// extra stuff at the end
// 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 35 0 0 0 1 0 0 0 0 0

// type BMPHeader struct {
//     content string
//     ID string
//     size string
//     bitmap_data_offeset string
// }

// type Header struct {
//     ID string
//     size int
//     bitmap_data_offeset int
// }

// func makeHeader () (header string) {
//     id := "BM"
//     // size := "42 4D" // BM
//     // id := "42 4D"
//     // id := "42 4D"

//     // header := "42 4D"
//     return id
// }

// func main() {
//     fmt.Println("Hello, World!")

//     data, err := os.ReadFile("./image.bmp")
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println("Image content:")
//     fmt.Println(data)
//     fmt.Println()

//     header := BMPHeader{
//         hex.EncodeToString(data[:14]),
//         hex.EncodeToString(data[:2]),
//         hex.EncodeToString(data[2:6]),
//         hex.EncodeToString(data[10:14]),
//     }

//     fmt.Println("BMP Header:")
//     fmt.Println(header.content)
//     fmt.Println()

//     fmt.Println(hex.EncodeToString([]byte("BM")))

//     // fmt.Println(data[118/4:])

//     // image := []byte{
//     //     66, 77,
//     //     70, 0, 0, 0,
//     //     0, 0,
//     //     0, 0,
//     //     54, 0, 0, 0,
//     //     40, 0, 0, 0,
//     //     2, 0, 0, 0,
//     //     2, 0, 0, 0,
//     //     1, 0,
//     //     4, 0,
//     //     0, 0, 0, 0,
//     //     10, 0, 0, 0,
//     //     18, 11, 0, 0,
//     //     18, 11, 0, 0,
//     //     0, 0, 0, 0,
//     //     0, 0, 0, 0,
//     //     0, 255, 255, 255,
//     //     0, 0, 0, 255,
//     //     0, 0, 255, 0,
//     //     0, 255, 0, 0,
//     // }

// }
