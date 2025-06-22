//brute force,TLE 
// class Solution {
//     public void rotate(int[] nums, int k) {
//         for(int i=0;i<k;i++)
//         {
//             int temp=nums[nums.length-1];
//             for(int j=nums.length-1;j>0;j--)
//             {
//                 nums[j]=nums[j-1];
//             }
//             nums[0]=temp;
//         }
//     }
// }

class Solution {
    public void rotate(int[] nums, int k) {
       if (nums.length <= 1 || k == 0) 
      System.out.println(Arrays.toString(nums)); // Edge case handling
        k=k%(nums.length) ;//doing this because if k=10 its basically k=1 and if k=50 k is 5
        reverse(nums,0,(nums.length)-k-1);
        reverse(nums,nums.length-k,nums.length-1);
        reverse(nums,0,nums.length-1);
    }
    public void reverse(int nums[],int start,int end)
    {
        while(start<end)
        {
            int temp=nums[start];
            nums[start]=nums[end];
            nums[end]=temp;
            start++;
            end--;
        }
        }
         }
