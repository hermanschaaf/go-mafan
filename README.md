Go Mafan
============

Mafan is a Go package for handling everything in Chinese text processing that just makes your life `mafan` (troublesome). It will split strings into words, do parts-of-speech tagging, handle file encodings and more.

Right now it only has `Split`, which splits a string of Chinese text into a slice of words, but it aims to achieve all of the above eventually. Use like so:

    import (
    	"fmt"
    	"github.com/hermanschaaf/go-mafan"
    )

    func main() {
    	fmt.Println(mafan.Split("上海十大接吻聖地")) // prints ["上海", "十大", "接吻", "聖地"]
    }
    
To use it, you will need to [download the dictionary file from here](https://github.com/fxsjy/jieba/blob/master/extra_dict/dict.txt.big) and add it to a new data/ directory inside the mafan package. I'll improve this situation soon (but trying to avoid git repo size bloat).

************

This package is the Golang rewrite (not a port) of the original Python parent-project, also called [Mafan](http://github.com/hermanschaaf/mafan).
