// Contoh Soal Competitive Programming dengan Python, Golang, dan JavaScript
// Soal:

// Diberikan sebuah string yang terdiri dari karakter 0 dan 1. Hitunglah jumlah substring yang memiliki jumlah karakter 0 dan 1 yang sama.

// Contoh:

// Input: 1101
// Output: 3
// Penjelasan:

// Substring dengan jumlah 0 dan 1 sama: 11, 01, dan 101.

function countSubstrings(s) {
    let count = 0;
    let zeros = 0;
    let ones = 0;
    let prevChar = '';
    for (const char of s) {
      if (char === '0') {
        zeros++;
      } else {
        ones++;
      }
      if (zeros <= ones && (prevChar !== '1' || char !== '0')) {
        count++;
      }
      prevChar = char;
    }
    return count;
  }
  
  const s = '1101';
  console.log(countSubstrings(s));
  