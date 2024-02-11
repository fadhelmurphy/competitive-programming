def is_palindrome(s):
  for i in range(len(s) // 2):
    if s[i] != s[-i-1]:
      return False
  return True

s = "katak"
if is_palindrome(s):
  print(s, "adalah palindrome")
else:
  print(s, "bukan palindrome")
