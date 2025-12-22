# str = "abcdef"
# n = len(str)
# print("Độ Dài Của Chuỗi:",n)

# print(-n+1, "==", str[-n+1])
# print(-n, "==", str[-n])
# print(str[0], str[1])

##############################################
##                Ghép Chuỗi                ##
##############################################
# str1 = "Hien"
# str2 = "Qua"
# str3 = "Dep Trai"
# str = str1 + str2 + str3
# print(str)

##############################################
##                 Lặp Chuỗi                ##
##############################################
# st = "Hien"
# str1 = st*5
# str2 = 5*st
# print(str1)
# print(str2)

##############################################
##              Lấy Chuỗi Con               ##
##############################################
# str = "HienQuaDepTrai"
# str1 = str[1:4]
# str2 = str[0:5]
# str3 = str[4:7]
# print("Chuỗi Con Chỉ Số 1 Đến 3:", str1)
# print("Chuỗi Con Chỉ Số 0 Đến 4:", str2)
# print("Chuỗi Con Chỉ Số 4 Đến 6:", str3)

##############################################
##               Dạng Tham Số               ##
##############################################
# st = "abcdefgh"
# str1 = st[0:7:2]
# str2 = st[0:6:2]
# str3 = st[6:0:-1]
# print(str1, str2)
# print(str3)

##############################################
##                  Ép Kiểu                 ##
##############################################
str1 = "Hello"
n = 40
str2 = str1 + str(n)
print(str2)