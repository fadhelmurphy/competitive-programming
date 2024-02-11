
# Contoh Soal Competitive Programming dengan Python, Golang, dan JavaScript
# Soal:

# Diberikan sebuah string yang terdiri dari karakter 0 dan 1. Hitunglah jumlah substring yang memiliki jumlah karakter 0 dan 1 yang sama.

# Contoh:

# Input: 1101
# Output: 3
# Penjelasan:

# Substring dengan jumlah 0 dan 1 sama: 11, 01, dan 101.

def count_substrings(s):
  count = 0
  zeros = ones = 0
  prev_char = ''
  for char in s:
    if char == '0':
      zeros += 1
    else:
      ones += 1
    if zeros <= ones and (prev_char != '1' or char != '0'):
      count += 1
    prev_char = char
  return count

s = '1101'
print(count_substrings(s))
