package proxy

import "fmt"

type Image interface {
    Draw()
}

type Bitmap struct {
    filename string
}

func (b *Bitmap) Draw() {
    fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
    fmt.Println("Loading image from", filename)
    return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
    fmt.Println("About to draw the image")
    image.Draw()
    fmt.Println("Done drawing the image")
}


//LazyBitmap is a virtual proxy because the underlying bitmap is only constructed once it is used
type LazyBitmap struct {
    filename string
    bitmap *Bitmap
}

func (l *LazyBitmap) Draw() {
    if l.bitmap == nil {
        l.bitmap = NewBitmap(l.filename)
    }
    l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
    return &LazyBitmap{filename: filename}
}

func VirtualProxy() {
    bmp := NewLazyBitmap("demo.png")
    DrawImage(bmp)
}
