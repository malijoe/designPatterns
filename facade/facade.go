package facade

type Buffer struct {
    width, height int
    buffer []rune
}

func NewBuffer(width int, height int) *Buffer {
    return &Buffer{width: width, height: height, buffer: make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
    return b.buffer[index]
}

type Viewport struct {
    buffer *Buffer
    offset int
}

func NewViewport(buffer *Buffer) *Viewport {
    return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
    return v.buffer.At(v.offset+index)
}

type Console struct {
    buffers []*Buffer
    viewports []*Viewport
    offset int
}

func NewConsole() *Console {
    b := NewBuffer(200, 150)
    v := NewViewport(b)
    return &Console{buffers: []*Buffer{b}, viewports: []*Viewport{v}, offset: 0}
}

func (c *Console) GetCharacterAt(index int) rune {
    return c.viewports[0].GetCharacterAt(index)
}

func Facade() {
    // Console is a simpler way to interact with the entire ecosystem created above however, the individual components are still available for use
    c := NewConsole()
    _ = c.GetCharacterAt(1)
}
