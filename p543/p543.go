package main

import "os"
import "fmt"
import "strconv"

func sieve(n int) []bool {
  s := make([]bool, n + 1)
  s[0] = true
  s[1] = true
  for i := 2; i <= n; i++ {
    if !s[i] {
      for x := i * i; x <= n; x += i {
        s[x] = true
      }
    }
  }
  return s
}

func main() {
  var n, a, b int
  n, _ = strconv.Atoi(os.Args[1])
  a, b = 1, 2
  pos := make([]int, 0)
  for i := 3; i <= n; i++ {
    a, b = b, a + b
    pos = append(pos, a)
  }

  s := sieve(a)
  cur := 0
  ans := 0
  j := 0
  for i := 1; i <= a; i++ {
    base := i / 2 - 2
    if base < 0 {
      base = 0
    }
    if !s[i] && i >= 2 && !s[i - 2] {
      base += 2
    } else if !s[i] || (i >= 2 && !s[i - 2]) || i % 2 == 0 {
      base += 1
    }
    for pos[j] < i {
      j += 1
    }
    cur += base
    if pos[j] == i {
      fmt.Println(i)
      fmt.Println(cur)
      ans += cur
    }
  }
  fmt.Println(ans)
}
