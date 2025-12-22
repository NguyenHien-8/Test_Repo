##############################################
##            Xác Nhận Bạn Hay Bồ           ##
##############################################

# Tuổi_Tôi = int(input("Nhập Số Tuổi Của Bạn: "))
# Tuổi_Bồ = int(input("Nhập Số Tuổi Của Bồ: "))

# if(Tuổi_Tôi >= Tuổi_Bồ):
#     print("2 Đứa Mình Yêu Nhau")
# else: 
#     print("2 Đứa Mình Coi Là Bạn")

##############################################
##              Xác Định Số Thứ             ##
##############################################

# NhapThu = int(input("Nhập Số Thứ: "))

# if(NhapThu == 1):
#     print("Thứ Chúa Nhật")
# elif(NhapThu == 2):
#     print("Thứ Hai")
# elif(NhapThu == 3):
#     print("Thứ Ba")
# elif(NhapThu == 4):
#     print("Thứ Tư")
# elif(NhapThu == 5):
#     print("Thứ Năm")
# elif(NhapThu == 6):
#     print("Thứ Sáu")
# elif(NhapThu == 7):
#     print("Thứ Bảy")
# else:
#     print("Bạn Nhập Sai Số Thứ")

##-------------------------Bài Tập 6--------------------------##
################################################################
##      Cho điểm thi Toán,Lý,Hóa,Sinh,Tin của học sinh        ##
##      Hãy xếp loại kết quả của học sinh theo tiêu chí       ##
################################################################

DiemToan = int(input("Nhập Điểm Toán: "))
DiemLy = int(input("Nhập Điểm Lý: "))
DiemHoa = int(input("Nhập Điểm Hóa: "))
DiemSinh = int(input("Nhập Điểm Sinh: "))
DiemTin = int(input("Nhập Điểm Tin: "))

DTB = (DiemToan + DiemLy + DiemHoa + DiemSinh + DiemTin) / 5
print("Điểm Trung Bình =", DTB)

if(9 <= DTB):
    print("Đạt Xuất Sắc")
elif(8 <= DTB < 9):
    print("Đạt Giỏi")
elif(7 <= DTB < 8):
    print("Đạt Khá")
elif(5 <= DTB < 7):
    print("Đạt Trung Bình")
elif(DTB < 5):
    print("Đạt Yếu")
