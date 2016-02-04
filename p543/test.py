#!/usr/bin/env python3

n = int(input())

isPrime = [False] * 2 + [True] * (n - 1)
primes = []
for p in range(2, n + 1):
  if isPrime[p]:
    primes.append(p)
    x = p * p
    while x <= n:
      isPrime[x] = False
      x += p

sP = [max(0, x // 2 - 2) for x in range(n + 1)]
for i in range(1, n + 1):
  if isPrime[i] and i >= 2 and isPrime[i - 2]:
    sP[i] += 2
  elif isPrime[i] or (i >= 2 and isPrime[i - 2]) or i % 2 == 0:
    sP[i] += 1

print(sum(sP))
