//takes 2ms 
//kadanes algo
class Solution {
    public int maxSubArray(int[] nums) {
        int MaxEnding=nums[0];
        int MaxSoFar=nums[0];
        if(nums.length==1)
        return nums[0];
        else
       { 
        for(int i=1;i<nums.length;i++)
        {
            MaxEnding=Math.max(nums[i],MaxEnding+nums[i]);
            MaxSoFar= Math.max(MaxSoFar,MaxEnding);
        }
        return MaxSoFar;
        }
    }
}

//takes 1ms
// class Solution {
//     public int maxSubArray(int[] nums) {
//         int max =nums[0];
//         int sum=0;
//         for(int i=0;i<nums.length;i++){
//             sum +=nums[i];
//             max=sum>max?sum:max;
//             if(sum < 0){
//                 sum = 0;
//             }
//         }
//         return max;
//     }
// }
