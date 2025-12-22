"""""""""""""""""""""
Kiểu Số Nguyên - int
"""""""""""""""""""""
a = 28
b = 3
luythua = a**b
print(a, "^", b, "=", luythua)


"""""""""""""""""""""
Kiểu Số Thực - float
"""""""""""""""""""""
x = 5.44
y = 4.11
print("Số Thực x =", x, "\nSố Thực y =", y)
i = 4e4
print("Giá Trị Của i: ", i)

"""""""""""""""""""""
Hàm Tạo Số Ngẫu Nhiên
"""""""""""""""""""""
from random import randint, uniform

x = randint(-10,5)
y = uniform(-3,10)
print("Số Nguyên Tạo Được:", x)
print("Số Thực Tạo Được:", y)