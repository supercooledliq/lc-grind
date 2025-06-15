class Solution {
    public int reverse(int x) {
        int rev=0;
        int rem;
        int min=Integer.MIN_VALUE;
        int max=Integer.MAX_VALUE;

        while (x!=0) {
            rem=x%10;

        //pls check for overflow/underflow before doing rev * 10 + rem
            if (rev>max/10||(rev==max/10&&rem>7)){
                return 0;
            }
            if (rev<min/10||(rev==min/10&&rem<-8)){
                return 0;
            }

            rev=rev*10+rem;
            x=x/10;
        }
        return rev;
    }
}
