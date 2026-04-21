for i in range(256):
    # print(hex(i).encode("cp1251").decode("cp1251"))
    try:
        print(bytes([i]).decode("cp1251"))
    except:
        print(i)
# print("привет".encode("cp1251").decode("cp1251"))
