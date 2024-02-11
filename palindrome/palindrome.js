function isPalindrome(s) {
    for (let i = 0; i < s.length / 2; i++) {
      if (s[i] !== s[s.length - 1 - i]) {
        return false;
      }
    }
    return true;
  }
  
  const s = "katak";
  if (isPalindrome(s)) {
    console.log(s, "adalah palindrome");
  } else {
    console.log(s, "bukan palindrome");
  }
  