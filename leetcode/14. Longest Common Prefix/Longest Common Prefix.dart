String longestCommonPrefix(List<String> strs) {
  var prefix = '';
  if (strs.isEmpty) return prefix;

  var firstWord = strs[0];
  for (var index = 0; index < firstWord.length; index++) {
    for (var i = 1; i < strs.length; i++) {
      var word = strs[i];
      if (index >= word.length || word[index] != firstWord[index]){
        return prefix;
      }
    }
    prefix += firstWord[index];
  }

  return prefix;
}

void main() {
  print(longestCommonPrefix(["flower", "flow", "flight"]));
  print(longestCommonPrefix(["dog", "racecar", "car"]));
  print(longestCommonPrefix(["a"]));
  print(longestCommonPrefix([]));
}
