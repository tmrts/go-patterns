# Fluent interface

Fluent interface is a method for constructing object oriented APIs
focused on better readability. Implemented by using method chaining

## Implementation

```go
package fluent

type Align int

const (
	Left = iota
	Center = iota
	Right = iota
)

type TextContainer struct {
    Font string
    FontSize int
    Color string
    Align Align
}

func NewTextContainer() *TextContainer {
    return &TextContainer{
        Font: "Helvetica",
        FontSize: 10,
        Color: "Red",
        Align: Left}
}

func (tc *TextContainer) SetFont (font string) *TextContainer {
    tc.Font = font
    return tc
}

func (tc *TextContainer) SetFontSize (fontSize int) *TextContainer {
    tc.FontSize = fontSize
    return tc
}

func (tc *TextContainer) SetColor (color string) *TextContainer {
    tc.Color = color
    return tc
}

func (tc *TextContainer) SetAlign (align Align) *TextContainer {
    tc.Align = align
    return tc
}
```

## Usage
```go
func getAlignText(align Align) string{
    switch align{
        case Left:
            return "Left"
        case Center:
            return "Center"
        case Right:
            return "Right"
        default:
            panic("Not determinated align")
    }
}

tContainer := NewTextContainer()
tContainer.SetFont("Calibri").SetFontSize(20).SetColor("Green").SetAlign(Center)

fmt.Printf("Font: %s\n", tContainer.Font)
fmt.Printf("Font size: %v\n", tContainer.FontSize)
fmt.Printf("Color: %s\n", tContainer.Color)
fmt.Printf("Align: %s\n", getAlignText(tContainer.Align))
```