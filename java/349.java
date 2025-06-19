class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        HashSet<Integer> set=new HashSet<>();
        HashSet<Integer> result=new HashSet<>();
        for(int i=0;i<nums1.length;i++)
        {
            set.add(nums1[i]);
        }
        for(int k=0;k<nums2.length;k++)
        {
            if(set.contains(nums2[k]))
            result.add(nums2[k]);
        }
        int[] output = new int[result.size()];
        int j = 0;
        for (int num : result) {
            output[j++] = num;
        }
        return output;
    }
}
