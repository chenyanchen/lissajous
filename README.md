# lissajous

Lissajous is an image convertor.

Convert your images into jpeg(jpg), current supported image formats is jpeg,
jpg, png, gif.

## Usage

| Parameter | Type   | Description                                                                                           | 
| --------- | ------ | ----------------------------------------------------------------------------------------------------- | 
| -dst      | string | set the destination directory path (default "tmp")                                                    |
| -env      |        | print runtime environment                                                                             |
| -h        |        | usage of lissajous                                                                                    |
| -q        | int    | set the quality of the converted image, ranges from 1 to 100 inclusive, higher is better (default 75) |
| -src      | string | set the source directory or image file path (default ".")                                             |
| -v        |        | print lissajous version                                                                               |

## Build

`make build`
