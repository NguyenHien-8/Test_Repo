##############################################
##                Tách Chuỗi                ##
##############################################
# str1 = "Hien Qua Dep Trai"
# L = str1.split()
# print("Chuỗi Được Tách Thành:", L)

##############################################
##         Thay Đổi Kí Tự Tách Chuỗi        ##
##############################################
# str1 = "Hien Qua, Dep, Trai"
# L = str1.split(",")
# print("Chuỗi Được Tách Thành:", L)

##############################################
##          Chuỗi Tối Đa Được Tách          ##
##############################################
# str1 = "Hien#Qua#Dep#Trai"
# L = str1.split("#", 2)
# print("Chuỗi Được Tách Thành:", L)

# str1 = "Hien#Qua#Dep#Trai"
# L = str1.split("#", -1) # Tham số mặc định -1 (nghĩa là tách tất cả)
# print("Chuỗi Được Tách Thành:", L)

##############################################
##    Chuyển Đổi Kí Tự Hoa, Kí Tự Thường    ##
##############################################
# st = "abCDeFgh456"
# st1 = "Sao Hien Dep Trai Qua Vay"
# st_hoa = st.upper()
# st_thuong = st.lower()
# st1_kitudauchuoiinhoa = st1.capitalize() # Chuyển đổi kí tự đầu tiên trong chuỗi thành kí tự in hoa
# st1_kitudauinhoa = st1.title() # chuyển đổi các kí tự đầu tiên của các từ thành kí tự in hoa
# st1_hoathuongthuonghoa = st1.swapcase() # chuyển đổi các kí tự hoa thành thường, kí từ thường thành hoa
# print("Ký Tự In Hoa:", st_hoa)
# print("Ký Tự In Thường:", st_thuong)
# print(st1_kitudauchuoiinhoa)
# print(st1_kitudauinhoa)
# print(st1_hoathuongthuonghoa)

##############################################
##               Phép Toán in               ##
##############################################
st1 = "Hello"
st2 = "Hello Hien"
st3 = "Tùng"
st4 = "Tùng Tùng Tùng"
if st1 in st2:
    print("st1 Là Con Của st2")
else:
    print("st1 Không Là Con Của st2")

if st4 in st3:
    print("st4 Là Con Của st3")
else:
    print("st4 Không Là Con Của st3")