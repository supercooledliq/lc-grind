class Solution {
    public boolean isPalindrome(String s) {
        s=s.toLowerCase();
        s=s.replaceAll("\\s+", "");
        s=s.replaceAll("[^a-zA-Z0-9]", "");
        StringBuffer sb=new StringBuffer(s);
        StringBuffer rev=new StringBuffer(s);
        rev=rev.reverse();
        System.out.println("reverse is: "+rev);
        if(sb.toString().equals(rev.toString())==true)
        return true;
        else return false;
    }
}
