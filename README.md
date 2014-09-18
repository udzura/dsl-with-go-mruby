dsl-with-go-mruby
=================

DSL test with go-mruby

## How to run

```console
$ make
//...

go run sample.go
src is:
foobar_block do
  buz  "hogehoge"
  fizz "fugafuga"
end

buz  "waiwai"
fizz "wakuwaku"

parsed data is:
map[buz:hogehoge fizz:fugafuga]
```

After a `libmruby.so` created:

```bash
$ go run sample.go
```

Please feel free to modify `Samplefile`.

## Copyright

Same as mruby(Golang is under BSD License, so it is reasonable).

```
The MIT License (MIT)

Copyright © 2014 Uchio KONDO

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in the
Software without restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
