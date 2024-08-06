"""
A[i] =  (x[0] + x[1] + ... + x[i-1] +  x[i]) / [i + 1]
"""

arr = [5, -1, 3, 7, -2]
p = []
sum = 0

for i in range(len(arr)):
    sum += arr[i]
    p.append(sum / (i + 1))

print(p)
