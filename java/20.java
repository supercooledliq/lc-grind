class Solution {
    public boolean isValid(String s) {
         Stack<Character> stack=new Stack<>();
        char[] ch=s.toCharArray();
        for(int i=0;i<ch.length;i++)
        {
            if(ch[i]== '(' || ch[i]=='[' || ch[i]=='{')
            stack.push(ch[i]);
            else if(ch[i]==')')
            {
                if(stack.isEmpty())
                return false;
                else if(stack.peek()!='(')
                return false;
                else
                stack.pop();
            }
              else if(ch[i]==']')
            {
                if(stack.isEmpty())
                return false;
                else if(stack.peek()!='[')
                return false;
                else 
                stack.pop();
            }
              else if(ch[i]=='}')
            {
                if(stack.isEmpty())
                return false;
                else if(stack.peek()!='{')
                return false;
                else
                stack.pop();
            }
        }
        if(stack.isEmpty())
        return true;
        else return false;
    }
}
