# imagefn
--
    import "vimagination.zapto.org/imagefn"


## Usage

#### func  FlipX

```go
func FlipX(i image.Image) image.Image
```

#### func  FlipY

```go
func FlipY(i image.Image) image.Image
```

#### func  Invert

```go
func Invert(i image.Image) image.Image
```

#### func  Rotate180

```go
func Rotate180(i image.Image) image.Image
```

#### func  Rotate270

```go
func Rotate270(i image.Image) image.Image
```

#### func  Rotate90

```go
func Rotate90(i image.Image) image.Image
```

#### func  Scale

```go
func Scale(i image.Image, xScale, yScale float64) image.Image
```

#### func  ScaleDimensions

```go
func ScaleDimensions(i image.Image, x, y int) image.Image
```

#### func  SmoothScale

```go
func SmoothScale(i image.Image, xScale, yScale float64, scaler Scaler) image.Image
```

#### func  SmoothScaleDimensions

```go
func SmoothScaleDimensions(i image.Image, x, y int, scaler Scaler) image.Image
```

#### func  SubImage

```go
func SubImage(i image.Image, r image.Rectangle) image.Image
```

#### type Scaler

```go
type Scaler interface {
	At(image.Image, int, int) color.Color
}
```
