//recursion
// class Solution {
//     public int fib(int n) {
//         if(n==0||n==1)
//         return n;
//         return fib(n-1)+fib(n-2);
//     }
// }

//with dynamic programming
// class Solution{
//     public int fib(int n){
//         if(n==0)
//         return 0;
//     int f[]=new int[n+1];
//     f[0] = 0;
//     f[1] = 1;
//     for(int i=2;i<=n;i++)
//     {
//         f[i]=f[i-1]+f[i-2];
//     }
//     return f[n];
// }
// }

class Solution{
    public int fib(int n)
    {
        if(n<=1)
        return n;
        int a=0,b=1,c=0;
        for(int i=2;i<=n;i++)
        {
            c=a+b;
            a=b;
            b=c;
        }
        return c;
    }
}
