# type hinting in PyCharm
Please pay attention to the # type comment in the following code:

```
# from typing import Tuple
i = 0 # type: int
a = (2, "4", 5)  # type: Tuple[int, str, int]

class C(object):
    foo = None # type: str

 def __init__(self, bar):
  self.bar = bar  # type: str
def f(x):
    # type: (str) -> Tuple[int, str]
  return len(x), "error"
```
Now PyCharm IDE can give you the type information when you press Ctrl and move the mouse to.

The above code is available in Python 2.7, that we can only declare the type as comment.

Beside comment, we can also define a file extended by "pyi" to declare the type information.

In Python 3.6+, we can declare the type directly in code instead of comment:
```
from typing import Tuple
i: int = 0
a: Tuple[int, str, int] = (2, "4", 5)
```
The above code is so similar with Go, Scala, Kotlin and TypeScript.

Reference: [Type Hinting in PyCharm](https://www.jetbrains.com/help/pycharm/type-hinting-in-product.html)

# How to define pyi file
Beside comment, we can also define a file extended by ".pyi" to declare the type information.

We can define pyi file in the python path, for example:
```
hello.py
hello.pyi
```
Here is hello.py, it defines f1 with type comment and f2 without type comment:
```
def f1(x):
    # type: (str) -> Tuple[int, str]
    return len(x), 56


def f2(x):
    return x.capitalize()
```
Now PyCharm can give type hint for f1, but not for f2. For test, we define another file hi.py and see the type hint when input code:
```
import hello

print hello.f1(1)
print hello.f2(2)
```
We can see f1 has type hint for f1, but not for f2, as expected.

Now we define f2 in pyi file. Here is hello.pyi:
```
def f2(x: str) -> str: ...
```
Now testing in hi.py, this time f2 has type hint, but f1 doesn't, and IDE doesn't know hello has f1 function. In hello.py, f2 still has no any type hint.

The above experiment shows type comment inline and pyi file are conflicted. As the testing result, I prefer type comment inline for the code that I can control, only use pyi file to the code that I cannot control. (We don't have to define pyi file in the same folder of py file, python will search in python path, so don't have to be the same folder, for more example you can refer to PyCharm's typings.py file)

# How to give warning if not match the type
In PyCharm, right click the python file, and select "Inspect Code...", then we can see the mis-match type.

