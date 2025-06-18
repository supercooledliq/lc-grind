//idea is to sort all words in ascending order
// ex:- eat->aet
//for tea also-> aet(sortedword)
// use sortedword as the key
//sort all the words in strs[]
//add original words to the list under that key
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String,List<String>> map=new HashMap<>();
        for(int i=0;i<strs.length;i++)
        {
            char[] chars=strs[i].toCharArray();
            Arrays.sort(chars);
            String sorted=new String(chars);
            if(map.containsKey(sorted))
            {
                map.get(sorted).add(strs[i]);
            }
            else
           {
            List<String>newList=new ArrayList<>();
            newList.add(strs[i]);
             map.put(sorted,newList);
           }
        }
        return new ArrayList<>(map.values());
    }
}
