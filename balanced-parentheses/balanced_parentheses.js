function isBalanced(myStr) {
    const stack = [];
    const openList = ["[", "{", "("];
    const closeList = ["]", "}", ")"];
  
    for (const char of myStr) {
      if (openList.includes(char)) {
        stack.push(char);
      } else if (closeList.includes(char)) {
        const idx = closeList.indexOf(char);
        if (stack.length === 0 || openList[idx] !== stack.pop()) {
          return "Unbalanced";
        }
      }
    }
  
    return stack.length === 0 ? "Balanced" : "Unbalanced";
  }
  
  const strings = ["{[]{()}}", "[{}{})(]", "((()"];
  for (const str of strings) {
    console.log(str, "-", isBalanced(str));
  }
  