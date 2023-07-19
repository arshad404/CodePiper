import java.util.ArrayList;
import java.util.List;

public class Para {
    public static void main(String[] args) {
        System.out.println("Hello World");
    }
}

class Solutions {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();   
        dfs(0, 0, res, "", n);
        return res;
    }

    public void dfs(int left, int right, List<String> res, String sb, int n){
        if(sb.length() == 2*n) {
            res.add(sb.toString());
            return;
        }

         if(n > left) {
            sb = sb + "(";
            dfs(left+1, right, res, sb, n);
        }

        if(left > right) {
            sb = sb + ")";
            dfs(left, right+1, res, sb, n);
        }
    }
}