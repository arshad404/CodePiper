import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.letterCombinations("23"));

    }
}

class Solution {
    List<String> store = new ArrayList<>(Arrays.asList("", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"));

    public List<String> letterCombinations(String digits) {
         StringBuilder sb = new StringBuilder();
         List<String> res = new ArrayList<>();

         dfs(0, digits, res, sb);
         return res;
    }

    public void dfs(int index, String digits, List<String> res, StringBuilder sb) {
        if(index >= digits.length()) {
            res.add(sb.toString());
            return;
        }

        String options = store.get(Character.getNumericValue(digits.charAt(index)));
        for(char option : options.toCharArray()) {
            sb.append(option);
            dfs(index+1, digits, res, sb);
            sb.deleteCharAt(sb.length() - 1);
        }
    }
}