# Using OpenCV Contrib

The OpenCV Contrib library contains experimental or non-free (aka patented) algorithms.

GoCV support for OpenCV Contrib can be found here in the "github.com/xuxiaoshuo/gocv/contrib" package.

For more information about OpenCV Contrib, please go to:

https://github.com/opencv/opencv_contrib

## How to use

If you have followed the installation instructions from the main README, then the OpenCV contrib modules have already been compiled and installed.

First, you must include the `contrib` subpackage:

```go
import (
    "github.com/xuxiaoshuo/gocv"
    "github.com/xuxiaoshuo/gocv/contrib"
)
```

Then you will be able to use the functions within the `contrib` subpackage. For example, this uses the `SIFT` feature identitification algorithm that is within the `xfeatures2d` module of OpenCV:

```go
si := contrib.NewSIFT()
kp := si.Detect(img)
```

Note that some of the features in this package require building OpenCV with the `OPENCV_ENABLE_NONFREE=ON` option. You can run `make build_nonfree` to build with this option.
