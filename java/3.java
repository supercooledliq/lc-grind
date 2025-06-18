//sliding window approach:
// Start with two pointers: left = 0, right = 0
// Use a Set or Map to track characters you've seen in the current window
// Slide right forward one step at a time
// If the character at right is not in the Set, add it and update maxLength(maxLength=Math.max(maxLength,(r-l)+1))
// If it's a duplicate, move the left pointer right until the duplicate is gone from the Set
// Repeat until you reach the end

//example:
// Input: "abcabcbb"
// a, ab, abc → length 3
// then you hit duplicate a, so slide the window to bca, then cab, etc

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int l=0, r=0, maxLength=0;
        HashSet<Character>set=new HashSet<Character>();
        for(r=0;r<s.length();r++)
        {
            while(set.contains(s.charAt(r)))
            {
                set.remove(s.charAt(l));
                l++;
            }
            set.add(s.charAt(r));
            maxLength=Math.max(maxLength,(r-l)+1);
        }
        return maxLength;
        }
    }
