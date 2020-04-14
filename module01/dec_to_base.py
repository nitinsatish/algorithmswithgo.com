import sys

def dec_to_base(num, base) :
    result = ""
    while num > 0 :
      rem = num%base
      result = "%X"%rem +result
      num = num // base
    return result

if __name__ == "__main__":
    num, base = sys.argv[1], sys.argv[2]
    output = dec_to_base(int(num),int(base))
    print(output)